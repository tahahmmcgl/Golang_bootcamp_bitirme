package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func apiConnection() {
	fmt.Println("api initilazing ....")
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/users_api", users_api) //if sign in post method you can add user
	r.HandleFunc("/information/{user}/{password}", userInfo)
	r.HandleFunc("/information/{user}/{password}/{process}", userInfo)
	r.HandleFunc("/informationMarka/{marka}/{idKullanici}/{idMarka}/{process}", markaInfo)
	r.HandleFunc("/informationPatent/{patent}/{idKullanici}/{idPatent}/{process}", patentInfo)

	r.HandleFunc("/adminLogin", adminLogin)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Write([]byte("İndexteyiz"))
}
func users_api(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)
	w.Write([]byte("Users apideyiz"))
	adminLogin(w, r)

	if r.Method == "POST" {

		fmt.Println("Post Metodu")
		w.Header().Set("Content-Type", "application/json")

		log.Println("usernameLogin :", r.FormValue("usernameLogin"))
		log.Println("passwordLogin :", r.FormValue("passwordLogin"))
		log.Println("emailLogin :", r.FormValue("emailLogin"))
		fmt.Println("ekleniyor")
		addKullanici(r.FormValue("usernameLogin"), r.FormValue("passwordLogin"), r.FormValue("emailLogin"))
	} else if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		jsonString, err := json.Marshal(kullaniciList)

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, string(jsonString))
	}

}
func adminLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	if r.Method == ("POST") {
		fmt.Fprintln(w, "admin kullanıcı adınızı ve şifrenizi girin")
		//log.Println("usernameLogin :", r.FormValue("usernameLogin"))
		//log.Println("passwordLogin :", r.FormValue("passwordLogin"))
		if admins(r.FormValue("usernameLogin"), r.FormValue("passwordLogin")) == true {
			fmt.Println("HOŞGELDİNİZ " + r.FormValue("usernameLogin"))
			w.Write([]byte(("HOŞGELDİNİZ " + r.FormValue("usernameLogin"))))
		} else {
			fmt.Fprintln(w, string("kullanıcı adı yada şifre yanlış"))
		}
	} else {
		w.Write([]byte("Post metoduyla giriş yapın"))
	}
}
func userInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)

		user := vars["user"]
		password := vars["password"]
		process := vars["process"]

		if process == "get_user" {

			w.Header().Set("Content-Type", "application/json")
			jsonString, err := json.Marshal(infoUsers(user, password))

			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(jsonString))

		} else if process == "delete_user" {

			fmt.Fprintln(w, deleteKullanici(user)+" kadar kişi silindi")
		} else {
			fmt.Fprintln(w, "diğer işlemler için post methodu kullanınız")
			fmt.Fprintln(w, "işlemler = update_user , add user ")

		}

	} else if r.Method == "POST" {

		vars := mux.Vars(r)

		user := vars["user"]
		password := vars["password"]

		w.Header().Set("Content-Type", "application/json")
		jsonString, err := json.Marshal(infoUsers(user, password))

		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, string(jsonString))

		process := vars["process"]
		if process == "update_user" {

			idUser, err := strconv.Atoi(r.FormValue("idLogin"))
			checkError(err)
			fmt.Fprintln(w, "kullanıcı güncellendi "+updateKullanici(idUser, user, r.FormValue("password"), r.FormValue("emailLogin")))
		} else if process == "add_user" {

			addKullanici(user, password, r.FormValue("emailLogin"))

			fmt.Fprintln(w, "eklendi")
			jsonString, err := json.Marshal(infoUsers(user, password))

			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(jsonString))
		}

	}

}
func markaInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	marka := vars["marka"]
	idKullanici, err := strconv.Atoi(vars["idKullanici"])
	checkError(err)
	idMarka, err := strconv.Atoi(vars["idMarka"])
	checkError(err)
	process := vars["process"]

	if r.Method == "GET" {
		if process == "get_marka" {
			w.Header().Set("Content-Type", "application/json")
			jsonString, err := json.Marshal(getonemarks(idMarka))

			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(jsonString))
		} else if process == "delete_marka" {
			fmt.Fprintf(w, "silinecek marka =")
			w.Header().Set("Content-Type", "application/json")
			jsonString, err := json.Marshal(getonemarks(idMarka))

			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(jsonString))

			fmt.Fprintln(w, "Silinen marka sayısı", deleteMark(idMarka))

		} else if r.Method == "POST" {
			if process == "add_marka" {

				idkullanici, err := strconv.Atoi(r.FormValue("idkullanici"))
				checkError(err)
				fmt.Fprintln(w, addMark(r.FormValue("markaname"), idkullanici), " adet marka eklendi")

			} else if process == "update_marka" {
				taha, err := strconv.Atoi(r.FormValue("idmarka"))
				checkError(err)
				fmt.Fprintln(w, "Marka Update edildi", updateMark(idKullanici, marka, taha))
			}

		}

	}
}
func patentInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patent := vars["patent"]
	idKullanici, err := strconv.Atoi(vars["idKullanici"])
	checkError(err)
	idpatent, err := strconv.Atoi(vars["idPatent"])
	checkError(err)
	process := vars["process"]

	if r.Method == "GET" {
		if process == "get_patent" {
			w.Header().Set("Content-Type", "application/json")
			jsonString, err := json.Marshal(getonepatent(idpatent))

			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(jsonString))
		} else if process == "delete_patent" {
			fmt.Fprintf(w, "silinecek patent =")
			w.Header().Set("Content-Type", "application/json")
			jsonString, err := json.Marshal(getonepatent(idpatent))

			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(jsonString))

			fmt.Fprintln(w, "Silinen patent sayısı", deletePatent(idpatent))

		} else if r.Method == "POST" {
			if process == "add_patent" {

				idkullanici, err := strconv.Atoi(r.FormValue("idkullanici"))
				checkError(err)
				fmt.Fprintln(w, addMark(r.FormValue("patentname"), idkullanici), " adet patent eklendi")

			} else if process == "update_patent" {
				taha, err := strconv.Atoi(r.FormValue("idpatent"))
				checkError(err)
				fmt.Fprintln(w, "patent Update edildi", updateMark(idKullanici, patent, taha))
			}

		}

	}
}
