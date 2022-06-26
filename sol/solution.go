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
