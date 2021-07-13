package main

import "fmt"//short form for format

// Go only uses  structs rather than using classes

type car struct {
	gas_pedal uint16 // min 0 max 65535, means unsigned
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}

func main() {
    a_car := car{gas_pedal: 22341, 
				brake_pedal: 0, 
				steering_wheel: 12343, 
				top_speed_kmh: 225.9} // longer way
	// a_car := car{22341, 0, 12343, 225.9} ///shorter way

	fmt.Println(a_car.gas_pedal)
}

