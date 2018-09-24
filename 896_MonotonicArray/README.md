# Monotonic Array
[LeetCode](https://leetcode.com/problems/monotonic-array/description/)

An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all `i <= j`, `A[i] <= A[j]`.  An array A is monotone decreasing if for all `i <= j`, `A[i] >= A[j]`.

Return true if and only if the given array A is monotonic.

 

Example 1:
```
Input: [1,2,2,3]
Output: true
```

Example 2:
```
Input: [6,5,4,4]
Output: true
```

Example 3:
```
Input: [1,3,2]
Output: false
```

Example 4:
```
Input: [1,2,4,5]
Output: true
```

Example 5:
```
Input: [1,1,1]
Output: true
```

Note:
```
1 <= A.length <= 50000
-100000 <= A[i] <= 100000
```

### Solution
- Traverse through `[1..n-1]` items of array.
- Mark the first `diff := a[i] - a[i-1]; diff != 0` as the sign of array (increasing or decreasing).
- If found another `diff != 0` with `(diff > 0 && sign < 0) || (diff < 0 && sign > 0)` then return `false`.  

Time complexity: `O(n)`.  
Space complexity: `O(1)`.
