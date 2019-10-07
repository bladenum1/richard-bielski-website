package logger

import (
	"errors"
	"richard-bielski-website/internal/javascript"
	"strings"
	"syscall/js"
)

type Listener struct {
	Scope string
	Function string
	Level string
}

// stores a slice of all listeners
var listeners []Listener

// AddListener ads a listener to the slice
func AddListener(listener string, level string) error {
	scope := strings.Split(listener,".")
	if len(scope) <= 1 {
		return errors.New("ERROR: " + "is not something we can call")
	}
	newListener := Listener {
		Scope: strings.Split(listener, "." + scope[len(scope)-1])[0],
		Function: scope[len(scope)-1],
		Level: level,
	}
	if javascript.Get(newListener.Scope + "." + newListener.Function) == js.Undefined() {
		return errors.New("ERROR: The listener function does not exist: " + newListener.Function )
	}
	listeners = append(listeners, newListener)
	return nil
}

// Logger fires events to all the listeners
func Logger(event string, level string) {
	for i := 0; i < len(listeners); i++ {
		javascript.Get(listeners[i].Scope).Call(listeners[i].Function, event)
	}
}
