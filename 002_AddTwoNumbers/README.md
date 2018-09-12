# Add Two Numbers
[LeetCode](https://leetcode.com/problems/add-two-numbers/description/)

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

### Solution
- My first solution is traversing through each linked list and convert the list values to actual number, them adds two numbers and convert it back to linked list.  
But this solution will failed test case with very large number (exceed math.MaxUint64). 
- Another proper solution is traverse simultaneously on both linked list until reached the end of both lists. At each node, add two values then push result to output list. If sum of two value exceed 10, then stores the dozens value and add up to the next node.
