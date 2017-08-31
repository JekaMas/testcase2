package progression

import (
	"errors"
	"math/big"
	"os"
	"strings"
)

func getArray() ([]*big.Float, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("No array given")
	}

	arrString := os.Args[1]

	//cleanup input
	arrString = strings.TrimSpace(arrString)
	arrString = strings.Trim(arrString, ",")

	parts := splitArray(arrString)
	for i := range parts {
		parts[i] = strings.Trim(parts[i], " ")

		parts[i] = strings.TrimLeft(parts[i], "0")

		if strings.Index(parts[i], ".") != -1 {
			parts[i] = strings.TrimRight(parts[i], "0")
		}
	}

	return castSlice(parts)
}

func splitArray(s string) []string {
	return strings.Split(s, ",")
}

func castSlice(parts []string) ([]*big.Float, error) {
	arr := make([]*big.Float, len(parts))

	var err error

	for i, s := range parts {
		bigNum := big.NewFloat(0)
		err = bigNum.UnmarshalText([]byte(s))
		if err != nil {
			return nil, errors.New("Not number given")
		}
		arr[i] = bigNum
	}

	return arr, nil
}
