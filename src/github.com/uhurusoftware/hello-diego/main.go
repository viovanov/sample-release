// hello-diego project main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/apcera/nats"
	"github.com/cloudfoundry-incubator/cf-lager"
	"github.com/cloudfoundry/gunk/diegonats"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"
	"github.com/uhurusoftware/hello-diego/viewer"
)

var natsAddresses = flag.String(
	"natsAddresses",
	"127.0.0.1:4222",
	"comma-separated list of NATS addresses (ip:port)",
)

var natsUsername = flag.String(
	"natsUsername",
	"nats",
	"Username to connect to nats",
)

var natsPassword = flag.String(
	"natsPassword",
	"nats",
	"Password for nats user",
)

var listenAddr = flag.String(
	"listenAddr",
	"0.0.0.0",
	"host to listen on for the HTTP viewer",
)

var listenPort = flag.Int(
        "listenPort",
        8080,
        "port to listen on for the HTTP viewer",
)

var subscription *nats.Subscription

func main() {
	logger := cf_lager.New("hello-diego")

	natsClient := diegonats.NewClient()
	natsClientRunner := diegonats.NewClientRunner(*natsAddresses, *natsUsername, *natsPassword, logger, natsClient)

	// This is the notification channel from NATS to the HTTP Viewer
	flowViewer := make(chan string, 100)

	group := grouper.NewOrdered(os.Interrupt, grouper.Members{
		{"nats-client", natsClientRunner},
		{"viewer-server", InitializeHTTPServer(logger, flowViewer)},
	})

	monitor := ifrit.Envoke(sigmon.New(group))

	ListenOnNATS(natsClient, flowViewer)

	<-monitor.Wait()

	fmt.Println("\nReceived an interrupt, stopping ...\n")
	StopListening()
}

func ListenOnNATS(client diegonats.NATSClient, flowViewer chan (string)) error {
	var err error
	subscription, err = client.Subscribe(">", func(msg *nats.Msg) {
		fmt.Printf("Received: %s\n", msg.Subject)
		flowViewer <- TagMessage(msg)
	})

	if err != nil {
		return err
	}

	return nil
}

func InitializeHTTPServer(logger lager.Logger, flowViewer chan (string)) ifrit.Runner {
	viewerHandler, err := viewer.NewServer(logger, flowViewer)

	if err != nil {
		panic("failed to initialize viewer server: " + err.Error())
	}
	return http_server.New(*listenAddr, viewerHandler)
}

func StopListening() error {
	return subscription.Unsubscribe()
}
