package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	//連接數據庫
	//連接數據庫 規定的字符拼接  用戶名稱:用戶密碼@tcp(127.0.0.1)/數據庫名稱
	connStr := "root:sofia25037245@tcp(127.0.0.1:3306)/ginsql"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	//DB CRUD
	//創建數據庫
	// person:id,name,age
	//_, err = db.Exec("create table person(\n\t\tid int auto_increment primary key, \n\t\tname varchar(12) not null,\n\t\tage int default 1 );")

	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("數據庫成功")
	}

	//插入數據到數據庫
	_, err = db.Exec("insert into person(name,age) VAlUES(\"davie\", 18);")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("數據插入成功")
	}

	//查詢數據庫
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//Davie 18
	//JACk 20
	//lili 25
scan:
	if rows.Next() {
		person := new(Person)
		rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}

}

type Person struct {
	Id   int
	Name string
	Age  int
}

