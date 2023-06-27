// Package wasmtalk contains types and helpers used for communicating over webassembly based stdio using json.
package wasmtalk

import "encoding/json"

type MessageType string

const (
	UnknownMessageType MessageType = ""
)

// Message is a message from or to a wasm wasi based script.
// Communicates over stdin/out via json messages using this format.
type Message struct {
	// Type is the type of message.
	Type MessageType `json:",omitempty"`

	Value json.RawMessage
}

// NewMessage creates a new message with the given type and value.
func NewMessage(msgType MessageType, value any) (*Message, error) {
	msg := &Message{
		Type: msgType,
	}

	var err error
	msg.Value, err = json.Marshal(value)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
