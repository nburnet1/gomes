package namespace

import (
	"encoding/json"
	"fmt"
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

func (e *NamespaceEngine) CreateNode(topic string, value any, eventHandler EventHandler) (node, error) {
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
		var err error
		parentNodePtr, err = e.checkForParent(parentTopic)
		if err != nil {
			return node{}, err
		}
	}

	var err error
	var jsonValue []byte

	if v, ok := value.([]byte); ok {
		if !json.Valid(v) {
			return node{}, fmt.Errorf("invalid JSON: %s", string(v))
		}
		jsonValue = v
	} else {
		jsonValue, err = json.MarshalIndent(value, "", "  ")
		if err != nil {
			return node{}, err
		}
	}

	e.nodes[topic] = &node{
		topic:        topic,
		name:         relativeTopic,
		parent:       parentNodePtr,
		value:        jsonValue,
		children:     make(map[string]*node),
		timeStamp:    time.Now().UTC(),
		eventHandler: eventHandler,
		channel:      make(chan []byte),
	}

	if e.nodes[topic].channel != nil {
		go func(ch chan []byte, val []byte) {
			ch <- val
		}(e.nodes[topic].channel, jsonValue)
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
		return node{}, fmt.Errorf("node %s doesn't exist", topic)
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

func (e *NamespaceEngine) SubNode(topic string) (chan []byte, error) {
	if !e.checkValidTopic(topic) {
		return nil, fmt.Errorf("invalid topic %s", topic)
	} else if !e.checkNodeExists(topic) {
		return nil, fmt.Errorf("node %s doesn't exist", topic)
	}
	readOnly := make(chan []byte)
	e.nodes[topic].channel = readOnly
	return readOnly, nil
}

func (e *NamespaceEngine) UpdateNode(topic string, value any) (node, error) {
	if !e.checkValidTopic(topic) {
		return node{}, fmt.Errorf("invalid topic %s", topic)
	}
	if !e.checkNodeExists(topic) {
		return node{}, fmt.Errorf("node %s doesn't exist", topic)
	}
	if string(e.nodes[topic].value) == "_folder" {
		return node{}, fmt.Errorf("can't update a node of type folder")
	}
	var err error
	var jsonValue []byte

	if v, ok := value.([]byte); ok {
		if !json.Valid(v) {
			return node{}, fmt.Errorf("invalid JSON: %s", string(v))
		}
		jsonValue = v
	} else {
		jsonValue, err = json.MarshalIndent(value, "", "  ")
		if err != nil {
			return node{}, err
		}
	}

	if string(e.nodes[topic].value) == string(jsonValue) {
		return *e.nodes[topic], nil
	}

	nodePtr := e.nodes[topic]
	oldValue := nodePtr.value
	nodePtr.value = jsonValue
	oldTimestamp := nodePtr.timeStamp
	nodePtr.timeStamp = time.Now().UTC()

	if nodePtr.channel != nil {
		// Send the new value in a goroutine to avoid blocking
		go func(ch chan []byte, val []byte) {
			ch <- val
		}(nodePtr.channel, jsonValue)
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

func (e *NamespaceEngine) checkForParent(parentTopic string) (*node, error) {
	if e.checkNodeExists(parentTopic) {
		parentNode := e.nodes[parentTopic]

		parentValue := string(parentNode.value)
		if parentValue != "\"_folder\"" {
			return nil, fmt.Errorf("invalid hierarchy found. parent {%s} must be a folder but is really of type {%s}", parentTopic, string(parentNode.value))
		}
		return parentNode, nil
	} else {
		_, err := e.CreateNode(parentTopic, "_folder", nil)
		if err != nil {
			return nil, err
		}
		return e.nodes[parentTopic], nil
	}
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
