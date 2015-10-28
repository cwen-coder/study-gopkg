###type FileInfo    
```go 
type FileInfo interface {
    Name() string       //base name of the file
    Size() int64        //length in bytes for regular files;system-dependent for others
    Mode() FileMode     //file mode bits
    ModTime() time.Time //modification time
    IsDir() bool        //abbreviation for Mode().IsDir()
    Sys() interface{}   //underlying data source (can return nil)
}
``` 
A FileInfo describle a file and is returned by Stat and Lstat

```go
//Lstat returns a FileInfo describing the named file.
//If the file is a symbolic link,the returned FileInfo describles the sysmblic link.
//Lstat makes no attempt to follow the link.   
//If there is an error ,it will be of tyoe *PathError  
func Lstat(name string) (FileInfo, error)  

//Stat returs a FileInfo describle the named file.   
//if there is an error, it will be of tyoe *PathError.  
func Stat(name string) (FileInfo, error)    
```