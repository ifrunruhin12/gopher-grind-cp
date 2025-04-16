package main

func main() {
	defer out.Flush()
	n := readInt()
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = readLine()
	}
	// ✅ You now have n strings from n lines
	fmt.Fprintln(out, "Strings:", strs)
}
