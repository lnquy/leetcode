Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.



Example 1:
![](https://assets.leetcode.com/uploads/2021/03/10/oddeven-linked-list.jpg)
```
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
```
Example 2:
![](https://assets.leetcode.com/uploads/2021/03/10/oddeven2-linked-list.jpg)
```
Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
```

Constraints:

- The number of nodes in the linked list is in the range [0, 10^4].
- -10^6 <= Node.val <= 10^6



--- 
```
# o = head, o = head.next
1(o) -> 2(e) -> 3 -> 4 -> 5 

# Loop#1.1: o.next = e.next
1(o) -> 3 -> 4 -> 5
        ^    
        |
       2(e)
    
# Loop#1.2: o = o.next
1 -> 3(o) -> 4 -> 5
     ^    
     |
     2(e) 

# Loop#1.3: e.next = o.next
1 -> 3(o) -> 4 -> 5
             ^    
             |
             2(e)

# Loop#1.4: e = e.next
1 -> 3(o) -> 4(e) -> 5
             ^    
             |
             2(headEvenNode)


# Loop#2.1: o.next = e.next
1 -> 3(o) -> 5
             ^
             |
             4(e)
             ^    
             |
             2(headEvenNode)


# Loop#2.2: o = o.next
1 -> 3 -> 5(o)
          ^
          |
          4(e)
          ^    
          |
          2(headEvenNode)

# Loop#2.3:e.next = o.next
1 -> 3 -> 5(o)
          
          4(e) -> nil
          ^    
          |
          2(headEvenNode)

# Loop#2.4:e = e.next
1 -> 3 -> 5(o)
          
          4 -> nil(e)
          ^    
          |
          2(headEvenNode)

# o.next = headEvenNode
1 -> 3 -> 5(o) -> 2(headEvenNode) -> 4 -> nil(e)
          
   
```