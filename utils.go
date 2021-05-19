package gobudins

func FilterInvestments(investments []Investment, test func(Investment) bool) []Investment {
	filtered := []Investment{}
	for _, investment := range investments {
		if test(investment) {
			filtered = append(filtered, investment)
		}
	}

	return filtered
}
