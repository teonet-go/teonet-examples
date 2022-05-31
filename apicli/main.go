package main

import (
	"fmt"
	"time"

	"github.com/teonet-go/teonet"
)

const (
	appName    = "Teonet api client sample application"
	appShort   = "teoapicli"
	appVersion = "0.0.1"
	apiServer  = "WXJfYLDEtg6Rkm1OHm9I9ud9rR6qPlMH6NE"
	sendDelay  = 3000
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

	// Connect to api server
	err = teo.ConnectTo(apiServer)
	if err != nil {
		panic("can't connect to server, error: " + err.Error())
	}

	// Create Teonet client API interface
	apicli, err := teo.NewAPIClient(apiServer)
	if err != nil {
		panic("can't create Api Client, error: " + err.Error())
	}

	// Send messages to api server and receive answer
	for {
		data := []byte("Developer!")
		fmt.Printf("send to  %s, \"%s\", len: %d\n", apiServer, data, len(data))
		_, err = apicli.SendTo("hello", data, func(data []byte, err error) {
			// Process error
			if err != nil {
				fmt.Println(err)
				return
			}
			// Print received message
			fmt.Printf("got from %s, \"%s\", len: %d\n\n",
				apicli.Address(), data, len(data),
			)
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Duration(sendDelay) * time.Millisecond)
	}
}
