package model

import (
	"fmt"
	"gomes/namespace"
	"time"
)

func init() {
	someStructNodes := []*namespace.NodeInstance{
		{
			Topic:        "Enterprise/Site/Area/SomeStruct",
			Value:        SomeStruct{Value: "1", AnotherField: "test"},
			EventHandler: SomeStructFunc,
		},
		{
			Topic:        "Enterprise/Site/Area1/SomeStruct",
			Value:        SomeStruct{Value: "2", AnotherField: "test2"},
			EventHandler: SomeStructFunc,
		},
	}
	namespace.RegisterNodes(someStructNodes...)

	namespace.RegisterModels(SomeStruct{})
}

type SomeStruct struct {
	Value        string
	AnotherField string
}

func SomeStructFunc(tagEngine *namespace.NamespaceEngine, path string, oldValue, newValue interface{}, oldTimestamp time.Time) {
	fmt.Println(oldValue, newValue)

}
