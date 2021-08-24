package main

var ldb *LocalDatabase

var lorem = `Lorem ipsum dolor sit amet.`

func main() {
	ldb = &LocalDatabase{}
	ldb.init("./chat.db")
	views()
}
