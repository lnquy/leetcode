# Median of Two Sorted Arrays
[LeetCode](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:
```
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
```

Example 2:
```
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```

### Solution
- The easiest way is to merge first `(nums1 + nums2) / 2` items to an external array, the last item (or 2 items) are the median.  
This solution (`findMedianSortedArrays_Slow`) has `O((m+n)/2)` and I couldn't come up with a `O(log(m+n))` solution, so I read the [LeetCode's solution](https://leetcode.com/problems/median-of-two-sorted-arrays/solution/) and re-implement it.
