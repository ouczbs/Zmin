package zproto

import (
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto/zpb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"strings"
)

var (
	pbMessageTypes  = make(map[TCmd]IReflectMessageType)
	CommandListName = zpb.CommandList_name
)

func PbMessageHandle(proxy *UClientProxy , packet * UPacket , cmd TCmd){
	wrapBytes := packet.MessagePayload()
	wrapMessage := &UWrapMessage{}
	Unmarshal(wrapBytes , wrapMessage)
	handle , globalHandle := proxy.GetRequestHandles(wrapMessage.Response , cmd)
	if cmd == 0 || (handle == nil && globalHandle == nil){
		return
	}
	pbMessage,_ := newPbMessage(cmd)
	request := GetRequestMessage(wrapMessage ,pbMessage)
	proxy.Then(handle , request).Then(globalHandle , request)
}

func GetRequestMessage(wrap *UWrapMessage, message IReflectMessage) *URequest {
	Unmarshal(wrap.Content, message)
	request := znet.NewRequest(_CMD_INVALID, _MT_INVALID, message)
	request.Code = wrap.Code
	request.Next = true
	return request
}
func newPbMessage(cmd TCmd) (IReflectMessage , error){
	messageType := pbMessageTypes[cmd]
	if messageType != nil {
		return messageType.New().Interface() ,nil
	}
	pbName := CommandListName[int32(cmd)]
	pbName = strings.Replace(pbName, "MT_" , "zpb." , 1)
	messageType,err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(pbName))
	if err != nil {
		return nil , err
	}
	pbMessageTypes[cmd] = messageType
	return messageType.New().Interface() ,err
}