package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	// "runtime"
)

const OFF = false
const ON = true

// App struct
type App struct {
	ctx context.Context
	conn net.Conn 
	alive bool
	recv bool
	filename string
}


// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	a.alive = OFF
	a.recv = OFF
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}


// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Connect(port string) string {

	if a.alive {
		return "请先断开连接"
	}

	conn, err := net.Dial("tcp", port)

	if err != nil {
		// log.Fatalf("dial failed :%s", err.Error())
		fmt.Println("net.Dial err:", err)
		return "连接失败"
	}

	a.conn = conn
	a.alive = ON
	fmt.Println("已连接...")
	return "连接成功"
}

//根据不同命令进行不同处理
func (a *App) Handle(msg string) {
		
	if !a.alive {
		// return "请先连接"
		a.windowPrint("请先连接")
		return
	}

	if a.recv {
		if msg != "" {
			  a.filename = msg
		}
		a.recvFile()
		a.recv = OFF
		a.windowPrint("接收完毕")
		return
	}

	// a.conn.Write([]byte(msg + "\n" ))
	msg = strings.Trim(msg, "\r\n")
	args := strings.Split(msg, " ")
	cmd := strings.TrimSpace(args[0])

	a.send(msg)
	servermsg, _ := bufio.NewReader(a.conn).ReadString('\n')
	// return fmt.Sprintln(servermsg)
	// 处理服务器传来的字符串
	servermsg = strings.Trim(servermsg, "\r\n")

	if cmd == "ls" {
		runtime.EventsEmit(a.ctx, "ls", servermsg)
		return
	}

	if cmd == "exit" {
		if servermsg == "Bye" {
			a.conn.Close()
			a.alive = OFF
			fmt.Println("已断开...")
		}
	}
	if cmd == "req" && len(args) == 2  {
		if strings.TrimSuffix(servermsg, "?") == args[1] {
			// a.recvFile(args[1])
			a.ackFile()
			a.filename = args[1]
			a.recv = ON
			a.windowPrint("给文件取个名吧")
			return
		}
	}
	// runtime.EventsEmit(a.ctx, "msg", servermsg)
	a.windowPrint(servermsg)

}

//发送msg
func (a *App) send(msg string) {
	a.conn.Write([]byte(msg + "\n" ))
}

//注册事件msg,将servermsg传给前端
func (a *App)windowPrint(servermsg string) {
	runtime.EventsEmit(a.ctx, "msg", servermsg)
} 

//确认接收的文件，向服务端发送"ACK"
func (a *App)ackFile() {
	fmt.Println("准备接收")
	_, err := a.conn.Write([]byte("ACK"))
	if err != nil {
			fmt.Println("Write error" , err)
	}

}

// 接收文件
func (a *App)recvFile() {

	 savePath := "D:\\recv"
	 file, err := os.Create(savePath + string(os.PathSeparator) + a.filename)
	 if err != nil {
			fmt.Println("os.Create err:" , err)
			a.windowPrint("创建文件失败")
		  return
	 }
	 buf := make([]byte, 4096)
	 a.windowPrint("正在接收")
	 for {
		 	fmt.Println(".")
		 	n, err := a.conn.Read(buf)

			if err != nil {
				fmt.Println("conn.Read err:", err)
				return
			}

			//如果不这样进行文件写入的结尾，下一次循环的conn.Read就会阻塞,本质上是由于粘包问题导致的
			//最后服务端conn.Write一次写入
			// if string(buf[n-3:n]) == "EOF" {
			//    fmt.Println("接收完毕")
			// 	 file.Write(buf[:n-3])
			// 	 return
			// }
			file.Write(buf[:n])
			//当n < 规定的buf长度，说明是最后一次接收和写入，直接return
			if n < 4096 {
				fmt.Println("接收完毕")
				// file.Close()
				return
			}
	 }
}