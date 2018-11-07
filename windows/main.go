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
	project := os.Getenv("GOOGLE_PROJECT")
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if project == "" {
		log.Fatalln("You must provide a project id from google cloud platform")
	}
	if credentials == "" {
		log.Fatalln("You must provide a path pointing to the credentials file from google cloud platform")
	}
	p, err := google.New(
		project,
		iaas.WithSO("windows-cloud/global/images/windows-server-1803-dc-core-for-containers-v20181009"),
		iaas.WithDiskSize(32),
	)
	if err != nil {
		log.Println(err)
	}
	buildOpts.Iaas = p

	stdout, _, err := gofn.Run(context.Background(), buildOpts, containerOpts)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Stdout: ", stdout)
}
