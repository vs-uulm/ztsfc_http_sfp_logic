package router

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	md "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/metadata"
	sfpl "github.com/vs-uulm/ztsfc_http_sfp_logic/internal/app/sfp_logic"
)

const (
	// Request URI for the API endpoint. Consists of version number and resource name.
	endpointName = "/v1/sfp"
)

type Router struct {
	frontend_tls_config *tls.Config
	frontend_server     *http.Server

	router_cert_shown_to_clients tls.Certificate
	certs_that_router_accepts    *x509.CertPool

	//    md         *metadata.Cp_metadata
}

func NewRouter() (*Router, error) {
	// Create new Router
	router := new(Router)

	// Initialize Certificates
	router.router_cert_shown_to_clients, _ = tls.LoadX509KeyPair("./certs/ztsfc_pdp_prototype.crt", "./certs/ztsfc_pdp_prototype_priv.key")
	router.certs_that_router_accepts = x509.NewCertPool()
	ca_cert, err := ioutil.ReadFile("./certs/bwnet_root.pem")
	if err != nil {
		return nil, fmt.Errorf("router: NewRouter(): could not read root CA certificate: %v", err)
	}
	ok := router.certs_that_router_accepts.AppendCertsFromPEM(ca_cert)
	if !ok {
		return nil, fmt.Errorf("router: NewRouter(): could not append root CA certificate to cert pool")
	}

	// Create TLS config for frontend server
	router.frontend_tls_config = &tls.Config{
		Rand:                   nil,
		Time:                   nil,
		MinVersion:             tls.VersionTLS13,
		MaxVersion:             tls.VersionTLS13,
		SessionTicketsDisabled: true,
		Certificates:           []tls.Certificate{router.router_cert_shown_to_clients},
		ClientAuth:             tls.RequireAndVerifyClientCert,
		ClientCAs:              router.certs_that_router_accepts,
	}

	// Create MUX server
	mux := http.NewServeMux()
	mux.Handle(endpointName, router)

	// Create HTTP frontend server
	router.frontend_server = &http.Server{
		Addr:      "10.4.0.52:8889",
		TLSConfig: router.frontend_tls_config,
		Handler:   mux,
	}

	return router, nil
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	md := new(md.Cp_metadata)

	md.ExtractMetadata(req)

	sfpl.TransformSFCintoSFP(md)

	// Encode metadata as json and set header respectively
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(md)
}

func (router *Router) ListenAndServeTLS() error {
	return router.frontend_server.ListenAndServeTLS("", "")
}
