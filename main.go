package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type BitcoinData struct {
    Time  time.Time `json:"time"`
    Price float64   `json:"price"`
}

func fetchBitcoinData() ([]BitcoinData, error) {
    apiUrl := "https://api.coingecko.com/api/v3/coins/bitcoin/market_chart?vs_currency=usd&days=365"
    resp, err := http.Get(apiUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Prices [][]float64 `json:"prices"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    var data []BitcoinData
    for _, item := range result.Prices {
        data = append(data, BitcoinData{
            Time:  time.Unix(int64(item[0]/1000), 0),
            Price: item[1],
        })
    }

    return data, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    data, err := fetchBitcoinData()
    if err != nil {
        http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
        return
    }

    jsonData, _ := json.Marshal(data)

    htmlContent := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Bitcoin Historical Data</title>
        <script type="text/javascript" src="https://cdn.canvasjs.com/canvasjs.stock.min.js"></script>
    </head>
    <body>
        <h1>Bitcoin Historical Data Visualization</h1>
        <div id="chartContainer" style="height: 370px; width: 100%;"></div>
        <script>
            var dataPoints = ` + string(jsonData) + `;

            var chartData = dataPoints.map(function(item) {
                return { x: new Date(item.time), y: item.price };
            });

            var chart = new CanvasJS.StockChart("chartContainer", {
                title: {
                    text: "Bitcoin Price History"
                },
                charts: [{
                    data: [{
                        type: "line",
                        dataPoints: chartData
                    }]
                }],
                rangeSelector: {
                    enabled: true
                }
            });
            chart.render();
        </script>
    </body>
    </html>`

    fmt.Fprint(w, htmlContent)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
