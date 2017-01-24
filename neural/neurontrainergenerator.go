package neural

import (
	"math/rand"
	"time"

	"github.com/johnsudaar/scalingo_autoscaling/config"
)

var (
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateSimpleRampUp(size int) []float64 {
	values := make([]float64, size)

	value := rng.Float64()

	for i := 0; i < size; i++ {
		values[i] = value

		value = value + ((1.0 - value) * rng.Float64())
	}

	return values
}

func GenerateSimpleRampDown(size int) []float64 {
	values := make([]float64, size)

	value := rng.Float64()

	for i := 0; i < size; i++ {
		values[i] = value

		value = value - (value * rng.Float64())
	}

	return values
}

func GenerateFlat(size int) []float64 {
	values := make([]float64, size)

	value := rng.Float64()*(config.NEURAL_HIGH-config.NEURAL_LOW) + config.NEURAL_LOW
	for i := 0; i < size; i++ {
		values[i] = value

		incr := (rng.Float64() - 0.5) / 100.0

		if value+incr < config.NEURAL_HIGH && value+incr > config.NEURAL_HIGH {
			value = value + incr
		}
	}

	return values
}

func GenerateHighPeak(size int) []float64 {
	values := GenerateFlat(size)

	index := rng.Int() % size

	values[index] = config.NEURAL_HIGH + (rng.Float64() * (1.0 - config.NEURAL_HIGH))

	return values
}

func GenerateLowPeak(size int) []float64 {
	values := GenerateFlat(size)
	index := rng.Int() % size
	values[index] = rng.Float64() * config.NEURAL_LOW

	return values
}

func GenerateLowActity(size int) []float64 {
	values := GenerateFlat(size)

	index := rng.Int()%(size-3) + 2

	for i := index; i < size; i++ {
		values[i] = rng.Float64() * config.NEURAL_LOW
	}

	return values
}

func GenerateHighActivity(size int) []float64 {
	values := GenerateFlat(size)

	index := rng.Int()%(size-3) + 2

	for i := index; i < size; i++ {
		values[i] = config.NEURAL_HIGH + (rng.Float64() * (1.0 - config.NEURAL_HIGH))
	}

	return values
}
