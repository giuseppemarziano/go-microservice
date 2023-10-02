package messagebus

type Message struct {
	RoutingKey string
	Payload    []byte
}
