package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Define a struct to hold your data
type RowData struct {
    Date     string  `json:"date"`
    Open     float64 `json:"open"`
    High     float64 `json:"high"`
    Low      float64 `json:"low"`
    Close    float64 `json:"close"`
    AdjClose float64 `json:"adjClose"`
    Volume   float64 `json:"volume"`
}

// Function to fetch and parse CSV data
func fetchCSVData(url string) ([]RowData, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    reader := csv.NewReader(resp.Body)
    var data []RowData

    // Skip the header row
    if _, err := reader.Read(); err != nil {
        return nil, err
    }

    // Read each record
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }

        open, _ := strconv.ParseFloat(record[1], 64)
        high, _ := strconv.ParseFloat(record[2], 64)
        low, _ := strconv.ParseFloat(record[3], 64)
        close, _ := strconv.ParseFloat(record[4], 64)
        adjClose, _ := strconv.ParseFloat(record[5], 64)
        volume, _ := strconv.ParseFloat(record[6], 64)

        data = append(data, RowData{
            Date:     record[0],
            Open:     open,
            High:     high,
            Low:      low,
            Close:    close,
            AdjClose: adjClose,
            Volume:   volume,
        })
    }
    return data, nil
}

func bitcoinDataHandler(w http.ResponseWriter, r *http.Request) {
    // Use the URL of your CSV data
    w.Header().Set("Access-Control-Allow-Origin", "*") // Cho phép tất cả các domain
    data, err := fetchCSVData("https://raw.githubusercontent.com/Chaudeptrai123456/DataViz/master/BTC-USD.csv")
    if err != nil {
        http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/api/bitcoin-history", bitcoinDataHandler)
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
