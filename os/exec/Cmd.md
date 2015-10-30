###type Cmd
```go
type Cmd struct {
    //Path is the path of the command to run .
    //This is the only field that must be set to non-zero value.
    //If Path is relative , it is evaluated relative to Dir.
    Path string 

    //Args holds command line arguments, including the command as Args[0].
    //If the Args field is empty or nil, Run uses {Path}.
    //In typical use, both Path and Args are set by calling Command.
    Args []string

    //Env specifies the environment of the process.
    //If Env is nil, Run uses the current process's environment.
    Env []string

    //Dir specifies the working directory of the command.
    //If Dir is the empty,Run runs the command in the 
    //calling process's current directory.
    Dir string  

    //Stdin specifies the process's standard input.
    //If Stdin is nil, the process reads from the null device(os.DevNull).
    //If Stdin is an *os.File, the process's input is connetcted 
    //directory to that file.
    //Otherwise, during the execution of the command a separate
    //gorountine reads from Stdin and delivers that data to the command 
    //over a pipe. In this case,Wait does not complete until the grountine
    //stops copying, either because it has reached the end of Stdin
    //(EOF or a read error) or because writing to the pipe returned an error.
    Stdin io.Reader

    // Stdout and Stderr specify the process's standard output and error.
    //
    // If either is nil, Run connects the corresponding file descriptor
    // to the null device (os.DevNull).
    //
    // If Stdout and Stderr are the same writer, at most one
    // goroutine at a time will call Write.
    Stdout io.Writer
    Stderr io.Writer

    // ExtraFiles specifies additional open files to be inherited by the
    // new process. It does not include standard input, standard output, or
    // standard error. If non-nil, entry i becomes file descriptor 3+i.
    //
    // BUG(rsc): On OS X 10.6, child processes may sometimes inherit unwanted fds.
    // https://golang.org/issue/2603
    ExtraFiles []*os.File

    // SysProcAttr holds optional, operating system-specific attributes.
    // Run passes it to os.StartProcess as the os.ProcAttr's Sys field.
    SysProcAttr *syscall.SysProcAttr

    // Process is the underlying process, once started.
    Process *os.Process

    // ProcessState contains information about an exited process,
    // available after a call to Wait or Run.
    ProcessState *os.ProcessState
    // contains filtered or unexported fields
}
```

Cmd represents an external command being prepared or run.   

```go
//CombinedOutput runs the command and returns its combined standard output and standard error
func (c *Cmd) CombinedOutput() ([]byte, error) 

//Output runs the command and returns its standard output
func (c *Cmd) Output() ([]byte, error)  

Example   
out, err := exec.Command("date").Output()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("The date is %s\n", out)

//Run starts the specified command and waits for it to complete.\
//The returned error is nil command runs, 
//has no problems copying stdin,stdout,and stderr, and exits with a zero exit status
//If the command fails to run or doesn't complete successfully, the error is of type *ExitError.
//Other error types may be returned fo I/O problems.
func (c *Cmd) Run() error

//Start starts the specifiesd command but does not wait for it to complete.
//The Wait method will return the exit code and release associated resources 
//once the command exits.
func (c *Cmd) Start() error  

func (c *Cmd) StderrPipe() (io.ReadCloser, error) 
func (c *Cmd) StdinPipe() (io.WriteCloser， error)
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
func (c *Cmd) Wait() error

```  