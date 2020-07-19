package data

type DataService struct {
	currencies map[string]float64
}

func New() *DataService {
	ds := &DataService{
		currencies: make(map[string]float64),
	}

	return ds
}
