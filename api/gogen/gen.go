package gogen

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	apiformat "github.com/shyandsy/shygoctl/api/format"
	"github.com/shyandsy/shygoctl/api/parser"
	"github.com/shyandsy/shygoctl/api/spec"
	apiutil "github.com/shyandsy/shygoctl/api/util"
	"github.com/shyandsy/shygoctl/config"
	"github.com/shyandsy/shygoctl/pkg/golang"
	"github.com/shyandsy/shygoctl/util"
	"github.com/shyandsy/shygoctl/util/pathx"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

const tmpFile = "%s-%d"

var (
	tmpDir = path.Join(os.TempDir(), "goctl")
	// VarStringDir describes the directory.
	VarStringDir string
	// VarStringAPI describes the API.
	VarStringAPI string
	// VarStringHome describes the go home.
	VarStringHome string
	// VarStringRemote describes the remote git repository.
	VarStringRemote string
	// VarStringBranch describes the branch.
	VarStringBranch string
	// VarStringStyle describes the style of output files.
	VarStringStyle  string
	VarBoolWithTest bool
	// VarBoolTypeGroup describes whether to group types.
	VarBoolTypeGroup bool
)

// GoCommand gen go project files from command line
func GoCommand(_ *cobra.Command, _ []string) error {
	apiFile := VarStringAPI
	dir := VarStringDir
	namingStyle := VarStringStyle
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	withTest := VarBoolWithTest
	if len(remote) > 0 {
		repo, _ := util.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}
	if len(apiFile) == 0 {
		return errors.New("missing -api")
	}
	if len(dir) == 0 {
		return errors.New("missing -dir")
	}

	return DoGenProject(apiFile, dir, namingStyle, withTest)
}

// DoGenProject gen go project files with api file
func DoGenProject(apiFile, dir, style string, withTest bool) error {
	api, err := parser.Parse(apiFile)
	if err != nil {
		return err
	}

	if err := api.Validate(); err != nil {
		return err
	}

	// load plugin
	p, err := plugin.Open("gozero_template_plugin.so")
	if err != nil {
		panic(fmt.Sprintf("插件加载失败: %v", err))
	}

	// 查找方法符号
	//DoGenProject
	method, err := p.Lookup("DoGenProject")
	if err != nil {
		panic(fmt.Sprintf("符号查找失败: %v", err))
	}

	doGenProject, ok := method.(func(api *spec.ApiSpec, dir, style string) error)
	if !ok {
		panic("函数签名不匹配，请检查参数类型")
	}

	if err := doGenProject(api, "demo", "goZero"); err != nil {
		panic("函数签名不匹配，请检查参数类型")
	}

	spec, err := json.Marshal(*api)
	if err != nil {
		panic("cannot marshal api specification")
	}
	if err := os.WriteFile("test.json", spec, 0666); err != nil {
		panic("fail to write api specification file")
	}

	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	logx.Must(pathx.MkdirIfNotExist(dir))
	rootPkg, err := golang.GetParentPackage(dir)
	if err != nil {
		return err
	}

	logx.Must(genEtc(dir, cfg, api))
	logx.Must(genConfig(dir, cfg, api))
	logx.Must(genMain(dir, rootPkg, cfg, api))
	logx.Must(genServiceContext(dir, rootPkg, cfg, api))
	logx.Must(genTypes(dir, cfg, api))
	logx.Must(genRoutes(dir, rootPkg, cfg, api))
	logx.Must(genHandlers(dir, rootPkg, cfg, api))
	logx.Must(genLogic(dir, rootPkg, cfg, api))
	logx.Must(genMiddleware(dir, cfg, api))
	if withTest {
		logx.Must(genHandlersTest(dir, rootPkg, cfg, api))
		logx.Must(genLogicTest(dir, rootPkg, cfg, api))
	}

	if err := backupAndSweep(apiFile); err != nil {
		return err
	}

	if err := apiformat.ApiFormatByPath(apiFile, false); err != nil {
		return err
	}

	fmt.Println(color.Green.Render("Done."))
	return nil
}

func backupAndSweep(apiFile string) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(2)
	_ = os.MkdirAll(tmpDir, os.ModePerm)

	go func() {
		_, fileName := filepath.Split(apiFile)
		_, e := apiutil.Copy(apiFile, fmt.Sprintf(path.Join(tmpDir, tmpFile), fileName, time.Now().Unix()))
		if e != nil {
			err = e
		}
		wg.Done()
	}()
	go func() {
		if e := sweep(); e != nil {
			err = e
		}
		wg.Done()
	}()
	wg.Wait()

	return err
}

func sweep() error {
	keepTime := time.Now().AddDate(0, 0, -7)
	return filepath.Walk(tmpDir, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		pos := strings.LastIndexByte(info.Name(), '-')
		if pos > 0 {
			timestamp := info.Name()[pos+1:]
			seconds, err := strconv.ParseInt(timestamp, 10, 64)
			if err != nil {
				// print error and ignore
				fmt.Println(color.Red.Sprintf("sweep ignored file: %s", fpath))
				return nil
			}

			tm := time.Unix(seconds, 0)
			if tm.Before(keepTime) {
				if err := os.RemoveAll(fpath); err != nil {
					fmt.Println(color.Red.Sprintf("failed to remove file: %s", fpath))
					return err
				}
			}
		}

		return nil
	})
}
