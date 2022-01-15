package config

import (
)

var (
    Config ConfigT
)

type sysLoggerT struct {
    LogLevel string `yaml:"system_logger_logging_level"`
    LogFilePath string `yaml:"system_logger_destination"`
    IfTextFormatter string `yaml:"system_logger_format"`
}

type ConfigT struct {
    SysLogger sysLoggerT `yaml:"system_logger"`
}
