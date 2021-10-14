package main
import "fmt"
import "net"
import "time"


func main(){
   service:=":1201"
   udpaddr,err:=net.ResolveUDPAddr("udp4",service)
   if err!=nil{
      fmt.Println("Error:",err)
      return
   }
   for{
     conn,err:=net.ListenUDP("udp",udpaddr)
     if err!=nil{
      fmt.Println("Error occured:",err)
      return
     }
     go handleClient(conn)
   }
}


func handleClient(conn *net.UDPConn){
   fmt.Println("Waiting for client")
   var buf [512]byte
   n,err:=conn.Read(buf[0:])
   if err!=nil{
      return
   }
   fmt.Println(string(buf[0:n]))
   daytime:=time.Now().String()
   fmt.Println("Current daytime is :",daytime)
   fmt.Fprintf(conn,daytime)
}