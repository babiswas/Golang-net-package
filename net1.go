package main
import "fmt"
import "os"
import "net"

func main(){
  if len(os.Args)!=2{
    fmt.Fprintf(os.Stderr,"Please provide ip adess:%s format",os.Args[0])
    os.Exit(1)
  }
  name:=os.Args[1]
  addr:=net.ParseIP(name)
  if addr==nil{
    fmt.Println("Invalid IP adreess")
  }else{
    fmt.Println("the adress is",addr.String())
  }
  os.Exit(0)
}