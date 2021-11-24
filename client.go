package ogmios

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	blocks chan json.RawMessage
	ch     chan struct{}
	conn   *websocket.Conn
	tip    chan struct{} // tip will be published to whenever
	group  *errgroup.Group
}

func New(ctx context.Context, logger *zap.Logger, endpoint string, pipeline int) (*Client, error) {
	logger = logger.With(zap.String("service", "ogmios"))

	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ogmios, %v: %w", endpoint, err)
	}

	group, ctx := errgroup.WithContext(ctx)
	client := &Client{
		blocks: make(chan json.RawMessage, 8),
		ch:     make(chan struct{}, 64),
		tip:    make(chan struct{}, 1),
		conn:   conn,
		group:  group,
	}

	group.Go(func() error {
		init := []byte(`{"type":"jsonwsp/request","version":"1.0","servicename":"ogmios","methodname":"FindIntersect","args":{"points":["origin"]},"mirror":{"step":"INIT"}}`)
		if err := conn.WriteMessage(websocket.TextMessage, init); err != nil {
			return fmt.Errorf("failed to write FindIntersect: %w", err)
		}

		next := []byte(`{"type":"jsonwsp/request","version":"1.0","servicename":"ogmios","methodname":"RequestNext","args":{}}`)
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-client.ch:
				if err := conn.WriteMessage(websocket.TextMessage, next); err != nil {
					return fmt.Errorf("failed to write RequestNext: %w", err)
				}
			}
		}
	})

	group.Go(func() error {
		for {
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return fmt.Errorf("failed to read message from ogmios: %w", err)
			}

			select {
			case client.ch <- struct{}{}:
			default:
			}

			if messageType == websocket.PingMessage {
				if err := conn.WriteMessage(websocket.PongMessage, nil); err != nil {
					return fmt.Errorf("failed to respond with pong to ogmios: %w", err)
				}
			}

			select {
			case <-ctx.Done():
				return nil
			case client.blocks <- data:
				// ok
			}
		}
	})

	for i := 0; i < pipeline; i++ {
		select {
		case client.ch <- struct{}{}:
		default:
		}
	}

	return client, nil
}

func (c *Client) Blocks() <-chan json.RawMessage {
	return c.blocks
}

func (c *Client) Close() error {
	defer c.group.Wait()
	return c.conn.Close()
}

func (c *Client) ReadNext(ctx context.Context) (json.RawMessage, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case block := <-c.blocks:
		return block, nil
	}
}

func (c *Client) Tip() <-chan struct{} {
	return c.tip
}
