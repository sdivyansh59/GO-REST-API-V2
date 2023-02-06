package main

import (
	"fmt"
)

// Run - is going to be responsible for
// the nstantiation and startup of our
// go application

func Run() error {
	fmt.Fprint("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REst Api V2")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}