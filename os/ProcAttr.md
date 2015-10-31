###type ProcAttr 
```go 
type ProcAttr struct {
    //If Dir is non-empty, the child changes into the directory before 
    //creating the process.
    Dir string
    //If Env is non-nil, it gives the environment variable for 
    //new process in the form returned by Environ.
    //If it is nil, the result of Environ will be used.
    Env []string
    //Files specifies the open files inherited by the new process. The 
    //first three entires correspond to standard input,standard output, and
    //standard error. An implementation may support additional entries,
    //depending on the underlying operating system, A nil entry corresponds 
    //to that file being closed when the process starts.
    Files []*File  
    //Operating system-specific process creation attributes.
    //Note that setting this field means that your program
    //may not execute properly or even compile on some
    //operating systems 
    Sys *syscall.SysProcAttr
}
```
ProcAttr holds the attributes that will be applied to a new process started by StartProcess.

