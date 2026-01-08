package auth ////we use package name as same as folder name ( it is a good practice), can be anything though

import "fmt"

//to import the funciton in another package , the function name should start with uppercase letter
//that why name of all the inbuit functions in golang start with uppercase letter
//ex: fmt.Println , os.Open , etc
//fmt is package name here, Println is function name
func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}
