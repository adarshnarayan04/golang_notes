package main

type MyInterface interface {
	Method1()
	Method2()
}

type MyStruct struct{}

func (o MyStruct) Method1()  {}
func (o *MyStruct) Method2() {}

func main() {
	var i MyInterface

	s := MyStruct{}
	i = s // not valid, because Method2 has pointer receiver
	// so for value type, only Method1 is implemented, but Method2 is not, so it will give error

	ps := &MyStruct{}
	i = ps // valid, because pointer type implements both Method1 and Method2
	// 	A pointer type (*MyStruct) has:
	// methods with value receivers
	// methods with pointer receivers
}
// Generally , we always use pointer receiver for methods, because it allows us to modify the original struct and 
// also avoids unnecessary copying of the struct when the method is called.