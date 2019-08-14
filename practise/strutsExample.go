
package main

import ( "fmt")


type student struct{
	   firstName string
	   lastname string
	   sid  int
	   age int
	   address address

}

type address struct{
	  street string
	  city string
	  state string
	  country string

}



func main(){
fmt.Println("hi this is structs example")
	
     s1 := student{
		  firstName :"phani",
		  lastname : "gurram",
		  sid : 91,
		  age: 24,
		  address:address{
			street : "n college drive",
			city :"maryville",
			state :"missouri",
			country:" usa",
		 },
	 }

     fmt.Println(s1.address)


}