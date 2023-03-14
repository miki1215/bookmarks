package main

import (
	"fmt"
)


type Bookmark struct{
	Name string
	Id int
	Url string
}

func (b Bookmark)AddDate() Bookmark{
return Bookmark{Name: "jj"}
}

func main() {
	var bbb Bookmark
	bbb.Id = 7 
	fmt.Print(bbb.AddDate().Id)
}
