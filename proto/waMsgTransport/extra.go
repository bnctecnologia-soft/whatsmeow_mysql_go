package waMsgTransport

import (
	"github.com/bnctecnologia-soft/whatsmeow_mysql_go/proto/armadilloutil"
	"github.com/bnctecnologia-soft/whatsmeow_mysql_gologia-soft/whatsmeow_mysql_go/proto/waMsgApplication"
)

const (
	MessageApplicationVersion = 2
)

func (msg *MessageTransport_Payload) Decode() (*waMsgApplication.MessageApplication, error) {
	return armadilloutil.Unmarshal(&waMsgApplication.MessageApplication{}, msg.GetApplicationPayload(), MessageApplicationVersion)
}

func (msg *MessageTransport_Payload) Set(payload *waMsgApplication.MessageApplication) (err error) {
	msg.ApplicationPayload, err = armadilloutil.Marshal(payload, MessageApplicationVersion)
	return
}
