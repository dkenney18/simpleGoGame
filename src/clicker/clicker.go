package clicker

import (
	"fmt"
	"userutills"
)

//main struct
type clicker struct {
	Name string
	Price int
	CPS int
	Level int
	Owned bool
	NumOwned int
	ID int
}

//Rate is the global clicks per clicks per sec
var Rate int = 0

//Money is the global money
var Money int = 100

//the list of clickers that you can use
var clickers = make(map[int]clicker)

//Init sets up clickers and game things
func Init() {
	clickers[0] = clicker{
		Name: "cookieClicker",
		Price: 10,
		CPS: 1,
		Level: 0,
		Owned: false,
		NumOwned: 0,
		ID: 0,
	}

	clickers[1] = clicker{
		Name: "epicClicker",
		Price: 60,
		CPS: 10,
		Level: 1,
		Owned: false,
		NumOwned: 0,
		ID: 1,
	}

	clickers[2] = clicker{
		Name: "spaceClicker",
		Price: 100,
		CPS: 100,
		Level: 2,
		Owned: false,
		NumOwned: 0,
		ID: 2,
	}

	clickers[3] = clicker{
		Name: "ultraClicker",
		Price: 1000,
		CPS: 1000,
		Level: 3,
		Owned: false,
		NumOwned: 0,
		ID: 3,
	}
}

//getters for struct

//GetName returns name for use in other functions
func GetName(c clicker) string {
	return c.Name
}

//GetPrice returns price for use in other functions
func GetPrice(c clicker) int {
	return c.Price
}

//GetCPS returns cps for use in other functions
func GetCPS(c clicker) int {
	return c.CPS
}

//GetLevel returns level for use in other functions
func GetLevel(c clicker) int {
	return c.Level
}

//GetOwned returns owned for use in other functions
func GetOwned(c clicker) bool {
	return c.Owned
}

//GetNumberOwned returns number of clickers owned for use in other functions
func GetNumberOwned(c clicker) int {
	return c.NumOwned
}

//GetID returns ID for use in other functions
func GetID(c clicker) int {
	return c.ID
}

//setters for struct

//SetName sets name of clicker
func SetName(c clicker, n string) {
	c.Name = n
}

//SetPrice sets price of clicker
func SetPrice(c clicker, p int) {
	c.Price = p
}

//SetCPS sets cps of clicker
func SetCPS(c clicker, cps int) {
	c.CPS = cps
}

//SetLevel sets level of clicker
func SetLevel(c clicker, l int) {
	c.Level = l
}

//SetOwned sets owned of clicker
func SetOwned(c clicker, o bool) {
	c.Owned = o
}

//SetNumberOwned sets number of clickers owned of clicker
func SetNumberOwned(c clicker, numown int) {
	c.NumOwned += numown
}

//SetID sets id of clicker
func SetID(c clicker, id int) {
	c.ID = id
}

//logic for clicker

//BuyClicker lets you buy clickers
func BuyClicker() {	
	buyMsg := "Enter ID number to buy the clicker: "

	displayClickersForSale()

	ans := userutills.GetIntInput(buyMsg)

	fmt.Println(ans)

	if (ans <= len(clickers) && Money >= clickers[ans].Price) {
		SetOwned(clickers[ans], true)
		SetNumberOwned(clickers[ans], 1)
		AddRate(clickers[ans].CPS)
		SubtractMoney(clickers[ans].Price)
		fmt.Println("Bought:", clickers[ans].Name, "for ", clickers[ans].Price)

	} else {
		fmt.Println("Number is too large.")
	}

}

//SetMoney adds rate to money
func SetMoney() {
	Money += Rate
}

//AddMoney adds int to money
func AddMoney(m int) {
	Money += m
}

//GetRate returns rate
func GetRate() int {
	return Rate
}

//GetMoney returns money
func GetMoney() int {
	return Money
}

//SubtractMoney subtracts int form money
func SubtractMoney(m int) {
	Money -= m
}

//AddRate adds CPS to rate
func AddRate(r int) {
	Rate += r
}



//assign unique id to clickers. Use this when you add a new clicker to maintain uniqueness amoung the clickers
func assignUniqueIDToClickers()  {
	for i := 0; i < len(clickers); i++ {
		SetID(clickers[i], i)
	}
}

//display clickers for sale
func displayClickersForSale()  {
	for i := 0; i < len(clickers); i++ {
		fmt.Println("Name:", clickers[i].Name, "Price:", clickers[i].Price, "ID:", clickers[i].ID)
	}
}
