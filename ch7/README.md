## Simple Hello HTTP server

A simple HTTP server based on the example presented in chapter 7 with minor modifications which responds with the appropriate greeting based on the user ID, presenting the following concepts from the book:

- More types: structs, user-defined types, pointers
- Methods and interfaces
- Dependency injection
- Basic error handling
- JSON marshaling and unmarshaling
- HTTP request and response

### API overview

#### `GET /ping`

Unconditionally responds with 200 OK and the following JSON payload:

```json
{
  "message": "pong"
}
```

#### `POST /hello`

Accepts a JSON payload of the form:

```json
{
  "user_id": 3
}
```

The `user_id` field can be any number \(integer\). On success, a 200 OK response is returned with a JSON payload of the following form based on the provided user ID:

```json
{
  "message": "Hello, Pat"
}
```

Other possible status codes based on invalid input:

- 405 Method Not Allowed: when the endpoint `/hello` is invoked with an HTTP verb other than `POST`
- 400 Bad Request: when the JSON payload is malformed or the provided `user_id` could not be found

In any case, the response body is always a JSON payload of the form:

```json
{
  "message": "reason for the error encountered"
}
```

### Developing

Follow the instructions in the project README and run `make`. You should not see any meaningful program output - instead, the terminal should appear to hang which indicates that the HTTP server is running as expected and listening at port `8080/tcp`.

While the HTTP server is running, open your web browser and visit the URL [http://localhost:8080/ping](http://localhost:8080/ping). You should see the following output:

```json
{
  "message": "pong"
}
```

For more extensive testing of the `/hello` endpoint involving POST requests with JSON payloads, you may wish to leverage the following additional tools:

- Within your web browser: leverage the developer console by right-clicking and selecting "Inspect Element" or similar
- GUI application: try [Postman](https://www.postman.com/downloads/) or similar tools
- Command line: [cURL](https://curl.se/) is your best friend \([wget](https://www.gnu.org/software/wget/) is also fine\)

For example, with cURL:

```bash
curl -i \
    -XPOST \
    -H'Content-Type: application/json' \
    -d '{"user_id":3}' \
    http://localhost:8080/hello
```

Return to your terminal session and press Ctrl+C to stop the HTTP server when done.

Note that the port at which the HTTP server listens to is configurable. For example, to make it listen on port `9090/tcp` instead:

```bash
make PORT=9090
```
