package main

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/Electra-project/electra-auth/src/helpers"
	"github.com/Electra-project/electra-auth/src/server"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	startDaemon()
	go waitForDaemonToBeReady()

	helpers.LogInfo("Starting server...")
	server.Start()
}

func startDaemon() {
	var binaryFilePath string
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		binaryFilePath = "./bin/electrad-win32-x64.exe"
	} else if runtime.GOOS == "darwin" {
		binaryFilePath = "./bin/electrad-darwin-x64"
	} else {
		binaryFilePath = "./bin/electrad-linux-x64"
	}

	helpers.LogInfo("Checking Electra daemon binary...")
	if _, err := os.Stat(binaryFilePath); err != nil {
		helpers.LogInfo("Downloading Electra daemon binary...")
		helpers.DownloadBinary()
	}

	if runtime.GOOS != "windows" {
		helpers.LogInfo("Changing binary rights...")
		cmd = exec.Command("chmod", "755", binaryFilePath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			helpers.LogErr("Error: " + err.Error())
		}
	}

	helpers.LogInfo("Starting Electra daemon...")
	cmd = exec.Command(binaryFilePath, "--rpcuser=user", "--rpcpassword=pass")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go runDaemon(cmd)
}

func runDaemon(c *exec.Cmd) {
	helpers.LogInfo("Running Electra daemon (in another thread)...")
	if err := c.Run(); err != nil {
		helpers.LogErr("Error: " + err.Error())
	}
}

func waitForDaemonToBeReady() {
	helpers.LogInfo("Waiting for Electra daemon to be ready...\n")

	sum := 1
	for sum < 1000 {
		daemonURI := "http://localhost:5788"

		data := bytes.NewBuffer([]byte(`{"jsonrpc":"2.0","method":"getinfo"}`))
		req, err := http.NewRequest("POST", daemonURI, data)
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth("user", "pass")

		client := &http.Client{}
		res, err := client.Do(req)
		if err == nil {
			defer res.Body.Close()

			if res.StatusCode == 200 {
				break
			}
		} else {
			helpers.LogInfo("The following error can be safely ignored.")
			helpers.LogWarn("Error: " + err.Error())
		}

		time.Sleep(2000 * time.Millisecond)
	}
}
