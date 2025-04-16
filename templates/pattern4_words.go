package main

func main() {
	defer out.Flush()
	n := readInt()       // Total number of words (optional, use if needed)
	words := readStrings()
	// ✅ All words read in one line
	fmt.Fprintln(out, "Words:", words)
}
