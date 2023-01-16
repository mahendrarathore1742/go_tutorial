package main;


import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"

)


type Course struct {

	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`

}

type Author struct {
	FullName string `json:"fullname"`
	Websiet string `json:"website"`
}


//fake db
var courses []Course;


func main(){


	 fmt.Println("")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "python", CoursePrice: 299, Author: &Author{FullName: "Mahendra Singh", Websiet: "google.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "go", CoursePrice: 199, Author: &Author{FullName: "Mahendra Singh", Websiet: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))


}

//middleware,helper- file
func (c *Course) isEmpty() bool {
	return c.CourseName == "";
}


//controller - File


//serve home router

func serveHome(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("<h1>Welcome to API"))
}  

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Get One Course");
	w.Header().Set("Content-Type", "applicatioan/json")

	//grab id from request

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return;
}

func createOneCourse(w http.ResponseWriter, r *http.Request){

		fmt.Println("Create One Course");
		w.Header().Set("Content-Type","application/json");

		//what if body is empty
		if r.Body == nil{
			json.NewEncoder(w).Encode("Please send some data ");
		}

		var course Course;


		_ = json.NewDecoder(r.Body).Decode(&course)

		if course.isEmpty(){
			json.NewEncoder(w).Encode("Please Send some data");
			return;
		}

		// genrate unique id,string;
		//append course into courses


		rand.Seed(time.Now().UnixNano())
		course.CourseId = strconv.Itoa(rand.Intn(100));

		courses = append(courses,course)

		json.NewEncoder(w).Encode(course);
		return;

}


func updateOneCourse(w http.ResponseWriter, r *http.Request){

		fmt.Println("update One Course");
		w.Header().Set("Content-Type","application/json");


		//first grab id from req
		params := mux.Vars(r);


		for index,course := range courses {

			if course.CourseId == params["id"] {

				courses = append(courses[:index],courses[index+1:]...);
				var course Course;
				_ = json.NewDecoder(r.Body).Decode(&course);
				course.CourseId = params["id"];
				json.NewEncoder(w).Encode(course);
				return;
			}
		}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){

		fmt.Println("Create One Course");
		w.Header().Set("Content-Type","application/json");

		params := mux.Vars(r);

		for index,course := range courses {

			if course.CourseId == params["id"] {

				courses = append(courses[:index],courses[index+1:]...);
				break;
			}
		}
}


 
