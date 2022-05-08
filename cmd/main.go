package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/k8s4u/vcluster-runasuser-plugin/syncers"
	"github.com/loft-sh/vcluster-sdk/plugin"

	"github.com/ElisaOyj/runasuser-admission-controller/pkg/controller"
)

const (
	tlsCertFile = `/tls/tls.crt`
	tlsKeyFile  = `/tls/tls.key`
)

func main() {
	fmt.Println("Starting webhook server")
	mux := http.NewServeMux()
	mux.Handle("/mutate", controller.AdmitFuncHandler(controller.ApplySecurityDefaults))
	server := &http.Server{
		Addr:    ":8765",
		Handler: mux,
	}

	go func() {
		log.Fatal(server.ListenAndServeTLS(tlsCertFile, tlsKeyFile))
	}()

	fmt.Println("Starting vcluster-runasuser-plugin")
	ctx := plugin.MustInit("runasuser-plugin")
	plugin.MustRegister(syncers.NewRegisterSyncer(ctx))
	plugin.MustStart()
}
