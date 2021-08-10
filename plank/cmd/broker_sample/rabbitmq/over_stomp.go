// Copyright 2019-2021 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package rabbitmq

import (
	"github.com/vmware/transport-go/bridge"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/plank/utils"
	"os"
	"time"
)

func ListenViaStomp(c2 chan os.Signal) {
	b := bus.GetBus()
	bus.EnableLogging(true)

	broker, err := b.ConnectBroker(&bridge.BrokerConnectorConfig{
		Username:     "guest",
		Password:     "guest",
		ServerAddr:   "localhost:61613",
		HeartBeatOut: 30 * time.Second,
		STOMPHeader: map[string]string{
			"access-token": "something",
		},
	})
	if err != nil {
		utils.Log.Fatalln("conn error", err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		broker.SendMessage("/topic/something.somewhere", "application/octet-stream", []byte("i can send too!"))
	}()

	subs, err := broker.Subscribe("/topic/something.somewhere")
	if err != nil {
		utils.Log.Fatalln(err)
	}
	c := subs.GetMsgChannel()
	go func() {
		for msg := range c {
			utils.Log.Infoln(msg)
		}
	}()

	utils.Log.Infoln("waiting for messages")
	<-c2
}
