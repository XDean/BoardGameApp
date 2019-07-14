package guobiao

import "fmt"

func CalcFan(hand Hand) Fans {
	max := Fans{}
	for _, h := range FindHu(hand) {
		fmt.Println(h)
		fan := CalcGroupFan(h)
		if fan.Sum() > max.Sum() {
			max = fan
		}
	}
	return max
}
