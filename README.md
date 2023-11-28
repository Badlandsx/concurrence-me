# concurrence-me

Project to understand the basics of Golang concurrency

## Documentation


## Commands

````bash
# Simple test with coverage pourcentage
go test -cover ./... -v

# With html 
go test -coverprofile='coverage.out' ./...
go tool cover -html='coverage.out'
````

## Structure

- [Basics goroutine](./01-example/)
  - Basics of goroutine execution
- [Goroutine / WaitGroup](./02-example/)
  - Use of sync.waitGroup
- [Exercising](./03-challenge/)
- [Mutex basics](./04-example/)

````bash
# Same result is given
go run main1.go
# Error occured with race because multiple goroutine access the same data
# We don't know which goroutine will finish first/last
go run -race main1.go
````

> Use mutex to solve this problem of accessing the same data

````bash
# Not the same result because one goroutine access the data after each other
go run main2.go
# No errors with mutex
go run -race main2.go
````



