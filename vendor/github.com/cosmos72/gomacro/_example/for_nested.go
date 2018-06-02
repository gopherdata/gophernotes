package main

func for_nested(n1, n2, n3 int) int {
	x := 0
	for i := 0; i < n1; i++ {
		for k := 0; k < n2; k++ {
			for j := 0; j < n3; j++ {
				x++
			}
		}
	}
	return x
}

func run_for_nested() {
	for_nested(2, 3, 4)
}

/*
(func (n1, n2, n3 int) int {
	x := 0
	for i := 0; i < n1; i++ {
		for k := 0; k < n2; k++ {
			for j := 0; j < n3; j++ {
				x++
			}
		}
	}
	return x
})(2,3,5)
*/
