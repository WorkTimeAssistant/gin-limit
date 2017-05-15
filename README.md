# Rate limiter [![Build Status](https://travis-ci.org/WorkTimeAssistant/gin-limit.svg?branch=master)](https://travis-ci.org/WorkTimeAssistant/gin-limit) [![GoDoc](https://godoc.org/github.com/WorkTimeAssistant/gin-limit?status.svg)](https://godoc.org/github.com/WorkTimeAssistant/gin-limit)
Simple rate limiting middleware for Gin.
You can set rates per endpoint or on the global router.

# Install
go get -u github.com/WorkTimeAssistant/gin-limit

# Usage
## Limit global router
```go
package main

import (
	gl github.com/WorkTimeAssistant/gin-limit
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gl.Limit(1000 /* max requests per second */, 10 /* max burst size*/))
	router.Run()
}
```

## Limit endpoint
```go
package main

import (
	gl github.com/WorkTimeAssistant/gin-limit
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1.GET("/test", gl.Limit(1000 /* max requests per second */, 10 /* max burst size*/), /* handler */)
	router.Run()
}
```