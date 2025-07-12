**Manacher’s Algorithm** is a clever linear-time algorithm (**O(n)**) to find the **longest palindromic substring** in a given string.

---

## 🔍 Problem

> Given a string `s`, find the **longest palindromic substring** in **O(n)** time.

Traditional solutions (like expand-around-center or DP) are **O(n²)**.
**Manacher's Algorithm** optimizes it to **O(n)**.

---

## 🧠 Core Idea

It avoids redundant comparisons by **reusing previously computed palindrome lengths** using **mirrored properties**.

### Key Concepts:

* **Preprocess the string** to handle both even and odd-length palindromes uniformly.

  * Insert `#` between characters and at both ends.
  * Example: `"abba"` → `"^#a#b#b#a#$"`

    * `^` and `$` are sentinels to avoid bounds checking.

* Maintain:

  * `C`: center of the current rightmost palindrome
  * `R`: right boundary (index) of that palindrome
  * `P[i]`: length of the palindrome centered at index `i` (in the transformed string)

---

## ⚙️ Steps

1. **Transform the string**:
   Input: `"abba"`
   Transformed: `"^#a#b#b#a#$"` (length = 11)

2. **Iterate through the transformed string**:
   For each position `i` in the transformed string:

   * Find mirror position: `i_mirror = 2*C - i`
   * If `i < R`, set `P[i] = min(R - i, P[i_mirror])`
   * Try to expand palindrome centered at `i`
   * If expanded beyond `R`, update `C` and `R`

3. **After the loop**, find the index with the maximum `P[i]` and extract the corresponding substring in the original string.

---

## ✅ Example

Original: `"abba"`
Transformed: `"^#a#b#b#a#$"`
`P = [0, 0, 1, 0, 3, 0, 1, 0, 1, 0, 0]`

* Max `P[i] = 4` at `i = 5`
* Longest palindromic substring length: `4 - 1 = 3` (because of `#`)
* Extract: center at `i=5` → original substring is `"abba"`

---

## 🧑‍💻 Golang Code (Optional)

Would you like a full Golang implementation with comments?

---

## 📌 Summary

| Aspect           | Value                         |
| ---------------- | ----------------------------- |
| Time Complexity  | **O(n)**                      |
| Space Complexity | **O(n)**                      |
| Handles even/odd | Uniformly with `#`            |
| Use case         | Longest palindromic substring |

