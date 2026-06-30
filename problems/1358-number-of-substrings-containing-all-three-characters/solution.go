// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/
package numberofsubstringscontainingallthreecharacters

func numberOfSubstrings(s string) int {
	var res int
	counter := [3]int{}

	left := 0
	for right := 0; right < len(s); right++ {
		counter[s[right]-'a']++

		for counter[0] > 0 && counter[1] > 0 && counter[2] > 0 {
			res += len(s) - right
			counter[s[left]-'a']--
			left++
		}
	}

	return res
}
