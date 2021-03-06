package zproto

import (
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zutil"
)

func sendPbMessage(proxy *UClientProxy , request * URequest, wrap *UWrapMessage)error{
	packet := znet.NewPacket()
	packet.WriteMessageType(request.MessageType)
	packet.WriteMessageCmd(request.Cmd)
	out , err := Marshal(request.ProtoMessage)
	if err != nil{
		return err
	}
	wrap.Content = out
	buf , err := Marshal(wrap)
	if err != nil{
		return err
	}
	packet.AppendBytes(buf)
	err = proxy.SendPacket(packet)
	packet.Release()
	return err
}
func SendPbMessage(proxy *UClientProxy , request * URequest)error{
	wrap := &UWrapMessage{}
	return sendPbMessage(proxy , request , wrap)
}
func RequestMessage(proxy *UClientProxy , request * URequest , handle FRequestHandle)error{
	sequence := zutil.IncSequence()
	wrap := &UWrapMessage{
		Request: sequence,
		Code:request.Code,
	}
	proxy.ReqHandleMaps[sequence] = handle
	return sendPbMessage(proxy , request , wrap)
}
func ResponseMessage(proxy *UClientProxy ,request * URequest)error{
	wrap := &UWrapMessage{
		Response: request.Request,
		Code:request.Code,
	}
	return sendPbMessage(proxy , request , wrap)
}
func MakePbMessagePacket(request * URequest)* UPacket{
	packet := znet.NewPacket()
	wrap := &UWrapMessage{
		Code:request.Code,
	}
	packet.WriteMessageType(request.MessageType)
	packet.WriteMessageCmd(request.Cmd)
	out , err := Marshal(request.ProtoMessage)
	if err != nil{
		zlog.Error(" MakePbMessagePacket Marshal message error " ,err)
		return nil
	}
	wrap.Content = out
	buf , err := Marshal(wrap)
	if err != nil{
		zlog.Error(" MakePbMessagePacket Marshal wrap message error " ,err)
		return nil
	}
	packet.AppendBytes(buf)
	return packet
}