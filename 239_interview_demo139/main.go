package main

func alwaysFalse() bool {
	return false
}

func main() {
	switch alwaysFalse()
	{
	case false:
		println(false)
	case true:
		println(true)
	}
}
