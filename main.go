package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() *big.Int {
	var a, b *big.Int = big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		a, b = b, a.Add(a, b)
		return a
	}
}

func main() {

	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		f := fib()

		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		w.Header().Set("X-Content-Type-Options", "nosniff")
		keys, ok := r.URL.Query()["position"]

		if !ok || len(keys[0]) < 1 {
			fmt.Fprintf(w, "Url Param 'position' is missing")
			return
		}

		// Query()["key"] will return an array of items,
		// we only want the single item.
		key := keys[0]

		position, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Error! %v", err)
			return
		}

		for i := 0; i <= int(position); i++ {
			fmt.Fprintf(w, "%d fib number is %v\n", i, f())
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
		}
	})

	log.Print("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
