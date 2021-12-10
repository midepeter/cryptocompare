package db

type Result struct {
	RAW struct {
		BTC struct {
			USD struct {
				CHANGE24HOUR    float64 `json:"change24hour"`
				CHANGEPCT24HOUR float64 `json:"changepcd24hour"`
				OPEN24HOUR      float64 `json:"open24hour"`
				VOLUME24HOUR    float64 `json:"volume24hour"`
				VOLUME24HOURTO  float64 `json:"volume24hourto"`
				LOW24HOUR       float64 `json:"low24hour"`
				HIGH24HOUR      float64 `json:"high24hour"`
				PRICE           float64 `json:"price"`
				SUPPLY          float64 `json:"supply"`
				MKTCAP          float64 `json:"mktcap"`
			}
		}
	}
	DISPLAY struct {
		BTC struct {
			USD struct {
				CHANGE24HOUR    string `json:"change24hour"`
				CHANGEPCT24HOUR string `json:"changepct24hour"`
				OPEN24HOUR      string `json:"open24hour"`
				VOLUME24HOUR    string `json:"volume24hour"`
				VOLUME24HOURTO  string `json:"volume24hourto"`
				HIGH24HOUR      string `json:"high24hour"`
				PRICE           string `json:"price"`
				FROMSYMBOL      string `json:"fromsymbol"`
				TOSYMBOL        string `json:"tosymbol"`
				LASTUPDATE      string `json:"lastupdate"`
				SUPPLY          string `json:"supply"`
				MKTCAP          string `json:"mktcap"`
			}
		}
	}
}
