package main
import ( "fmt"
		"net/http"
		"io/ioutil"
);


const url = "https://lco.dev/";

func main(){

	fmt.Println("welcome to web request");

	response,err := http.Get(url);

	if err != nil{
		panic(err);
	}

	fmt.Println(response);

	defer response.Body.Close();

	file,err := ioutil.ReadAll(response.Body);

	if err != nil{
		panic(err);
	}

	content := string(file);

	fmt.Println(content);




}