package ezrpc

import (
	"bytes"
	"errors"
	"strings"
	"time"

	"github.com/Wuvist/go-thrift/thrift"
	"github.com/nats-io/nats"
)

type Client struct {
	Conn      *nats.Conn
	Serice    string
	DirectKey string
}

func NewClient(service string, conn *nats.Conn) *Client {
	return &Client{
		Serice: service,
		Conn:   conn,
	}
}

func (c *Client) Call(method string, request interface{}, response interface{}) error {
	buf := &bytes.Buffer{}
	w := thrift.NewCompactProtocolWriter(buf)
	thrift.EncodeStruct(w, request)

	var subject string
	if strings.HasPrefix(method, "Direct") {
		if c.DirectKey == "" {
			return errors.New("client DirectKey is empty")
		}
		subject = c.DirectKey + "." + c.Serice + "." + method
	} else {
		subject = c.Serice + "." + method
	}

	if response == nil {
		return c.Conn.Publish(subject, buf.Bytes())
	}

	msg, err := c.Conn.Request(subject, buf.Bytes(), 10*time.Second)
	if err != nil {
		return err
	}
	r := thrift.NewCompactProtocolReader(bytes.NewReader(msg.Data))

	return thrift.DecodeStruct(r, response)
}
