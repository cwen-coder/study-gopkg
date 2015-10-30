// Command returns the Cmd struct to execute the named program with the given arguments.
// It sets only the Path and Args in the returned structure.
// If name contains no path separators. Command users LookPath to resolve tht path a complete name if possible
// OtherWise it users name directly.
// The returned Cmd's Args field is constructed from the command name follwed by the elements of args ,so
// arg should not include the command name itself.For example ,Command("echo","hello")
package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
