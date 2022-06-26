# golang_course_schedule_v3

There are `n` different online courses numbered from `1` to `n`. You are given an array `courses` where `courses[i] = [durationi, lastDayi]` indicate that the `ith` course should be taken **continuously** for `durationi` days and must be finished before or on `lastDayi`.

You will start on the `1st` day and you cannot take two or more courses simultaneously.

Return *the maximum number of courses that you can take*.

## Examples

**Example 1:**

```
Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
Output: 3
Explanation:
There are totally 4 courses, but you can take 3 courses at most:
First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day.
Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day.
The 4th course cannot be taken now, since you will finish it on the 3300th day, which exceeds the closed date.

```

**Example 2:**

```
Input: courses = [[1,2]]
Output: 1

```

**Example 3:**

```
Input: courses = [[3,2],[4,3]]
Output: 0

```

**Constraints:**

- `1 <= courses.length <= 104`
- `1 <= durationi, lastDayi <= 104`

## 解析

題目給定一個整數矩陣 courses , 每個 entry courses[i] = $[duration_i, lastday_i]$ 代表該課程所會花費的時間 還有最少需要在哪一天上完

要求寫一個演算法來計算在給定的 courses 條見下 最多可以修多少課

假設給定有兩堂課 個別花的時間與最後時能上的時間如下

(a, x), (b,y)

假設 x < y

考慮可能上的情況 也就是 a < x, b < y 因為如果 a> x 或 b > y 就會是 花的時間比起始時間長 無法上課

以下有幾種狀況

 **a + b ≤  x**

![](https://i.imgur.com/RTqmHrM.png)



在這個情況下， 兩個 course 都可以上

**a+ b > x,  a > b , a + b < y**

![](https://i.imgur.com/1ilFf5l.png)


在這個情況下， 如果先上b 就會造成 a 超過 lastDay

所以必須先上消耗時間大的課程， 才能達到最大化

**a + b > x, a < b,  a + b < y**

![](https://i.imgur.com/m1nZgMM.png)

在這個情況下, 必須要讓 lastDay 愈小愈先上, 才能達到最大化

因為 b > a, 如果  b 先上 會讓 a 超過 lastDay

**a + b > y**

![](https://i.imgur.com/fY5Thdf.png)

在這種情況下，不論使用哪一種順序最多都只能上種課

從以上的策略可以發現

要達到最大化選課的策略

首先是先照最後期限遇快到的先選

然後再從消耗時間最小的先選

具體作法就是先把傳入的  courses 依照 lastDay 由小到大作排序

然後在依序檢查當下累計時間 + 當下課程 duration 有沒有超過 課程 LastDay

如果超過則從 priorityQueue 拿出目前 duration 消耗多的出來做替換

如果沒超過 則把當下課程的duration 加入累計時間，並且放入 PriorityQueue

這樣每次找最大的時間是 O(logn)

然後要找 n 次

所以時間複雜度是 O(nlogn)

## 程式碼
```go
package sol

import (
	"container/heap"
	"sort"
)

type Course struct {
	Duration, LastDay int
}
type ByLastDay []Course

func (a ByLastDay) Len() int           { return len(a) }
func (a ByLastDay) Less(i, j int) bool { return a[i].LastDay < a[j].LastDay }
func (a ByLastDay) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type DurationMaxHeap []int

func (h *DurationMaxHeap) Len() int           { return len(*h) }
func (h *DurationMaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *DurationMaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *DurationMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *DurationMaxHeap) Push(value interface{}) {
	*h = append(*h, value.(int))
}
func scheduleCourse(courses [][]int) int {
	sortedCourses := []Course{}
	for _, course := range courses {
		sortedCourses = append(sortedCourses, Course{Duration: course[0], LastDay: course[1]})
	}
	sort.Sort(ByLastDay(sortedCourses))
	priorityQueue := DurationMaxHeap{}
	heap.Init(&priorityQueue)
	time := 0
	for _, course := range sortedCourses {
		if time+course.Duration <= course.LastDay {
			time += course.Duration
			heap.Push(&priorityQueue, course.Duration)
		} else if priorityQueue.Len() > 0 && priorityQueue[0] > course.Duration {
			duration := heap.Pop(&priorityQueue).(int)
			time += course.Duration - duration
			heap.Push(&priorityQueue, course.Duration)
		}
	}
	return priorityQueue.Len()
}
```
## 困難點

1. 需要找出一個可以最佳花上最多課的方法

## Solve Point

- [x]  先透過 sort 將輸入 courses 做排序
- [x]  把當下符合的 courses 放入 priorityQueue
- [x]  遇到不符合的 course 找出當下最大的 courses 做替換