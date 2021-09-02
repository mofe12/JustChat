package main

func (me *post) storeMessage() {

	ldb.stmtInsertMessage.Exec(me.globUsername, me.time, me.message, me.userID)
}

func getNote() (result []post) {
	sqlReadall := `
	SELECT Name,Time,Message FROM chats WHERE UserID = ?
	`

	rows, err := ldb.db.Query(sqlReadall, globUserid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		post := post{}
		err = rows.Scan(&post.globUsername, &post.time, &post.message)
		if err != nil {
			panic(err)
		}
		result = append(result, post)
	}
	return result
}

func createUser(user string) {
	ldb.stmtCreateUser.Exec(user)
}
func getCurrentUser(userid int, check bool) map[int]string {
	result := map[int]string{}
	switch check {
	case true:
		sqlReadID := `
		SELECT * FROM users WHERE user_id = ?
		`
		row, err := ldb.db.Query(sqlReadID, userid)
		if err != nil {
			panic(err)
		}
		defer row.Close()
		var id int
		var user string
		for row.Next() {
			err = row.Scan(&id, &user)
			if err != nil {
				panic(err)
			}
		}
		result[id] = user
		return result

	case false:
		sqlReadID := `
		SELECT * FROM users ORDER BY rowid DESC LIMIT 1
		`
		row, err := ldb.db.Query(sqlReadID)
		if err != nil {
			panic(err)
		}
		defer row.Close()
		var id int
		var user string
		for row.Next() {
			err = row.Scan(&id, &user)
			if err != nil {
				panic(err)
			}
		}
		result[id] = user
		return result
	}
	// Should not get here
	return nil
}
