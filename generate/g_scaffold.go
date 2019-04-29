package generate

import (
	"strings"

	"github.com/goasana/asana/cmd/commands/migrate"
	asanaLogger "github.com/goasana/asana/logger"
	"github.com/goasana/asana/utils"
)

func GenerateScaffold(sname, fields, currpath, driver, conn string) {
	asanaLogger.Log.Infof("Do you want to create a '%s' model? [Yes|No] ", sname)

	// Generate the model
	if utils.AskForConfirmation() {
		GenerateModel(sname, fields, currpath)
	}

	// Generate the controller
	asanaLogger.Log.Infof("Do you want to create a '%s' controller? [Yes|No] ", sname)
	if utils.AskForConfirmation() {
		GenerateController(sname, currpath)
	}

	// Generate the views
	asanaLogger.Log.Infof("Do you want to create views for this '%s' resource? [Yes|No] ", sname)
	if utils.AskForConfirmation() {
		GenerateView(sname, currpath)
	}

	// Generate a migration
	asanaLogger.Log.Infof("Do you want to create a '%s' migration and schema for this resource? [Yes|No] ", sname)
	if utils.AskForConfirmation() {
		upsql := ""
		downsql := ""
		if fields != "" {
			dbMigrator := NewDBDriver()
			upsql = dbMigrator.GenerateCreateUp(sname)
			downsql = dbMigrator.GenerateCreateDown(sname)
		}
		GenerateMigration(sname, upsql, downsql, currpath)
	}

	// Run the migration
	asanaLogger.Log.Infof("Do you want to migrate the database? [Yes|No] ")
	if utils.AskForConfirmation() {
		migrate.MigrateUpdate(currpath, driver, conn, "")
	}
	asanaLogger.Log.Successf("All done! Don't forget to add  asana. Router(\"/%s\" ,&controllers.%sController{}) to routers/route.go\n", sname, strings.Title(sname))
}
