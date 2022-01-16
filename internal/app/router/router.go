package router

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

    logger "github.com/vs-uulm/ztsfc_http_logger"
	md "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/metadata"
	sfpl "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/sfp_logic"
	"github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/config"
)

const (
	// Request URI for the API endpoint. Consists of version number and resource name.
	endpointName = "/v1/sfp"
)

type Router struct {
	frontend_tls_config *tls.Config
	frontend_server     *http.Server

    sysLogger *logger.Logger
}

func NewRouter(sysLogger *logger.Logger) *Router {
	// Create new Router
	router := new(Router)

    router.sysLogger = sysLogger

	// Create TLS config for frontend server
	router.frontend_tls_config = &tls.Config{
		Rand:                   nil,
		Time:                   nil,
		MinVersion:             tls.VersionTLS13,
		MaxVersion:             tls.VersionTLS13,
		SessionTicketsDisabled: true,
		Certificates:           []tls.Certificate{config.Config.Sfpl.X509KeyPairShownBySfplToPep},
		ClientAuth:             tls.RequireAndVerifyClientCert,
		ClientCAs:              config.Config.Sfpl.CaCertPoolSfplAcceptsFromPep,
	}

	// Create MUX server
	mux := http.NewServeMux()
	mux.Handle(endpointName, router)

	// Create HTTP frontend server
	router.frontend_server = &http.Server{
		Addr:      config.Config.Sfpl.ListenAddr,
		TLSConfig: router.frontend_tls_config,
		Handler:   mux,
	}

	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	md := new(md.Cp_metadata)

	if err := md.ExtractMetadata(req); err != nil {
        router.sysLogger.Errorf("router: ServeHTTP(): %v", err)
    }

	sfpl.TransformSFCintoSFP(router.sysLogger, md)

    router.sysLogger.Debugf("metadata after transformation: %v", md)

	// Encode metadata as json and set header respectively
    // TODO: Do not encode the whole md but only the SFP part
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(md)
}

func (router *Router) ListenAndServeTLS() error {
	return router.frontend_server.ListenAndServeTLS("", "")
}
