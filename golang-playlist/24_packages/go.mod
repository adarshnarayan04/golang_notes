module github.com/codersgyan/podcast

go 1.22.5

require github.com/fatih/color v1.17.0

//indirect dependencies means these packages are not directly used in our code but are required by other packages that we are using
//as we used color package which internally uses these packages
//but in our code we never used these packages directly 
//so they are indirect dependencies
require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.18.0 // indirect
)
