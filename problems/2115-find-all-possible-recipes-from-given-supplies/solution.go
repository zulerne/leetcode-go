// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/description/
package findallpossiblerecipesfromgivensupplies

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	res := make([]string, 0)
	inDegree := make(map[string]int, len(recipes))
	g := make(map[string][]string, len(supplies)+len(recipes))

	for i, r := range recipes {
		inDegree[r] = len(ingredients[i])
		for _, ing := range ingredients[i] {
			g[ing] = append(g[ing], r)
		}
	}

	queue := make([]string, len(supplies))
	copy(queue, supplies)
	for len(queue) > 0 {
		sup := queue[0]
		queue = queue[1:]

		for _, v := range g[sup] {
			inDegree[v]--

			if inDegree[v] == 0 {
				queue = append(queue, v)
				res = append(res, v)
			}
		}
	}

	return res
}
