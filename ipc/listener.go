package ipc

import (
	"fmt"
	"net"
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/logging"
)

// Exported function for the Listen
func Listen() {
	daemon.listen()
}

func (d *ipcDaemon) listen() {
	logging.Info("Daemon Started Listening on " + config.Cfg.SockPath)
	for {

		conn, err := d.listener.Accept()
		if err != nil {
			logging.Error("Error in Accepting Incoming Connection: " + err.Error())
			continue
		}

		logging.Info("New Connection Established at the local address: " + conn.LocalAddr().String())
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {

	// close the connection from the server side if the signalling is success
	defer conn.Close()

	buff := make([]byte, 4096)

	n, err := conn.Read(buff)
	if err != nil {
		logging.Error("Error in Reading from the established Connection: " + err.Error())
		return
	}

	logging.Info(fmt.Sprintf("%d bytes of data read from the connection at: ", n) + conn.LocalAddr().String())

	// TODO: Should be removed once we have the schema and further signal processing
	conn.Write(buff[:n])
}
