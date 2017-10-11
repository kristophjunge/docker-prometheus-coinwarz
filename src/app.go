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
var sha256HashRate string
var sha256Power string
var sha256PowerCost string
var scryptHashRate string
var scryptPower string
var scryptPowerCost string
var x11HashRate string
var x11Power string
var x11PowerCost string
var quarkHashRate string
var quarkPower string
var quarkPowerCost string
var groestlHashRate string
var groestlPower string
var groestlPowerCost string
var blake256HashRate string
var blake256Power string
var blake256PowerCost string
var neoscryptHashRate string
var neoscryptPower string
var neoscryptPowerCost string
var lyra2rev2HashRate string
var lyra2rev2Power string
var lyra2rev2PowerCost string
var cryptonightHashRate string
var cryptonightPower string
var cryptonightPowerCost string
var ethashHashRate string
var ethashPower string
var ethashPowerCost string
var equihashHashRate string
var equihashPower string
var equihashPowerCost string

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
    return strconv.FormatInt(value, 10)
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
    if sha256HashRate != "" {
        url += "&sha256HashRate=" + sha256HashRate
    }
    if sha256Power != "" {
        url += "&sha256Power=" + sha256Power
    }
    if sha256PowerCost != "" {
        url += "&sha256PowerCost=" + sha256PowerCost
    }
    if scryptHashRate != "" {
        url += "&scryptHashRate=" + scryptHashRate
    }
    if scryptPower != "" {
        url += "&scryptPower=" + scryptPower
    }
    if scryptPowerCost != "" {
        url += "&scryptPowerCost=" + scryptPowerCost
    }
    if x11HashRate != "" {
        url += "&x11HashRate=" + x11HashRate
    }
    if x11Power != "" {
        url += "&x11Power=" + x11Power
    }
    if x11PowerCost != "" {
        url += "&x11PowerCost=" + x11PowerCost
    }
    if quarkHashRate != "" {
        url += "&quarkHashRate=" + quarkHashRate
    }
    if quarkPower != "" {
        url += "&quarkPower=" + quarkPower
    }
    if quarkPowerCost != "" {
        url += "&quarkPowerCost=" + quarkPowerCost
    }
    if groestlHashRate != "" {
        url += "&groestlHashRate=" + groestlHashRate
    }
    if groestlPower != "" {
        url += "&groestlPower=" + groestlPower
    }
    if groestlPowerCost != "" {
        url += "&groestlPowerCost=" + groestlPowerCost
    }
    if blake256HashRate != "" {
        url += "&blake256HashRate=" + blake256HashRate
    }
    if blake256Power != "" {
        url += "&blake256Power=" + blake256Power
    }
    if blake256PowerCost != "" {
        url += "&blake256PowerCost=" + blake256PowerCost
    }
    if neoscryptHashRate != "" {
        url += "&neoscryptHashRate=" + neoscryptHashRate
    }
    if neoscryptPower != "" {
        url += "&neoscryptPower=" + neoscryptPower
    }
    if neoscryptPowerCost != "" {
        url += "&neoscryptPowerCost=" + neoscryptPowerCost
    }
    if lyra2rev2HashRate != "" {
        url += "&lyra2rev2HashRate=" + lyra2rev2HashRate
    }
    if lyra2rev2Power != "" {
        url += "&lyra2rev2Power=" + lyra2rev2Power
    }
    if lyra2rev2PowerCost != "" {
        url += "&lyra2rev2PowerCost=" + lyra2rev2PowerCost
    }
    if cryptonightHashRate != "" {
        url += "&cryptonightHashRate=" + cryptonightHashRate
    }
    if cryptonightPower != "" {
        url += "&cryptonightPower=" + cryptonightPower
    }
    if cryptonightPowerCost != "" {
        url += "&cryptonightPowerCost=" + cryptonightPowerCost
    }
    if ethashHashRate != "" {
        url += "&ethashHashRate=" + ethashHashRate
    }
    if ethashPower != "" {
        url += "&ethashPower=" + ethashPower
    }
    if ethashPowerCost != "" {
        url += "&ethashPowerCost=" + ethashPowerCost
    }
    if equihashHashRate != "" {
        url += "&equihashHashRate=" + equihashHashRate
    }
    if equihashPower != "" {
        url += "&equihashPower=" + equihashPower
    }
    if equihashPowerCost != "" {
        url += "&equihashPowerCost=" + equihashPowerCost
    }

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
    if !jsonData.Success {
        up = 0
    }
    io.WriteString(w, formatValue("coinwarz_up", "", integerToString(up)))

    for _, Coin := range jsonData.Data {
        up = 1
        if Coin.HealthStatus != "Healthy" {
            log.Print("Status of '" + Coin.CoinName + "' is '" + Coin.HealthStatus + "': " + Coin.Message);
            up = 0
        }

        io.WriteString(w, formatValue("coinwarz_coin_up", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", integerToString(up)))
        io.WriteString(w, formatValue("coinwarz_difficulty", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", floatToString(Coin.Difficulty, 8)))
        io.WriteString(w, formatValue("coinwarz_block_reward", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", integerToString(Coin.BlockReward)))
        io.WriteString(w, formatValue("coinwarz_block_count", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", integerToString(Coin.BlockCount)))
        io.WriteString(w, formatValue("coinwarz_profit_ratio", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", floatToString(Coin.ProfitRatio, 15)))
        io.WriteString(w, formatValue("coinwarz_avg_profit_ratio", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", floatToString(Coin.AvgProfitRatio, 15)))
        io.WriteString(w, formatValue("coinwarz_block_time_in_seconds", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", integerToString(Coin.BlockTimeInSeconds)))
        io.WriteString(w, formatValue("coinwarz_exchange_rate", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", floatToString(Coin.ExchangeRate, 8)))
        io.WriteString(w, formatValue("coinwarz_exchange_volume", "name=\"" + Coin.CoinName + "\",symbol=\"" + Coin.CoinTag + "\"", floatToString(Coin.ExchangeVolume, 8)))
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
    sha256HashRate = os.Getenv("HA256_HASHRATE")
    sha256Power = os.Getenv("SHA256_POWER")
    sha256PowerCost = os.Getenv("SHA256_POWERCOST")
    scryptHashRate = os.Getenv("SCRYPT_HASHRATE")
    scryptPower = os.Getenv("SCRYPT_POWER")
    scryptPowerCost = os.Getenv("SCRYPT_POWERCOST")
    x11HashRate = os.Getenv("X11_HASHRATE")
    x11Power = os.Getenv("X11_POWER")
    x11PowerCost = os.Getenv("X11_POWERCOST")
    quarkHashRate = os.Getenv("QUARK_HASHRATE")
    quarkPower = os.Getenv("QUARK_POWER")
    quarkPowerCost = os.Getenv("QUARK_POWERCOST")
    groestlHashRate = os.Getenv("GROESTL_HASHRATE")
    groestlPower = os.Getenv("GROESTL_POWER")
    groestlPowerCost = os.Getenv("GROESTL_POWERCOST")
    blake256HashRate = os.Getenv("BLAKE256_HASHRATE")
    blake256Power = os.Getenv("BLAKE256_POWER")
    blake256PowerCost = os.Getenv("BLAKE256_POWERCOST")
    neoscryptHashRate = os.Getenv("NEOSCRYPT_HASHRATE")
    neoscryptPower = os.Getenv("NEOSCRYPT_POWER")
    neoscryptPowerCost = os.Getenv("NEOSCRYPT_POWERCOST")
    lyra2rev2HashRate = os.Getenv("LYRA2REV2_HASHRATE")
    lyra2rev2Power = os.Getenv("LYRA2REV2_POWER")
    lyra2rev2PowerCost = os.Getenv("LYRA2REV2_POWERCOST")
    cryptonightHashRate = os.Getenv("CRYPTONIGHT_HASHRATE")
    cryptonightPower = os.Getenv("CRYPTONIGHT_POWER")
    cryptonightPowerCost = os.Getenv("CRYPTONIGHT_POWERCOST")
    ethashHashRate = os.Getenv("ETHASH_HASHRATE")
    ethashPower = os.Getenv("ETHASH_POWER")
    ethashPowerCost = os.Getenv("ETHASH_POWERCOST")
    equihashHashRate = os.Getenv("EQUIHASH_HASHRATE")
    equihashPower = os.Getenv("EQUIHASH_POWER")
    equihashPowerCost = os.Getenv("EQUIHASH_POWERCOST")

    log.Print("CoinWarz exporter listening on " + LISTEN_ADDRESS)
    http.HandleFunc("/", index)
    http.HandleFunc("/metrics", metrics)
    http.ListenAndServe(LISTEN_ADDRESS, nil)
}
