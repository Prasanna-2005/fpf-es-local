package main

import (
	"os"
	"os/signal"
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/db"
	"quantrium-tech/form_pf_extractor_service/ipc"
	"quantrium-tech/form_pf_extractor_service/logging"
	"syscall"
)

func init() {

	logging.InitialiseLogger()
	logging.Info("Logger Initialised")

	config.InitialiseConfig()
	logging.Info("Config Initialised")

	err := db.InitialseDB()
	if err != nil {
		panic(err)
	}
	logging.Info("Database Connection Initialised")

	err = ipc.InitialiseDaemon()
	if err != nil {
		panic(err)
	}
	logging.Info("IPC Daemon Initialised")

	// start the active listener
	go ipc.Listen()

}

func handleSignal(c chan os.Signal) {
	<-c
	logging.Warn("Got a Kill/Interupt Signal")
	if err := ipc.TearDownSocket(); err != nil {
		logging.Error("Error in Tearing Down: " + err.Error() + " Sock File May Prevail. Manual Cleanup is Mandatory")
	}
	logging.Info("Teardown of Sock file is completed")
	os.Exit(1)

}

func main() {

	// Create the signal handler and register with a channel
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go handleSignal(c)

	// Main Forwarder/Handler from the Listener
	for {
	}

}
