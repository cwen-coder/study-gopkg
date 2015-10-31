###type UnknownUserError  
+ func (e UnknownUserError) Error() string  
###type UnknownUserIdError  
+ func (e UnknownUserIdError) Error() string  
###type User  
+ func Current() (*User, error)  
+ func Lookup(username string) (*User, error)  
+ func LookupId(uid string) (*User, error)  
