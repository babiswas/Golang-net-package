package main
import "net"
import "fmt"
import "io/ioutil"
import "os"


func main(){
  if len(os.Args)!=2{
    fmt.Fprintf(os.Stderr,"Usage %s host:port",os.Args[0])
    os.Exit(1)
  }
  service:=os.Args[1]
  tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)
  if err!=nil{
    fmt.Println("Error occured:",err)
    return
  }
  conn,err:=net.DialTCP("tcp",nil,tcpAddr)
  if err!=nil{
    fmt.Println("Error occured:",err)
    return
  }
  _,err=conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
  if err!=nil{
    fmt.Println("Error occured:",err)
    return
  }
  result,err:=ioutil.ReadAll(conn)
  if err!=nil{
    fmt.Println("Error Occured:",err)
    return 
  }
  fmt.Println(string(result))
  os.Exit(0)
}