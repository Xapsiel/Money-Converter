package domain

type Convertion interface {
	GetCoefficient(string, string) (float64, error)
	GetCoefficientByDate(string, string, int, int, int) (float64, error)

	ReadData() ([]byte, error)
	UpdateDB(string) error
}
