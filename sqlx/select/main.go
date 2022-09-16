
// 4 查询数据单条
var music struct {
	Id     int
	Name   string
	Year   string
	SignId int
}
//QueryRow后一定要调用Scan方法，否则持有的数据库链接不会被释放
err = DB.QueryRow("select * from music where id=?", 2).Scan(&music.Id, &music.Name, &music.Year, &music.SignId)
if err != nil {
	fmt.Println("查询出错：", err)
}
fmt.Println(music)

// 5 查询多条
sqlStr := "select * from music where id > ?"
rows, err := DB.Query(sqlStr, 1)
if err != nil {
	fmt.Println("查询出错：", err)
	return
}
// 关闭rows释放持有的数据库链接
defer rows.Close()
// 循环读取结果集中的数据
for rows.Next() {
	var m struct {
		Id     int
		Name   string
		Year   string
		SignId int
	}
	err := rows.Scan(&m.Id, &m.Name, &m.Year, &m.SignId)
	if err != nil {
		fmt.Println("遍历出错", err)
		return
	}
	fmt.Println(m)
}

type Music struct {
	Id     int
	Name   string
	Year   string
	SignId int `db:"sign_id"`
}
var music Music
err = DB.Get(&music, "select * from music where id=?", 2)
if err != nil {
	fmt.Println("查询出错", err)
}
fmt.Println(music)

// Select 查询
type Music struct {
	Id     int
	Name   string
	Year   string
	SignId int `db:"sign_id"`
}
var music []Music
err = DB.Select(&music, "select * from music where id > ?", 2)
if err != nil {
	fmt.Println("查询出错", err)
}
fmt.Println(music)

sqlStr := "insert into music(name, year,sign_id) values (:name,:year,:sign_id)"
_, err = DB.NamedExec(sqlStr,
	map[string]interface{}{
		"name":    "好汉歌",
		"year":    2024,
		"sign_id": 1,
	})

// 使用map作为查询名
type Music struct {
	Id     int
	Name   string
	Year   string
	SignId int `db:"sign_id"`
}
sqlStr := "SELECT * FROM music WHERE name=:name"
// 使用map做命名查询
rows, _ := DB.NamedQuery(sqlStr, map[string]interface{}{"name": "好汉歌"})
defer rows.Close()
for rows.Next() {
	var m Music
	rows.StructScan(&m)
	fmt.Println(m)
}

// 使用结构体作为查询名
type Music struct {
	Id     int
	Name   string
	Year   string
	SignId int `db:"sign_id"`
}
sqlStr := "SELECT * FROM music WHERE name=:name"
// 使用结构体做命名查询
var music = Music{Name: "好汉歌"}
rows, _ := DB.NamedQuery(sqlStr, music)
defer rows.Close()
for rows.Next() {
	var m Music
	rows.StructScan(&m)
	fmt.Println(m)
}
