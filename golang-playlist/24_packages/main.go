package main

import (
	"fmt"

	"github.com/codersgyan/podcast/auth"
	"github.com/codersgyan/podcast/user"
	"github.com/fatih/color"
)
//to import our own package
// first we need to initialize a module using "go mod init <module-name>"
//generally module name is the repository path where the code will be hosted
// then we can import our own package using the module name followed by the package path
// ex: module name is "github.com/adarshnarayan04/learninggo" and package path is "auth"
// so we can import the package using "github.com/adarshnarayan04/learninggo/auth"

//it will create go.mod  file in the current folder
//like npm init in nodejs

//to install the external package use "go get <package-name>"
//ex: go get github.com/fatih/color
//it will download the package and add the dependency in go.mod and go.sum file

//go.sum is created when we use external packages to ensure the integrity of the packages
//it contains the checksums of the packages

//to automatically download all the dependencies mentioned in go.mod file use "go mod tidy"
// like npm install in nodejs
func main() {
	auth.LoginWithCredentials("codersgyan", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}
