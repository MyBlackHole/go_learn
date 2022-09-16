
	// 2 修改数据
	sqlStr := "update music set name= ? where id = ?"
	ret, err := DB.Exec(sqlStr, "长大后我就成了你", 4)
	if err != nil {
		fmt.Println("更新失败", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("获取影响的行数失败：", err)
		return
	}
	fmt.Println("更新成功，影响行数为：", n)
