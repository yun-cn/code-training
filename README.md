#### Leetcode 1 (Easy)
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

#### 解题思路:
 主要有2两种解题思路, 
 第一种: 较为暴力的解决方法, 数组进行两两对比,然后最终匹配成功.
 第二种: 对数据进行循环, 再循环中,通过减法找到找到匹配的数字.然后返回数组.