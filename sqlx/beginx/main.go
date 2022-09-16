
	tx, err := DB.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("开启事务错误:%v\n", err)
		return
	}

	_, err = tx.Exec("insert into music(name, year,sign_id) values (?,?,?)", "听奶奶的话", 2023, 1)
	if err != nil {
		tx.Rollback() // 出错就回滚
		fmt.Println("出错回滚")
		return
	}
	_, err = tx.Exec("insert into music(name, year,sign_id) values (?,?,?)", "听爷爷的话")
	if err != nil {
		tx.Rollback() // 出错就回滚
		fmt.Println("出错回滚")
		return
	}
	tx.Commit()
