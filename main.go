package main

type Profile struct {
	Profile_Name string
}

func main() {
	Myself := Profile{"Joseph Theiss"}
	println("Hello! My name is " + Myself.Profile_Name)
}
