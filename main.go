package main

import (
	"net/http"

	"github.com/AdjectiveAllison/tasks/app"
	"github.com/syumai/workers"
)

func main() {
	http.Handle("/", app.NewTaskHandler())
	workers.Serve(nil) // use http.DefaultServeMux
}
