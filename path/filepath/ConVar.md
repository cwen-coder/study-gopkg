###Constants 
```go 
const (
    Separator = os.PathSeparator
    ListSeparator = os.PathListSeparator
)
```

###Variables 
```go 
//ErrBadPattern indicates a globbing pattern was malformed.
var ErrBadPattern = errors.New("syntax error in pathern")

//SkipDir is used as a return value from WalkFuncs to indicate that the 
//directory named in the call is to be skipped.
//It is not returned as an error by any function.
var SkipDir = errors.New("skip this directory")
```