package main

import (
	"fmt"
	"gomes/model"
	_ "gomes/admin/model"
	_ "gomes/admin/api"
	"gomes/namespace"

	"github.com/gin-gonic/gin"
)
func main() {
	/*
		Read configs
		Start tag engine
		Start web server
		Listen for requests
	*/
	var err error

	err = namespace.MigrateFromRegistry()
	if err != nil {
		fmt.Println(err)
	}

	err = namespace.Engine.StartFromRegistry()
	if err != nil {
		fmt.Println(err)
	}

	namespace.Engine.UpdateNode("Enterprise/Site/Area/SomeStruct", model.SomeStruct{Value: "325", AnotherField: "test3"})

	namespace.Engine.CreateNode("Enterprise/someTag", 63, nil)

	namespace.Engine.CreateNode("anotherTag", "test", nil)

	namespace.Engine.CreateNode("SomeNodeToTest", 45, nil)

	namespace.Engine.CreateNode("Enterprise/Site/Area/Machine1/AnotherNode", 4510, nil)

	r := gin.Default()
	r.LoadHTMLGlob("admin/templates/*")

	namespace.EnableEndpointsFromEngine(r)

	// fmt.Println(namespace.Engine)
	fmt.Println(namespace.GetEndpointRegistry())

	err = r.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println("Error setting trusted proxies:", err)
	}

	readOnlyCh, err := namespace.Engine.SubNode("Enterprise/Site/Area/Machine1/Asset")
	if err != nil {
		fmt.Println("Error listening to node:", err)
	}


	go func() {
		for data := range readOnlyCh {
			fmt.Println("Received update from node 'Enterprise/Site/Area/Machine1/Asset':", data)
		}
		fmt.Println("Read-only channel closed for node 'Enterprise/Site/Area/Machine1/Asset'")
	}()

	r.Run()

}
