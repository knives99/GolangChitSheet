package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {

	//1 創建數據庫引擎對象
	//連接數據庫 規定的字符拼接  用戶名稱:用戶密碼@tcp(127.0.0.1)/數據庫名稱
	engine, err := xorm.NewEngine("mysql", "root:sofia25037245@tcp(127.0.0.1:3306)/ginsql")
	if err != nil {
		panic(err.Error())
	}
	//2 設置映射規則 (要駝峰式還是怎樣的拼法)
	engine.SetMapper(core.SnakeMapper{})
	engine.ShowSQL(true)

	//3.同步數據庫表格
	engine.Sync2(new(PersonTable))
	//4.判斷person表格是否存在
	personExist, err := engine.IsTableExist(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personExist {
		fmt.Println("人員表存在")
	} else {
		fmt.Println("人員表不存在")
	}
	//5.判斷person表格是否為空
	personEmpty, err := engine.IsTableEmpty(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personEmpty {
		fmt.Println("表格沒有內容")
	} else {
		fmt.Println("表格有內容")
	}

	// 二.條件查詢

	// 1.ID 查詢
	var person PersonTable
	//select * from person_table where id = 1
	result, err := engine.Id(1).Get(&person)

	fmt.Println(result, err)
	fmt.Println(person.PersonName)

	//2. Where 條件語句查詢
	//var person1 Persontable
	//因為不知道有多少數據會回來所以要放Slice 不過只想回來單一一個的話可PersonTable
	var person2 []PersonTable

	//Select * form person_table where person_age = 26 amd perspn_sex = 2
	err = engine.Where("person_age = ? and person_sex = ?", 26, 2).Find(&person2)
	//find為查詢多數 get只查一個
	fmt.Println(person2)

	/// 3.AND條件查詢
	var persons []PersonTable
	//Select * form person_table where person_age = 26 amd perspn_sex = 2
	err = engine.Where("person_age = ? ", 26).And("person_sex = ?", 2).Find(&persons)
	//And可隨意連接任意數量
	fmt.Println(persons)

	// 4. Or條件查詢
	var personArr []PersonTable
	//Select * form person_table where person_age = 26 amd perspn_sex = 2
	err = engine.Where("person_age = ? ", 26).Or("person_sex = ?", 2).Find(&personArr)
	//Or可隨意連接任意數量
	fmt.Println(persons)

	// 5. 原生SQL查詢 like 語法
	var personNative []PersonTable
	err = engine.SQL("select * from person_table where person_name like 't%'").Find(&personNative)
	fmt.Println(personNative)

	//6. 排序條件查詢
	var personOrderBy []PersonTable
	err = engine.OrderBy("person_age desc").Find(&personOrderBy)
	fmt.Println(personOrderBy)

	//7.  查詢特定字串 select後面不放*
	var personsCols []PersonTable
	err = engine.Cols("person_name", "person_age").Find(&personsCols)
	//fmt.Println(personsCols)
	for _, col := range personsCols {
		fmt.Println(col)
	}

	// 三. ＣＲＵＤ
	//1.Create

	personInsert := PersonTable{PersonAge: 18, PersonName: "Hello", PersonSex: 1}

	rowNum, err := engine.Insert(&personInsert)
	//rowNum, err := engine.InsertOne(&personInsert)
	fmt.Println(rowNum) //RowNum = 插入幾條數據

	//2. Delete
	rowNum, err = engine.Delete(&personInsert)
	fmt.Println(rowNum)

	//3. Update
	rowNum, err = engine.Id(7).Update(&personInsert)
	fmt.Println(rowNum)

	//4.統計 count
	count, err := engine.Count(new(PersonTable))
	fmt.Println("personTable表總計路條數", count)

	//7.事務(session)操作 以防多條數據執行時發生錯誤
	personsArray := []PersonTable{
		PersonTable{PersonName: "Jave", PersonSex: 1, PersonAge: 28},
		PersonTable{PersonName: "XX", PersonAge: 27, PersonSex: 0},
		PersonTable{PersonName: "XX", PersonAge: 27, PersonSex: 0},
	}

	session := engine.NewSession()
	session.Begin()
	for i := 0; i < len(personsArray); i++ {
		_, err = session.Insert(personsArray[i])
		if err != nil {
			session.Rollback() //整個事務回滾 以免前面更 新的有問題
			session.Close()
		}
		err = session.Commit()
		session.Close()
		if err != nil {
			panic(err.Error())
		}
	}

}

type PersonTable struct {
	Id         int64  `xorm:"pk autoincr"`
	PersonName string `xorm:"varchar(24)"`
	PersonAge  int    `xorm:"int default 0"`
	PersonSex  int    `xorm:"notnull -"`
}
