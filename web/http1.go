
package main

import   "io"
import   "log"
import   "net/http"


func main(){
  mux:=http.NewServeMux()
  mux.Handle("/",&myHandler{})

  err:=http.ListenAndServe(":8090",mux)//default while the second parameter is nil
  if err!=nil{
  log.Fatal(err)
  }
}

type myHandler struct{}

func (this *myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){//must have
   io.WriteString(w,"URL:"+r.URL.String())  
}
