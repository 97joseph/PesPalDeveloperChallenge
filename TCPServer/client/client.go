package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/DurgeshBabal/TCP-Messaging/keygen"
	"github.com/DurgeshBabal/TCP-Messaging/models"
	"github.com/fatih/color"
)

type client struct {
	pvtKey string
	pubKey string
}

func NewClient(pvtKey string, pubKey string) *client {
	return &client{
		pvtKey: pvtKey,
		pubKey: pubKey,
	}
}

const serverPort = ":5000"
const serverIP = "127.0.0.1"

func main() {
	// Generate new key pair
	pubKey, pvtKey, err := keygen.NewKey()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("\n", pubKey)

	// Connect to server
	client := NewClient(pvtKey, pubKey)

	c, err := net.Dial("tcp", serverIP+serverPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("TCP Client online!")

	// Register client with server
	m := models.Message{
		Operation: "RegisterClient",
		Source:    client.pubKey,
	}

	writeMessage(m, c)

	go func() {
		serverReader := bufio.NewReader(c)

		for {
			m, err := readServerResponse(serverReader)
			if err != nil {
				if err == io.EOF {
					c.Close()
					os.Exit(0)
					return
				}

				log.Println(err)
				continue
			}

			printResponse(m)

			client.handleOperation(m, c)
		}
	}()

	userReader := bufio.NewReader(os.Stdin)

	for {
		m, err := readUserInput(userReader)
		if err != nil {
			log.Println(err)
			continue
		}

		writeMessage(m, c)
	}

}

func readUserInput(r *bufio.Reader) (m models.Message, err error) {
	input, err := r.ReadString('~')
	if err != nil {
		return m, err
	}

	input = strings.TrimSuffix(input, "~")

	err = json.Unmarshal([]byte(input), &m)

	return m, err
}

func readServerResponse(r *bufio.Reader) (m models.Message, err error) {
	netData, err := r.ReadString('~')
	if err != nil {
		return m, err
	}

	netData = strings.TrimSuffix(netData, "~")

	err = json.Unmarshal([]byte(netData), &m)

	return m, err
}

func writeMessage(m models.Message, c net.Conn) {
	b := m.Bytes()
	b = append(b, '~')

	fmt.Fprintf(c, "%s", b)
}

func printResponse(m models.Message) {
	if (m == models.Message{}) {
		return
	}

	color.Blue("\nServer Response:")
	if m.Operation != "" {
		color.Cyan("Operation:")
		color.Red(m.Operation)
	}
	if m.Value != "" {
		color.Cyan("Value:")
		color.Red(m.Value)
	}
	if m.Source != "" {
		color.Cyan("Source Client PubKey:")
		color.Red(m.Source)
	}
	fmt.Print("\n>> ")
}

func (client *client) handleOperation(m models.Message, c net.Conn) {
	switch m.Operation {
	case "ForwardMessage":
		msg := models.Message{
			Operation: "ClientResponse",
			Source:    client.pubKey,
			Value:     "Sending a random payload",
			Target:    m.Source,
		}

		writeMessage(msg, c)
	}
}
