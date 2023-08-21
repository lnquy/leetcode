# Longest Palindromic Substring
[LeetCode](https://leetcode.com/problems/longest-palindromic-substring/description/)

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

Example 2:
```
Input: "cbbd"
Output: "bb"
```

### Solution
- Traverse through string, at each index, expanding to the left and right of index to find the length of palindrome substring. Update the start and end of longest palindrome substring if required.  
Time complexity: `O(n^2)`  
Space complexity: `O(1)`
