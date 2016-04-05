package handler

import (
	"bytes"

	"golang.org/x/net/context"

	proto "github.com/olivere/greeter/proto/greeter"
)

// Greeter handles RPC calls on the server-side.
type Greeter struct {
	ID string
}

// Say implements the Say rpc call.
func (g Greeter) Say(ctx context.Context, req *proto.SayRequest, res *proto.SayResponse) error {
	var buf bytes.Buffer
	buf.WriteString("Hello ")
	if req.Name != "" {
		buf.WriteString(req.Name)
	} else {
		buf.WriteString("stranger")
	}
	if g.ID != "" {
		buf.WriteString(" (reply from ")
		buf.WriteString(g.ID)
		buf.WriteString(")")
	}
	// Don't initialize a new response: Simply set its properties.
	res.Reply = buf.String()
	return nil
}
