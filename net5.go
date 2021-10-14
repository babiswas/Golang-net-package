package main
import "fmt"
import "os"
import "net"

func main(){
  if len(os.Args)!=3{
    fmt.Fprintf(os.Stderr,"Usage is %s network service\n",os.Args[0])
    os.Exit(1)
  }
  networkType:=os.Args[1]
  service:=os.Args[2]
  port,err:=net.LookupPort(networkType,service)
  if err!=nil{
    fmt.Println("Error:",err.Error())
    os.Exit(2)
  }
  fmt.Println("Service Port",port)
  os.Exit(0)
}