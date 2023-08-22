# 735. Asteroid Collision (bad problem!)
[LeetCode](https://leetcode.com/problems/asteroid-collision/)

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.



Example 1:
```
Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
```
Example 2:
```
Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.
```
Example 3:
```
Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.
```

Constraints:

- 2 <= asteroids.length <= 104
- -1000 <= asteroids[i] <= 1000
- asteroids[i] != 0

---
This is a poorly described problem. I'm including another discussion here which will provide more info to solve this problem correctly:
[https://leetcode.com/problems/asteroid-collision/description/comments/1842348](https://leetcode.com/problems/asteroid-collision/description/comments/1842348)
```
In addition to the question description.
Assume the asteroids index is there position in space and negative signed asteroids are moving towards 0th index and positive asteroids are moving towards lasts index.

For example:
1. [1,-2,10,-5]
here, 1 and -2 will collide and 10 and -5 will collide but -2 and 10 wont collide.
result -> [-2, 10]
2. [-4, 6]
here, -4 will move left and 6 to its right and will never collide
result -> [-4, 6]
but in
3. [4, -6] 4 and -6 will collide and stronger asteroids will remain, as 6 is stronger than 4 , so -6 will remain resulting [-6].