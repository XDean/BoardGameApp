package guobiao

func CalcFan(hand Hand) Fans {
	max := Fans{}
	for _, h := range FindHu(hand) {
		fan := CalcGroupFan(h)
		if fan.Sum() > max.Sum() {
			max = fan
		}
	}
	return max
}
