// https://leetcode.com/problems/gas-station/description/
package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	total, tank, start := 0, 0, 0
	for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}

	return start
}
