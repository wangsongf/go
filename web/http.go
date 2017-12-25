
package main

import   "io"
import   "log"
import   "net/http"


func sayHello(w http.ResponseWriter,r *http.Request){
  io.WriteString(w,"Hello world,this is version one.")   

}


func main(){
  //set route
  http.HandleFunc("/",sayHello)
  err:=http.ListenAndServe(":8090",nil)//default while the second parameter is nil
  if err!=nil{
  log.Fatal(err)
  }
}
