package auth//we use package name as same as folder name ( it is a good practice), can be anything though

func extractSession() string {
	return "loggedin"
}

func GetSession() string {
	return extractSession()
}

func checking()  {
	LoginWithCredentials("hello","world");//can access the function in another file of same package

}
