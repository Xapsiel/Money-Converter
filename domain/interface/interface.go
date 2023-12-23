package domain

type Convertion interface {
	GetCoefficient(string, string, int, int, int) (float64, error)

	ReadData() ([]byte, error)
	UpdateDB(string) error
}
