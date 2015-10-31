###type SyscallError 
```go 
type SyscallError struct {
    Syscall string
    Err error
}
```     
SyscallError records an error from a special system call

```go 
func (e *SyscallError) Error() string
```