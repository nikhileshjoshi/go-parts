package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nikhileshjoshi/go-ml/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "Handler Started")
	defer log.Println(ctx, "Handler Ended")

	select {
	case <-time.After(5 * time.Second): //time taken to complete the work by the handler
		fmt.Fprintln(w, "Hello from the context!!")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
