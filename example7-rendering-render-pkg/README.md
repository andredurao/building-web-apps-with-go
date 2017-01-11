# Example 7 - Using the render package

This example looks simple and served as a very simple introduction to the render package.
An interesting fact, when rendering json or other data types it automatically sets the headers on the response:

```
$ curl localhost:8080/json -D -

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 11 Jan 2017 10:01:52 GMT
Content-Length: 16

{"hello":"json"}
```

I've also added a layout as suggested on the example exercises.
