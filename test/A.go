package main

func main() {
	bytes := make([]byte, 1)
	cbytes := bytes

	bytes[0] = 1
	println(bytes[0])
	println(cbytes[0])
	cbytes[0] = 2
	println(bytes[0])
	println(cbytes[0])
}
