package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/DurgeshBabal/TCP-Messaging/models"
)

type server struct {
	ip          string
	port        string
	connections map[string]net.Conn
	clients     map[int]string
	ctr         int
}

func NewServer(ip string, port string) *server {
	return &server{
		ip:          ip,
		port:        port,
		connections: make(map[string]net.Conn),
		clients:     make(map[int]string),
		ctr:         1,
	}
}

func (s *server) getClientList() (resp string) {
	for k, v := range s.clients {
		resp += "Client ID: " + strconv.Itoa(k) + "\n" + v + "\n"
	}

	return resp
}

func (s *server) registerClient(m models.Message, c net.Conn) (resp string) {
	pubKey := m.Source

	s.connections[pubKey] = c
	s.clients[s.ctr] = pubKey
	s.ctr++

	resp = "Client registered!"

	return resp
}

// Assuming the source and target fiels in message will use the clientID
// displayed when using the 'ClientList' operation to make forwarding
// messages easier. forwardMessage() is same as clientResponse() except
// that the former is used by users while the latter is used by the client
// to send automatic responses
func (s *server) forwardMessage(m models.Message, c net.Conn) (resp string) {
	var sourcePubKey string

	for k, v := range s.connections {
		if v == c {
			sourcePubKey = k
			break
		}
	}

	clientID, err := strconv.Atoi(m.Target)
	if err != nil {
		log.Println(err)
		return "Target Invalid"
	}

	targetPubKey := s.clients[clientID]
	targetConn := s.connections[targetPubKey]
	if targetConn == nil {
		return "Targed Invalid"
	}

	msg := models.Message{
		Operation: m.Operation,
		Value:     m.Value,
		Source:    sourcePubKey,
	}

	writeMessage(msg, targetConn)

	return "Message Forwarded!"
}

func (s *server) clientResponse(m models.Message, c net.Conn) (resp string) {
	targetConn := s.connections[m.Target]
	if targetConn == nil {
		return "Targed Invalid"
	}

	msg := models.Message{
		Operation: m.Operation,
		Value:     m.Value,
		Source:    m.Source,
	}

	writeMessage(msg, targetConn)

	return "Response Sent!"
}

func (s *server) handleOperation(c net.Conn, m models.Message) {
	var resp string
	op := m.Operation

	switch m.Operation {
	case "ClientList":
		resp = s.getClientList()

	case "RegisterClient":
		resp = s.registerClient(m, c)

	case "ForwardMessage":
		resp = s.forwardMessage(m, c)
		op = ""

	case "ClientResponse":
		resp = s.clientResponse(m, c)
		op = ""

	default:
		op = "Invalid Operation"
	}

	msg := models.Message{
		Operation: op,
		Value:     resp,
	}

	writeMessage(msg, c)
}

func writeMessage(m models.Message, c net.Conn) {
	b := m.Bytes()
	b = append(b, '~')

	fmt.Fprintf(c, "%s", b)
}

func readMessage(r *bufio.Reader) (m models.Message, err error) {
	netData, err := r.ReadString('~')
	if err != nil {
		return m, err
	}

	netData = strings.TrimSuffix(netData, "~")

	err = json.Unmarshal([]byte(netData), &m)

	return m, err
}

func (s *server) cleanUp(c net.Conn) {
	var pubKey string

	for k, v := range s.connections {
		if v == c {
			pubKey = k
			delete(s.connections, k)
			break
		}
	}

	for k, v := range s.clients {
		if v == pubKey {
			delete(s.clients, k)
			break
		}
	}

	c.Close()
}

func (s *server) handleConnection(c net.Conn) {
	defer s.cleanUp(c)

	r := bufio.NewReader(c)

	for {
		message, err := readMessage(r)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			continue
		}

		s.handleOperation(c, message)
	}
}

func main() {
	s := NewServer("127.0.0.1", "5000")

	l, err := net.Listen("tcp", s.ip+":"+s.port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("TCP Server online!")

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go s.handleConnection(c)
	}
}
