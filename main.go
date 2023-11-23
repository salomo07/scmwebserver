package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		// Serve the Flutter web app from the "./build/web" directory
		path := string(ctx.Path())
		if path == "/" {
			path = "/index.html"
		}
		filePath := "./build/web" + path

		// Serve the file
		fasthttp.ServeFile(ctx, filePath)
	}

	// Start the fasthttp server
	addr := ":8080"
	fmt.Printf("Server is running on %s...\n", addr)
	if err := fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
