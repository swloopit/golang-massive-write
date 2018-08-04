# golang-massive-write
save a massive number of files in text form without thinking about giving a unique name

example:

'''golang
package main

import(
"fmt"
"github.com/swloopit/golang-massive-write"
"encoding/json"
)

func main(){
     type s struct{
          B string
          L string
     }

     k:=new(s)
     k.B="CIAO"
     k.L="BELLA"
     b, _ := json.Marshal(k)
     fmt.Println(massivewrite.BuildRandomWrite("./trydir1/",string(b),"startA",".json"))
     fmt.Println(massivewrite.BuildRandomWrite("./trydir2/","ciao","startB",".json"))

}
''''
