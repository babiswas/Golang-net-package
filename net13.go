package main
import "net"
import "fmt"
import "os"
import "bufio"


func main(){
  if len(os.Args)!=2{
    fmt.Fprintf(os.Stderr,"Usage %s host:port",os.Args[0])
    os.Exit(1)
  }
  service:=os.Args[1]
  tcpaddr,err:=net.ResolveTCPAddr("tcp4",service)
  if err!=nil{
   fmt.Println("Error:",err)
   return
  }
  conn,err:=net.DialTCP("tcp",nil,tcpaddr)
  if err!=nil{
     fmt.Println("Error:",err)
     return
  }
  _,err=conn.Write([]byte("Hello Server"))
  if err!=nil{
     fmt.Println("Error:",err)
     return
  }
  fmt.Println("Message sent to the server")
  for{
     var buf [512]byte
     n,err:=conn.Read(buf[0:])
     fmt.Println(string(buf[0:n]))
     scanner:=bufio.NewScanner(os.Stdin)
     fmt.Println("Enter data")
     scanner.Scan()
     _,err=conn.Write([]byte(scanner.Text()))
     if err!=nil{
        fmt.Println("Error:",err)
        return
     }
  }
  os.Exit(2)
}