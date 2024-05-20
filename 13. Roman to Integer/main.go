package main

func romanToInt(s string) int {
	sum := 0
	sub := 0

	var m = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for _, char := range s {
		value := m[string(char)]

		if sub != 0 && sub < value {
			sum = sum + (value - 2*sub)
			sub = 0
			continue
		}
		sub = value
		sum += value
	}
	return sum
}

func main() {
	println(romanToInt("MCMXCIV"))
}
