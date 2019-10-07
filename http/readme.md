# Http

## BasicAuth

This is a utility tool to make HTTP Basic Auth easy & relatively safe.

It wraps around http handler functions: 
```go
http.HandleFunc("/", auth.BasicAuth(MyHandlingFunction, "my_username", "my_password", "Please log in"))
```
While Basic Auth tends to use hardcoded passwords in plaintext, this wrapper uses hashing purely to prevent timing attacks.

## AddParams

Adds Parameters to a URL. 