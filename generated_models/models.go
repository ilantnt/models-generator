type Person struct {
	Name string `json:"name"`
	Age float64 `json:"age"`
	Addresses []Address `json:"addresses"`
type Address struct {
	City string `json:"city"`
	Country string `json:"country"`
type Student struct {
	Classes []string `json:"classes"`
	Grade float64 `json:"grade"`
