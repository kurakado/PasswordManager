package main

import (
	"flag"
	"fmt"
	"os"
)

/*
var (
		intOpt  = flag.Int("i", 1234, "help message for i option")
		boolOpt = flag.Bool("b", false, "help message for b option")
	url = flag.String("u", "http://example.com", "enter a URL of target.")
)
*/

func command_parser() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
    %s TARGET PASSWORD [OPTIONS]
Optins
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var url = fs.String("u", "http://example.com", "enter a URL of target.")
	target := os.Args[1]
	password := os.Args[2]
	if len(os.Args) >= 3 {
		fs.Parse(os.Args[3:])
	}
	fmt.Println(flag.Args())
	fmt.Println("target: ", target)
	fmt.Println("password: ", password)
	fmt.Println("url: ", *url)
}

func main() {
	command_parser()
}
