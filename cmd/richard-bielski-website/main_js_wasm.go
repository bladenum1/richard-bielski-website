package main

import (
    "richard-bielski-website/internal/javascript"
    "richard-bielski-website/internal/logger"
    "syscall/js"
)
func main() {
    namespace := "window.namespace"
    space := javascript.Get(namespace)
    err := logger.AddListener(namespace + "." + "listener", "debug")
    if err != nil {
        return
    }
    logger.Logger("namespace " + namespace, "debug")

    err = javascript.Set(space, "test", "test")
    if err != nil {
        logger.Logger(err.Error(), "debug")
        return
    }
    str := javascript.JSON("stringify", space)
    logger.Logger(str.String(), "debug")

    err, _ = javascript.NewFunction(namespace, "added", func(this js.Value, args []js.Value) interface{} {
        return "added this"
    })
    if err != nil {
        logger.Logger(err.Error(), "debug")
        return
    }

    err, _ = javascript.NewFunction(namespace, "log", func(this js.Value, args []js.Value) interface{} {
        logger.Logger(javascript.JSON("stringify",args[0]).String(), "debug")
        return nil
    })
    if err != nil {
        logger.Logger(err.Error(), "debug")
        return
    }

    err, _ = javascript.AddTrustedListener("window", namespace, "blur", func(this js.Value, args []js.Value) interface{} {
        logger.Logger("blur", "debug")
        return nil
    })
    if err != nil {
        logger.Logger(err.Error(), "debug")
        return
    }
}