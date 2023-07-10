Threads (threads.net) Go API wrapper

### Installation

```bash
$ go get github.com/antonprokopovich/threadsnet
```
### Initialization

Initialize the instance of `Threads` type to use its public methods that correspond to Threads API endpoints:
```go
import (
    "github.com/antonprokopovich/threadsnet"
)


threads, err := threadsnet.NewThreads()
if err != nil {
    fmt.Println("Error:", err)
}
```

### Examples
See [/examples](./examples) folder for a runnable example of every available API method.
