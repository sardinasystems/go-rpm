/*
A native implementation of the RPM file specification in Go.

	package main

	import (
		"fmt"
		"github.com/jfrog/go-rpm"
	)

	func main() {
		p, err := rpm.OpenPackageFile("my-package.rpm")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Loaded package: %v", p)
	}

*/
package rpm
