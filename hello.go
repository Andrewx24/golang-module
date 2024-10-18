// main.go
package main

import (
    "golang-module/mypackage"  // Importing the mypackage
)
func main() {
    println("Hello, World!")
    mypackage.New()  // Calling the New function from mypackage
}
