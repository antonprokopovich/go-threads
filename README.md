<div align="center">
Threads (threads.net) Go API wrapper
<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#initialization">Initialization</a> •
  <a href="#examples">Examples</a> •
  <a href="#other-languages">Other languages</a>
</p>

</div>

## Installation

```bash
$ go get github.com/antonprokopovich/threadsnet
```
## Initialization

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
## API
### User
#### Get User ID by username
Pass a user's username string as an argument.
```go
id, err := threads.GetUserID("zuck")
```

#### Get User by ID
Pass a user's ID number as an argument.
```go
user, err := threads.GetUser(314216)
```

#### Get User's replies by user's ID
Pass user's ID number as an argument.
```go
replies, err := threads.GetUserReplies(314216)
```

#### Get User's threads by user's ID
Pass user's ID number as an argument.
```go
threads, err := threads.GetUserThreads(314216)
```
### Post
#### Get Post by ID
Pass a post's ID number as an argument.
```go
post, err := threads.GetPost(3141002295235099165)
```

## Run examples
See [/examples](./examples) folder for a runnable example of every available API method.
