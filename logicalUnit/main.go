package main

import "./services"

func main() {
	services.SetupLog()
	/*var iteration = 100
	f, err := os.OpenFile("tmp/orders.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(" Simulation API Called")
	start := time.Now()
	for i := 0; i < iteration; i++ {
		fmt.Println("PARALLEL")
		fmt.Println("I: ", i)
		start1 := time.Now()
		simulator.SimulateParallel()
		end1 := time.Now()
		elapsed1 := end1.Sub(start1)
		fmt.Println("TIME: ", elapsed1)
		fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("TIME: ", elapsed/100)
	start3 := time.Now()
	for i := 0; i < iteration; i++ {
		fmt.Println("NORMAL")
		fmt.Println("I: ", i)

		start1 := time.Now()
		simulator.SimulateParallel2()
		//simulator.Simulate()
		end1 := time.Now()
		elapsed1 := end1.Sub(start1)
		fmt.Println("TIME: ", elapsed1)
		fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
	}
	end3 := time.Now()
	elapsed3 := end3.Sub(start3)




	log.Println("PARALLEL TIME: ", elapsed/100," (ITERATION: ", iteration, "NUMBER OF FOOD SOURCES: ", globals.NumberOfFoodSources,
		" NUMBER OF EPOCHES: ", globals.FinalEpoch, ")")
	log.Println(" TIME: ", elapsed3/100	," (ITERATION: ", iteration, "NUMBER OF FOOD SOURCES: ", globals.NumberOfFoodSources,
		" NUMBER OF EPOCHES: ", globals.FinalEpoch, ")")
	fmt.Println("ERROR PARALLEL: ", globals.ErrorParallel)
	fmt.Println("ERROR PARALLEL2: ", globals.ErrorNormal)*/
	initRouter()
}
