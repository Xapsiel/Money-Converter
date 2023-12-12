package domain

type Convertion interface {
	GetCoefficient(string, string) float64
	GetAllCoefficient(string, string) map[string]float64
	ReadData() []byte
	UpdateDB(string)
}
