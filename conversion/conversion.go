package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var result []float64

	for _, str := range strings {
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float64")
		}
		result = append(result, floatVal)
	}

	return result, nil
}
