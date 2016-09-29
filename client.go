package main

import (
        "fmt"
        "net"
)

func main() {      
      var port string
      fmt.Println("Ingrese el puerto del servidor a conectarse")
      fmt.Scanln(&port)     
     
    for{
        conn,err := net.Dial("tcp", "127.0.0.1:"+port)
        if(err!=nil){
    fmt.Println("Ingrese el puerto correcto")
    break
   }
    fmt.Print("Ingrese su mensaje: ") 
       var msg string
        
       fmt.Scanln(&msg)
       if (msg=="cerrar"){
                break
                }
       fmt.Fprintf(conn,msg+"\n")
         }
        
}
