###type Error
```go
type Error struct {
    Name string
    Err  error
}
```
Error records the name of a binary that failed to be executed and the reason it failed.
```go
func (e *Error) Error() string
```

###type ExitError
```go
type ExitError struct {
    *os.ProcessState
}
```
An ExitError reports an unsuccessful exit by a command.
```go
func (e *ExitError) Error() string
```