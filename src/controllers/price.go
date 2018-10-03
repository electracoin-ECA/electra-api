package controllers

import (
	"net/http"

	"github.com/Electra-project/electra-api/src/helpers"
	"github.com/gin-gonic/gin"
	cache "github.com/patrickmn/go-cache"
)

type requestResponseData []requestResponseDataEntry

type requestResponseDataEntry struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	Rank                 string `json:"rank"`
	PriceUSD             string `json:"price_usd"`
	PriceBTC             string `json:"price_btc"`
	DayVolumeUSD         string `json:"24h_volume_usd"`
	MarketCapUSD         string `json:"market_cap_usd"`
	AvailableSupply      string `json:"available_supply"`
	TotalSupply          string `json:"total_supply"`
	MaxSupply            string `json:"max_supply"`
	PercentChangeOneHour string `json:"percent_change_1h"`
	PercentChangeOneDay  string `json:"percent_change_24h"`
	PercentChangeOneWeek string `json:"percent_change_7d"`
	LastUpdated          string `json:"last_updated"`
}

type responseData struct {
	price    string
	priceBtc string
}

// PriceController describes a price controller.
type PriceController struct{}

// Get the current CoinMarketCap fiat price of ECA.
func (p PriceController) Get(c *gin.Context) {

	coin := c.Param("coin")

	if len(coin) < 1 {
		coin = "electra"
	}

	cacheKey := c.Param("coin") + "-price-" + c.Param("currency")

	cacheData, found := helpers.CacheInstance.Get(cacheKey)
	if found {
		c.JSON(http.StatusOK, gin.H{
			"price":    cacheData.(*responseData).price,
			"priceBtc": cacheData.(*responseData).priceBtc,
		})

		return
	}

	url := "https://api.coinmarketcap.com/v1/ticker/" + coin + "/?convert=" + c.Param("currency")
	var inputData requestResponseData
	hasError := helpers.GetJSON(url, &inputData)
	if hasError {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error."})

		return
	}

	data := &responseData{
		price:    inputData[0].PriceUSD,
		priceBtc: inputData[0].PriceBTC,
	}
	helpers.CacheInstance.Set(cacheKey, data, cache.DefaultExpiration)

	c.Header("X-Version", "1.0")
	c.JSON(http.StatusOK, gin.H{
		"price":    data.price,
		"priceBtc": data.priceBtc,
	})
}
