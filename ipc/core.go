package ipc

import (
	"errors"
	"fmt"
	"net"
	"os"
	"quantrium-tech/form_pf_extractor_service/config"
)

// ipcDaemon represents the core daemon that listens and forwards the inbound messages
type ipcDaemon struct {
	listener net.Listener
}

// inboundMessage represents the structure that holds the information passed from the backend
type inboundMessage struct {
	FilingID string `json:"filing_id"`
	UserID   string `json:"user_id"`
	Series   string `json:"series"`
}

var InboundChannel = make(chan inboundMessage)

var (
	daemon *ipcDaemon
)

// InitialiseDaemon opens a new uds socket and attach the listener to the daemon
// if socket is already exists, it throws the error
func InitialiseDaemon() error {

	l, err := net.Listen("unix", config.Cfg.SockPath)
	if err != nil {
		return fmt.Errorf("creating listener: " + err.Error())
	}

	daemon = &ipcDaemon{
		listener: l,
	}

	return nil

}

// TearDownSocket removes the socket from the system
func TearDownSocket() error {
	if err := os.Remove(config.Cfg.SockPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("removing the file: " + err.Error())
	}

	return nil

}
