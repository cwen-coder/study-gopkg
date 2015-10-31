###type Process 
```go 
type Process struct {
    Pid int
    //contains filtered or unexported fields
}
``` 
Process stores the information about a process created by StartProcess

```go 
//FindProcess looks for a running process by its pid.
//The Process it returns can be used to obtain information about 
//the underlying operation system process.
func FindProcess(pid int) (p *Precess, err error)

//StartProcess starts a new process with the program,
//arguments and attributes specified by name,argv and attr.
//StartProcess is a low-level interface. The os/exec package provides higher-level interfaces.
//If there is an error, it will be of type *PathError.
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

//Kill causes the Process to exit immediately 
func (p *Process) Kill() error

//Release releases any resources associated with the Process p,
//rending it unusable in the future.
//Release only needs to be called if Wait is not
//这个函数主要是释放一个进程相关的资源
func (p *Process) Release() error  

//Signal sends a signal to the Process.
//Sending Interrupt on Windows is not implemented.
//给一个进程发送一个信号
func (p *Process) Signal(sig Signal) error

//Wait waits for the Process to exit, 
//and then returns a ProcessState describing its status and an error, if any.
//Wait releases any resource associated with the Process.
//On most operating systems,
//the Process must be a child of the current process or an error will be returned.
//等待一个进程的执行完成
func (p *Process) Wait() (*ProcessState, err) 
```