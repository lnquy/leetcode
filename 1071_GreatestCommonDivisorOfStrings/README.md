# 1071. Greatest Common Divisor of Strings
[LeetCode](https://leetcode.com/problems/greatest-common-divisor-of-strings)

For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:
```
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
```
Example 2:
```
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
```
Example 3:
```
Input: str1 = "LEET", str2 = "CODE"
Output: ""
```

Constraints:
- 1 <= str1.length, str2.length <= 1000
- str1 and str2 consist of English uppercase letters.


--- 
### Solution:
This problem description is quite confusing.  
The most important one is the example#2 where we only output `AB`, not `ABAB`.
Since `AB` is divisible by both str1 and str2, while `ABAB` will only be divisible by str1.

Idea: 
1. Since the output string must be divisible by both str1 and str2, the maximum length of output string is the greatest common divisor (GCD): `maxGCDLen := gcd(len(str1), len(str2))`
2. Since we only care about the longest GCD string, iterate from the maxGCDLen to 1. At each iterate, check if both str1 and str2 can divide the GCD string or not. The first match is the result.

