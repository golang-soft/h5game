package main

import (
	"github.com/magicsea/ganet/network"
	"net"

	//	"io/ioutil"
	"log"
	//	"net/http"
	//	"github.com/magicsea/ganet/network/protobuf"
	//	"github.com/gogo/protobuf/proto"
	"encoding/json"
)

//JsData
type JsData struct {
	Id  string `json:Id`
	Msg string `json:Msg`
}

func newAgent(conn network.TCPConn) network.Agent {
	Client := new(Agent)
	Client.conn = conn
	return Client
}

//Agent
type Agent struct {
	conn      network.TCPConn
	msgHandle func(channel byte, msgId interface{}, data []byte)
}

//Run
func (a *Agent) Run() {
	log.Println("Agent.run")
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Println("read message: ", err)
			break
		}
		var jd = new(JsData)
		errUm := json.Unmarshal(data, jd)
		if errUm != nil {
			log.Println("Unmarshal fail:", errUm)
			break
		}
		a.msgHandle(0, jd.Id, []byte(jd.Msg))
	}
}

//OnClose
func (a *Agent) OnClose() {}

//WriteMsg
func (a *Agent) WriteMsg0(channel byte, msgId byte, msg []byte) {

	data := []byte{channel, msgId}
	data = append(data, msg...)
	err := a.conn.WriteMsg(data)
	if err != nil {
		log.Println("write message error:", err)
	}

}

//WriteMsg
func (a *Agent) WriteMsg(msgID interface{}, rawmsg []byte) {

	var jd = JsData{Id: msgID.(string), Msg: string(rawmsg)}
	//m := map[string]interface{}{JsonIdName: msgID,JsonMsgName:rawmsg}
	data, _ := json.Marshal(jd)

	// data := []byte{msgId, 0, 0}
	// data = append(data, msg...)
	serr := a.conn.WriteMsg(data)
	if serr != nil {
		log.Println("write message error:", serr)
	}

}

//LocalAddr
func (a *Agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

//RemoteAddr
func (a *Agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

//Close
func (a *Agent) Close() {
	a.conn.Close()
}

//Destroy
func (a *Agent) Destroy() {
	a.conn.Destroy()
}
