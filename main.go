package main

import (
	"fmt"
	"io/ioutil"
	"os"
)



func readFile(){
	data, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}

func writeFile() {
	data := []byte("say hello to you")
	err := ioutil.WriteFile("love.txt", data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func appendFile(){
	// 打开一个文件
	f,err:= os.OpenFile("test.txt",os.O_APPEND|os.O_WRONLY,0600)
	if err != err {
		panic(err)
	}
	// 然后往里面写东西
	if _,err= f.WriteString("new new new , i am new "); err!= nil {
		panic(err)
	}
}




func main() {
	writeFile()
	readFile()
	appendFile()
}
