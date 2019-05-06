---
title: go-kit
date: 2019-05-06T15:05:00+01:00
categories:
	- go
url: /post/uuid/e9631120-573d-5971-bea2-2d044e4ba18d
---

## Design — How is a Go kit microservice modeled?

Putting all these concepts together, we see that Go kit microservices are modeled like an onion, with many layers. The layers can be grouped into our three domains. The innermost **service** domain is where everything is based on your specific service definition, and where all of the business logic is implemented. The middle **endpoint** domain is where each method of your service is abstracted to the generic [endpoint.Endpoint](https://godoc.org/github.com/go-kit/kit/endpoint#Endpoint), and where safety and antifragile logic is implemented. Finally, the outermost **transport** domain is where endpoints are bound to concrete transports like HTTP or gRPC.

You implement the core business logic by defining an interface for your service and providing a concrete implementation. Then, you write service middlewares to provide additional functionality, like logging, analytics, instrumentation — anything that needs knowledge of your business domain.

Go kit provides endpoint and transport domain middlewares, for functionality like rate limiting, circuit breaking, load balancing, and distributed tracing — all of which are generally agnostic to your business domain.

In short, Go kit tries to enforce strict **separation of concerns** through studious use of the **middleware** (or decorator) pattern.

![Go kit service diagram](http://gokit.io/faq/onion.png)

