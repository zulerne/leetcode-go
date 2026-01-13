// Package stringcompression
// https://leetcode.com/problems/string-compression
package stringcompression

import "strconv"

func compressFirst(chars []byte) int {
	acc := chars[0]
	count := 1
	idx := 0

	compressChar := func() {
		chars[idx] = acc
		idx++
		if count > 1 {
			byteCount := []byte(strconv.Itoa(count))
			for i := 0; i < len(byteCount); i++ {
				chars[idx] = byteCount[i]
				idx++
			}
		}
	}

	for i := 1; i < len(chars); i++ {
		b := chars[i]

		if b == acc {
			count++
			continue
		}

		compressChar()

		acc = b
		count = 1
	}
	compressChar()

	return idx
}

func compress(chars []byte) int {
	w := 0
	r := 0

	for r < len(chars) {
		cur := chars[r]
		count := 0

		for r < len(chars) && chars[r] == cur {
			count++
			r++
		}

		chars[w] = cur
		w++

		if count > 1 {
			byteCount := []byte(strconv.Itoa(count))
			for i := 0; i < len(byteCount); i++ {
				chars[w] = byteCount[i]
				w++
			}
		}
	}

	return w
}
