package main

type Foo struct{
	X int
	Y string
}

func (foo Foo) sayHello() string {
	return "hello"
}

func (foo Foo) sayGoodBye() string {
	return "goodbye"
}

// This lets the users of this module access Foo by the name Bar
// type alias
type Bar = Foo