package generated_models

type Student struct {
	Classes []string
	Grade   float64
}

type Person struct {
	Name      string
	Age       float64
	Addresses []Address
}

type Address struct {
	City    string
	Country string
}
