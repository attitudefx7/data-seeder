package contract

type ISeed interface {
	Seed(int, interface{}) interface{}
}
