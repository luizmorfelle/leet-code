package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i, v := range flowerbed {
		if v == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}

func main() {
	println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}
