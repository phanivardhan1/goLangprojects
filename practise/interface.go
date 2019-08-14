
package main

import("fmt")

type vehicle interface{
     
	getmodel()  
	getYear()
}	

type car struct{

	model string
	year string
} 

type truck struct{
	model string
	year string
	wheels int
	maxspeed int
}

func (r car) getmodel(){
	
	 fmt.Println(r.model)

}

func (r car) getYear( ){
	
	fmt.Println(r.year)

}


func (r truck) getmodel(){
	
	fmt.Println(r.model , r.wheels)

}

func (r truck) getYear( ){
   
   fmt.Println(r.year,r.maxspeed)

}


func main(){
  
 c:= car{
		model : "ford",
		year : "2008",
 }

 c.getmodel()

 a := car{
	model : "toyota",
	year : "2020",
 }
	a.getYear()
	
    t:= truck{
		model : "toyota",
		year : "2020",
		wheels:10,
		maxspeed:100,
	}

  t.getmodel()
  t.getYear()



}