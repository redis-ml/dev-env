import sys
import unittest

from solution import Solution


class SolutionTest(unittest.TestCase):
  def test_solution(self):
    s = Solution()
    nums = [1, 2, 3]
    target = 5
    ret = s.twoSum(nums, target)
    print("result:", ret)
    print(sys.version_info)
    assert ret == [1, 2]

def test_solution():
  assert 1 == 0


if __name__ == "__main__":
    unittest.main()
