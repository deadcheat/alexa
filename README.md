# corner-cut library for alexa-skills on lambda

## what's this ?

To cut the corner about developing alexa-skills, built on amazon-lambda using Go.

## How to use

### With using github.com/aws/aws-lambda-go

first of all, run `go get github.com/aws/aws-lambda-go`

and get this through `go get github.com/deadcheat/alexa`

then,

```
func main() {
	// create lambda handler
	h := alexa.NewLambdaHandler()

	// assign handler to request and intent
	h.HandleLaunch(func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
		return alexa.EmptyResponse, alexa.ErrNoHandler
	})
	h.HandleEnd(func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
		return alexa.EmptyResponse, alexa.ErrNoHandler
	})

	/// assign intent handlers
	h.HandleIntent([]string{""}, func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
		return alexa.EmptyResponse, alexa.ErrNoHandler
	})

	lambda.Start(h.Handle)
}
```

Ofcourse, you all should rewrite code above, each func return not-empty response struct.

### Without using lambda-go

- you can use this as you like,
- you can use struct definitions in this repo,
- you can use LambdaHandler with unmarshalling http.Request.Body

