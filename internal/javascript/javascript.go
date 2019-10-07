package javascript

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"syscall/js"
)

// GoJsObj holds the reference string and js object
type GoJsObj struct {
	ref string
	val js.Value
}

// Get returns JS object
func Get(attribute string) js.Value {
	obj := js.Global()

	structure := strings.Split(attribute,".")

	for i := 0; i < len(structure); i++  {
		if obj == js.Undefined() {
			return obj
		}
		obj = obj.Get(structure[i])
	}
	return obj
}

// Set sets an item on a JS object
func Set(obj js.Value, attribute string, value interface{}) error {
	if obj == js.Undefined() {
		return errors.New("ERROR: Can not set attribute to an undefined object")
	}
	obj.Set(attribute, value)
	return nil
}

// Calls JSON in javascript
func JSON(attribute string, value interface {}) js.Value {
	JSON := Get("window.JSON")
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		fmt.Println(value)
	}
	return JSON.Call(attribute, value)
}

// NewFunction ads a go function to an exported javascript scope
func NewFunction(scope string, attribute string, function func(this js.Value, args []js.Value) interface{}) (error, js.Value) {
	obj := Get(scope)
	if obj == js.Undefined() {
		return errors.New("ERROR: The scope provided " + scope + " does not exits"), js.Undefined()
	}
	obj.Set(attribute,js.FuncOf(function))
	return nil, obj.Get(attribute)
}

// AddListener attaches a listener to an event
func AddTrustedListener(objToListenTo string, namespace string, event string, function func(this js.Value, args []js.Value) interface{}) (error, js.Value) {
	rng := strconv.FormatInt(rand.Int63n(math.MaxInt64), 16)
	objVal := Get(objToListenTo)
	if objVal == js.Undefined() {
		return errors.New("ERROR: The scope provided " + objToListenTo + " does not exits"), js.Undefined()
	}
	err, obj := NewFunction(namespace, event + rng, function)
	if err !=nil {
		return err, js.Undefined()
	}
	objVal.Call("addEventListener", event, obj)
	return nil, objVal
}