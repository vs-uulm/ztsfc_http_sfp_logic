package config

import (
    "crypto/x509"
    "crypto/tls"
)

var (
    Config ConfigT
)

type ConfigT struct {
    SysLogger sysLoggerT `yaml:"system_logger"`
    Sfpl sfpLogicT `yaml:"sfp_logic"`
}

type sysLoggerT struct {
    LogLevel string `yaml:"system_logger_logging_level"`
    LogFilePath string `yaml:"system_logger_destination"`
    IfTextFormatter string `yaml:"system_logger_format"`
}

type sfpLogicT struct {
    ListenAddr string `yaml:"listen_addr"`
    CertsSfplAcceptsWhenShownByPep []string `yaml:"certs_sfpl_accepts_when_shown_by_pep"`
    CertShownBySfplToPep string `yaml:"cert_shown_by_sfpl_to_pep"`
    PrivkeyForCertShownBySfplToPep  string  `yaml:"privkey_for_cert_shown_by_sfpl_to_pep"`

    CaCertPoolSfplAcceptsFromPep *x509.CertPool
    X509KeyPairShownBySfplToPep  tls.Certificate
}
