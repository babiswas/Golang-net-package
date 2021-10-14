package main
import "fmt"
import "net"
import "time"

func main(){
  service:=":1200"
  tcpaddr,err:=net.ResolveTCPAddr("tcp4",service)
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  listener,err:=net.ListenTCP("tcp",tcpaddr)
  if err!=nil{
    fmt.Println("Error occured")
    return
  }
  for{
    conn,err:=listener.Accept()
    if err!=nil{
       continue
    }
    daytime:=time.Now().String()
    conn.Write([]byte(daytime))
    conn.Close()
  }
}