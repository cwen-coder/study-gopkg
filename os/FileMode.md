###type FileMode  
```go
type FileMode uint32    
```
A FileMode respresents a file's mode and permission bits. The bits have the same definition on all systems,    
so that information about files can be removed from one system to another portably.     
Not all bits apply to all systems.The only required bit is ModeDir fot directories      

```go   
const (
    // The single letters are the abbreviations
    // used by the String method's formatting.
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
    ModeAppend                                     // a: append-only
    ModeExclusive                                  // l: exclusive use
    ModeTemporary                                  // T: temporary file (not backed up)
    ModeSymlink                                    // L: symbolic link
    ModeDevice                                     // D: device file
    ModeNamedPipe                                  // p: named pipe (FIFO)
    ModeSocket                                     // S: Unix domain socket
    ModeSetuid                                     // u: setuid
    ModeSetgid                                     // g: setgid
    ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
    ModeSticky                                     // t: sticky

    // Mask for the type bits. For regular files, none will be set.
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

    ModePerm FileMode = 0777 // Unix permission bits
)
```     
The defined file mode bits are the most significant bits of the FileMode.   
The nine least-significant bits are the standard Unix rwxrwxrwx permissions.   
The values of these bits should be considered part of the public API and may be used in wire protocols or disk representations: they must not be changed, although new bits might be added.     

```go
//IsDir reports whether m describles a regular file.  
//That is ,it tests that no mode type bits are set.
func (m FileMode) IsDir() bool

//IsRegular reports whether m describle a regular file.
//That is, it tests that no mode type bits are set.
func (m FileMode) IsRegular() bool 

//Perm returns the Unix permission bits in m.
func (m FileMode) Perm() FileMode

//返回文件模式的字符串
func (m FileMode) String() string

```