// package main

// import (
// 	"context"
// 	"flag"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
// 	"github.com/hashicorp/terraform-provider-kubernetes/kubernetes"
// )

// func main() {
// 	debugFlag := flag.Bool("debug", false, "Start provider in stand-alone debug mode.")
// 	flag.Parse()

// 	serveOpts := &plugin.ServeOpts{
// 		ProviderFunc: kubernetes.Provider,
// 	}
// 	if debugFlag != nil && *debugFlag {
// 		plugin.Debug(context.Background(), "registry.terraform.io/hashicorp/kubernetes", serveOpts)
// 	} else {
// 		plugin.Serve(serveOpts)
// 	}
// }

package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	tf5server "github.com/hashicorp/terraform-plugin-go/tfprotov5/server"
	tfmux "github.com/hashicorp/terraform-plugin-mux"

	"github.com/hashicorp/terraform-provider-kubernetes/kubernetes"
	kubernetesalphaprovider "github.com/hashicorp/terraform-provider-kubernetes/manifest/provider"
)

func main() {

	// TODO add debug flag

	ctx := context.Background()

	kprov := kubernetes.Provider().GRPCProvider
	kprovalpha := kubernetesalphaprovider.Provider()

	factory, err := tfmux.NewSchemaServerFactory(ctx, kprov, kprovalpha)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	tf5server.Serve("registry.terraform.io/hashicorp/kubernetes", func() tfprotov5.ProviderServer {
		return factory.Server()
	})
}
