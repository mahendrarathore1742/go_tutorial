package main
import ("fmt"
		"net/url"
)

const myurl string ="https://lco.dev:3000/learn?coursename=reactjs";

func main(){

	fmt.Println("welcome to url handling");
	fmt.Println(myurl);

	// result, _ := url.Parse(myurl);


	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)


	partofUrl := &url.URL{
		Scheme: "https",
		Host : "https://lco.dev",
		Path:"/learn",
		RawPath :"User=hitesh",

	}

	anotherUrl  := partofUrl.String()
	fmt.Println(anotherUrl);
}