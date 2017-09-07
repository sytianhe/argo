// Copyright 2015-2016 Applatix, Inc. All rights reserved.
package axrc

import (
	"os"
	"time"

	"applatix.io/axdb/axdbcl"
	axopsutils "applatix.io/axops/utils"
	"applatix.io/axrc/utils"
	"applatix.io/common"
	"applatix.io/restcl"
)

var AX_NAMESPACE string
var AX_VERSION string

func Init(database string, kafka string, axops string) {
	utils.InitLoggers("axrc")

	if len(database) != 0 {
		utils.DbCl = axdbcl.NewAXDBClientWithTimeout(database, 5*time.Minute)
		axopsutils.Dbcl = utils.DbCl
	}
	if len(axops) != 0 {
		utils.AxopsCl = restcl.NewRestClientWithTimeout(axops, 5*time.Minute)
	}

	AX_NAMESPACE = common.GetAxNameSpace()
	AX_VERSION = common.GetAxVersion()
	appName := common.GetApplicationName()

	if appName != "" && appName != "${APPLICATION_NAME}" {
		utils.APPLICATION_NAME = appName
	}

	if AX_NAMESPACE == "" {
		utils.ErrorLog.Printf("AX_NAMESAPCE is not available from environment variables. Abort.")
		os.Exit(1)
	}

	if AX_VERSION == "" {
		utils.ErrorLog.Printf("AX_VERSION is not available from environment variables. Abort.")
		os.Exit(1)
	}

	utils.InfoLog.Printf("AXRC AX_NAMESPACE:%s AX_VERSION:%s\n", AX_NAMESPACE, AX_VERSION)
}
