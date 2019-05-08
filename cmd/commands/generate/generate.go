// Copyright 2019 asana authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package generate

import (
	"os"
	"strings"

	"github.com/goasana/asanacli/cmd/commands"
	"github.com/goasana/asanacli/cmd/commands/version"
	"github.com/goasana/asanacli/config"
	"github.com/goasana/asanacli/generate"
	"github.com/goasana/asanacli/generate/swaggergen"
	"github.com/goasana/asanacli/logger"
	"github.com/goasana/asanacli/utils"
)

var CmdGenerate = &commands.Command{
	UsageLine: "generate [command]",
	Short:     "Source code generator",
	Long: `▶ {{"To scaffold out your entire application:"|bold}}

     $ asana generate scaffold [scaffoldname] [-fields="title:string,body:text"] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]

  ▶ {{"To generate a Model based on fields:"|bold}}

     $ asana generate model [modelname] [-fields="name:type"]

  ▶ {{"To generate a controller:"|bold}}

     $ asana generate controller [controllerfile]

  ▶ {{"To generate a CRUD view:"|bold}}

     $ asana generate view [viewpath]

  ▶ {{"To generate a migration file for making database schema updates:"|bold}}

     $ asana generate migration [migrationfile] [-fields="name:type"]

  ▶ {{"To generate swagger doc file:"|bold}}

     $ asana generate docs

  ▶ {{"To generate a test case:"|bold}}

     $ asana generate test [routerfile]

  ▶ {{"To generate appcode based on an existing database:"|bold}}

     $ asana generate appcode [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]
`,
	PreRun: func(cmd *commands.Command, args []string) { version.ShowShortVersionBanner() },
	Run:    GenerateCode,
}

func init() {
	CmdGenerate.Flag.Var(&generate.Tables, "tables", "List of table names separated by a comma.")
	CmdGenerate.Flag.Var(&generate.SQLDriver, "driver", "Database SQLDriver. Either mysql, postgres or sqlite.")
	CmdGenerate.Flag.Var(&generate.SQLConn, "conn", "Connection string used by the SQLDriver to connect to a database instance.")
	CmdGenerate.Flag.Var(&generate.Level, "level", "Either 1, 2 or 3. i.e. 1=models; 2=models and controllers; 3=models, controllers and routers.")
	CmdGenerate.Flag.Var(&generate.Fields, "fields", "List of table Fields.")
	CmdGenerate.Flag.Var(&generate.DDL, "ddl", "Generate DDL Migration")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdGenerate)
}

func GenerateCode(cmd *commands.Command, args []string) int {
	currPath, _ := os.Getwd()
	if len(args) < 1 {
		asanaLogger.Log.Fatal("Command is missing")
	}

	gps := utils.GetGOPATHs()
	if len(gps) == 0 {
		asanaLogger.Log.Fatal("GOPATH environment variable is not set or empty")
	}

	gopath := gps[0]

	asanaLogger.Log.Debugf("GOPATH: %s", utils.FILE(), utils.LINE(), gopath)

	gcmd := args[0]
	switch gcmd {
	case "scaffold":
		scaffold(cmd, args, currPath)
	case "docs":
		swaggergen.GenerateDocs(currPath)
	case "appcode":
		appCode(cmd, args, currPath)
	case "migration":
		migration(cmd, args, currPath)
	case "controller":
		controller(args, currPath)
	case "model":
		model(cmd, args, currPath)
	case "view":
		view(args, currPath)
	default:
		asanaLogger.Log.Fatal("Command is missing")
	}
	asanaLogger.Log.Successf("%s successfully generated!", strings.Title(gcmd))
	return 0
}

func scaffold(cmd *commands.Command, args []string, currPath string) {
	if len(args) < 2 {
		asanaLogger.Log.Fatal("Wrong number of arguments. Run: asana help generate")
	}

	_ = cmd.Flag.Parse(args[2:])
	if generate.SQLDriver == "" {
		generate.SQLDriver = utils.DocValue(config.Conf.Database.Driver)
		if generate.SQLDriver == "" {
			generate.SQLDriver = "mysql"
		}
	}
	if generate.SQLConn == "" {
		generate.SQLConn = utils.DocValue(config.Conf.Database.Conn)
		if generate.SQLConn == "" {
			generate.SQLConn = "root:@tcp(127.0.0.1:3306)/test"
		}
	}
	if generate.Fields == "" {
		asanaLogger.Log.Hint("Fields option should not be empty, i.e. -Fields=\"title:string,body:text\"")
		asanaLogger.Log.Fatal("Wrong number of arguments. Run: asana help generate")
	}
	sname := args[1]
	generate.GenerateScaffold(sname, generate.Fields.String(), currPath, generate.SQLDriver.String(), generate.SQLConn.String())
}

func appCode(cmd *commands.Command, args []string, currPath string) {
	_ = cmd.Flag.Parse(args[1:])
	if generate.SQLDriver == "" {
		generate.SQLDriver = utils.DocValue(config.Conf.Database.Driver)
		if generate.SQLDriver == "" {
			generate.SQLDriver = "mysql"
		}
	}
	if generate.SQLConn == "" {
		generate.SQLConn = utils.DocValue(config.Conf.Database.Conn)
		if generate.SQLConn == "" {
			if generate.SQLDriver == "mysql" {
				generate.SQLConn = "root:@tcp(127.0.0.1:3306)/test"
			} else if generate.SQLDriver == "postgres" {
				generate.SQLConn = "postgres://postgres:postgres@127.0.0.1:5432/postgres"
			}
		}
	}
	if generate.Level == "" {
		generate.Level = "3"
	}
	asanaLogger.Log.Infof("Using '%s' as 'SQLDriver'", generate.SQLDriver)
	asanaLogger.Log.Infof("Using '%s' as 'SQLConn'", generate.SQLConn)
	asanaLogger.Log.Infof("Using '%s' as 'Tables'", generate.Tables)
	asanaLogger.Log.Infof("Using '%s' as 'Level'", generate.Level)
	generate.GenerateAppcode(generate.SQLDriver.String(), generate.SQLConn.String(), generate.Level.String(), generate.Tables.String(), currPath)
}

func migration(cmd *commands.Command, args []string, currPath string) {
	if len(args) < 2 {
		asanaLogger.Log.Fatal("Wrong number of arguments. Run: asanacli help generate")
	}
	_ = cmd.Flag.Parse(args[2:])
	mname := args[1]

	asanaLogger.Log.Infof("Using '%s' as migration name", mname)

	upsql := ""
	downsql := ""
	if generate.Fields != "" {
		dbMigrator := generate.NewDBDriver()
		upsql = dbMigrator.GenerateCreateUp(mname)
		downsql = dbMigrator.GenerateCreateDown(mname)
	}
	generate.GenerateMigration(mname, upsql, downsql, currPath)
}

func controller(args []string, currPath string) {
	if len(args) == 2 {
		cname := args[1]
		generate.GenerateController(cname, currPath)
	} else {
		asanaLogger.Log.Fatal("Wrong number of arguments. Run: asanacli help generate")
	}
}

func model(cmd *commands.Command, args []string, currPath string) {
	if len(args) < 2 {
		asanaLogger.Log.Fatal("Wrong number of arguments. Run: asanacli help generate")
	}
	cmd.Flag.Parse(args[2:])
	if generate.Fields == "" {
		asanaLogger.Log.Hint("Fields option should not be empty, i.e. -Fields=\"title:string,body:text\"")
		asanaLogger.Log.Fatal("Wrong number of arguments. Run: asanacli help generate")
	}
	sname := args[1]
	generate.GenerateModel(sname, generate.Fields.String(), currPath)
}

func view(args []string, currPath string) {
	if len(args) == 2 {
		cname := args[1]
		generate.GenerateView(cname, currPath)
	} else {
		asanaLogger.Log.Fatal("Wrong number of arguments. Run: asanacli help generate")
	}
}
