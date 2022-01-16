package init

import (
    "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/config"
)

// InitSysLoggerParams() sets default values for the system logger parameters
// The function should be called before the system logger creation!
// Error handling in the case of unsupported options is done by logrus itself
func InitSysLoggerParams() {
        // Set a default value of a logging level parameter
        if config.Config.SysLogger.LogLevel == "" {
            config.Config.SysLogger.LogLevel = "info"
        }

        // Set a default value of a log messages destination parameter
        if config.Config.SysLogger.LogFilePath == "" {
                config.Config.SysLogger.LogFilePath = "stdout"
        }

        // Set a default value of a log messages formatter parameter
        if config.Config.SysLogger.IfTextFormatter == "" {
                config.Config.SysLogger.IfTextFormatter = "json"
        }
}
