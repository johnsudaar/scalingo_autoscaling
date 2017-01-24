package neural

import "log"

func TestOR() {
	n := NewNeuron(2)

	tests := [][]float64{
		{0.0, 0.0},
		{0.0, 1.0},
		{1.0, 0.0},
		{1.1, 1.1},
	}

	predicates := []bool{false, true, true, true}

	cont := true

	for cont {
		cont = false
		log.Println("New iteration")
		log.Println(n.Weights)
		for i := 0; i < 4; i++ {
			res, err := n.Learn(tests[i], predicates[i])
			if err != nil {
				log.Fatal(err)
			}

			if !res {
				cont = true
			}
		}
	}

	log.Println("Done")
	log.Println(n.Weights)
}
