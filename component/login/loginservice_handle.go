package login

import (
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto"
	"github.com/ouczbs/Zmin/engine/zproto/zpb"
)

func (service *ULoginService) MessageLoop() {
	for {
		select {
		case message := <-service.MessageQueue:
			proxy := message.Proxy
			messageType := message.MessageType
			packet := message.Packet
			switch messageType {
			case zconf.MT_TO_SERVER, zconf.MT_BROADCAST,zconf.MT_TO_ALL:
				zproto.PbMessageHandle(proxy, packet,message.Cmd)
			case zconf.MT_TO_CLIENT:
				service.ForwardToClient(packet)
			default:
				if messageType > zconf.MT_TO_GAME_START && messageType < zconf.MT_TO_GAME_END {
					service.ForwardToGame(proxy , messageType , packet)
				}
			}
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *ULoginService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT)] = service.AddEngineComponent
	reqHandleMaps[TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK)] = service.AddEngineComponentAck
	service.ReqHandleMaps[TCmd(zpb.CommandList_MT_SYNC_PROXY_PROPERTY)] = service.SyncProxyProperty
}
func (service *ULoginService) ConnectToCenter() {
	centerProxy = service.MakeCenterProxy()
	if centerProxy == nil {
		service.Close()
	}
	message := &zpb.ADD_ENGINE_COMPONENT{
		ComponentId: service.Config.ComponentId,
		Type:       zpb.COMPONENT_TYPE_LOGIN,
		ListenAddr: service.Config.ListenAddr,
	}
	request := znet.NewRequest(TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_TO_SERVER,message)
	zproto.SendPbMessage(centerProxy, request)
	request.Release()
}
func (service *ULoginService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, int32(message.ComponentId))
	proxy.SetProperty(zattr.Int32ComponentType, int32(zpb.COMPONENT_TYPE_CENTER))
	for _, login := range message.ComponentList {
		service.MakeClientProxy(string(login.ListenAddr),zpb.COMPONENT_TYPE_GAME)
	}
}
func (service *ULoginService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	service.MakeClientProxy(string(message.ListenAddr),zpb.COMPONENT_TYPE_GAME)
	zlog.Debug("AddEngineComponent " , message.Type , message.ListenAddr)
}
func (service *ULoginService) ForwardToGame(proxy * UClientProxy,messageType TMessageType , packet * UPacket) {
	id ,ok := proxy.GetProperty(zattr.Int32ComponentId).(int32)
	if !ok {
		return
	}
	packet.AppendComponentId(id)
	game := gameMessageMaps[messageType]
	if game != nil {
		game.ForwardPacket(packet)
	}
}
func (service *ULoginService) ForwardToClient(packet * UPacket) {
	id := packet.SubtractComponentId()
	proxy := clientProxyMaps[id]
	if proxy != nil {
		proxy.ForwardPacket(packet)
	}
}
func (service *ULoginService) SyncProxyProperty(proxy *UClientProxy, request * URequest) {
	service.UService.SyncProxyProperty(proxy ,request)
	componentType , ok := proxy.GetProperty(zattr.Int32ComponentType).(int32)
	if ok && componentType == int32(zpb.COMPONENT_TYPE_GAME){
		if messageType , ok := proxy.GetProperty(zattr.Int32MessageType).(int32); ok{
			gameMessageMaps[TMessageType(messageType)] = proxy
		}
	}
}