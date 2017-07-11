package main

func main() {
	for i := 1; i < 101; i++ {
		n, m := i%3 == 0, i%5 == 0
		if !(n || m) {
			print(i)
		} else {
			if n {
				print("Fizz")
			}
			if m {
				print("Buzz")
			}
		}
		println()
	}
}
