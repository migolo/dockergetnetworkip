package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/docker/engine-api/client"
	"golang.org/x/net/context"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Fprintln(os.Stderr, "No network name specified")
		os.Exit(1)
	}

	networkName := argsWithoutProg[0]
	var (
		cmdOut     []byte
		errCommand error
	)
	cmdName := "cat"
	cmdArgs := []string{"/proc/self/cgroup"}
	if cmdOut, errCommand = exec.Command(cmdName, cmdArgs...).Output(); errCommand != nil {
		fmt.Fprintln(os.Stderr, "There was an error running command: ", errCommand)
		os.Exit(1)
	}
	re := regexp.MustCompile(`cpu:\/docker\/([0-9a-f]+)`)
	sha := string(cmdOut)
	found := re.FindStringSubmatch(sha)
	containerID := found[1]

	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}

	information, errInspect := cli.ContainerInspect(context.Background(), containerID)
	if errInspect != nil {
		panic(errInspect)
	}

	networkSelected, networkKeyExists := information.NetworkSettings.Networks[networkName]
	if !networkKeyExists {
		fmt.Fprintln(os.Stderr, "No network found")
		os.Exit(1)
	}
	fmt.Println(networkSelected.IPAddress)
}
