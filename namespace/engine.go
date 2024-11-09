package namespace

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var Engine = NewNamespaceEngine()

type NamespaceEngine struct {
	nodes map[string]*node
}

func NewNamespaceEngine() *NamespaceEngine {
	return &NamespaceEngine{
		nodes: make(map[string]*node),
	}
}

func (e *NamespaceEngine) Start() error {
	return nil
}

func (e *NamespaceEngine) StartFromRegistry() error {
	for _, nodeInstance := range nodeRegistry {
		_, err := e.CreateNode(nodeInstance.Topic, nodeInstance.Value, nodeInstance.EventHandler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *NamespaceEngine) CreateNode(topic string, value interface{}, eventHandler EventHandler) (node, error) {
	if !e.checkValidTopic(topic) {
		return node{}, fmt.Errorf("invalid topic %s", topic)
	} else if e.checkNodeExists(topic) {
		return node{}, fmt.Errorf("node %s already exists", topic)
	}
	parentTopic, relativeTopic := e.getSplitTopic(topic)

	var parentNodePtr *node

	if parentTopic == "" {
		parentNodePtr = nil
	} else {
		parentNode, err := e.checkForParent(parentTopic)
		if err != nil {
			return node{}, err
		}
		parentNodePtr = &parentNode
	}

	e.nodes[topic] = &node{
		topic:        topic,
		name:         relativeTopic,
		parent:       parentNodePtr,
		value:        value,
		children:     make(map[string]*node),
		timeStamp:    time.Now().UTC(),
		eventHandler: eventHandler,
		channel:      make(chan interface{}),
	}

	if e.nodes[topic].channel != nil {
		go func(ch chan interface{}, val interface{}) {
			ch <- val
		}(e.nodes[topic].channel, value)
	}

	if parentNodePtr != nil {
		nodeValue := e.nodes[topic]
		parentNodePtr.children[relativeTopic] = nodeValue
	}
	return *e.nodes[topic], nil
}

func (e *NamespaceEngine) ReadNode(topic string) (node, error) {
	if !e.checkValidTopic(topic) {
		return node{}, fmt.Errorf("invalid topic %s", topic)
	} else if !e.checkNodeExists(topic) {
		return node{}, fmt.Errorf("node %s doesn'e exist", topic)
	}
	return *e.nodes[topic], nil
}

func (e *NamespaceEngine) BrowseNodes() map[string]*node {
	return e.nodes
}

func (e *NamespaceEngine) BrowseRootNodes() map[string]*node {
	rootNodes := make(map[string]*node)
	for k, v := range e.nodes {
		if v.parent == nil {
			rootNodes[k] = v
		}
	}
	return rootNodes
}

func (e *NamespaceEngine) SubNode(topic string) (chan interface{}, error) {
	if !e.checkValidTopic(topic) {
		return nil, fmt.Errorf("invalid topic %s", topic)
	} else if !e.checkNodeExists(topic) {
		return nil, fmt.Errorf("node %s doesn't exist", topic)
	}
	readOnly := make(chan interface{})
	e.nodes[topic].channel = readOnly
	return readOnly, nil
}

func (e *NamespaceEngine) UpdateNode(topic string, value interface{}) (node, error) {
	if !e.checkValidTopic(topic) {
		return node{}, fmt.Errorf("invalid topic %s", topic)
	}
	if !e.checkNodeExists(topic) {
		return node{}, fmt.Errorf("node %s doesn't exist", topic)
	}
	if _, ok := e.nodes[topic].value.(Folder); ok {
		return node{}, fmt.Errorf("can't update a node of type folder")
	}
	if e.nodes[topic].value == value {
		return *e.nodes[topic], nil
	}

	nodePtr := e.nodes[topic]
	oldValue := nodePtr.value
	nodePtr.value = value
	oldTimestamp := nodePtr.timeStamp
	nodePtr.timeStamp = time.Now().UTC()

	if nodePtr.channel != nil {
		// Send the new value in a goroutine to avoid blocking
		go func(ch chan interface{}, val interface{}) {
			ch <- val
		}(nodePtr.channel, value)
	}

	if nodePtr.eventHandler == nil {
		return *e.nodes[topic], nil
	}

	go func() {
		nodePtr.eventHandler(e, nodePtr.topic, oldValue, value, oldTimestamp)
	}()

	return *e.nodes[topic], nil
}

func (e *NamespaceEngine) DeleteNode(topic string) error {
	if !e.checkValidTopic(topic) {
		return fmt.Errorf("invalid topic %s", topic)
	} else if !e.checkNodeExists(topic) {
		return fmt.Errorf("node %s already exists", topic)
	}
	if e.nodes[topic].parent != nil {
		nodeToDelete := e.nodes[topic]
		delete(e.nodes[topic].parent.children, nodeToDelete.name)
	}
	if len(e.nodes[topic].children) > 0 {
		for k := range e.nodes[topic].children {
			e.DeleteNode(e.nodes[topic].children[k].topic)
		}

	}
	delete(e.nodes, topic)

	return nil
}

func (e *NamespaceEngine) checkValidTopic(topic string) bool {
	regex := `^[a-zA-Z0-9_]+(\/[a-zA-Z0-9_]+)*\/?$`
	r := regexp.MustCompile(regex)

	return r.MatchString(topic)
}

func (e *NamespaceEngine) checkNodeExists(topic string) bool {
	_, ok := e.nodes[topic]
	return ok
}

func (e *NamespaceEngine) checkForParent(parentTopic string) (node, error) {
	var parent node

	if e.checkNodeExists(parentTopic) {
		parent = *e.nodes[parentTopic]

		if _, ok := parent.value.(Folder); !ok {
			return node{}, fmt.Errorf("invalid hierarchy found. parent {%s} must be a folder but is really of type {%s}", parentTopic, reflect.TypeOf(parent.value).String())
		}
	} else {
		parent, _ = e.CreateNode(parentTopic, Folder{}, nil)
	}
	return parent, nil
}

func (e *NamespaceEngine) getSplitTopic(topic string) (string, string) {
	delim := "/"
	index := strings.LastIndex(topic, delim)
	if index != -1 {
		return topic[:index], topic[index+1:]
	}
	return "", topic
}

func (e *NamespaceEngine) String() string {
	var str string
	for k, v := range e.nodes {
		str += fmt.Sprintf("%s \n\t%v\n\n", k, v)
	}
	return str
}
