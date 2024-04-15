package main

import (

	"context"

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
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"CONTAINER NAME", "LOGS"})
	for _, container := range containers {
		row := []string{container.Names[0], ""}

		options := container.LogsOptions{ShowStdout: true, Timestamps: true, Details: true}
		out, err := cli.ContainerLogs(ctx, container.ID, options)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		// Read container logs
		logs, err := io.ReadAll(out)
		if err != nil {
			panic(err)
		}

		row[1] = string(logs)

		table.Append(row)
	}

	table.Render()
}


