package main

import (
  "github.com/paulbellamy/mango"
  "os"
)

// Our default handler
func Hello(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 200, mango.Headers{}, mango.Body("<b>Hi Esther!</b>")
}

// Our handler for /goodbye
func Goodbye(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 200, mango.Headers{}, mango.Body("<b>Goodbye Esther!</b>")
}

func main() {
	stack := new(mango.Stack)
	stack.Address = ":"+os.Getenv("PORT")

	// Route all requests for /goodbye to the Goodbye handler
	routes := map[string]mango.App{"/goodbye(.*)": Goodbye}
	stack.Middleware(mango.Routing(routes))

	// Hello handles all requests not sent to Goodbye
	stack.Run(Hello)
}
