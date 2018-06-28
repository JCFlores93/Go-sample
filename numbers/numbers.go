package numbers

import "errors"

//return three numbers int
func GetVariables() (int, int32, int64) {
	//2147000000 = max. value for int32
	return 1, 2147000000, 211231231231231233
}

//return sum
func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}

//return 2 floats
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}
