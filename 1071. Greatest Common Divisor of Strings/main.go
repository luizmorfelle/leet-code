package main

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	if len(str1) > len(str2) {
		return str1[:gcd(len(str1), len(str2))]
	}

	return str2[:gcd(len(str1), len(str2))]

}

func gcd(a, b int) int {
	println(a, b)
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	println(gcdOfStrings("NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM", "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"))
}
