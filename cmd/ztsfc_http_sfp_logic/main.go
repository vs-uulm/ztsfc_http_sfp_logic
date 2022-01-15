package main

import (
	"net/http"
    "flag"
    "log"

    "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/router"
    "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/config"
    logger "github.com/vs-uulm/ztsfc_http_logger"
    yt "github.com/leobrada/yaml_tools"
)

var (
    sysLogger *logger.Logger
)

func init() {
    var confFilePath string

    flag.StringVar(&confFilePath, "c", "./config/conf.yml", "Path to user defined yaml config file")
    flag.Parse()

    // Loading the YAML config File
    err := yt.LoadYamlFile(confFilePath, &config.Config)
    if err != nil {
        log.Fatalf("main: init(): could not load yaml file: %v", err)
    }

    // Creating the Logger
    sysLogger, err = logger.New(config.Config.SysLogger.LogFilePath,
            config.Config.SysLogger.LogLevel,
            config.Config.SysLogger.IfTextFormatter,
            logger.Fields{"type": "system"},
    )
    if err != nil {
            log.Fatalf("main: init(): could not initialize logger: %v", err)
    }

    sysLogger.Info("main: init(): initialization process successfully completed")
}

func main() {
	router, err := router.NewRouter()
	if err != nil {
		sysLogger.Fatalf("THIS IS JUST A PLACEHOLDER")
	}

	http.Handle("/", router)

	err = router.ListenAndServeTLS()
	if err != nil {
		sysLogger.Fatalf("main: main(): listen and serve failed: %w", err)
	}
}
