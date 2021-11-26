# Authorizer application

## Code challenge of Nubank

This application is in charge of authorizing the transactions in the accounts.

## Discussing regarding the technical and architectural decisions

For my solution I choose the **Go** language because I'm feel comfortable with it. It's a very powerful 
and performant language and very nice to implements microservices.

I choose to apply the **hexagonal architecture** (a.k.a. DDD or Domain-Driven Design) because I've been working with this 
architecture for 2+ years, and I think if it is well applied it gives solutions that scale correctly and each layer has your
own logic and responsibilities well separated from the rest. That prevents possible bad practices (e.g. spaghetti code) 
and if in the future you want to change a component, it is very easy to do. 

## Reasoning about the frameworks used (if any framework/library was used);

I don't use any extra framework for my solution, I use built-in libraries except for [zap logger](https://github.com/uber-go/zap) 
(logger of Uber) because this library is a logger with steroids. Anyway I've created a [wrapper](./pkg/log/log.go) to avoid adding external 
dependencies or libraries to my core application, and if in the future I want to change my logger I only have to modify my pkg.

I use **[Go modules](https://go.dev/blog/using-go-modules)** to manage my dependencies. 

## How to compile and run the project

#### Requirements

- Go 1.15
- Operations file to send via stdin

#### Run

- If it's your first running, you must download the dependencies. Run at the root project:

```
-> go mod tidy
```

- To run main package:

``` 
-> go run main.go < {your_operations_file} 
```

## Additional comments

I did a diagram to show us graphically the flow of the application

![activity diagram authorizer flow](./docs/activiry_diagram_authorizer_flow.svg)

I hope to meet your expectations and have created a simple and elegant application.

Matias Bravo :wolf: