# Single Number
[LeetCode](https://leetcode.com/problems/single-number/description/)

Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
```
Input: [2,2,1]
Output: 1
```

Example 2:
```
Input: [4,1,2,1,2]
Output: 4
```

### Solution
 - a ^ a = 0
 - So just have to XOR all items in array. 
