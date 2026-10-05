package main

import (
	"errors"
	"fmt"
)
var b float64 = float64(a)+ c
const a = 8
var c float64
var d int
var e bool
var f string
func main(){
	 s:=34.56
r:="I go, therefore I am"
t:= true
fmt.Printf("%f\n",b)
fmt.Printf("%f\n",s)
fmt.Printf("%t\n",t)
fmt.Printf("%f\n",c)
fmt.Printf("%d\n",a)
fmt.Printf("%t\n",e)
fmt.Printf("%s\n",f)
fmt.Printf("%s\n",r)
fmt.Printf("%d\n",d)
realm2()
realm3()
realm4()
realm5()}

//:= is an imperative statement and requires order to work whereas declatraion using "var" can be both imperative as well as flow independent depending upon the scopes it is put into 
// more specifically := works only into function scopes and var can be used any where thru out the code
// := infers type from the literal so for giving a:= 8 as a float value it must be written as 8.0 or should be deifnesd using var
//constants are looked after at compile time and are not assigned any memory address unlike variables which are assigned a memory address along with a value (default or user input) at run time
//constans stay constant through out the course and can be used in places where the value has to remain constant such as scientific constants


//Go has a unique feature of assigning default 0 values to the variables it havent recievd values from the programmer becaus majority of the languages just keep it undefined or print garbage value which again creates inconsistency within the code 
// it also avoides errors caused due to forgot to assign a value to although it comes at a trade off. It becomes indistinguishable for user to know whether the zero value is explicitly assigned or is given by the compiler itself.




func realm2 (){
	var age int 
fmt.Println("Enter the age")
fmt.Scanf("%d",&age)
if age>18{
	fmt.Println("You are allowed to vote")

   } else if (age>16) {
    fmt.Println("wait for one more year")
	} else
	 {fmt.Println("Grow up first")
	}
// the else if ladder here including the normal if else statement doesnt require brackets for the conditions and go also supports short one liner tasks such as calling another function which would provide a vlaue which is then used in the if condition
var ratings int
fmt.Println("Enter the rating")
fmt.Scanf("%d", &ratings)
	switch ratings {
		 case 5: fmt.Println("Excellent")
		case 4: fmt.Println("very good")
		case 3: fmt.Println("good")
		case 2: fmt.Println("ok")
		case 1: fmt.Println("bad")
		default : fmt.Println("give valid readings")
		}
		
// switch in go does not require an explicit break command. unlike other languages the fallthrough command has to be explicitly written and the break comand is by default executed once a successful case is met
		switch {
			 case age>18 : fmt.Println("You are allowed to vote")
			case age>16  : fmt.Println("wait for one more year")
			default: fmt.Println("Grow up first")
			}
// switch can also be condition matching based rather than value matching based there can be multiple cases possible in a particular value but the first one that becomes true is considered
	for i:=0; i<5; i++ {// regular for loop
	  fmt.Printf("%d\n",i)
		}	// prints 0 1 2 3 4
		i:=0 
	for i<10 {// while style for loop 
	  fmt.Printf("%d\n",i)
	 
		i++
		} // prints 0 to 9
		
		for { fmt.Printf("%d \n",i)
		     i++
			 if i>15 { break;}
			} // no condition for loop. runs indefinitely untill hit by a manual break command
			

		array:=[]int{10,11,12}
		for i,v := range array {
		 fmt.Printf("%d,%d \n",i,v)
			} //ranging over slice
		// if index is not required it can be replaced by - but it has to be replaced orelse the declared but not used error shall hit as the range compulsorily returns two values one the index and the second one actually the value that is stired in that index


		ab:=map[string]int {"a":1, "b":2,}
		for k,v:= range ab {
		fmt.Printf("%s,%d \n",k,v)
			} //randing over a map

		} // rangeing over a map gives two values the key and the value itself. ranging over maps is by design made random so that any coincidental pattern in the mapping sequence is not taken as guaranteed pattern and is not used to build logic around that pattern.

		//logic behhind every loop is similar that run the loop untill the ocndition is true. so why keep multiple syntax for the same logic
		// thus in go as we keep dropping the clauses the for loop starts changing into so called different loops like dropping the init and post conditions makes for behave like a while loop
		//whereas dropping all 3 factors make it an endless loop 




		func realm3() {
			 fmt.Printf("%s\n", ageverification(32))
		   st, in := helios(44)
		   fmt.Printf("%s,%d", st, in)
		   numslice:=[]int{23,24,56}
		   variadic(numslice...)
            fun:= closure()
			fmt.Printf("%d", fun())
		}

		func ageverification(ag int) string {  // has only one parameter and one return type string. A function can easily have multiple parameters
			//and can have multiple return types. like leiterally anything could be the parameter and anything could be the return type be it a function itself., a function can pass a function as a return type.
			if ag>18  {return "You are eligibler"
		    } else { return "Grow up"}
		
		}

		func helios(num int) (name string, serial int) {

		 if num>50 {
		 name="heilos senior"
	      serial=50
		   return 
		} else 
		{ name="heilos junior"
	      serial=10
		   return}
		}

		func variadic(numslice ...int) {
				 fmt.Printf("%v\n",numslice) }

		func closure() func() int { // function that returns function iteself. here when the return statement is hit it returns the immediate unnamed function as it is.
	
			fmt.Println("This is the first closure call")
			return func() int { fmt.Println("This is consecutive inner function call")
		                    return 0} 
		}

// The go developers delibrately dropped the try and catch block because it, by its design, afer hitting an error will skip the commands until it hits a catch block which could be multiple levels away
//this makes tracing the error very difficult and also skips multiple commands 
// go wanted the error to be treated at the very point it originates so that the error does not interfer and stop the execution of other staements

type People struct {
	aj int
	naim string
} 
func (p People) Greet() string {
    greeting:="Hello"+p.naim
	return greeting
}
func realm4() {

slice :=[]int{2,4,5}  //slice is dynamic in size as go automatically increases or decreases the size of slice based upon the new entries that has been made by fully coping the existing slice plus the new elements into a bigger memory block. It is basically a mechanisim that uses arrays and creates new arrays to adjust the sizes and changes
slice = append(slice,6)
array :=[4]int{45,67,7,4} //arrays are  fixed in nature you cannot change the size of an  array
fmt.Printf("%v \n", slice)
fmt.Printf("%v \n", array)

mapp:=make(map[string]int) // map declaration has 2 ways one is to just initialize the map using make keyword and other is to create and add entrties on the spot without using the make keyword. Just initializing the map without make keyword or the literals like the other varibales will give it a nil value which will create a panic when trying to add entries in the map later on
mapp["Hello"]=3

var p People= People{ aj:45, naim:"Ajay"} // adding values to the structure created outside the function
fmt.Printf("%s", p.Greet())  //function with value reciever
p.Ageincrease() // function with pointer reciever


}

func (p *People) Ageincrease() { 
	p.aj=p.aj+1
	fmt.Printf("\nThe age is now %d\n",p.aj)
}

// pointer is passed on so that the age changes could actually reflect in the structure as otherwise the structure copy is passed and the chnages are made into the copy and nothing actually reflects the changes.
// value and pointer reciever can be used differently at places where a copy of the vaue is needed to work upon so that the real data is not altered and the pointer is used where the real data needs to be updated so the funct runs dirctly on the memory location itself rather than a cpoy of it	



func realm5() { 
	go sayhi() // go routines are basically a way that go developers usedto enable concept similar to multitreading where the traditional line by line execution is skipped and the code blocks which can run independently are run simultaneously to maximise efficiency and make full use of resources available 
	defer fmt.Println("This is bye from the fuction") // defer function allows us to run a particular command right before the function ends. such as closing an open folder or saving the changes the code might have done irrespective of which way the function ends ( in case of multiple return conditions)
	fmt.Println("This is hello from realm5")
ch:= <- channel // channel is the tool used to create some level of syncronization between code blocks running simultaneously. basically it ensures that the block where a certain function was called does not complete execution and return values without waiting for the simultaneous running function.
fmt.Printf("The new channel value is %d \n",ch)
result, err := safeDivide(10, 0)
if err != nil {  // go skips try and catch block as explained in the realm above.
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}


}
var channel = make(chan int)
func sayhi() {
	fmt.Printf("This is hello from goroutine\n")
 channel <- 1}

func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
// The goroutines entirely eliminate the fuss created by sharing the memory between two function/variables at the same time. this can create confusion regarding the rewriting or reading the wrong value/old value. Although locks handle that issue well but locks are very manual and if not placed correctly they might endup being locked forever. 
// the channel solves this issue by not allowing memory sharing between two variables aat the same time