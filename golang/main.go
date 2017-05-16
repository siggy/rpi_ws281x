package main

import (
        "fmt"
	"github.com/jgarff/rpi_ws281x/golang/ws2811"

	"encoding/binary"
)

//ws2811_led_t dotcolors_rgbw[] =
//{
//    0x00200000,  // red
//    0x10200000,  // red + W
//    0x00002000,  // green
//    0x10002000,  // green + W
//    0x00000020,  // blue
//    0x10000020,  // blue + W
//    0x00101010,  // white
//    0x10101010,  // white + W
//};

const (
	LED_COUNT = 30
	GPIO_PIN = 18
	LOOPS = 100
)

var (
	red    = binary.LittleEndian.Uint32([]byte{0x00, 0x20, 0x00, 0x00})
	redw   = binary.LittleEndian.Uint32([]byte{0x10, 0x20, 0x00, 0x00})
        green  = binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x20, 0x00})
        greenw = binary.LittleEndian.Uint32([]byte{0x10, 0x00, 0x20, 0x00})
        blue   = binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x20})
        bluew  = binary.LittleEndian.Uint32([]byte{0x10, 0x00, 0x00, 0x20})
        white  = binary.LittleEndian.Uint32([]byte{0x00, 0x10, 0x10, 0x10})
        whitew = binary.LittleEndian.Uint32([]byte{0x10, 0x10, 0x10, 0x10})

	colors = []uint32{red, redw, green, greenw, blue, bluew, white, whitew}
)

func main() {
	err := ws2811.Init(GPIO_PIN, LED_COUNT, 255)
        if err != nil {
		fmt.Printf("ws2811.Init failed: %+v\n", err)
		panic(err)
	}

	defer ws2811.Fini()

	fmt.Printf("calling Clear()\n")
	ws2811.Clear()

	err = ws2811.Render()
        if err != nil {
                fmt.Printf("ws2811.Render failed: %+v\n", err)
                panic(err)
        }
	err = ws2811.Wait()
        if err != nil {
                fmt.Printf("ws2811.Wait failed: %+v\n", err)
                panic(err)
        }

	color := 0
	fmt.Printf("cycle LEDs\n")
	for l := 0; l < LOOPS; l++ {
	for i := 0; i < LED_COUNT; i++ {
   		ws2811.SetLed(i, colors[color % len(colors)])
		color++

		err = ws2811.Render()
		if err != nil {
			fmt.Printf("ws2811.Render failed: %+v\n", err)
			panic(err)
		}
		err = ws2811.Wait()
		if err != nil {
			fmt.Printf("ws2811.Wait failed: %+v\n", err)
			panic(err)
        	}
	}
	}

        fmt.Printf("calling Clear()\n")
        ws2811.Clear()

        err = ws2811.Render()
        if err != nil {
                fmt.Printf("ws2811.Render failed: %+v\n", err)
                panic(err)
        }
        err = ws2811.Wait()
        if err != nil {
                fmt.Printf("ws2811.Wait failed: %+v\n", err)
                panic(err)
        }
}

