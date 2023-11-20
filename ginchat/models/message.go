package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"strconv"

	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

// 消息
type Message struct {
	gorm.Model
	FormId   int64  //发送者
	TargetId int64  //接受者
	Type     int    //消息类型  群聊，私聊，广播
	Media    int    ///消息类型  文字 图片 音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {

	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 需要：发送者ID 接收者ID 消息类型 发送的内容 发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	//1.获取参数并校验合法性	校验token
	//token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	//msgtype := query.Get("Type")
	isvalida := true
	conn, err := (&websocket.Upgrader{ //.Upgrader在这个依赖中github.com/gorilla/websocket
		//token校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	获取连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//3、用户关系
	//	4、userid 跟node绑定 并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	//	5、完成发送逻辑
	go sendProc(node)

	//	6、完成接受逻辑
	go recvProc(node)
	sedMsg(userId, []byte("欢迎进入聊天系统"))
}
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws]<<<<<", string(data))
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

// 完成udp数据发送协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		sedMsg(msg.TargetId, data)
		//case "2"://群发
		//	sedGroupMsg()
		//case "3"://广播
		//	sedAllMsg()
		//case 4:

	}
}
func sedMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
