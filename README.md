# concurrence-me

Project to understand the basics of Golang concurrency

## Documentation


## Commands

````bash
// Simple test with coverage pourcentage
go test -cover ./... -v

// With html 
go test -coverprofile='coverage.out' ./...
go tool cover -html='coverage.out'
````

## Structure

- [Example n°1](./1-example/)
  - Basics of goroutine execution
- [Example n°2](./2-example/)
  - Use of sync.waitGroup
- [Exercise n°1](./3-challenge/)


