package main

func (me *post) storeMessage() {

	ldb.stmtInsertMessage.Exec(me.username, me.time, me.message)
}

// func getNote() (result []notesStruct) {
// 	sqlReadall := `
// 	SELECT notes, time FROM notesWritten
// 	`

// 	rows, err := ldb.db.Query(sqlReadall)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		notes := notesStruct{}
// 		err = rows.Scan(&notes.Notes, &notes.TimeNow)
// 		if err != nil {
// 			panic(err)
// 		}
// 		result = append(result, notes)
// 	}
// 	log.Printf("\nThese are the results: %s\n ", result)

// 	return result
// }
