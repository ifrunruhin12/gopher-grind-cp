//https://www.codechef.com/problems/KBOXES

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monster struct {
	level, coin, index int
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, _ := in.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInts() []int {
	line := readLine()
	parts := strings.Fields(line)
	nums := make([]int, len(parts))
	for i, s := range parts {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func solveOne(n, k int, levels, coins []int) []int {
	monsters := make([]Monster, n)
	for i := 0; i < n; i++ {
		monsters[i] = Monster{
			level: levels[i],
			coin:  coins[i],
			index: i,
		}
	}

	sort.Slice(monsters, func(i, j int) bool {
		return monsters[i].level < monsters[j].level
	})

	ans := make([]int, n)
	h := &MinHeap{}
	heap.Init(h)
	sum := 0

	for _, m := range monsters {
		ans[m.index] = sum
		heap.Push(h, m.coin)
		sum += m.coin
		if h.Len() > k {
			removed := heap.Pop(h).(int)
			sum -= removed
		}
	}

	return ans
}

func main() {
	defer out.Flush()
	t := 0
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		params := readInts()
		n, k := params[0], params[1]
		levels := readInts()
		coins := readInts()
		result := solveOne(n, k, levels, coins)
		for _, val := range result {
			fmt.Fprintf(out, "%d ", val)
		}
		fmt.Fprintln(out)
	}
}

