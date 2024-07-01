class Solution {
  fun kidsWithCandies(candies: IntArray, extraCandies: Int): List<Boolean> {
    return candies.map { it + extraCandies >= candies.max() }
  }
}

fun main() {
  Solution().kidsWithCandies(intArrayOf(2, 3, 5, 1, 3), 3)
}
