package progression

import (
	"errors"
	"math/big"
)

const (
	minProgressionSize = 3
	precession = 15

	IsProgression    = "Your sequence of numbers is definitely a progression"
	IsNotProgression = "Your sequence of numbers isn`t a progression"
)

type (
	floatSlice []*big.Float
	restriction func(floatSlice) error
)

func newSlice(slice []*big.Float, restrictions ...restriction) (floatSlice, error) {
	res := floatSlice(slice)

	var err error
	for _, r := range restrictions {
		// we want OR restrictions
		if err = r(res); err == nil {
			return res, nil
		}
	}

	return nil, errors.New("restrictions failed")
}

func isArithmeticProgression(arr floatSlice) error {
	if len(arr) < minProgressionSize {
		return errors.New("is not arithmetic progression")
	}

	step := big.NewFloat(0)

	prevNegative := big.NewFloat(0).Neg(arr[0])
	step.Add(arr[1], prevNegative)

	// array bounds check optimization
	arr = arr[:len(arr)]
	for i := 2; i < len(arr); i++ {
		prevNegative := big.NewFloat(0).Neg(arr[i-1])
		newStep := big.NewFloat(0).Add(arr[i], prevNegative)

		if !isEqual(step, newStep) {
			return errors.New("is not arithmetic progression")
		}
	}

	return nil
}

func isEqual(x, y *big.Float) bool {
	return x.Text('g', precession) == y.Text('g', precession)
}

func isGeometricProgression(arr floatSlice) error {
	if len(arr) < minProgressionSize {
		return errors.New("is not geometric progression")
	}

	scaleFactor := big.NewFloat(0).Quo(arr[1], arr[0])
	newScaleFactor := big.NewFloat(0)

	// array bounds check optimization
	arr = arr[:len(arr)]
	for i := 2; i < len(arr); i++ {
		newScaleFactor.Quo(arr[i], arr[i-1])

		if !isEqual(scaleFactor, newScaleFactor) {
			return errors.New("is not geometric progression")
		}
	}

	return nil
}

func IsArgsProgression() bool {
	arr, err := getArray()
	if err != nil {
		return false
	}

	_, err = newSlice(arr, isArithmeticProgression, isGeometricProgression)
	if err != nil {
		return false
	}

	return true
}
