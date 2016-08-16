// +build windows
package main

import "log"

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}
