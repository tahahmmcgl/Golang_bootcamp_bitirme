package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Admin struct {
	AdminId       int
	AdminName     string
	AdminPassword string
}

type PatentStr struct {
	PatentId   int
	PatentName string
}
type MarkaStr struct {
	MarkaId   int
	MarkaName string
}

type Kullanici struct {
	IdLogin       int
	UsernameLogin string
	PasswordLogin string
	EmailLogin    string
	Marka         []MarkaStr
	Patent        []PatentStr
}

var db *sql.DB

func DatabaseConnection() {
	fmt.Println("databasedeyiz")
	var err error
	db, err = sql.Open("mysql", "root:19981998aA.@(127.0.0.1:3306)/graduationschemas?parseTime=true")
	checkError(err)
	_ = db

}

func tabloCreate(db *sql.DB) {
	fmt.Println("tablo createdeyiz")
	db.Exec("create table if not exists users(username text, password text)") //eğer yoksa bir tablo oluşturduk

}

var kullaniciList []Kullanici
var markaList []MarkaStr
var patentList []PatentStr

/*
users
*/
func getUsers() {
	fmt.Println("get usersdeyiz")
	kullaniciList = nil
	rows, err := db.Query("Select * from graduationschemas.tbllogin")
	if err == nil {
		for rows.Next() {

			var idLogin int
			var usernameLogin string
			var passwordLogin string
			var emailLogin string

			var err = rows.Scan(&idLogin, &usernameLogin, &passwordLogin, &emailLogin)
			if err == nil {

				kullaniciList = append(kullaniciList, Kullanici{IdLogin: idLogin, UsernameLogin: usernameLogin, PasswordLogin: passwordLogin, EmailLogin: emailLogin, Marka: getmarks(idLogin), Patent: getpatents(idLogin)})
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
func infoUsers(user, password string) Kullanici {
	fmt.Println("info usersdeyiz")
	kullaniciList = nil
	rows, err := db.Query("Select * from graduationschemas.tbllogin")
	if err == nil {
		for rows.Next() {

			var idLogin int
			var usernameLogin string
			var passwordLogin string
			var emailLogin string

			var err = rows.Scan(&idLogin, &usernameLogin, &passwordLogin, &emailLogin)
			if err == nil {

				kullaniciList = append(kullaniciList, Kullanici{IdLogin: idLogin, UsernameLogin: usernameLogin, PasswordLogin: passwordLogin, EmailLogin: emailLogin, Marka: getmarks(idLogin), Patent: getpatents(idLogin)})
			} else {
				fmt.Println(err)

			}
		}
	} else {
		fmt.Println(err)
	}
	rows.Close()
	var zero Kullanici
	for _, k := range kullaniciList {
		if user == k.UsernameLogin && password == k.PasswordLogin {
			return k
		}

	}
	return zero
}

func addKullanici(usernameLogin, passwordLogin, emailLogin string) {

	fmt.Println("add  kullanıcdayız")
	stmt, err := db.Prepare("INSERT INTO  graduationschemas.tbllogin(usernameLogin, passwordLogin, emaillogin) values (?, ?,?)")
	checkError(err)
	res, err := stmt.Exec(usernameLogin, passwordLogin, emailLogin)
	checkError(err)
	id, err := res.LastInsertId()
	checkError(err)
	fmt.Println("son eklenen kayıdın id : ", id)
	getUsers()
	fmt.Println("eleman eklendi")
}

func deleteKullanici(usernamelogin string) string {

	fmt.Println("delete  kullanıcdayız")
	stmt, err := db.Prepare("DELETE FROM  graduationschemas.tbllogin where usernameLogin=?")
	checkError(err)
	res, err := stmt.Exec(usernamelogin)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	getUsers()
	return strconv.Itoa(int(id))

}

func updateKullanici(idLogin int, usernameLogin, passwordLogin, emailLogin string) string {

	fmt.Println("update  kullanıcdayız")
	stmt, err := db.Prepare("update graduationschemas.tbllogin set usernameLogin=?, passwordLogin=?, emailLogin=? where idLogin=?")
	checkError(err)
	res, err := stmt.Exec(usernameLogin, passwordLogin, emailLogin, idLogin)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	fmt.Println("veri güncellendi : ", id)
	getUsers()
	return strconv.Itoa(int(id))
}

func getmarks(kullaniciId int) []MarkaStr {
	taha := strconv.Itoa(kullaniciId)
	markaList = nil
	rows, err := db.Query("Select idmarka, markaName from tblmarka where KullaniciId=" + taha)
	checkError(err)
	if err == nil {
		for rows.Next() {

			var idmarka int
			var markaName string

			var err = rows.Scan(&idmarka, &markaName)

			if err == nil {
				markaList = append(markaList, MarkaStr{MarkaId: idmarka, MarkaName: markaName})
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)

	}
	rows.Close()
	return markaList
}
func getonemarks(idMarka int) []MarkaStr {
	markaList = nil
	taha := strconv.Itoa(idMarka)
	rows, err := db.Query("Select idmarka, markaName from tblmarka where idmarka=" + taha)
	checkError(err)
	if err == nil {
		for rows.Next() {

			var idmarka int
			var markaName string

			var err = rows.Scan(&idmarka, &markaName)

			if err == nil {
				markaList = append(markaList, MarkaStr{MarkaId: idmarka, MarkaName: markaName})
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)

	}
	rows.Close()
	return markaList
}

/*
MARKA EKLEME ÇIKARMA GÜNCELLEME
*/

func addMark(markaname string, kullaniciId int) int {

	stmt, err := db.Prepare("INSERT INTO  graduationschemas.tblmarka(markaName, kullaniciId) values (?, ?)")
	checkError(err)
	res, err := stmt.Exec(markaname, kullaniciId)
	checkError(err)
	id, err := res.LastInsertId()
	checkError(err)
	fmt.Println("son eklenen kayıdın id : ", id)

	return int(id)

}
func deleteMark(idmarka int) int {

	fmt.Println("delete  marka")
	stmt, err := db.Prepare("DELETE FROM  graduationschemas.tblmarka where idmarka=?")
	checkError(err)
	res, err := stmt.Exec(idmarka)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	fmt.Println("Silinen Kayıdın id : ", id)
	//getmarks()
	return int(id)
}

func updateMark(kullaniciId int, markaname string, idmarka int) int {

	fmt.Println("update  marka")
	stmt, err := db.Prepare("update graduationschemas.tblmarka set kullaniciId=?, markaName=? where idmarka=?")
	checkError(err)
	res, err := stmt.Exec(kullaniciId, markaname, idmarka)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	fmt.Println("veri güncellendi : ", id)
	getUsers()
	fmt.Println("eleman eklendi")
	return int(id)
}

/*


	Patent ekleme çıkarma güncelleme


*/

func getpatents(kullaniciid int) []PatentStr {
	taha := strconv.Itoa(kullaniciid)
	patentList = nil
	rows, err := db.Query("Select idPatent,patentName from tblpatent where userId=" + taha)
	checkError(err)
	if err == nil {
		for rows.Next() {

			var idpatent int
			var patentname string

			var err = rows.Scan(&idpatent, &patentname)

			if err == nil {
				patentList = append(patentList, PatentStr{PatentId: idpatent, PatentName: patentname})
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)

	}
	rows.Close()
	return patentList
}
func getonepatent(idpatent int) []PatentStr {
	patentid := strconv.Itoa(idpatent)
	patentList = nil
	rows, err := db.Query("Select idPatent,patentName from tblpatent where idPatent=" + patentid)
	checkError(err)
	if err == nil {
		for rows.Next() {

			var idpatent int
			var patentname string

			var err = rows.Scan(&idpatent, &patentname)

			if err == nil {
				patentList = append(patentList, PatentStr{PatentId: idpatent, PatentName: patentname})
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)

	}
	rows.Close()
	return patentList
}
func addPatent(patentName string, userId int) int {

	stmt, err := db.Prepare("INSERT INTO  graduationschemas.tblpatent(patentName, userId) values (?, ?)")
	checkError(err)
	res, err := stmt.Exec(patentName, userId)
	checkError(err)
	id, err := res.LastInsertId()
	checkError(err)
	fmt.Println("son eklenen kayıdın id : ", id)
	return int(id)
}
func deletePatent(idpatent int) int {

	fmt.Println("delete  patent")
	stmt, err := db.Prepare("DELETE FROM  graduationschemas.tblpatent where idPatent=?")
	checkError(err)
	res, err := stmt.Exec(idpatent)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	fmt.Println("Silinen Kayıdın id : ", id)
	getonepatent(idpatent)
	return int(id)

}

func updatePatent(userid int, patentname string, idpatent int) {

	fmt.Println("update  kullanıcdayız")
	stmt, err := db.Prepare("update graduationschemas.tblpatent set userId=?, patentName=? where idPatent=?")
	checkError(err)
	res, err := stmt.Exec(userid, patentname, idpatent)
	checkError(err)
	id, err := res.RowsAffected()
	checkError(err)
	fmt.Println("veri güncellendi : ", id)
	getUsers()
}

/*
	admins

*/
var adminList []Admin

func admins(name_admin string, password_admin string) bool {
	rows, err := db.Query("select * from tbladmin")
	if err == nil {
		for rows.Next() {
			var adminId int
			var adminName string
			var adminPassword string
			var err = rows.Scan(&adminId, &adminName, &adminPassword)
			if err == nil {
				adminList = append(adminList, Admin{AdminId: adminId, AdminName: adminName, AdminPassword: adminPassword})
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
	}
	rows.Close()
	for _, k := range adminList {
		if name_admin == k.AdminName && password_admin == k.AdminPassword {
			return true
		}
	}
	return false
}
