###type LinkError
```go 
type LinkError struct {
    Op string
    Old string
    New string
    Err error
}
```     
LinkError records an error during a link or symlink or rename system call and the paths caused it.      

```go 
//返回LinkError的字符串形式
func (e *LinkError) Error() string
```