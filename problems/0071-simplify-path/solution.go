// https://leetcode.com/problems/simplify-path/description/
package simplifypath

import (
	"strings"
)

func simplifyPath(path string) string {
	result := []string{}

	for _, v := range strings.Split(path, "/") {
		switch v {
		case "", ".":
			continue
		case "..":
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		default:
			result = append(result, v)
		}
	}

	return "/" + strings.Join(result, "/")
}
