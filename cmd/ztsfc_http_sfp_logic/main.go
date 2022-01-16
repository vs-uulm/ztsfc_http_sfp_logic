package main

import (
	"net/http"
    "flag"
    "log"
    "crypto/x509"

    "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/router"
    "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/policies"
    "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/config"
    logger "github.com/vs-uulm/ztsfc_http_logger"
    yt "github.com/leobrada/yaml_tools"
    confInit "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/init"
)

var (
    sysLogger *logger.Logger
)

func init() {
    var confFilePath string
    var policiesFilePath string

    flag.StringVar(&confFilePath, "c", "./config/conf.yml", "Path to user defined yaml config file")
    flag.StringVar(&policiesFilePath, "p", "./policies/policies.yml", "Path to user defined yaml policy file")
    flag.Parse()

    // Loading the YAML config file
    err := yt.LoadYamlFile(confFilePath, &config.Config)
    if err != nil {
        log.Fatalf("main: init(): could not load yaml file: %v", err)
    }

    // Loading the YAML policy file
    err = yt.LoadYamlFile(policiesFilePath, &policies.Policies)
    if err != nil {
        log.Fatalf("main: init(): could not load yaml file: %v", err)
    }

    // Creating the Logger
    confInit.InitSysLoggerParams()
    sysLogger, err = logger.New(config.Config.SysLogger.LogFilePath,
            config.Config.SysLogger.LogLevel,
            config.Config.SysLogger.IfTextFormatter,
            logger.Fields{"type": "system"},
    )
    if err != nil {
            log.Fatalf("main: init(): could not initialize logger: %v", err)
    }

    // Load sfp_logic parameters
    config.Config.Sfpl.CaCertPoolSfplAcceptsFromPep = x509.NewCertPool()
    if err = confInit.InitSfplParams(); err != nil {
        sysLogger.Fatalf("main: init(): could not initialize sfp_logic params: %v", err)
    }

    if err = confInit.InitPolicyParams(); err != nil {
        sysLogger.Fatalf("main: init(): could not initialize policy params: %v", err)
    }

    sysLogger.Info("main: init(): initialization process successfully completed")
}

func main() {
	router := router.NewRouter(sysLogger)

	http.Handle("/", router)

	err := router.ListenAndServeTLS()
	if err != nil {
		sysLogger.Fatalf("main: main(): listen and serve failed: %v", err)
	}
}
