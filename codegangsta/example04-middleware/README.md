# Running the examples

Run the server, like `go run main.go`
1. Test the URL without the password, it results on 401 error
```
curl localhost:8080
```

2. Add the password parameter
```
curl localhost:8080?password=secret123
```
