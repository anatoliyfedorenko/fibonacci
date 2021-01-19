# Fibonacci sequence implementation

This is a simple Fibonacci sequence implemented via HTTP handler that takes an argument and responds with a sequence up to that argument. The data is sent in chuncs 

To run 
```
go run main.go
```
To use, open a separate terminal and send CURL request such as 
```
curl -X GET "http://localhost:8080/fib?position=30000"
``` 
Expected output: 
```
0 fib number is 1
1 fib number is 1
2 fib number is 2
3 fib number is 3
4 fib number is 5
5 fib number is 8
6 fib number is 13
7 fib number is 21
8 fib number is 34
9 fib number is 55
10 fib number is 89
... 
```