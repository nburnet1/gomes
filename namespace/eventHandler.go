package namespace

import "time"

type EventHandler func(engine *NamespaceEngine, path string, oldValue, newValue interface{}, oldTimestamp time.Time)
