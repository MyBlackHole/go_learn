
// 1 增加数据
sqlStr := "insert into music(name, year,sign_id) values (?,?,?)"
ret, err := DB.Exec(sqlStr, "听爸爸的话", 2023, 1)
if err != nil {
	fmt.Println("插入出错", err)
	return
}
theID, err := ret.LastInsertId() // 新插入数据的id
if err != nil {
	fmt.Println("获取插入的id错误：", err)
	return
}
fmt.Println("插入成功，id为：", theID)
