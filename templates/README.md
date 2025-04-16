# 🧩 Golang Competitive Programming Templates

This folder contains modular Go templates for handling various input formats frequently used in competitive programming platforms like Codeforces, AtCoder, LeetCode, and CSES.

All templates use `bufio` for fast I/O and are designed for quick adaptation during contests.

---

## 🛠️ File Guide

### `base.go`
- Core file with:
  - Fast input/output setup (`bufio`)
  - Helper functions:
    - `readLine()`
    - `readInt()`
    - `readInts()`
    - `readStrings()`
- All other templates import and use these helpers.

---

### `pattern1_array.go`
- 📘 Pattern:  
```
x
a1 a2 a3 ... ax
```
- 🧠 Usage: When an integer is followed by that many integers.
- ✅ Example:
```
5 
3 5 7 2 11
```

---

### `pattern2_testcases.go`
- 📘 Pattern:
```
t 
n1 
a1 a2 a3 ... an1 
n2 
b1 b2 b3 ... bn2 ...
```
- 🧠 Usage: For multiple test cases, where each test has its own size and array.
- ✅ Example:
```
3 
2 
10 20 
3 
5 6 7 
1 
42
```

---

### `pattern3_lines.go`
- 📘 Pattern:
```
n 
line1 line2 ... linen
```
- 🧠 Usage: For reading multiple lines of strings (like names, sentences).
- ✅ Example:
```
4 
apple
orange 
grape 
banana
```

---

### `pattern4_words.go`
- 📘 Pattern:
```
n 
word1 word2 word3 ... wordn
```
- 🧠 Usage: For space-separated strings (often in one line).
- ✅ Example:
```
5 
apple orange grape watermelon banana
```

---

## 🧪 How to Use

1. Copy any pattern file you need.
2. Combine it with `base.go` in the same folder or merge into a single `main.go`.
3. Modify the logic section as per the problem.
4. 🚀 Run it in your local dev setup or submit!

---

## 🧡 Pro Tip

Keep a shortcut/snippet to these templates in your editor (like Neovim snippets, or VSCode `user-snippets`) for even faster access during contests.

---

Happy Hacking! 🚀💻