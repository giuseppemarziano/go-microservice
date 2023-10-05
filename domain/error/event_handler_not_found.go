package error

type HandlerNotFound struct {
	RoutingKey string
}

func (e HandlerNotFound) Error() string {
	return "handler not found for routing key: " + e.RoutingKey
}
