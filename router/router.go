package router

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	metadata "local.com/leobrada/ztsfc_http_sfpLogic/metadata"
	sfpl "local.com/leobrada/ztsfc_http_sfpLogic/sfp_logic"
)

type Router struct {
	frontend_tls_config *tls.Config
	frontend_server     *http.Server

	router_cert_shown_to_clients tls.Certificate
	certs_that_router_accepts    *x509.CertPool

	//    md         *metadata.Cp_metadata
}

func NewRouter() *Router {

	// Create new Router
	router := new(Router)

	// Initialize Certificates
	router.router_cert_shown_to_clients, _ = tls.LoadX509KeyPair("./certs/srv2_basicAuth.crt", "./certs/srv2_basicAuth_priv.key")
	router.certs_that_router_accepts = x509.NewCertPool()
	ca_cert, err := ioutil.ReadFile("./certs/ca_pool.pem")
	if err != nil {
		fmt.Print("[Router.makeCAPool]: ReadFile: ", err)
		return nil
	}
	ok := router.certs_that_router_accepts.AppendCertsFromPEM(ca_cert)
	if !ok {
		fmt.Print("[Router.makeCAPool]: AppendCertsFromPEM: ", err)
		return nil
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
	mux.Handle("/", router)

	// Create HTTP frontend server
	router.frontend_server = &http.Server{
		Addr:      "10.4.0.52:8889",
		TLSConfig: router.frontend_tls_config,
		Handler:   mux,
	}

	// router.md = new(metadata.Cp_metadata)

	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//fmt.Printf("%+v\n", req.Header)

	md := new(metadata.Cp_metadata)

	md.ExtractMetadata(req)

	sfpl.TransformSFCintoSFP(md)

	//    fmt.Printf("%s\n", md.SFP)
	w.Header().Set("sfp", md.SFP)
	fmt.Fprintf(w, "")
}

func (router *Router) ListenAndServeTLS() error {
	return router.frontend_server.ListenAndServeTLS("", "")
}
