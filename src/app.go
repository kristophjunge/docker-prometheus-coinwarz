package main

import (
    "io"
    "net/http"
    "log"
    "os"
    "strconv"
    "io/ioutil"
    "encoding/json"
    "errors"
)

const LISTEN_ADDRESS = ":9206"
const API_URL = "http://www.coinwarz.com/v1/api/profitability/"

var apiKey string
var testMode string

type CoinWarzMiningProfitability struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data []struct {
        CoinName string `json:"CoinName"`
        CoinTag string `json:"CoinTag"`
        Algorithm string `json:"Algorithm"`
        Difficulty float64 `json:"Difficulty"`
        BlockReward int64 `json:"BlockReward"`
        BlockCount int64 `json:"BlockCount"`
        ProfitRatio float64 `json:"ProfitRatio"`
        AvgProfitRatio float64 `json:"AvgProfitRatio"`
        Exchange string `json:"Exchange"`
        ExchangeRate float64 `json:"ExchangeRate"`
        ExchangeVolume float64 `json:"ExchangeVolume"`
        IsBlockExplorerOnline bool `json:"IsBlockExplorerOnline"`
        IsExchangeOnline bool `json:"IsExchangeOnline"`
        Message string `json:"Message"`
        BlockTimeInSeconds int64 `json:"BlockTimeInSeconds"`
        HealthStatus string `json:"HealthStatus"`
    } `json:"data"`
}

func integerToString(value int64) string {
    return strconv.FormatInt(value, 16)
}

func floatToString(value float64, precision int64) string {
    return strconv.FormatFloat(value, 'f', int(precision), 64)
}

func booleanToString(value bool) string {
    if value {
        return "1"
    }
    return "0"
}

func stringToInteger(value string) int64 {
    if value == "" {
        return 0
    }
    result, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
        log.Fatal(err)
    }
    return result
}

func stringToFloat(value string) float64 {
    if value == "" {
        return 0
    }
    result, err := strconv.ParseFloat(value, 64)
    if err != nil {
        log.Fatal(err)
    }
    return result
}

func formatValue(key string, meta string, value string) string {
    result := key;
    if (meta != "") {
        result += "{" + meta + "}";
    }
    result += " "
    result += value
    result += "\n"
    return result
}

func queryData() (string, error) {
    var err error

    // Build URL
    url := API_URL + "?apikey=" + apiKey + "&algo=all"

    // Perform HTTP request
    resp, err := http.Get(url);
    if err != nil {
        return "", err;
    }

    // Parse response
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        return "", errors.New("HTTP returned code " + integerToString(int64(resp.StatusCode)))
    }
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    if err != nil {
        return "", err;
    }

    return bodyString, nil;
}

func getTestData() (string, error) {
    dir, err := os.Getwd()
    if err != nil {
        return "", err;
    }
    body, err := ioutil.ReadFile(dir + "/test.json")
    if err != nil {
        return "", err;
    }
    return string(body), nil
}

func metrics(w http.ResponseWriter, r *http.Request) {
    log.Print("Serving /metrics")

    var up int64 = 1
    var jsonString string
    var err error

    if (testMode == "1") {
        jsonString, err = getTestData()
    } else {
        jsonString, err = queryData()
    }
    if err != nil {
        log.Print(err)
        up = 0
    }

    // Parse JSON
    jsonData := CoinWarzMiningProfitability{}
    json.Unmarshal([]byte(jsonString), &jsonData)

    // Output
    io.WriteString(w, formatValue("coinwarz_up", "", integerToString(up)))

    log.Print(jsonData)

    for _, Coin := range jsonData.Data {

        //Coin.CoinName
        //Coin.CoinTag
        //Coin.Algorithm
        floatToString(Coin.Difficulty, 8)
        //integerToString(Coin.BlockReward)
        //integerToString(Coin.BlockCount)
        floatToString(Coin.ProfitRatio, 15)
        floatToString(Coin.AvgProfitRatio, 15)
        //Coin.Exchange
        //floatToString(Coin.ExchangeRate, 8)
        //floatToString(Coin.ExchangeVolume, 8)
        //booleanToString(Coin.IsBlockExplorerOnline)
        //booleanToString(Coin.IsExchangeOnline)
        //Coin.Message
        //integerToString(Coin.BlockTimeInSeconds)
        //Coin.HealthStatus

        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))
        io.WriteString(w, formatValue("coinwarz_", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", ))

    }
}

func index(w http.ResponseWriter, r *http.Request) {
    log.Print("Serving /index")
    html := `<!doctype html>
<html>
    <head>
        <meta charset="utf-8">
        <title>CoinWarz Exporter</title>
    </head>
    <body>
        <h1>CoinWarz Exporter</h1>
        <p><a href="/metrics">Metrics</a></p>
    </body>
</html>`
    io.WriteString(w, html)
}

func main() {
    testMode = os.Getenv("TEST_MODE")
    if (testMode == "1") {
        log.Print("Test mode is enabled")
    }

    apiKey = os.Getenv("API_KEY")

    log.Print("CoinWarz exporter listening on " + LISTEN_ADDRESS)
    http.HandleFunc("/", index)
    http.HandleFunc("/metrics", metrics)
    http.ListenAndServe(LISTEN_ADDRESS, nil)
}
