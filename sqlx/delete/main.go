
	// 3 删除数据
	sqlStr := "delete from music where id = ?"
	ret, err := DB.Exec(sqlStr, 1)
	if err != nil {
		fmt.Println("删除出错：", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Println("获取操作影响的行数出错：", err)
		return
	}
	fmt.Println("删除成功，影响的行数为：", n)
