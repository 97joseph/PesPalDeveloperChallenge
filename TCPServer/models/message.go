package models

import (
	"encoding/json"
	"log"
)

type Message struct {
	Operation string `json:"operation"`
	Value     string `json:"value"`
	Source    string `json:"source"`
	Target    string `json:"target"`
}

func (m *Message) String() string {
	return string(m.Bytes())
}

func (m *Message) Bytes() []byte {
	s, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return s
}
