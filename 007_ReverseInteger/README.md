# Reverse Integer
[LeetCode](https://leetcode.com/problems/reverse-integer/description/)

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
```
Input: 123
Output: 321
```

Example 2:
```
Input: -123
Output: -321
```

Example 3:
```
Input: 120
Output: 21
```

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: `[−2^31,  2^31 − 1]`. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

### Solution
- Continuously get the modulo of `x%10` to separated input's number characters, then recombine those number characters to actual reversed int32.  
Time complexity: O(n)
Space complexity: O(1)
