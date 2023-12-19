package domain

type Convertion interface {
	GetCoefficient(string, string) (float64, error)
}
