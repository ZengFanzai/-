package main

func main() {

	//for i:=0;i<10 ;i++  {
	//loop:
	//	println(i)
	//}
	//goto loop

	var i int
loop:
	println(i)
	if i < 10 {
		i++
		goto loop
	}

}
