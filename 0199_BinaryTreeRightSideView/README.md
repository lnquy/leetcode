# 199. Binary Tree Right Side View
[LeetCode](https://leetcode.com/problems/binary-tree-right-side-view/)

Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.



Example 1:
![](https://assets.leetcode.com/uploads/2021/02/14/tree.jpg)
```
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
```
Example 2:
```
Input: root = [1,null,3]
Output: [1,3]
```
Example 3:
```
Input: root = []
Output: []
```

Constraints:

- The number of nodes in the tree is in the range [0, 100].
- -100 <= Node.val <= 100

---
This is another bad described / exampled problem.  
Actually, the problem want us to find all the last node (right-est) of each level.
![](https://assets.leetcode.com/users/images/cef92daf-88dd-46b5-a329-b179916c6482_1618278364.1240458.png)
In this example, we need to return [1, 3, 6, 4].