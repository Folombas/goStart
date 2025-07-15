package main

import "project/wallet"

func main() {
	myWallet := &wallet.Wallet{Cash: 100}
	wallet.Buy(myWallet)
	// Мы создали кошелёк и можем через него что-то купить


var myMoney wallet.Payer 
myMoney = &wallet.Card{Balance: 100, Cardholder: "Gosha Golang"}
wallet.Buy(myMoney)

myMoney = &wallet.ApplePay{Money: 9}
wallet.Buy(myMoney)
}