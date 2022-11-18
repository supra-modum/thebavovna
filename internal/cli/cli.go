package cli

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Print(`This program runs Bavovna backend server.
 
Usage:
 
bavovna [arguments]
 
Supported arguments:
 
`)
	flag.PrintDefaults()
	os.Exit(1)
}

func Parse() {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	fmt.Println(*env)
}
