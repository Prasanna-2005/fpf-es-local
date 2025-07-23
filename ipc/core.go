package ipc

import (
	"errors"
	"fmt"
	"net"
	"os"
	"quantrium-tech/form_pf_extractor_service/config"
)

type ipcDaemon struct {
	listener net.Listener
}

var (
	daemon *ipcDaemon
)

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

func TearDownSocket() error {
	if err := os.Remove(config.Cfg.SockPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("removing the file: " + err.Error())
	}

	return nil

}
