package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

type PingData struct {
	IP         string    `json:"ip"`
	Time       int       `json:"time"`
	LastActive time.Time `json:"last_active"`
}

func initDB() {
	var err error
	connStr := "postgres://user:password@postgres:5432/user?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Проверяем подключение к БД
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка пинга БД:", err)
	}

	query := `CREATE TABLE IF NOT EXISTS pings (
		ip TEXT PRIMARY KEY, 
		time INT, 
		last_active TIMESTAMP
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
}

func savePing(w http.ResponseWriter, r *http.Request) {
	var p PingData
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec(`
		INSERT INTO pings (ip, time, last_active) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (ip) DO UPDATE 
		SET time = $2, last_active = $3
	`, p.IP, p.Time, p.LastActive)

	if err != nil {
		http.Error(w, "Ошибка сохранения в БД: "+err.Error(), http.StatusInternalServerError)
	}
}

func getPings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT ip, time, last_active FROM pings`)
	if err != nil {
		http.Error(w, "Ошибка запроса к БД: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pings []PingData
	for rows.Next() {
		var p PingData
		err := rows.Scan(&p.IP, &p.Time, &p.LastActive)
		if err != nil {
			http.Error(w, "Ошибка сканирования строки: "+err.Error(), http.StatusInternalServerError)
			return
		}
		pings = append(pings, p)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(pings)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	initDB()
	defer db.Close() // Закрываем подключение при завершении

	r := mux.NewRouter()
	r.HandleFunc("/ping", savePing).Methods("POST")
	r.HandleFunc("/pings", getPings).Methods("GET")

	log.Println("Backend running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
