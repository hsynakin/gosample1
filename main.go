package main

func main() {

	switch foo := 1; foo {
	case 1:
		println("one")
	case 2:
		println("two")
	}

	if zoo := 1; zoo == 1 {
		println("çubuk")
	} else {
		println("demir  çubuk :)")
	}

	switch ttoo := 4; {
	case ttoo == 1:
		println("one")
	case ttoo > 2:
		println("two")
	}

	for i := 0; i < 3; i++ {
		println(i)
	}
}
