package main

import (
	"fmt"

	"github.com/teonet-go/teonet"
)

const (
	appName    = "Teonet api server application"
	appShort   = "teoechoapiserve"
	appLong    = "This description of Teonet api server application"
	appVersion = "0.0.1"
)

func main() {

	// Teonet application logo
	teonet.Logo(appName, appVersion)

	// Start Teonet client
	teo, err := teonet.New(appShort)
	if err != nil {
		panic("can't init Teonet, error: " + err.Error())
	}

	// Create Teonet API
	api := teo.NewAPI(appName, appShort, appLong, appVersion)

	// Create API command "hello"
	cmdApi := teonet.MakeAPI2()
	cmdApi.
		SetCmd(api.Cmd(129)).                 // Command number cmd = 129
		SetName("hello").                     // Command name
		SetShort("get 'hello name' message"). // Short description
		SetUsage("<name string>").            // Usage (input parameter)
		SetReturn("<answer string>").         // Return (output parameters)
		// Command reader (execute when command received)
		SetReader(func(c *teonet.Channel, p *teonet.Packet, data []byte) bool {
			fmt.Printf("got from %s, \"%s\", len: %d, id: %d, tt: %6.3fms\n",
				c, data, len(data), p.ID(),
				float64(c.Triptime().Microseconds())/1000.0,
			)
			data = append([]byte("Hello "), data...)
			api.SendAnswer(cmdApi, c, data, p)
			return true
		}).SetAnswerMode(teonet.DataAnswer)
	api.Add(cmdApi)

	// Add API reader
	teo.AddReader(api.Reader())

	// Print API
	fmt.Printf("API description:\n\n%s\n\n", api.Help())

	// Connect to Teonet
	err = teo.Connect()
	if err != nil {
		panic("can't connect to Teonet, error: " + err.Error())
	}

	// Print application address
	addr := teo.Address()
	fmt.Println("Connected to teonet, this app address:", addr)

	// Wait forever
	select {}
}
