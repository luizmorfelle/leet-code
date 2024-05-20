class Solution(object):
    def twoSum(self, nums, target):
        response = []
        index = 0
        response.append(index)

        while index < len(nums) and len(response) == 1:
            response[0] = index
            for i in range(len(nums)):
                if nums[i] + nums[index] == target and i != index:
                    response.append(i)
                    break

            index += 1
        return response


sol = Solution().twoSum([2, 7, 11, 15], 9)
print(sol)
