package main

import (
	"clicker"
	"fmt"
	"userutills"
)

func main() {
	clicker.Init()

	for true {

		//logic to close
		ans := userutills.GetStringInput("Buy clickers: b\nContinue (y/n): ")
		if ans == "n" {
			break
		} else if ans == "b" {
			clicker.BuyClicker()
		} else {
			fmt.Println("money: ", clicker.GetMoney())
			fmt.Println("rate: ", clicker.GetRate())
			clicker.SetMoney()
			continue
		}
	}
}
