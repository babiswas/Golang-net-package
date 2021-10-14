package main
import "fmt"
import "net"

func main(){
  service:=":1201"
  tcpaddr,err:=net.ResolveTCPAddr("tcp4",service)
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  listener,err:=net.ListenTCP("tcp",tcpaddr)
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  for{
    conn,err:=listener.Accept()
    if err!=nil{
      continue
    }
    handleClient(conn)	
    conn.Close()
  }
}


func handleClient(conn net.Conn){
   fmt.Println("Recieved client connection\n")
   _,err:=conn.Write([]byte("Connected to server\n"))
   if err!=nil{
       fmt.Println("Error occured:",err)
   }
   var buf [512]byte
   for{
     n,err:=conn.Read(buf[0:])
     if err!=nil{
         return
     }
     fmt.Println(string(buf[0:]))
     _,err2:=conn.Write(buf[0:n])
     if err2!=nil{
          return
     }
   }
}