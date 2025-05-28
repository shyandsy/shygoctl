package gogen

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
	apiformat "github.com/shyandsy/shygoctl/api/format"
	"github.com/shyandsy/shygoctl/api/parser"
	apiutil "github.com/shyandsy/shygoctl/api/util"
	"github.com/shyandsy/shygoctl/util"
	"github.com/shyandsy/shygoctl/util/pathx"
	"github.com/spf13/cobra"
)

const tmpFile = "%s-%d"

const apiSpecificationFileName = "api_specification.json"

var (
	tmpDir = path.Join(os.TempDir(), "goctl")
	// VarStringDir describes the directory.
	VarStringDir string
	// TemplatePluginName template plugin name
	TemplatePluginName string
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
	templatePluginName := TemplatePluginName
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

	return DoGenProject(apiFile, templatePluginName, dir, namingStyle, withTest)
}

// DoGenProject gen go project files with api file
func DoGenProject(apiFile, templatePluginName, dir, style string, withTest bool) error {
	api, err := parser.Parse(apiFile)
	if err != nil {
		return err
	}

	if err := api.Validate(); err != nil {
		return err
	}

	specification, err := json.Marshal(*api)
	if err != nil {
		panic("cannot marshal api specification")
	}

	if err := os.WriteFile(apiSpecificationFileName, specification, 0666); err != nil {
		panic("fail to write api specification file")
	}

	// execute template code gen
	cmd := exec.Command(templatePluginName, apiSpecificationFileName)
	stdout, _ := cmd.StdoutPipe()
	_ = cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	_ = cmd.Wait()

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
