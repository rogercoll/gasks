package gasks

import (
        "fmt"
        "net"
)

func Start(CONNECT string) {

        s, err := net.ResolveUDPAddr("udp4", CONNECT)
        c, err := net.DialUDP("udp4", nil, s)
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
        defer c.Close()

        data := []byte("Hellow from systemd" + "\n")
        _, err = c.Write(data)

        if err != nil {
           fmt.Println(err)
           return
        }
}
