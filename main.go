package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-go-provider/infer"
)

// version is the component provider version. It is overridden at publish time.
const version = "0.1.0"

func main() {
	prov, err := infer.NewProviderBuilder().
		WithNamespace("minecraft").
		WithComponents(
			infer.ComponentF(NewMinecraftServer),
		).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error building provider: %s\n", err.Error())
		os.Exit(1)
	}

	if err := prov.Run(context.Background(), "minecraft", version); err != nil {
		fmt.Fprintf(os.Stderr, "error running provider: %s\n", err.Error())
		os.Exit(1)
	}
}
