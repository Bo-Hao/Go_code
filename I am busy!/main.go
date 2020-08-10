package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	for {
		log.Printf("Process Initial!")
		for percent := 1; percent < 101; percent++ {
			subprocess := rand.Intn(500)
			for p := 1; p < subprocess; p++ {
				if rand.Float64() < 0.01 {
					log.Printf("Error detected!")
					err_time := rand.Intn(3)

					for s := 0; s < err_time; s++ {
						time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
						log.Printf("Debuging...")
						time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
						log.Printf("Recomplie the process...")
					}
					log.Printf("Error solved!")
					time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
					log.Printf("Return Parsing...")
					time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

				}
				time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

				fmt.Println("Process:", percent, "%(", p, "/", subprocess, "); Loss: ", rand.Intn(20), "%; Parsing...")
			}
		}
		log.Printf("Process Over!")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		log.Printf("Testing...")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		log.Printf("Saving...")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		log.Printf("Updating...")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		log.Printf("Reset Process")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	}

}
