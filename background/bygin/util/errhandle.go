package util

import "fmt"

func Errhandle(err error){
	if err!=nil{
		fmt.Println(err)
	}
}