package main
import "net"
import "os"
import "fmt"


func main(){
  if len(os.Args)!=2{
    fmt.Fprintf(os.Stderr,"Usage %s host:port",os.Args[0])
    os.Exit(1)
  }
  service:=os.Args[1]
  udpaddr,err:=net.ResolveUDPAddr("udp4",service)
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  conn,err:=net.DialUDP("udp",nil,udpaddr)
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  fmt.Println("Communicating with server")
  fmt.Fprintf(conn,"Hello Server")
  var buf [512]byte
  n,err:=conn.Read(buf[0:])
  if err!=nil{
     fmt.Println("Error:",err)
     return
  }
  fmt.Println(string(buf[0:n]))
  os.Exit(0)
}