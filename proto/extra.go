package armadillo

import (
	"google.golang.org/protobuf/proto"

	"github.com/bnctecnologia-soft/whatsmeow_mysql_go/proto/waArmadilloApplication"
	"github.com/bnctecnologia-soft/whatsmeow_mysql_gologia-soft/whatsmeow_mysql_go/proto/waCommon"
	"github.com/bnctecnologia-soft/whatsmeow_mysql_gologia-soft/whatsmeow_mysql_go/proto/waConsumerApplication"
	"github.com/bnctecnologia-soft/whatsmeow_mysql_gologia-soft/whatsmeow_mysql_go/proto/waMultiDevice"
)

type MessageApplicationSub interface {
	IsMessageApplicationSub()
}

type RealMessageApplicationSub interface {
	MessageApplicationSub
	proto.Message
}

type Unsupported_BusinessApplication waCommon.SubProtocol
type Unsupported_PaymentApplication waCommon.SubProtocol
type Unsupported_Voip waCommon.SubProtocol

var (
	_ MessageApplicationSub = (*waConsumerApplication.ConsumerApplication)(nil) // 2
	_ MessageApplicationSub = (*Unsupported_BusinessApplication)(nil)           // 3
	_ MessageApplicationSub = (*Unsupported_PaymentApplication)(nil)            // 4
	_ MessageApplicationSub = (*waMultiDevice.MultiDevice)(nil)                 // 5
	_ MessageApplicationSub = (*Unsupported_Voip)(nil)                          // 6
	_ MessageApplicationSub = (*waArmadilloApplication.Armadillo)(nil)          // 7
)

func (*Unsupported_BusinessApplication) IsMessageApplicationSub() {}
func (*Unsupported_PaymentApplication) IsMessageApplicationSub()  {}
func (*Unsupported_Voip) IsMessageApplicationSub()                {}
