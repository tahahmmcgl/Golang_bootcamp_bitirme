package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Kullanici struct {
	IdLogin       int
	UsernameLogin string
	PasswordLogin string
}

func DatabaseConnection() {
	fmt.Println("databasedeyiz")

	db, err := sql.Open("mysql", "root:19981998aA.@(127.0.0.1:3306)/graduationschemas?parseTime=true")
	checkError(err)
	_ = db

	//tabloCreate(db)
	//addKullanici(db, "Rana", "12345")
	//tabloCreate(db)
	//getUsers(db)
	//deleteKullanici(db, 4)
	getUsers(db)
	updateKullanici(db, "mustafa", "79965", 6)

}

func tabloCreate(db *sql.DB) {
	fmt.Println("tablo createdeyiz")
	db.Exec("create table if not exists users(username text, password text)") //eğer yoksa bir tablo oluşturduk

}

var kullaniciList []Kullanici

func getUsers(db *sql.DB) {
	fmt.Println("get usersdeyiz")
	kullaniciList = nil
	rows, err := db.Query("Select * from graduationschemas.tbllogin")
	if err == nil {
		for rows.Next() {
			var idLogin int
			var usernameLogin string
			var passwordLogin string
			err = rows.Scan(&idLogin, &usernameLogin, &passwordLogin)
			if err == nil {
				kullaniciList = append(kullaniciList, Kullanici{IdLogin: idLogin, UsernameLogin: usernameLogin, PasswordLogin: passwordLogin})
			} else {
				fmt.Println(err)
				return
			}
		}
	} else {
		fmt.Println(err)
	}
	rows.Close()
	fmt.Println(kullaniciList)
}

func addKullanici(db *sql.DB, usernameLogin, passwordLogin string) {

	fmt.Println("add  kullanıcdayız")
	stmt, err := db.Prepare("INSERT INTO  graduationschemas.tbllogin(usernameLogin, passwordLogin) values (?, ?)")
	checkError(err)
	res, err := stmt.Exec(usernameLogin, passwordLogin)
	checkError(err)
	id, err := res.LastInsertId()
	checkError(err)
	fmt.Println("son eklenen kayıdın id : ", id)
	getUsers(db)
	fmt.Println("eleman eklendi")
}

func deleteKullanici(db *sql.DB, idLogin int) {

	fmt.Println("delete  kullanıcdayız")
	stmt, err := db.Prepare("DELETE FROM  graduationschemas.tbllogin where idLogin=?")
	checkError(err)
	res, err := stmt.Exec(idLogin)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	fmt.Println("Silinen Kayıdın id : ", id)
	getUsers(db)
	fmt.Println("eleman eklendi")
}

func updateKullanici(db *sql.DB, usernameLogin, passwordLogin string, idLogin int) {

	fmt.Println("update  kullanıcdayız")
	stmt, err := db.Prepare("update graduationschemas.tbllogin set usernameLogin=?, passwordLogin=? where idLogin=?")
	checkError(err)
	res, err := stmt.Exec(usernameLogin, passwordLogin, idLogin)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	fmt.Println("veri güncellendi : ", id)
	getUsers(db)
	fmt.Println("eleman eklendi")
}
