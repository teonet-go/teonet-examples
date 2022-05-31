package main

import (
	"fmt"
	"time"

	"github.com/teonet-go/teonet"
)

const (
	appName       = "Teonet echo command client sample application"
	appShort      = "teoechocommandcli"
	appVersion    = "0.0.1"
	echoComServer = "WXJfYLDEtg6Rkm1OHm9I9ud9rR6qPlMH6NE"
	sendDelay     = 3000
)

func main() {
	// Teonet application logo
	teonet.Logo(appName, appVersion)

	// Start Teonet client
	teo, err := teonet.New(appShort)
	if err != nil {
		panic("can't init Teonet, error: " + err.Error())
	}

	// Connect to Teonet
	err = teo.Connect()
	if err != nil {
		teo.Log().Debug.Println("can't connect to Teonet, error:", err)
		panic("can't connect to Teonet, error: " + err.Error())
	}

	// Connect to echo command server
	err = teo.ConnectTo(echoComServer,

		// Get messages from echo command server
		func(c *teonet.Channel, p *teonet.Packet, e *teonet.Event) (proc bool) {
			// Skip not Data Events
			if e.Event != teonet.EventData {
				return
			}

			// Print received message
			fmt.Printf("got from %s, \"%s\", len: %d, id: %d, tt: %6.3fms\n\n",
				c, p.Data(), len(p.Data()), p.ID(),
				float64(c.Triptime().Microseconds())/1000.0,
			)
			proc = true

			return
		},
	)
	if err != nil {
		panic("can't connect to server, error: " + err.Error())
	}

	// Send messages to echo command server
	for {
		data := "Teonet developer!"
		fmt.Printf("send to  %s, \"%s\", len: %d\n", echoComServer, data,
			len(data))
		_, err = teo.Command(129, data).SendTo(echoComServer)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Duration(sendDelay) * time.Millisecond)
	}
}
