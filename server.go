package main

import (
        "fmt"
        "net"
)

func main() {
      var port string
      fmt.Println("Ingrese el puerto con el que desea abrir el servidor")
      fmt.Scanln(&port)
      ln,_ := net.Listen("tcp", ":"+port)
for {
    conn,_ := ln.Accept()
    var cmd []byte
    fmt.Fscan(conn, &cmd)
    fmt.Println("Recibimos:", string(cmd))
}
}

