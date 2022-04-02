class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        sorted_nums = []
        for idx, num in enumerate(nums):
          sorted_nums.append((num, idx,))
        sorted_nums.sort(key=lambda x: x[0])

        right = len(sorted_nums) - 1
        for left, num_arr in enumerate(sorted_nums):
          while left < right and num_arr[0] + sorted_nums[right][0] > target:
            right -= 1
          if left < right and num_arr[0] + sorted_nums[right][0] == target:
            return sorted([num_arr[1], sorted_nums[right][1]])
          
        return [0, 0]
        

