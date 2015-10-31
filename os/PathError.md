###type PathError  
```go 
type PathError stuct {
    Op string
    Path string
    Err error
}
```
PathError records an error and the operation and file path that caused it

```go 
func (e *PathError) Error() string
```