# Rate limiter [![Build Status](https://travis-ci.org/WorkTimeAssistant/gin-limit.svg?branch=master)](https://travis-ci.org/WorkTimeAssistant/gin-limit) [![GoDoc](https://godoc.org/github.com/WorkTimeAssistant/gin-limit?status.svg)](https://godoc.org/github.com/WorkTimeAssistant/gin-limit)
Simple rate limiting middleware for Gin.
You can set rates per endpoint or on the global router.

# Install
```bash
go get -u github.com/WorkTimeAssistant/gin-limit
```

# Usage
Limit global or endpoint's router.
```go
package main

import (
	gl github.com/WorkTimeAssistant/gin-limit
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gl.Limit(1000 /* max requests per second */, 20 /* max burst size*/))
	router.GET("/test", gl.Limit(100, 10), /* handler */)
	router.Run()
}
```
