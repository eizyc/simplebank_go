package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	RUB = "RUB"
	PLN = "PLN"
	JPY = "JPY"
	INR = "INR"
	GBP = "GBP"
	CHF = "CHF"
	SEK = "SEK"
	DKK = "DKK"
	CNY = "CNY"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, RUB, PLN, JPY, INR, GBP, CHF, SEK, DKK, CNY:
		return true
	}
	return false
}
