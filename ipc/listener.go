package ipc

import (
	"encoding/json"
	"fmt"
	"net"
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/logging"
)

// Exported function for the Listen
func Listen() {
	daemon.listen()
}

// listen actively listens for incoming connections and will handle the established connection in seperate goroutine
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

// handleConnection takes the established connection and reads the transferred bytes and handle it appropriately
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
	go handleMessage(buff[:n])
}

// handleMessage takes the incoming buffer and tries to forward it to the inbound channel if found to be valid
func handleMessage(buff []byte) {

	var message inboundMessage
	err := json.Unmarshal(buff, message)
	if err != nil {
		logging.Error("Error in unmarshaling the inbound data: " + err.Error())
		return
	}

	if message.FilingID == "" {
		logging.Error("Recieved a empty filing ID or filing ID not present on the signal")
		return
	}

	InboundChannel <- message

	logging.Info("Inbound Message forwarded to the main channel")
}
