package guobiao

import "fmt"

func CalcFan(hand Hand) Fans {
	max := Fans{}
	for _, h := range FindHu(hand) {
		fmt.Println(h)
		fan := CalcGroupFan(h)
		if fan.Fan() > max.Fan() {
			max = fan
		}
	}
	return max
}
