package asanafix

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/goasana/asanacli/cmd/commands"
	"github.com/goasana/asanacli/cmd/commands/version"
	asanaLogger "github.com/goasana/asanacli/logger"
	"github.com/goasana/asanacli/logger/colors"
)

var CmdFix = &commands.Command{
	UsageLine: "fix",
	Short:     "Fixes your application by making it compatible with newer versions of Asana",
	Long: `As of {{"Asana 1.0"|bold}}, there are some backward compatibility issues.

  The command 'fix' will try to solve those issues by upgrading your code base
  to be compatible  with Asana version 1.0+.
`,
}

func init() {
	CmdFix.Run = runFix
	CmdFix.PreRun = func(cmd *commands.Command, args []string) { version.ShowShortVersionBanner() }
	commands.AvailableCommands = append(commands.AvailableCommands, CmdFix)
}

func runFix(cmd *commands.Command, args []string) int {
	output := cmd.Out()

	asanaLogger.Log.Info("Upgrading the application...")

	dir, err := os.Getwd()
	if err != nil {
		asanaLogger.Log.Fatalf("Error while getting the current working directory: %s", err)
	}

	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), ".exe") {
			return nil
		}
		err = fixFile(path)
		_, _ = fmt.Fprintf(output, colors.GreenBold("\tfix\t")+"%s\n", path)
		if err != nil {
			asanaLogger.Log.Errorf("Could not fix file: %s", err)
		}
		return err
	})
	asanaLogger.Log.Success("Upgrade Done!")
	return 0
}

var rules = []string{
	"asana.AppName", "asana.BConfig.AppName",
	"asana.RunMode", "asana.BConfig.RunMode",
	"asana.RecoverPanic", "asana.BConfig.RecoverPanic",
	"asana.RouterCaseSensitive", "asana.BConfig.RouterCaseSensitive",
	"asana.AsanaServerName", "asana.BConfig.ServerName",
	"asana.EnableGzip", "asana.BConfig.EnableGzip",
	"asana.ErrorsShow", "asana.BConfig.EnableErrorsShow",
	"asana.CopyRequestBody", "asana.BConfig.CopyRequestBody",
	"asana.MaxMemory", "asana.BConfig.MaxMemory",
	"asana.Graceful", "asana.BConfig.Listen.Graceful",
	"asana.HttpAddr", "asana.BConfig.Listen.HTTPAddr",
	"asana.HttpPort", "asana.BConfig.Listen.HTTPPort",
	"asana.ListenTCP4", "asana.BConfig.Listen.ListenTCP4",
	"asana.EnableHttpListen", "asana.BConfig.Listen.EnableHTTP",
	"asana.EnableHttpTLS", "asana.BConfig.Listen.EnableHTTPS",
	"asana.HttpsAddr", "asana.BConfig.Listen.HTTPSAddr",
	"asana.HttpsPort", "asana.BConfig.Listen.HTTPSPort",
	"asana.HttpCertFile", "asana.BConfig.Listen.HTTPSCertFile",
	"asana.HttpKeyFile", "asana.BConfig.Listen.HTTPSKeyFile",
	"asana.EnableAdmin", "asana.BConfig.Listen.EnableAdmin",
	"asana.AdminHttpAddr", "asana.BConfig.Listen.AdminAddr",
	"asana.AdminHttpPort", "asana.BConfig.Listen.AdminPort",
	"asana.UseFcgi", "asana.BConfig.Listen.EnableFcgi",
	"asana.HttpServerTimeOut", "asana.BConfig.Listen.ServerTimeOut",
	"asana.AutoRender", "asana.BConfig.WebConfig.AutoRender",
	"asana.ViewsPath", "asana.BConfig.WebConfig.ViewsPath",
	"asana.StaticDir", "asana.BConfig.WebConfig.StaticDir",
	"asana.StaticExtensionsToGzip", "asana.BConfig.WebConfig.StaticExtensionsToGzip",
	"asana.DirectoryIndex", "asana.BConfig.WebConfig.DirectoryIndex",
	"asana.FlashName", "asana.BConfig.WebConfig.FlashName",
	"asana.FlashSeperator", "asana.BConfig.WebConfig.FlashSeparator",
	"asana.EnableDocs", "asana.BConfig.WebConfig.EnableDocs",
	"asana.XSRFKEY", "asana.BConfig.WebConfig.XSRFKey",
	"asana.EnableXSRF", "asana.BConfig.WebConfig.EnableXSRF",
	"asana.XSRFExpire", "asana.BConfig.WebConfig.XSRFExpire",
	"asana.TemplateLeft", "asana.BConfig.WebConfig.TemplateLeft",
	"asana.TemplateRight", "asana.BConfig.WebConfig.TemplateRight",
	"asana.SessionOn", "asana.BConfig.WebConfig.Session.SessionOn",
	"asana.SessionProvider", "asana.BConfig.WebConfig.Session.SessionProvider",
	"asana.SessionName", "asana.BConfig.WebConfig.Session.SessionName",
	"asana.SessionGCMaxLifetime", "asana.BConfig.WebConfig.Session.SessionGCMaxLifetime",
	"asana.SessionSavePath", "asana.BConfig.WebConfig.Session.SessionProviderConfig",
	"asana.SessionCookieLifeTime", "asana.BConfig.WebConfig.Session.SessionCookieLifeTime",
	"asana.SessionAutoSetCookie", "asana.BConfig.WebConfig.Session.SessionAutoSetCookie",
	"asana.SessionDomain", "asana.BConfig.WebConfig.Session.SessionDomain",
	"Ctx.Input.CopyBody(", "Ctx.Input.CopyBody(asana.BConfig.MaxMemory",
	".UrlFor(", ".URLFor(",
	".ServeJson(", ".ServeJSON(",
	".ServeXml(", ".ServeXML(",
	".ServeJsonp(", ".ServeJSONP(",
	".XsrfToken(", ".XSRFToken(",
	".CheckXsrfCookie(", ".CheckXSRFCookie(",
	".XsrfFormHtml(", ".XSRFFormHTML(",
	"asana.UrlFor(", "asana.URLFor(",
	"asana.GlobalDocApi", "asana.GlobalDocAPI",
	"asana.Errorhandler", "asana.ErrorHandler",
	"Output.Jsonp(", "Output.JSONP(",
	"Output.Json(", "Output.JSON(",
	"Output.Xml(", "Output.XML(",
	"Input.Uri()", "Input.URI()",
	"Input.Url()", "Input.URL()",
	"Input.AcceptsHtml()", "Input.AcceptsHTML()",
	"Input.AcceptsXml()", "Input.AcceptsXML()",
	"Input.AcceptsJson()", "Input.AcceptsJSON()",
	"Ctx.XsrfToken()", "Ctx.XSRFToken()",
	"Ctx.CheckXsrfCookie()", "Ctx.CheckXSRFCookie()",
	"session.SessionStore", "session.Store",
	".TplNames", ".TplName",
	"swagger.ApiRef", "swagger.APIRef",
	"swagger.ApiDeclaration", "swagger.APIDeclaration",
	"swagger.Api", "swagger.API",
	"swagger.ApiRef", "swagger.APIRef",
	"swagger.Infomation", "swagger.Information",
	"toolbox.UrlMap", "toolbox.URLMap",
	"logs.LoggerInterface", "logs.Logger",
	"Input.Request", "Input.Context.Request",
	"Input.Params)", "Input.Params())",
	"httplib.AsanaHttpSettings", "httplib.AsanaHTTPSettings",
	"httplib.AsanaHttpRequest", "httplib.AsanaHTTPRequest",
	".TlsClientConfig", ".TLSClientConfig",
	".JsonBody", ".JSONBody",
	".ToJson", ".ToJSON",
	".ToXml", ".ToXML",
	"asana.Html2str", "asana.HTML2str",
	"asana.AssetsCss", "asana.AssetsCSS",
	"orm.DR_Sqlite", "orm.DRSqlite",
	"orm.DR_Postgres", "orm.DRPostgres",
	"orm.DR_MySQL", "orm.DRMySQL",
	"orm.DR_Oracle", "orm.DROracle",
	"orm.Col_Add", "orm.ColAdd",
	"orm.Col_Minus", "orm.ColMinus",
	"orm.Col_Multiply", "orm.ColMultiply",
	"orm.Col_Except", "orm.ColExcept",
	"GenerateOperatorSql", "GenerateOperatorSQL",
	"OperatorSql", "OperatorSQL",
	"orm.Debug_Queries", "orm.DebugQueries",
	"orm.COMMA_SPACE", "orm.CommaSpace",
	".SendOut()", ".DoRequest()",
	"validation.ValidationError", "validation.Error",
}

func fixFile(file string) error {
	rp := strings.NewReplacer(rules...)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	fixed := rp.Replace(string(content))

	// Forword the RequestBody from the replace
	// "Input.Request", "Input.Context.Request",
	fixed = strings.Replace(fixed, "Input.Context.RequestBody", "Input.RequestBody", -1)

	// Regexp replace
	pareg := regexp.MustCompile(`(Input.Params\[")(.*)("])`)
	fixed = pareg.ReplaceAllString(fixed, "Input.Param(\"$2\")")
	pareg = regexp.MustCompile(`(Input.Data\[\")(.*)(\"\])(\s)(=)(\s)(.*)`)
	fixed = pareg.ReplaceAllString(fixed, "Input.SetData(\"$2\", $7)")
	pareg = regexp.MustCompile(`(Input.Data\[\")(.*)(\"\])`)
	fixed = pareg.ReplaceAllString(fixed, "Input.Data(\"$2\")")
	// Fix the cache object Put method
	pareg = regexp.MustCompile(`(\.Put\(\")(.*)(\",)(\s)(.*)(,\s*)([^\*.]*)(\))`)
	if pareg.MatchString(fixed) && strings.HasSuffix(file, ".go") {
		fixed = pareg.ReplaceAllString(fixed, ".Put(\"$2\", $5, $7*time.Second)")
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)
		if err != nil {
			panic(err)
		}
		// Print the imports from the file's AST.
		hasTimepkg := false
		for _, s := range f.Imports {
			if s.Path.Value == `"time"` {
				hasTimepkg = true
				break
			}
		}
		if !hasTimepkg {
			fixed = strings.Replace(fixed, "import (", "import (\n\t\"time\"", 1)
		}
	}
	// Replace the v.Apis in docs.go
	if strings.Contains(file, "docs.go") {
		fixed = strings.Replace(fixed, "v.Apis", "v.APIs", -1)
	}
	// Replace the config file
	if strings.HasSuffix(file, ".yaml") {
		fixed = strings.Replace(fixed, "HttpCertFile", "HTTPSCertFile", -1)
		fixed = strings.Replace(fixed, "HttpKeyFile", "HTTPSKeyFile", -1)
		fixed = strings.Replace(fixed, "EnableHttpListen", "HTTPEnable", -1)
		fixed = strings.Replace(fixed, "EnableHttpTLS", "EnableHTTPS", -1)
		fixed = strings.Replace(fixed, "EnableHttpTLS", "EnableHTTPS", -1)
		fixed = strings.Replace(fixed, "AsanaServerName", "ServerName", -1)
		fixed = strings.Replace(fixed, "AdminHttpAddr", "AdminAddr", -1)
		fixed = strings.Replace(fixed, "AdminHttpPort", "AdminPort", -1)
		fixed = strings.Replace(fixed, "HttpServerTimeOut", "ServerTimeOut", -1)
	}
	err = os.Truncate(file, 0)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, []byte(fixed), 0666)
}
