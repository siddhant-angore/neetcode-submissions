class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        length = len(nums)
        ans = [0] * (2 * length)
        for index in range(length):
            ans[index] = ans[index + length] = nums[index]
        return ans
        