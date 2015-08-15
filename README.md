# gore/rego

Testing URL redirection in go

## Usage

* Start server: `go run src/gore.go`
* Generate an encode URL fragment: `REQ=$(go run req.go http://www.example.com)`
* Print it out: `echo $REQ`
* Request redirect URL: `curl -vi localhost:12345/r/${REQ}`
