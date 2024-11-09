package namespace

import "github.com/gin-gonic/gin"

type Method string

const (
	GET Method = "GET"
	POST Method = "POST"
	PUT Method = "PUT"
	DELETE Method = "DELETE"
	PATCH Method = "PATCH"
)

type Endpoint struct {
	Method Method
	Url string
	Handlers []gin.HandlerFunc
}
var endpointRegistry = []Endpoint{}

func RegisterEndpoints(endpoints ...Endpoint) {
	endpointRegistry = append(endpointRegistry, endpoints...)
}

func GetEndpointRegistry() []Endpoint {
	return endpointRegistry
}

func EnableEndpointsFromEngine(r *gin.Engine) {
	for _, endpoint := range endpointRegistry {
		r.Handle(string(endpoint.Method), endpoint.Url, endpoint.Handlers...)
	}
}