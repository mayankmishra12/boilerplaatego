package main

import (
	"go.uber.org/zap"
	"usermvc/routes"
	logger2 "usermvc/utility/logger"

	Config "usermvc/config"

)


func main() {

	loggerMgr := logger2.InitLogger()
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync() // flushes buffer, if any
	logger := loggerMgr.Sugar()
	logger.Debug("START!")

	Config.LoadConfig()
  routes.SetupRouter()
	
}
