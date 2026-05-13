// https://leetcode.com/problems/minimum-genetic-mutation/description/
package minimumgeneticmutation

func minMutation(startGene string, endGene string, bank []string) int {
	var res int

	seen := make(map[string]bool)
	seen[startGene] = true

	bankMap := make(map[string]bool)
	for _, g := range bank {
		bankMap[g] = true
	}

	gens := []byte{'A', 'C', 'G', 'T'}
	neighs := make([]string, 24)
	updateNeighs := func(cur string) {
		k := 0
		b := []byte(cur)

		for i := range b {
			for _, gen := range gens {
				curGen := cur[i]
				if curGen == gen {
					continue
				}
				b[i] = gen
				neighs[k] = string(b)
				k++
				b[i] = cur[i]
			}
		}
	}

	queue := []string{startGene}
	for len(queue) > 0 {
		levelSize := len(queue)
		for range levelSize {
			cur := queue[0]
			queue = queue[1:]

			if cur == endGene {
				return res
			}

			updateNeighs(cur)
			for _, n := range neighs {
				if !seen[n] && bankMap[n] {
					seen[n] = true
					queue = append(queue, n)
				}
			}
		}
		res++
	}

	return -1
}
