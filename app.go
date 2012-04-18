package main

import (
    "github.com/hoisie/web"
    "log"
    "os"
)

func hello(val string) string { return "hello " + val }
func adam(val string) string { return "adam " + val } 

func main() {
    f, err := os.Create("server.log")
    if err != nil {
      // println(err.String())
      //      return
    }
    logger := log.New(f, "", log.Ldate|log.Ltime)
    web.Get("/(.*)", hello)
    web.Get("/adam(.*)", adam)
    web.SetLogger(logger)
    web.Run("0.0.0.0:56148")
}