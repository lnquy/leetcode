# Binary Gap (LeetCode)
[LeetCode](https://leetcode.com/problems/binary-gap/description/)

Given a positive integer N, find and return the longest distance between two consecutive 1's in the binary representation of N.

If there aren't two consecutive 1's, return 0.

Example 1:
```
Input: 22
Output: 2
Explanation: 
22 in binary is 0b10110.
In the binary representation of 22, there are three ones, and two consecutive pairs of 1's.
The first consecutive pair of 1's have distance 2.
The second consecutive pair of 1's have distance 1.
The answer is the largest of these two distances, which is 2.
```

Example 2:
```
Input: 5
Output: 2
Explanation: 
5 in binary is 0b101.
```

Example 3:
```
Input: 6
Output: 1
Explanation: 
6 in binary is 0b110.
```

Example 4:
```
Input: 8
Output: 0
Explanation: 
8 in binary is 0b1000.
There aren't any consecutive pairs of 1's in the binary representation of 8, so we return 0.
```

Note:
- `1 <= N <= 10^9`


# Binary Gap (Codility)
[Codility](https://app.codility.com/demo/results/trainingECGAX7-EVD/#)

A *binary gap* within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation `1001` and contains a binary gap of length 2. The number 529 has binary representation `1000010001` and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation `10100` and contains one binary gap of length 1. The number 15 has binary representation `1111` and has no binary gaps. The number 32 has binary representation `100000` and has no binary gaps.

Write a function: `func Solution(N int) int`

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation `10000010001` and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation `100000` and thus no binary gaps.

Write an efficient algorithm for the following assumptions:
- N is an integer within the range `[1..2,147,483,647]`.
