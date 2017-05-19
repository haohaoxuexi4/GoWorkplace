package main
import(
	"fmt"
	"log"
	"net/http"
	"sync"
)
var count int
var mu sync.Mutex
func handler(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	count++
	fmt.Println("handle")
	//fmt.Fprintln(count)
	mu.Unlock()
	fmt.Fprintf(w,"urlpath=%q\n",r.URL.Path)
}
func counter(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	fmt.Fprintf(w,"count=%d\n",count)
	mu.Unlock()
}
func main(){

	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}