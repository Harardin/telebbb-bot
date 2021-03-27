package telebbb

import (
	"fmt"
	"net/http"
	"time"
)

// NewBot Starts new bot service
// To run new bot service please provide
// telebbb.BotConfig type with data
func NewBot(c BotConfig) (*TbBot, error) {
	cli := &http.Client{
		Timeout: time.Second * 10,
	}

	// Create Connection and start webhook or
	b := &TbBot{
		client:   cli,
		token:    c.Token,
		Incoming: make(chan interface{}),
		Errors:   make(chan error, 1),
	}

	// Start Bot update listner
	switch c.Type {
	case "webhook":
		go b.ServeHook(c.Port) // webhook listner
	case "local":
		go b.LocalListen() // local listner
	default:
		return nil, fmt.Errorf("unknown bot type selected, please chose from 'webhook', or 'local'")
	}
	return b, nil
}
