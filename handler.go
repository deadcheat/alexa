package alexa

// LambdaHandler request handler works with github.com/aws/aws-lambda-go
type LambdaHandler interface {
	Handle(RequestEnvelope) (ResponseEnvelope, error)
	HandleIntent(intentNames []string, h Handler)
	HandleLaunch(Handler)
	HandleEnd(Handler)
}

// Handler intent handler func
type Handler func(RequestEnvelope) (ResponseEnvelope, error)

// DefaultLambdaHandler default implement for LambdaHandler
type DefaultLambdaHandler struct {
	intentHandlers map[string]Handler
	launchHandler  Handler
	endHandler     Handler
}

// NewLambdaHandler return new lambda handler implement
func NewLambdaHandler() LambdaHandler {
	return &DefaultLambdaHandler{
		intentHandlers: make(map[string]Handler),
		launchHandler:  nil,
		endHandler:     nil,
	}
}

// Handle handle all lambda requests
func (d *DefaultLambdaHandler) Handle(req RequestEnvelope) (ResponseEnvelope, error) {
	if d.launchHandler == nil || d.endHandler == nil {
		return EmptyResponse, ErrNecessaryHandlersAreNotAssigned
	}
	switch req.Request.Type {
	case TypeLaunchRequest:
		return d.launchHandler(req)
	case TypeSessionEndedRequest:
		return d.endHandler(req)
	case TypeIntentRequest:
		h, ok := d.intentHandlers[req.Intent.Name]
		if ok {
			return h(req)
		}
		return EmptyResponse, ErrNoIntentToHandle
	}
	return EmptyResponse, ErrNoHandler
}

// HandleIntent put intent handler to internal map
func (d *DefaultLambdaHandler) HandleIntent(intentNames []string, h Handler) {
	for i := range intentNames {
		name := intentNames[i]
		d.intentHandlers[name] = h
	}
}

// HandleLaunch put request launch handler to internal map
func (d *DefaultLambdaHandler) HandleLaunch(h Handler) {
	d.launchHandler = h
}

// HandleEnd put request sessionended handler to internal map
func (d *DefaultLambdaHandler) HandleEnd(h Handler) {
	d.endHandler = h
}
