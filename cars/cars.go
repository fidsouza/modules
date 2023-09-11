package cars

type Car struct {
	Brand string
}

func (c Car) Run() string {
	return "Car is moving"
}
