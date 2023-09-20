package main
import (
	"os"
)

port:=os.Getenv("PORT")
if port ==""{
	port="8080"
}

http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hi")
})
fmt.Println("Starting server on port %s...\n", port)
err:=http.ListenAndServe(":"+port, nil)
if err!=nil{
	fmt.Printf("Error: %s\n", err)
}