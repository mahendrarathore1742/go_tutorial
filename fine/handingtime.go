package main

import ("fmt"
		"time"
);

func main(){



	PresentTime :=  time.Now();
	
	fmt.Println(PresentTime);

	fmt.Println(PresentTime.Format("1-02-2006 Monday"));
	

}