
//一个简单httpserver

package main

import(
	"net/http"
	"log"
	"fmt"
	"strings"
)
func sayhi(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_log"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hi say my name")
}

func main(){
http.HandleFunc("/",sayhi)
err:=http.ListenAndServe(":8080",nil)
if err!=nil{
	log.Fatal("listerandserver:",err);
}

}