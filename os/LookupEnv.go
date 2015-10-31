/*
func LookupEnv(key string) (string, bool)
LookupEnv retrieves the value of the environment variable named by the key.
If the variable is present in the environment the value (which may be empty) is returned and the boolean is true.
Otherwise the returned value will be empty and the boolean will be false.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.LookupEnv("LANG"))
	fmt.Println(os.LookupEnv("ggg"))
}
