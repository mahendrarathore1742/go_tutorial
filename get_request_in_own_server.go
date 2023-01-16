package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func PerformGetReqest(){

	const url  = "http://localhost:3000/";

	response,err :=  http.Get(url);

	if err != nil{
			panic(err)
	}

	defer response.Body.Close();

	fmt.Println(response.StatusCode);
	fmt.Println(response.ContentLength);

	content,_ := ioutil.ReadAll(response.Body);

	fmt.Println(string(content));

}

func main(){

	PerformGetReqest();

}