package main

import (
	"fmt"
	"go.uber.org/zap"
	Config "usermvc/config"
	"usermvc/routes"
)


func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"myproject.log",
	}
	return cfg.Build()
}
func main() {
	loggerMgr,err := NewLogger()
	if err != nil {
		fmt.Println("erroe ", err)
	}
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync() // flushes buffer, if any
	logger := loggerMgr.Sugar()
	logger.Info("START!")
	Config.LoadConfig()
	routes.SetupRouter().Run(":8080")

}
