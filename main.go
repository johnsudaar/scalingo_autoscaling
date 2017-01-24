package main

import "github.com/johnsudaar/scalingo_autoscaling/neural"

func main() {
	//scaler := threshold.NewThresholdScaler(50, 80, 50, 80, 1*time.Minute)
	//scaler := neural.NewTrainerNeuralScaler(2 * time.Minute)
	//runner := runner.NewRunner(scaler, "ensiie-test-1", 10)

	//runner.Start()

	neural.TestOR()
}
