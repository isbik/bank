package util

const (
	USD = "USD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR:
		return true
	default:
		return false
	}
}
