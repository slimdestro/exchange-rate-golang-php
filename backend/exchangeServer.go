/**
	- syncRates: call exchange api
	- frontend: fetches data stored in mysql table by syncRate api
	- Author: Mukul(https://github.com/slimdestro) | https://www.modcode.dev
*/
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	apiKey   = "dd3f9615af335d61738bf3eb7017d4ae"

	// api key is free one so it wont support "https" as per their doc
	baseURL  = "http://api.currencylayer.com/historical"
	tableDDL = `
	CREATE TABLE IF NOT EXISTS exchange_rates (
		id INT AUTO_INCREMENT PRIMARY KEY,
		date DATE,
		currency VARCHAR(10),
		rate FLOAT,
		timestamp INT
	);
	`
)

type ExchangeRatesResponse struct {
	Success   bool              `json:"success"`
	Date      string            `json:"date"`
	Timestamp int64             `json:"timestamp"`
	Source    string            `json:"source"`
	Quotes    map[string]float64 `json:"quotes"`
}

func createTableIfNotExists(db *sql.DB) error {
	_, err := db.Exec(tableDDL)
	return err
}

func fetchExchangeRates(date time.Time) (*ExchangeRatesResponse, error) {
	url := fmt.Sprintf("%s?access_key=%s&date=%s", baseURL, apiKey, date.Format("2006-01-02"))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ratesResponse ExchangeRatesResponse
	err = json.Unmarshal(body, &ratesResponse)
	if err != nil {
		return nil, err
	}

	return &ratesResponse, nil
}

func storeExchangeRates(db *sql.DB, date time.Time, rates map[string]float64, timestamp int64) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO exchange_rates (date, currency, rate, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var rowsAffected int64
	for currency, rate := range rates {
		_, err := stmt.Exec(date.Format("2006-01-02"), currency, rate, timestamp)
		if err != nil {
			return 0, err
		}
		rowsAffected++
	}

	return rowsAffected, nil
}

func syncRates(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:Tetra@2021@tcp(localhost:3306)/akcommodities")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	err = createTableIfNotExists(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// remove all old data before it syncs
	/**
		- reason to truncate the table is to make the script do less memory work. 
		- instead checking each row against id then update, its better to make new entries in this scenario 
		- where data is realtime and only usecase is to show it on front
	*/
	_, err = db.Exec("TRUNCATE exchange_rates")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// >> last 10 days
	for i := 0; i < 10; i++ {
		date := time.Now().AddDate(0, 0, -i)
		ratesResponse, err := fetchExchangeRates(date)
		if err != nil {
			log.Printf("Error fetching exchange rates for %s: %s\n", date.Format("2006-01-02"), err)
			continue
		} 

		rowsAffected, err := storeExchangeRates(db, date, ratesResponse.Quotes, ratesResponse.Timestamp)
		if err != nil {
			log.Printf("Error storing exchange rates for %s: %s\n", date.Format("2006-01-02"), err)
		} else {
			log.Printf("Exchange rates for %s stored successfully. Rows affected: %d\n", date.Format("2006-01-02"), rowsAffected)
		}
	}

	fmt.Fprintln(w, "Table now synced with latest data since 10 days")
}

func frontend(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:Tetra@2021@tcp(localhost:3306)/akcommodities")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT date, currency, rate, timestamp FROM exchange_rates ORDER BY date DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rates []map[string]interface{}
	for rows.Next() {
		var dateStr string

		var currency string
		var rate float64
		var timestamp int64

		err := rows.Scan(&dateStr, &currency, &rate, &timestamp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rates = append(rates, map[string]interface{}{
			"date":      date.Format("2006-01-02"),
			"currency":  currency,
			"rate":      rate,
			"timestamp": timestamp,
		})
	}

	response, err := json.Marshal(rates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}


func main() {
	// syncRates: this will sync table with lates data since last 10 days
	http.HandleFunc("/syncRates", syncRates)
	// frontend: frontend/api.php using this endpoint
	http.HandleFunc("/frontend", frontend)

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
