// https://leetcode.com/problems/ransom-note/description/
package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	letters := [26]int{}

	for i := range magazine {
		letters[magazine[i]-'a']++
	}

	for i := range ransomNote {
		idx := ransomNote[i] - 'a'

		letters[idx]--
		if letters[idx] < 0 {
			return false
		}
	}

	return true
}
