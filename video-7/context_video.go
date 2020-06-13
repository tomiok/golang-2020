package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("server started!")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greeting", middleware(greetingHandler))
	http.ListenAndServe(":9191", nil)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	traceID := r.Context().Value("traceID")
	w.Write([]byte(traceID.(string)))
}

func middleware(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "traceID", "a1s2d3f4g5")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer fmt.Println("after signal interrupt")

	select {
	case <-time.After(10 * time.Second):
	case <-ctx.Done():
		err = ctx.Err()
	}

	if err != nil {
		fmt.Printf("error is %s \n", err.Error())
	}
}

func testGetContext() {
	// goroutine main is the holder
	ctx := getContext()
	value := ctx.Value(1)
	fmt.Println(value)
}

func getContext() context.Context {
	ctx := context.Background()
	//	ctxDone := context.TODO()
	ctxWithValue := context.WithValue(ctx, 1, 999)
	return ctxWithValue
}
