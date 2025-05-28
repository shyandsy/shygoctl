package api

import (
	"github.com/spf13/cobra"

	"github.com/shyandsy/shygoctl/api/apigen"
	"github.com/shyandsy/shygoctl/api/docgen"
	"github.com/shyandsy/shygoctl/api/format"
	"github.com/shyandsy/shygoctl/api/gogen"
	"github.com/shyandsy/shygoctl/api/new"
	"github.com/shyandsy/shygoctl/api/swagger"
	"github.com/shyandsy/shygoctl/api/validate"
	"github.com/shyandsy/shygoctl/config"
	"github.com/shyandsy/shygoctl/internal/cobrax"
	"github.com/shyandsy/shygoctl/plugin"
)

var (
	// Cmd describes an api command.
	Cmd = cobrax.NewCommand("api", cobrax.WithRunE(apigen.CreateApiTemplate))
	//dartCmd   = cobrax.NewCommand("dart", cobrax.WithRunE(dartgen.DartCommand))
	docCmd    = cobrax.NewCommand("doc", cobrax.WithRunE(docgen.DocCommand))
	formatCmd = cobrax.NewCommand("format", cobrax.WithRunE(format.GoFormatApi))
	genCmd    = cobrax.NewCommand("gen", cobrax.WithRunE(gogen.GoCommand))
	newCmd    = cobrax.NewCommand("new", cobrax.WithRunE(new.CreateServiceCommand),
		cobrax.WithArgs(cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)))
	validateCmd = cobrax.NewCommand("validate", cobrax.WithRunE(validate.GoValidateApi))
	//ktCmd       = cobrax.NewCommand("kt", cobrax.WithRunE(ktgen.KtCommand))
	pluginCmd  = cobrax.NewCommand("plugin", cobrax.WithRunE(plugin.PluginCommand))
	swaggerCmd = cobrax.NewCommand("swagger", cobrax.WithRunE(swagger.Command))
)

func init() {
	var (
		apiCmdFlags = Cmd.Flags()
		//dartCmdFlags     = dartCmd.Flags()
		docCmdFlags    = docCmd.Flags()
		formatCmdFlags = formatCmd.Flags()
		genCmdFlags    = genCmd.Flags()
		//ktCmdFlags       = ktCmd.Flags()
		newCmdFlags      = newCmd.Flags()
		pluginCmdFlags   = pluginCmd.Flags()
		validateCmdFlags = validateCmd.Flags()
		swaggerCmdFlags  = swaggerCmd.Flags()
	)

	apiCmdFlags.StringVar(&apigen.VarStringOutput, "o")
	apiCmdFlags.StringVar(&apigen.VarStringHome, "home")
	apiCmdFlags.StringVar(&apigen.VarStringRemote, "remote")
	apiCmdFlags.StringVar(&apigen.VarStringBranch, "branch")

	//dartCmdFlags.StringVar(&dartgen.VarStringDir, "dir")
	//dartCmdFlags.StringVar(&dartgen.VarStringAPI, "api")
	//dartCmdFlags.BoolVar(&dartgen.VarStringLegacy, "legacy")
	//dartCmdFlags.StringVar(&dartgen.VarStringHostname, "hostname")
	//dartCmdFlags.StringVar(&dartgen.VarStringScheme, "scheme")

	docCmdFlags.StringVar(&docgen.VarStringDir, "dir")
	docCmdFlags.StringVar(&docgen.VarStringOutput, "o")

	formatCmdFlags.StringVar(&format.VarStringDir, "dir")
	formatCmdFlags.BoolVar(&format.VarBoolIgnore, "iu")
	formatCmdFlags.BoolVar(&format.VarBoolUseStdin, "stdin")
	formatCmdFlags.BoolVar(&format.VarBoolSkipCheckDeclare, "declare")

	genCmdFlags.StringVar(&gogen.VarStringDir, "dir")
	genCmdFlags.StringVar(&gogen.TemplatePluginName, "tpn")
	genCmdFlags.StringVar(&gogen.VarStringAPI, "api")
	genCmdFlags.StringVar(&gogen.VarStringHome, "home")
	genCmdFlags.StringVar(&gogen.VarStringRemote, "remote")
	genCmdFlags.StringVar(&gogen.VarStringBranch, "branch")
	genCmdFlags.BoolVar(&gogen.VarBoolWithTest, "test")
	genCmdFlags.BoolVar(&gogen.VarBoolTypeGroup, "type-group")
	genCmdFlags.StringVarWithDefaultValue(&gogen.VarStringStyle, "style", config.DefaultFormat)

	//ktCmdFlags.StringVar(&ktgen.VarStringDir, "dir")
	//ktCmdFlags.StringVar(&ktgen.VarStringAPI, "api")
	//ktCmdFlags.StringVar(&ktgen.VarStringPKG, "pkg")

	newCmdFlags.StringVar(&new.VarStringHome, "home")
	newCmdFlags.StringVar(&new.VarStringRemote, "remote")
	newCmdFlags.StringVar(&new.VarStringBranch, "branch")
	newCmdFlags.StringVarWithDefaultValue(&new.VarStringStyle, "style", config.DefaultFormat)

	pluginCmdFlags.StringVarP(&plugin.VarStringPlugin, "plugin", "p")
	pluginCmdFlags.StringVar(&plugin.VarStringDir, "dir")
	pluginCmdFlags.StringVar(&plugin.VarStringAPI, "api")
	pluginCmdFlags.StringVar(&plugin.VarStringStyle, "style")

	swaggerCmdFlags.StringVar(&swagger.VarStringAPI, "api")
	swaggerCmdFlags.StringVar(&swagger.VarStringDir, "dir")
	swaggerCmdFlags.StringVar(&swagger.VarStringFilename, "filename")
	swaggerCmdFlags.BoolVar(&swagger.VarBoolYaml, "yaml")

	validateCmdFlags.StringVar(&validate.VarStringAPI, "api")

	// Add sub-commands
	Cmd.AddCommand(docCmd, formatCmd, genCmd, newCmd, pluginCmd, validateCmd, swaggerCmd)
}
