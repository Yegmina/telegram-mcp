package tg

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/gotd/td/tg"
	mcp "github.com/metoro-io/mcp-golang"
	"github.com/pkg/errors"
)

type SendTextArguments struct {
	Name string `json:"name" jsonschema:"required,description=Name of the dialog"`
	Text string `json:"text" jsonschema:"required,description=Plain text of the message"`
}

type SendTextResponse struct {
	Success bool `json:"success"`
}

func randomMessageID() (int64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}

// SendText delivers a plain text message to the peer (messages.sendMessage).
func (c *Client) SendText(args SendTextArguments) (*mcp.ToolResponse, error) {
	var ok bool
	client := c.T()
	if err := client.Run(context.Background(), func(ctx context.Context) (err error) {
		api := client.API()

		inputPeer, err := getInputPeerFromName(ctx, api, args.Name)
		if err != nil {
			return fmt.Errorf("get inputPeer from name: %w", err)
		}

		rid, err := randomMessageID()
		if err != nil {
			return fmt.Errorf("random id: %w", err)
		}

		_, err = api.MessagesSendMessage(ctx, &tg.MessagesSendMessageRequest{
			Peer:       inputPeer,
			Message:    args.Text,
			RandomID:   rid,
			ClearDraft: true,
		})
		if err != nil {
			return fmt.Errorf("send message: %w", err)
		}

		ok = true
		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "send text")
	}

	jsonData, err := json.Marshal(SendTextResponse{Success: ok})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal response")
	}

	return mcp.NewToolResponse(mcp.NewTextContent(string(jsonData))), nil
}
