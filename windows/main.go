package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofn/gofn"
	"github.com/gofn/gofn/iaas"
	"github.com/gofn/gofn/iaas/google"
	"github.com/gofn/gofn/provision"
)

func main() {
	buildOpts := &provision.BuildOptions{
		ImageName: "felipeweb/windowstest",
		RemoteURI: "",
		StdIN:     `Hello Windows`,
	}
	containerOpts := &provision.ContainerOptions{}
	stdout, _, err := gofn.Run(context.Background(), buildOpts, containerOpts)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Stdout: ", stdout)
}
