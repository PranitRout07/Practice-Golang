package main

import (

	"context"
	"fmt"

	"io"
	"os"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	//Listing all the containers and their logs
	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Println("CONTAINER NAME: ", container.Names)
		fmt.Println("------------------------------")
		options := containertypes.LogsOptions{ShowStdout: true,Timestamps: true , Details: true}
		out, err := cli.ContainerLogs(ctx, container.ID, options)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		// Copy container logs directly to standard output
		if _, err := io.Copy(os.Stdout, out); err != nil {
			panic(err)
		}

	}

}
