package hprose

import (
	"os"

	"fmt"
	"path"
	"strings"

	"github.com/goasana/asanacli/cmd/commands"
	"github.com/goasana/asanacli/cmd/commands/api"
	"github.com/goasana/asanacli/cmd/commands/version"
	"github.com/goasana/asanacli/generate"
	asanaLogger "github.com/goasana/asanacli/logger"
	"github.com/goasana/asanacli/utils"
)

var CmdHproseapp = &commands.Command{
	// CustomFlags: true,
	UsageLine: "hprose [appname]",
	Short:     "Creates an RPC application based on Hprose and Asana frameworks",
	Long: `
  The command 'hprose' creates an RPC application based on both Asana and Hprose (http://hprose.com/).

  {{"To scaffold out your application, use:"|bold}}

      $ asana hprose [appname] [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]

  If 'conn' is empty, the command will generate a sample application. Otherwise the command
  will connect to your database and generate models based on the existing tables.

  The command 'hprose' creates a folder named [appname] with the following structure:

	    ├── main.go
	    ├── {{"conf"|foldername}}
	    │     └── app.yaml
	    └── {{"models"|foldername}}
	          └── object.go
	          └── user.go
`,
	PreRun: func(cmd *commands.Command, args []string) { version.ShowShortVersionBanner() },
	Run:    createhprose,
}

func init() {
	CmdHproseapp.Flag.Var(&generate.Tables, "tables", "List of table names separated by a comma.")
	CmdHproseapp.Flag.Var(&generate.SQLDriver, "driver", "Database driver. Either mysql, postgres or sqlite.")
	CmdHproseapp.Flag.Var(&generate.SQLConn, "conn", "Connection string used by the driver to connect to a database instance.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdHproseapp)
}

func createhprose(cmd *commands.Command, args []string) int {
	output := cmd.Out()

	if len(args) != 1 {
		asanaLogger.Log.Fatal("Argument [appname] is missing")
	}

	curPath, _ := os.Getwd()
	if len(args) > 1 {
		_ = cmd.Flag.Parse(args[1:])
	}
	appPath, packpath, err := utils.CheckEnv(args[0])
	if err != nil {
		asanaLogger.Log.Fatalf("%s", err)
	}
	if generate.SQLDriver == "" {
		generate.SQLDriver = "mysql"
	}
	asanaLogger.Log.Info("Creating Hprose application...")

	_ = os.MkdirAll(appPath, 0755)
	_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", appPath, "\x1b[0m")
	_ = os.Mkdir(path.Join(appPath, "conf"), 0755)
	_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "conf"), "\x1b[0m")
	_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "conf", "app.yaml"), "\x1b[0m")
	utils.WriteToFile(path.Join(appPath, "conf", "app.yaml"), strings.Replace(generate.Hproseconf, "{{.Appname}}", args[0], -1))

	if generate.SQLConn != "" {
		asanaLogger.Log.Infof("Using '%s' as 'driver'", generate.SQLDriver)
		asanaLogger.Log.Infof("Using '%s' as 'conn'", generate.SQLConn)
		asanaLogger.Log.Infof("Using '%s' as 'tables'", generate.Tables)
		generate.GenerateHproseAppcode(string(generate.SQLDriver), string(generate.SQLConn), "1", string(generate.Tables), path.Join(curPath, args[0]))

		_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "main.go"), "\x1b[0m")
		maingoContent := strings.Replace(generate.HproseMainconngo, "{{.Appname}}", packpath, -1)
		maingoContent = strings.Replace(maingoContent, "{{.DriverName}}", string(generate.SQLDriver), -1)
		maingoContent = strings.Replace(maingoContent, "{{HproseFunctionList}}", strings.Join(generate.HproseAddFunctions, ""), -1)
		if generate.SQLDriver == "mysql" {
			maingoContent = strings.Replace(maingoContent, "{{.DriverPkg}}", `_ "github.com/go-sql-driver/mysql"`, -1)
		} else if generate.SQLDriver == "postgres" {
			maingoContent = strings.Replace(maingoContent, "{{.DriverPkg}}", `_ "github.com/lib/pq"`, -1)
		}
		utils.WriteToFile(path.Join(appPath, "main.go"),
			strings.Replace(
				maingoContent,
				"{{.conn}}",
				generate.SQLConn.String(),
				-1,
			),
		)
	} else {
		_ = os.Mkdir(path.Join(appPath, "models"), 0755)
		_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "models"), "\x1b[0m")

		_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "models", "object.go"), "\x1b[0m")
		utils.WriteToFile(path.Join(appPath, "models", "object.go"), apiapp.APIModels)

		_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "models", "user.go"), "\x1b[0m")
		utils.WriteToFile(path.Join(appPath, "models", "user.go"), apiapp.APIModels2)

		_, _ = fmt.Fprintf(output, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path.Join(appPath, "main.go"), "\x1b[0m")
		utils.WriteToFile(path.Join(appPath, "main.go"),
			strings.Replace(generate.HproseMaingo, "{{.Appname}}", packpath, -1))
	}
	asanaLogger.Log.Success("New Hprose application successfully created!")
	return 0
}
