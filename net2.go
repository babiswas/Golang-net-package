package main
import "net"
import "fmt"
import "os"


func main(){
  if len(os.Args)!=2{
    fmt.Fprintf(os.Stderr,"Usage %s dotted ip adress\n",os.Args[0])
    os.Exit(1)
  }
  dotAddr:=os.Args[1]
  addr:=net.ParseIP(dotAddr)
  if addr==nil{
    fmt.Println("Invalid IP adress")
    os.Exit(1)
  }
  mask:=addr.DefaultMask()
  network:=addr.Mask(mask)
  ones,bits:=mask.Size()
  fmt.Println("Adress is ",addr.String(),"DEfault mask length is ",bits,"Leading ones count is",ones,"Mask in hex is",mask.String(),"Network is:",network.String())
  os.Exit(0)
}