# Running the examples

Run the server, like `go run main.go`
1. Home
```
curl localhost:8080
```

2. Posts Index 
```
curl -d t=1 localhost:8080/posts
posts create
```

3. Show Post
```
curl localhost:8080/posts/1
showing post 1
```

4. Post Update
```
curl -X PUT localhost:8080/posts/1
post update
```

5. Post Edit
```
curl localhost:8080/posts/1/edit
post edit
```

