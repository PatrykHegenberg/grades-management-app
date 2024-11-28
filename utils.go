package main

func setNote(prozent float64) float64 {
	switch {
	case prozent <= 22:
		return 6.00
	case prozent <= 49:
		return 5.00
	case prozent <= 64:
		return 4.00
	case prozent <= 79:
		return 3.00
	case prozent <= 94:
		return 2.00
	default:
		return 1.00
	}
}

func checkGewichtung(lv, hv float64) bool {
	sum := hv/100 + lv/100
	return sum == 1
}
