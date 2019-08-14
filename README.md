# gpio
//linux gpio (sysfs). Forked github.com/lab409/go-artik/gpio
//Cannot fork normally, author's repo was broken
//Thanks for author (lab409)
//replaced log.Fatal() to log.Println(), for case the sysfs files has been created before. 
//add NewPinNoInit() for the same case

package main

import "fmt"
import "time"
import "github.com/adenis78/gpio"

func main() {
	fmt.Println("\nExample. Blink gpio43")
	fmt.Println("For exit press Ctrl+C\n")

	/* Create new pin */
	pin0 := gpio.NewPin(43, gpio.OUT)
	pin1 := gpio.NewPin(45, gpio.OUT)

	/* Just to be on the safe side - pull-up the pin */
	pin0.Set()
	pin1.Set()

	/* Infinite loop for blinking */
	for {
		/* Blink mat completed by Set/Clear logic ... */
		pin1.Clear()
		pin0.Set()
		time.Sleep(200 * time.Millisecond)
		pin1.Set()
		time.Sleep(200 * time.Millisecond)

		pin1.Set()
		pin0.Set()
		time.Sleep(200 * time.Millisecond)
		pin0.Clear()
		time.Sleep(200 * time.Millisecond)

		// /* Or by Toggle logic */
		// pin.Toggle()
		// time.Sleep(200 * time.Millisecond)
	}
}

