package main
import "net"
import "os"
import "fmt"
import "io/ioutil"

func main(){
  if len(os.Args)!=2{
    fmt.Fprintf(os.Stderr,"Usage %s host:port",os.Args[0])
    os.Exit(1)
  }
  service:=os.Args[1]
  tcpaddr,err:=net.ResolveTCPAddr("tcp4",service)
  if err!=nil{
    fmt.Println("Error occured:",err)
    return
  }
  conn,err:=net.DialTCP("tcp",nil,tcpaddr)
  if err!=nil{
    fmt.Println("Error :",err)
    return
  }
  result,err:=ioutil.ReadAll(conn)
  if err!=nil{
    fmt.Println("Error:",err)
    return
  }
  fmt.Println(string(result))
  os.Exit(0)
}