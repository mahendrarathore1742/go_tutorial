package main;

import ("fmt"
		"encoding/json"
);


// Varibale name should be Uppercase;
type courses struct{
	Course string 	`json:"coursesname"`
	Price int
	Platform string `json:"website"`
	Password string `json:"_"`
	Tags []string 	`json:"tags,omitempty"`
}

func main(){
	fmt.Println("Welcome to Json data")
	Decodejson();
}

func Encoding(){

	LearnCode := []courses{

		{"python",1500,"google.com","msr12345",[]string{"py","python"}},
		{"java",1500,"google.com","msr12345",[]string{"java","java"}},
		{"Javascript",1500,"google.com","msr12345",nil},

	}

	//package this data as Json
	finalJson,err := json.MarshalIndent(LearnCode,"","\t");

	if err !=nil{
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)
}


func Decodejson(){

	jsonDataFromWeb := []byte(`
	{
		"coursesname": "java",
		"Price": 1500,
		"website": "google.com",
		"_": "msr12345",
		"tags": ["java","java"]
	}
		`)

		// var LearnCourse courses

		// checkValie := json.Valid(jsonDataFromWeb)

		// if checkValie {

		// 	fmt.Println("Json was valid");
		// 	json.Unmarshal(jsonDataFromWeb,&LearnCourse);
		// 	fmt.Printf("%#v\n",LearnCourse)

		// } else {
		// 	fmt.Println("Json is not valie")
		// }




		// alternative way

		var myOnlineData map[string]interface{};

		json.Unmarshal(jsonDataFromWeb,&myOnlineData);
		// fmt.Printf("%#v\n",myOnlineData)


		for key,value := range myOnlineData{

			fmt.Printf("Key is %v and value is  %v and type is %T\n ",key,value,value)

		}

}