package neural

import (
	"github.com/johnsudaar/scalingo_autoscaling/config"
	"gopkg.in/errgo.v1"
)

func GenerateTestWithPredicate(size int, tests int) (bool, bool, []float64) {
	test := rng.Int() % tests

	var values []float64
	var predicateUp bool
	var predicateDown bool
	switch test {
	case 0:
		values = GenerateSimpleRampUp(size)
		if values[size-2] > config.NEURAL_HIGH {
			predicateUp = true
		} else {
			predicateUp = false
		}

		if values[size-2] < config.NEURAL_LOW {
			predicateDown = true
		} else {
			predicateDown = false
		}
	case 1:
		values = GenerateSimpleRampDown(size)
		if values[size-2] > config.NEURAL_HIGH {
			predicateUp = true
		} else {
			predicateUp = false
		}

		if values[size-2] < config.NEURAL_LOW {
			predicateDown = true
		} else {
			predicateDown = false
		}
	case 2:
		values = GenerateFlat(size)
		predicateDown = false
		predicateUp = false
	case 3:
		values = GenerateHighActivity(size)
		predicateUp = true
		predicateDown = false
	case 4:
		values = GenerateLowActity(size)
		predicateUp = false
		predicateDown = true
	case 5:
		values = GenerateHighPeak(size)
		predicateDown = false
		predicateUp = false
	case 6:
		values = GenerateLowPeak(size)
		predicateDown = false
		predicateUp = false
	}

	return predicateUp, predicateDown, values
}

func GenerateTest(size int, up *Neuron, down *Neuron) (bool, bool, error) {

	predicateUp, predicateDown, values := GenerateTestWithPredicate(size, 5)
	resUp, errUp := up.Learn(values, predicateUp)
	resDown, errDown := down.Learn(values, predicateDown)

	if errUp != nil {
		return false, false, errgo.Mask(errUp)
	}

	if errDown != nil {
		return false, false, errgo.Mask(errDown)
	}

	// log.Println("VALUES : ")
	// log.Println(values)
	// log.Println("Pred UP : ")
	// log.Println(predicateUp)
	// log.Println("Pred Down : ")
	// log.Println(predicateDown)
	return resUp, resDown, nil
}
