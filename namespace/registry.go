package namespace

import (
	"fmt"
	"reflect"
	"runtime"
)

// Global Registry for Nodes
var nodeRegistry = []*NodeInstance{}

func RegisterNodes(nodeInstances ...*NodeInstance) {
	nodeRegistry = append(nodeRegistry, nodeInstances...)
}

type NodeInstance struct {
	Topic        string
	Value        interface{}
	EventHandler EventHandler
}

// Global Registry for Event Handlers
var eventHandlerRegistry = make(map[string]EventHandler)

func RegisterEventHandlers(eventHandlers ...EventHandler) {
	for _, eventHandler := range eventHandlers {
		funcName := getFunctionName(eventHandler)
		eventHandlerRegistry[funcName] = eventHandler
		fmt.Println(funcName)
	}

}

func getFunctionName(handler EventHandler) string {
	return runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
}

// Global Registry for Models
var modelRegistry = make(map[string]interface{})

func RegisterModels(models ...interface{}) {
	for _, model := range models {
		modelRegistry[reflect.TypeOf(model).String()] = model
	}
}

func GetModelRegistry() map[string]interface{} {
	return modelRegistry
}
