package namespace

import (
	"fmt"
	"time"
)

type node struct {
	topic        string 
	name         string
	value        []byte
	timeStamp    time.Time
	parent       *node
	children     map[string]*node
	eventHandler EventHandler
	channel chan []byte
}

func (n node) GetName() string {
	return n.name
}

func (n node) GetTopic() string {
	return n.topic
}

func (n node) GetValue() []byte {
	return n.value
}

func (n node) GetTimeStamp() time.Time {
	return n.timeStamp
}

func (n node) GetParent() *node {
	return n.parent
}

func (n node) GetParentTopic() string {
	if n.parent == nil {
		return ""
	}
	return n.parent.topic
}

func (n node) GetChildren() map[string]*node {
	return n.children
}

func (n node) GetEventHandler() EventHandler {
	return n.eventHandler
}

func (n node) String() string {
	var parentName, childrenInfo string

	if n.parent == nil {
		parentName = "nil"
	} else {
		parentName = n.parent.name
	}

	if n.children == nil {
		childrenInfo = "nil"
	} else {
		childrenInfo = fmt.Sprintf("%v", n.children)
	}

	return fmt.Sprintf("Tag(Topic: %s, Name: %s, Value: %v, Timestamp: %v, Parent: %s, Children: %s)",
		n.topic, n.name, n.value, n.timeStamp, parentName, childrenInfo)
}
