```go
func Ignore(sig ...os.Signal)
func Notify(c chan<- os.Signal, sig ...os.Signal)
func Reset(sig ...os.Signal)
func Stop(c chan<- os.Signal)
```