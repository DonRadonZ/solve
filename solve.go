package main

import (
	"encoding/base64"
	"fmt"
)



func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
   
	return string(runes)
}

//Use For Answer
func main() {

	 var whatIsIt string
   
	   secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
   
	   sd, _ := base64.StdEncoding.DecodeString(secret)
   
	   whatIsIt = reverse(string(sd))
   
	   fmt.Println(whatIsIt)
   
   }