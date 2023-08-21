# Zigzag Conversion
[LeetCode](https://leetcode.com/problems/zigzag-conversion/description/)

The string `PAYPALISHIRING` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: `PAHNAPLSIIGYIR`

Write the code that will take a string and make this conversion given a number of rows:
```
string convert(string s, int numRows);
```

Example 1:
```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

Example 2:
```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
```

### Solution
I noticed the pattern for character of each zigzag row as below:
- First and last row: `2*numRows - 2 + prevIdx`
- Second row to (numRows-1)th row: `(2*numRows-2+prevIdx) - 2*i`  
with:
  - numRows: Number of rows to generate zigzag string.
  - i: Row index, `[1..numRows-2]`.
  - prevIdx: Index of previous character in current row, `[i..len(s)-1]`
