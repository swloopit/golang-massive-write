package massivewrite

import (
"time"
"os"
"github.com/swloopit/golang-random-key"
)

func BuildRandomWrite(position_write,data,defined_fixed_name,extension_file string) (string,bool){
     os.MkdirAll(position_write,0777)
     s:=GenerateFileName(defined_fixed_name,extension_file)
     for ;CheckFileExists(position_write,s)==true;{
         s=GenerateFileName(defined_fixed_name,extension_file)
     }
     return s,WriteFile(position_write,s,data)

}


func CheckFileExists(path,filename string) bool{
     s:=path
     s+=filename
     if _, err := os.Stat(s); os.IsNotExist(err) {
	return false
     }
     return true
}

func GenerateFileName(defined_fixed_name,extension_file string) string{
     tmp_name:=time.Now()
     tmp_random_key:=randomkey.CreateRandomKey(32)
     return defined_fixed_name+tmp_name.String()+tmp_random_key+extension_file
}

func WriteFile(path,filename,content string) bool{
     file,err :=os.Create(path+filename)
     defer file.Close()
     if err == nil {
     	file.Write([]byte(content))
	return true
     }
     return false
}
