package main

import "github/adarshnarayan04/learninggo/auth"
//to import our own package
// first we need to initialize a module using "go mod init <module-name>"
//generally module name is the repository path where the code will be hosted
// then we can import our own package using the module name followed by the package path
// ex: module name is "github.com/adarshnarayan04/learninggo" and package path is "auth"
// so we can import the package using "github.com/adarshnarayan04/learninggo/auth"

//it will create go.mod and go.sum file in the current folder
//like npm init in nodejs
func main() {
	auth.Login("user", "password")
}
