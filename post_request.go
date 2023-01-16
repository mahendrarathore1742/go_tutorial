package main;
import ("fmt"
		"net/http"
		"io/ioutil"
		"strings"
		"net/url"
);


func main(){
	fmt.Println("Welcome to Post Request")
	perfromPostRequest2();
}

func performPostRequest(){

	const myurl = "http://localhost:3000/post";

	requestBody := strings.NewReader(`
			{
				"Name" :"Mahendea",
				"Age" : 12,
				"course" :"python",
				"price" : 0
			}
		`)


		response,err := http.Post(myurl,"application/json",requestBody);

		if err != nil{
			panic(err)
		}

		defer response.Body.Close();

		content,_ := ioutil.ReadAll(response.Body);

		fmt.Println(string(content))

}


func perfromPostRequest2(){

	const myurl = "http://localhost:8000/post";
	data:= url.Values{}

	data.Add("Fiist Name","mahendra")
	data.Add("Mid Name","Singh")
	data.Add("Last Name","rathor")
	data.Add("Age","18")

	response,err := http.PostForm(myurl,data);

	if err != nil{
		panic(err)
	}


	defer response.Body.Close();

	content,_ := ioutil.ReadAll(response.Body);
	fmt.Println(string(content));
}



