###type ProcessState  
```go 
type ProcessState struct {
    //contains filtered or unexported fields
}
``` 
ProcessState stores information about a process, as reported by Wait.  

```go 
//Exited reports whether the program has exited
//判断一个进程是否已退出
func (p *ProcessState) Exited() bool

//Pid returns the process id of the exited process 
//获取已退出进程的pid
func (p *ProcessState) Pid() int 

//获取进程状态的字符串
func (p *ProcessState) String() string

//Success reports whether the program exited successfully, such as with exit 0 on Unix
func (p *ProcessState) Success() bool  

//Sys returns system-dependent exit information about the exited process.
//Convert it to the appropriate underlying type,
//such as *syscall.WaitStatus on Unix, to access its contents.
//获取进程的退出信息
func (p *ProcessState) Sys() interface{}

//获取进程资源使用情况
func (p *ProcessState) SysUsage() interface{}

//SystemTime returns the system CPU time of the exited process and its children
func (p *ProcessState) SystemTime() time.Duration

//UserTime return the user CPU time of the exited process and its children.
func (p *ProcessState) UserTime() time.Duration
```
