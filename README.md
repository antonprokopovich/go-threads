<div align="center">
Threads (threads.net) Go API wrapper
<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#initialization">Initialization</a> •
  <a href="#examples">Examples</a> •
  <a href="#other-languages">Other languages</a>
</p>

</div>

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

### Other languages

Threads.net API implementations in other languages:

JavaScript: 
- https://github.com/threadsjs/threads.js

Python:
- https://github.com/dmytrostriletskyi/threads-net
