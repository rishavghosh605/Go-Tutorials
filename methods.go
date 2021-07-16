//Two types of methods value receivers and pointer receivers

// If you want to change a value in a struct then use a pointer receivers

package main

import "fmt"

const u16bitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 // min 0 max 65535, means unsigned
	brake_pedal uint16
	steering_wheel int16 // -32k to +32k
	top_speed_kmh float64
}

//Creating value receivers
func (c car) kmh() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/u16bitmax/kmh_multiple)
}
func main() {
    a_car := car{gas_pedal: 22341, 
				brake_pedal: 0, 
				steering_wheel: 12343, 
				top_speed_kmh: 225.9} // longer way

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh()) // You can think of it as a method of the struct car
	fmt.Println(a_car.mph())	
}