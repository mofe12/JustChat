package main

import "fmt"

func main() {
	ldb = &LocalDatabase{}
	ldb.init("./chat.db")
	views()
	fmt.Println(globUsername, globUserid)

}
