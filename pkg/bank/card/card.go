package card

import "github.com/DamirYoqubov/types/pkg/bank/types"

// Issue создает экземпляр карты с предопределёнными полями
func Issue(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}
}

// Withdraw снимает деньги с карты
func Withdraw(card *types.Card, amount types.Money)  {
	const withdrawLimit = 20_000_00
	if amount < 0 {
		return
	}
	if amount > withdrawLimit {
		return
	} 
	if !card.Active {
		return
	}
	if card.Balance < 0 {
		return
	}

	card.Balance -= amount
}