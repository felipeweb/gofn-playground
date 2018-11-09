package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofn/gofn"
	"github.com/gofn/gofn/provision"
)

func main() {
	buildOpts := &provision.BuildOptions{
		ImageName: "felipeweb/windowstest",
		RemoteURI: "",
		StdIN:     `Hello Windows`,
	}
	containerOpts := &provision.ContainerOptions{}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	stdout, stderr, err := gofn.Run(context.Background(), buildOpts, containerOpts)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Stdout: ", stdout)
	fmt.Println("St: ", stderr)
	<-c
}
