# golang_single_number

Given a **non-empty** array of integers `nums`, every element appears *twice* except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

## Examples

**Example 1:**

```
Input: nums = [2,2,1]
Output: 1

```

**Example 2:**

```
Input: nums = [4,1,2,1,2]
Output: 4

```

**Example 3:**

```
Input: nums = [1]
Output: 1

```

**Constraints:**

- `1 <= nums.length <= 3 * 104`
- `3 * 104 <= nums[i] <= 3 * 104`
- Each element in the array appears twice except for one element which appears only once.

## 解析

給定一個整數陣列 nums ,

這個陣列只有一個元素只出現一次， 其他都成對出現

要求寫一個演算法來找出唯一出現的值

已知 XOR有以下特性

a XOR a = 0

0 XOR a = a

a XOR a XOR b = b XOR a XOR a

會發現當同一個數字在連續 XOR 的數字出現兩次時就會回到原本數值

所以只要對陣列每個數值 XOR 起來

最後的值就是只出現一次的值

時間複雜度為 O(n)

空間複雜度為 O(1)

## 程式碼
```go
package sol

func singleNumber(nums []int) int {
	result := nums[0]
	lenNums := len(nums)
	for pos := 1; pos < lenNums; pos++ {
		result ^= nums[pos]
	}
	return result
}

```
## 困難點

1. 需要理解 XOR 特性

## Solve Point

- [x]  初始化 result = nums[0]
- [x]  從 pos = 1.. len(nums)-1 , result ^= nums[pos]
- [x]  回傳 result