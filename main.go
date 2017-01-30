package main

import (
	"time"

	"github.com/johnsudaar/scalingo_autoscaling/runner"
	"github.com/johnsudaar/scalingo_autoscaling/threshold"
)

func main() {
	scaler := threshold.NewThresholdScaler(50, 80, 50, 80, 1*time.Minute)
	//scaler := neural.NewTrainerNeuralScaler(2 * time.Minute)
	runner := runner.NewRunner(scaler, "ensiie-test-1", 10)

	runner.Start()

	// network := neuralnetwork.NewNeuralNetworkScaler(10, 10)
	// network.Learn()
}
