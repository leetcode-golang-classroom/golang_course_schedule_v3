package sol

import (
	"container/heap"
	"sort"
)

type CourseMaxHeap []int

func (h *CourseMaxHeap) Len() int {
	return len(*h)
}
func (h *CourseMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *CourseMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *CourseMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *CourseMaxHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

type ByLastDay [][]int

func (a ByLastDay) Len() int           { return len(a) }
func (a ByLastDay) Less(i, j int) bool { return a[i][1] < a[j][1] }
func (a ByLastDay) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func scheduleCourse(courses [][]int) int {

	sort.Sort(ByLastDay(courses))

	time := 0
	// priorityQueue
	priorityQueue := CourseMaxHeap{}
	heap.Init(&priorityQueue)
	for _, course := range courses {
		if time+course[0] <= course[1] {
			heap.Push(&priorityQueue, course[0])
			time += course[0]
		} else if priorityQueue.Len() != 0 && priorityQueue[0] > course[0] {
			duration := heap.Pop(&priorityQueue).(int)
			time += course[0] - duration
			heap.Push(&priorityQueue, course[0])
		}
	}
	return priorityQueue.Len()
}
