package main

import (
	"log"
	"os"
)
func main(){

/* Exercise 1
Create a new file in the current working directory called info.txt and then close the file.
If the file cannot be created (no permissions, wrong path etc) then print out the error and
stop the program (do error handling).
*/

newFile, err := os.Create("info.txt")
if err!=nil{
	log.Fatal(err)
}
newFile.Close()


/* 
Exercise 2	
Rename the file created at Exercise #1 to data.txt
Check if file exists before renaming it. If it doesn't exist print a message and stop the program.
*/

var fileInfo os.FileInfo
fileInfo, err = os.Stat("info.txt")
if err!=nil{
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
}
var newFileName = "data.txt"
err = os.Rename(fileInfo.Name(), newFileName)
if err!=nil{
	log.Fatal(err)
}

/*
Exercise 3
Remove the file created at exercise #1. 
Take care that the file is now called data.txt (it has been renamed at exercise #2).
*/

err = os.Remove(newFileName)
if err!=nil {
	log.Fatal(err)
}


}