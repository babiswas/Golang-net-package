package main
import "fmt"
import "net"
import "os"

func main(){
  if len(os.Args)!=2{
     fmt.Fprintf(os.Stderr,"Usage %s hostname\n",os.Args[0])
     os.Exit(1)
  }
  name:=os.Args[1]
  addr,err:=net.ResolveIPAddr("ip",name)
  if err!=nil{
    fmt.Println("resolution error",err.Error())
    os.Exit(1)
  }
  fmt.Println("Resolved adress is:",addr.String())
  os.Exit(0)
}