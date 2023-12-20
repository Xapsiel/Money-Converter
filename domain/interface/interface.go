package domain

type Convertion interface {
	GetCoefficient(string, string) (float64, error)
	GetAllCoefficient(string, string) (map[string]float64, error)
	ReadData() ([]byte, error)
	UpdateDB(string) error
}
