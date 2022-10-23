package znet

import (
	"fmt"
	"go/src/zinx/ziface"
	"net"
)

// 服务器实体类
type Server struct {
	// 服务器名称
	Name string
	// 服务器IP版本
	IPVersion string
	// 服务器IP
	Ip string
	// 服务器端口
	Port int
	// 路由：当前的Server添加一个router, sever注册的链接对应的处理业务
	Router ziface.IRouter
}

// 服务器启动
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP: %s, Port: %d, is starting", s.Ip, s.Port)
	// 获取一个TCP的地址
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.resolveIPAddr err ", err)
		return
	}
	// 监听服务器的地址
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("net.ListenTCP err ", err)
		return
	}

	fmt.Println("start Zinx server succ, ", s.Name, " succ, Listenning")
	var cid uint32
	cid = 0

	// 阻塞等待客户端连接，处理客户端连接业务（读写）
	for {
		// 如果有客户端连接进来，阻塞会返回
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}
		dealConn := NewConntion(conn, cid, s.Router)
		cid++

		// 启动当前的链接业务处理
		go dealConn.Start()
	}
}

// 服务器停止
func (s *Server) Stop() {
	// TODO 将服务器的资源，状态，或者一些已经开辟的连接信息，进行停止或者释放
}

// 服务器运行
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务后的额外业务

	// 阻塞状态
	select {}

}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router succ!")
}

// 初始化Server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      9999,
		Router:    nil,
	}
	return s
}
