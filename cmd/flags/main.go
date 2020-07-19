package main

import (
	"os"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func ip(command string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("ip", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

}

func main() {
	// won't work for windows, exit
	if runtime.GOOS == "windows" {
		log.Fatal("this tool wasn't made for windows, bummer")
	}

	log.Info("started")
}
