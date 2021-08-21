package main

var ldb *LocalDatabase

func main() {
	ldb = &LocalDatabase{}
	ldb.init("./chat.db")

	chat()
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

}
