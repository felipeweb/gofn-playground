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
		iaas.WithName("instance-1"),
		iaas.IsReused(true),
		iaas.WithRegion("us-east1-b"),
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
