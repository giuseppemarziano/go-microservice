package messagebus

import "context"

type Message struct {
	RoutingKey string
	Payload    []byte
	Ctx        context.Context
}
