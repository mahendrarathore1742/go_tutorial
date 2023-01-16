package main;

import( "fmt"
		"os"
		"io"
		
);


func main(){

	fmt.Println("welcome to fine sections")


	content := "Hi my name is Mahendra singh Rathore "
	file,error := os.Create("./msr.txt");



	if error != nil{
		panic(error)
	}


	lenght , error := io.WriteString(file,content);

	if error != nil{
		panic(error)
	}

	fmt.Println(lenght)
	defer file.Close();

	readFile("./msr.txt")
}

func readFile(filename string){

	dataType,err := os.ReadFile(filename)

	if err != nil{
		panic(err);
	}


	fmt.Println(string(dataType))

}