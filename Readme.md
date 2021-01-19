# Fibonacci sequence implementation

This is a simple Fibonacci sequence implemented via HTTP handler that takes an argument and responds with a sequence up to that argument. The data is sent in chuncs 

To run 
`go run main.go`
To use, open a separate terminal and send CURL request such as 
```
curl -X GET "http://localhost:8080/fib?position=30000"
``` 
