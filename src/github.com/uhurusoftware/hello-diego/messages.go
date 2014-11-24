package main

import (
	"strings"

	"github.com/apcera/nats"
)

var messages map[string]string

func initMessageMap() {
	if messages == nil {
		messages = make(map[string]string)

	}
}

func TagMessage(msg *nats.Msg) string {
	tags := ""
	subject := strings.ToLower(msg.Subject)
	body := strings.ToLower(string(msg.Data))

	switch subject {
	case "router.register":
		if strings.Contains(body, "fib-cpu") {
			tags += "router;route_emmiter;"
		}
	case "diego.desire.app":
		tags += "cloud_controller_ng;nsync;converger"

	case "cell_z1-0.bid-for-start-auction":
		tags += "auctioneer;rep;auction;"

	case "cell_z1-0.rebid-then-tentatively-reserve":
		tags += "auctioneer;rep;auction;"

	case "cell_z1-0.run":
		tags += "executor;garden;warden-linux"
	}

	return tags
}
