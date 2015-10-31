###type Signal 
```go 
type Signal interface {
    String() string
    Signal() //to distinguish from other Stringers
}
```      
A Signal represents an operating system signal.   
The usual underlying implementation is operating system-dependent: on Unix it is syscall.Signal.   

```go 
var (
    Interrupt Signal = syscall.SIGINT
    Kill Signal = syscall.SIGKILL
)
```     
The only signal values guaranteed to be present on all systems are Interrupt (send the process an interrupt) and Kill (force the process to exit).
