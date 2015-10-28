###type File   
```go   
type File struct {
    //contains filtered or unexported fields
}
```     
File represents an open file descriptor.    

```go   
//Chdir changes the current working directory to the file ,
//while must be a directory,
//if there is an error,it will be of type *PathError.
func (f *File) Chdir() err  
    
//Chmod changes the mode of the file to mode,
//if there is an error,it will be of type *PathError    
func (f *File) Chmod(mode FileMode) error   
    
//Chown changes the numeric uid and gid of the name file,   
//if there is an error,it will be of type *PathError    
func (f *File) Chown(uid,gid int) error     
    
//Close closes the File,
//rendering it unusable for I/O.
//It return an error,if any.
func (f *File) Close() error

//Fd 返回文件句柄
//Fd returns the interger  Unix file descriptor referencint the open file.
//The file descriptor is valid only until f.Close  id callled or f is garbage collected.
func (f *File) Fd() uintptr

//Name returns the name of the file as presented to 
func (f *File) Name string

//Reaf reads up to len(b) bytes from the File.
//It returns the number of bytes read and an error,if any .
//EOF is signaled by a zero count with err set to io.EOF.
func (f *File) Read(b []byte) (n int, err error)

//ReadAt reads len(b) bytes from the File starting at bytes offset off.
//It return the number of bytes read anf the error,if any.
//ReadAt always returns a non-nil error when n < len(b).
//At end of file,that error is io.EOF.
func (f *File) ReadAt(n int) (fi []FileInfo, er error)

//Readdir reads the contents of the directory associated with file and reurn a sile of up n FileInfo values,
//as would be returned by Lstat, in directory order/
//Subsequent calls on the same file will yield future FileInfos
//if n > 0, Readdir returns at most n FileInfo structures.
//In this case, if Readdir return an empty slice.it will return a non-nil error explaining why,
//At the end of a directory, the error is io.EOF
//if n <= 0,Readdir returns all the FileInfo from the directory in a single slice.
//In this case, if Readdir succeedd(read all the way to the direcory), it return an slice and a nil error.
//Id it encounter an error before the end of the directory,
//Readdir returns the FileInfo read until that point and a non-nil error.
func (f *File) Readdir(n int)(fi []FielInfo, err error)

//Readdirnames reads and returns a slice of names from the directory f.
//If n > 0, Readdirnames returns at most n names.
// In this case, if Readdirnames returns an empty slice, it will return a non-nil error explaining why. 
//At the end of a directory, the error is io.EOF.
//If n <= 0, Readdirnames returns all the names from the directory in a single slice. 
//In this case, if Readdirnames succeeds (reads all the way to the end of the directory), 
//it returns the slice and a nil error. 
//If it encounters an error before the end of the directory, 
//Readdirnames returns the names read until that point and a non-nil error.
func (f *File) Readdirnames(n int)(name []string,err error)

//Seek set the offset for the next Read or Write on file to offset,
//interpreted according to whence: 0 means relative relative to the origin of the file,
//i means relative to the current offset, and 2 means relative to the end.
//It returns the offset and an error, id any.
func (f *File) Seek(offset int64, whence int) (ret int64, err error)

//Stat returns the FileInfo structure describint file/
//If there is an error, it will be of type *PathError.
func (f *File) Stat() (FileInfo, error)

//Sync commits the current contents of the file to stable storage.
//Typically, this means flushing the file system's inmemeory copy of recently written data to disk.
func (f *File) Sync() error

//Truncate changes the size of the file.
//It dose not change the I/O offset.
//If there is an error, it will be of type *PathError.
func (f *File) Truncate(size int64) error

//Write writes len(b) bytes to the File.
//It return the number of bytes written and an error ,if any.
//Write returns a non-nil error when n != len(b)
func (f *File) Write(b []byte) (n int, err error)

//WriteAr writes len(b) bytes to the File .
//File starting at byte offset off. It returns the number of bytes written and an error, if any. 
//WriteAt returns a non-nil error when n != len(b) 
func (f *File) WriteAt(b []byte, off int64) (n int, err error)

//WriteString is like Write, but writes the contents of string s rather than a slice of bytes.
func (f *File) WriteString(s string) (n int, err error)
```