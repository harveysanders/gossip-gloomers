package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type response struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func main() {
	n := server()

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func server() *maelstrom.Node {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		id, err := generateRandID()
		if err != nil {
			return fmt.Errorf("generateRandID: %w", err)
		}

		respMsg := response{
			Type: "generate_ok",
			ID:   id.String(),
		}

		return n.Reply(msg, respMsg)
	})

	return n
}

func generateRandID() (uuid.UUID, error) {
	return uuid.NewUUID()
}
