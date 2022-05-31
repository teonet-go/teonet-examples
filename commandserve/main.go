package main

import (
	"fmt"

	"github.com/teonet-go/teonet"
)

const (
	appName    = "Teonet echo command server application"
	appShort   = "teoechocommandserve"
	appVersion = "0.0.1"
)

func main() {

	// Teonet application logo
	teonet.Logo(appName, appVersion)

	// Start Teonet client
	teo, err := teonet.New(appShort,
		// Main application reader - receive and process incoming messages
		func(teo *teonet.Teonet, c *teonet.Channel, p *teonet.Packet,
			e *teonet.Event) bool {

			// Skip not Data events
			if e.Event != teonet.EventData {
				return false
			}

			// Skip not server mode messages
			if !c.ServerMode() {
				return false
			}

			// Parse command
			cmd := teo.Command(p.Data())

			// Process command
			switch cmd.Cmd {

			case 129:
				// Print received message
				fmt.Printf("got from %s, \"%s\", len: %d, id: %d, cmd: %d, "+
					"tt: %6.3fms\n",
					c, cmd.Data, len(cmd.Data), p.ID(), cmd.Cmd,
					float64(c.Triptime().Microseconds())/1000.0,
				)

				// Send answer
				answer := []byte("Teonet answer to " + string(cmd.Data))
				c.Send(answer)

			default:
				fmt.Println("got wrong command")
			}

			return true
		},
	)
	if err != nil {
		panic("can't init Teonet, error: " + err.Error())
	}

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
