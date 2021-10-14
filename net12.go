package main
import "net"
import "fmt"


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
       go handleClient(conn)
   }
}


func handleClient(conn net.Conn){
    defer conn.Close()
    var buf [512]byte
    for{
       n,err:=conn.Read(buf[0:])
       fmt.Println(string(buf[0:n]))
       if err!=nil{
          return
       }
       _,err2:=conn.Write(buf[0:n])
       if err2!=nil{
          return
       }
    }
}