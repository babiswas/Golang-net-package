package main
import "net"
import "os"
import "fmt"


func main(){
   if len(os.Args)!=2{
     fmt.Fprintf(os.Stderr,"Usage is %s hostname\n",os.Args[0])
     os.Exit(1)
   }
   name:=os.Args[1]
   addrs,err:=net.LookupHost(name)
   if err!=nil{
     fmt.Println("Error:",err.Error())
     os.Exit(2)
   }
   for _,s:=range addrs{
     fmt.Println(s)
   }
   os.Exit(0)
}