### Constans    
```go 
const(
    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
    O_RDWR   int = syscall.O_RDWR   // open the file read-write.
    O_APPEND int = syscall.O_APPEND // append data to the file when writing.
    O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
    O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
    O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
    O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
)
```     
Flags to Open wrapping those of the underlying system. Not all flags may be implemented on a given system.      

```go   
    const (
    SEEK_SET int = 0 // seek relative to the origin of the file
    SEEK_CUR int = 1 // seek relative to the current offset
    SEEK_END int = 2 // seek relative to the end
)
```     
Seek whence values.     
```go   
const(
    PathSeparator     = '/' // OS-specific path separator
    PathListSeparator = ':' // OS-specific path list separator
)
```     
```go   
const DevNull = "/dev/null"
```     
DevNull is the name of the operating system's “null device.” On Unix-like systems, it is "/dev/null"; on Windows, "NUL".  
### Variables       
```go   
var (
    ErrInvalid    = errors.New("invalid argument")
    ErrPermission = errors.New("permission denied")
    ErrExist      = errors.New("file already exists")
    ErrNotExist   = errors.New("file does not exist")
)
``` 
Portable analogs of some common system call errors.     
```go 
var (
    Stdin = NewFile(uintptr(syscall.Stdin),"/dev/stdin")
    Stdout = NewFile(uintptr(syscall.stdout),"/dev/stdout")
    Stderr = NewFile(uintptr(syscall.stderr),"/dev/stderr")
)
```     
Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.      
```go 
var Args []string
```     
Args hold the command-line arguments, starting with the program name.
