package controller

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/continentalbreakfast17/lighting-controller/src/model"
)

/*func TestConnect(t *testing.T) {
	serial, err := NewSerial("/dev/ttyACM0", 9600)
	if err != nil {
		panic(err)
	}

	if err := serial.Shutdown(); err != nil {
		panic(err)
	}
}*/

/*func TestSendStatic(t *testing.T) {
	serial, err := NewSerial("/dev/ttyACM0", 9600)
	if err != nil {
		panic(err)
	}

	static := model.SttcConfig{
		Color: [3]byte{255,0,255},
		Pins:  [3]int{9,10,11},
		Mode:  1,
		Wait:  5,
	}

	if resp, err := serial.Send(&static); err != nil {
		panic(err)
	} else {
		fmt.Printf("%d\n", resp)
	}

	if err := serial.Shutdown(); err != nil {
		panic(err)
	}
}*/

func TestSendStatic(t *testing.T) {
	var channel1 model.Channel
	if err := json.Unmarshal([]byte(raw1), &channel1); err != nil {
		panic(err)
	}
	var channel2 model.Channel
	if err := json.Unmarshal([]byte(raw2), &channel2); err != nil {
		panic(err)
	}
	addr := model.AddrConfig{[]model.Channel{channel1, channel2}}

	serial, err := NewSerial("/dev/ttyACM0", 115200)
	if err != nil {
		panic(err)
	}

	if err := serial.write(&addr); err != nil {
		panic(err)
	}

	if resp, err := serial.read(); err != nil {
		panic(err)
	} else {
		fmt.Printf("%d ", resp)
	}

	if err := serial.Shutdown(); err != nil {
		panic(err)
	}
}

const raw1 = `
{
	"pin": 4,
	"wait": 50,
	"ledCount": 16,
	"seqSize": 32,
	"sequence": [
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 5000,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[0,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[0,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[0,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[0,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[0,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 5000,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		}
	]
}
`

const raw2 = `
{
	"pin": 3,
	"wait": 50,
	"ledCount": 16,
	"seqSize": 32,
	"sequence": [
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 5000,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[255,0,255],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[255,0,255]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[255,0,255],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[255,0,255],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[255,0,255],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[255,0,255],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 50,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[255,0,255],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254]]
		},
		{
			"wait": 5000,
			"strip": [[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254],[0,0,254]]
		}
	]
}
`