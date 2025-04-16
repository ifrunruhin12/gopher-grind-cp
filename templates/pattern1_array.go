package main

func main() {
	defer out.Flush()
	n := readInt()
	arr := readInts()
	// ✅ n is number of elements, arr holds the values
	fmt.Fprintln(out, "n =", n, "arr =", arr)
}