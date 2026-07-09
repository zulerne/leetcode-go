// https://leetcode.com/problems/minimum-genetic-mutation/description/
package minimumgeneticmutation

func minMutation(startGene string, endGene string, bank []string) int {
	var res int
	mutations := []byte{'A', 'C', 'G', 'T'}

	bankMap := make(map[string]bool, len(bank))
	for _, gene := range bank {
		bankMap[gene] = true
	}

	seen := make(map[string]bool)
	seen[startGene] = true

	queue := []string{startGene}
	geneBytes := make([]byte, 8)
	for len(queue) > 0 {
		lvlSize := len(queue)

		for range lvlSize {
			gene := queue[0]
			queue = queue[1:]

			if gene == endGene {
				return res
			}

			copy(geneBytes, gene)

			for i := range geneBytes {
				for _, b := range mutations {
					if gene[i] == b {
						continue
					}

					geneBytes[i] = b
					newGene := string(geneBytes)

					if !seen[newGene] && bankMap[newGene] {
						seen[newGene] = true
						queue = append(queue, newGene)
					}
				}

				geneBytes[i] = gene[i]
			}
		}

		res++
	}

	return -1
}
