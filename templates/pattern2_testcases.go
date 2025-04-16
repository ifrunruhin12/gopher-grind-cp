package main

func main() {
	defer out.Flush()
	t := readInt()
	for i := 0; i < t; i++ {
		k := readInt()
		arr := readInts()
		// ✅ Each test case has k followed by k integers
		fmt.Fprintf(out, "Test #%d: k = %d, arr = %v\n", i+1, k, arr)
	}
}
