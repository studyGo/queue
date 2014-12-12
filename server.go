package main

import (
    "fmt"
    "syscall"
    "os"
    "strconv"
    "./queue"
)

func main () {
    server, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_IP)
    q := queue.Create()
    if err != nil {
        fmt.Println("create socket server error")
        return
    }

    var serveraddr syscall.SockaddrInet4
    s := make(chan string, 1000)

    serveraddr.Addr = [4]byte{127,0,0,1}
    serveraddr.Port = 6000

    err = syscall.Bind(server, &serveraddr)
    if err != nil {
        fmt.Println("bind err")
        return
    }

    err = syscall.Listen(server, syscall.SOMAXCONN)
    if err != nil {
        fmt.Println("listen err")
        return
    }

    for {
        fd, _, err := syscall.Accept(server)
        if err != nil {
            continue
        }
        go accept(fd, s, *q)
    }
}
func accept(fd int, sess chan string, q queue.Queue) {
    buffer := make([]byte, 1000)
    for {
        length, err := syscall.Read(fd, buffer)
        if err != nil {
            fmt.Println("read error")
        }
        f, _ := os.OpenFile("socket.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0660)
        f.WriteString(string(length))
        f.Close()

        switch string(buffer[:length-2]) {
            case "get":
                s := get(q)
                syscall.Write(fd, []byte(s))
                break;
            case "len":
                s := lenght(q)
                syscall.Write(fd,[]byte(s))
                break;
            default:
                set(string(buffer[:length-2]), q)
                break;
        }
    }
}

func get(q queue.Queue) string {

    return q.LPop().(string)
}

func set(s string, q queue.Queue) {
    q.RPush(s)
}

func lenght(q queue.Queue) string {
    return strconv.Itoa(q.Len())
}
