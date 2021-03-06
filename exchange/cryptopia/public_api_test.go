package cryptopia

import (
	"testing"
)

func TestRequestSignature(t *testing.T) {
	key := "abababababababababababababababab"
	secret := "YWJhYmFiYWJhYmFiYWJhYmFiYWJhYmFiYWJhYmFiYWI="
	nonce := "3"
	requrl := apiroot

	requrl.Path += "getbalance"
	var want = "amx abababababababababababababababab:QRB4yf+QkSxxzPg6JLDeNFdAsTu24wpiDozHNQZ3Jkc=:3"
	if expected := header(key, secret, nonce, requrl, []byte("{}")); want != expected {
		t.Fatal("invalid request signature")
	}
}

func Test_getCurrencyID(t *testing.T) {
	c := &Client{}

	btcID, err := c.GetCurrencyID("btc")
	if err != nil {
		t.Fatal(err)
	}
	if btcID != 1 {
		t.Errorf("Incorrect BTC id %d, want %d", btcID, 1)
	}
	ltcID, err := c.GetCurrencyID("ltc")
	if ltcID != 3 {
		t.Errorf("Incorrect BTC id %d, want %d", ltcID, 3)
	}
	if err != nil {
		t.Fatal(err)
	}
	skyID, err := c.GetCurrencyID("sky")
	if err != nil {
		t.Fatal(err)
	}
	if skyID != 504 {
		t.Errorf("Incorrect BTC id %d, want %d", skyID, 504)
	}
}

func Test_getMarketID(t *testing.T) {
	c := &Client{}

	btcltc, err := c.GetMarketID("ltc_btc")
	if err != nil || btcltc != 101 {
		t.Fatal(err, btcltc)
	}
}

func TestGetCurrencies(t *testing.T) {
	c := &Client{}

	_, err := c.GetCurrencies()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTradePairs(t *testing.T) {
	c := &Client{}

	_, err := c.GetTradePairs()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMarkets(t *testing.T) {
	c := &Client{}

	mkts, err := c.GetMarkets("ALL", -1)
	if err != nil {
		t.Fatal(err)
	}
	if len(mkts) < 1 {
		t.Fatal("empty")
	}
}

func TestGetMarket(t *testing.T) {
	c := &Client{}

	mkt, err := c.GetMarket("LTC/BTC", -1)
	if err != nil {
		t.Fatal(err)
	}
	if mkt.TradePairID != 101 {
		t.Fatal("API error", "want 101 TradePairID")
	}
}

func TestGetMarketHistory(t *testing.T) {
	c := &Client{}

	hst, err := c.GetMarketHistory("LTC/BTC", -1)
	if err != nil {
		t.Fatal(err)
	}
	if len(hst) < 1 {
		t.Fatal("empty")
	}
}

func TestGetMarketOrders(t *testing.T) {
	c := &Client{}

	orders, err := c.GetMarketOrders("LTC/BTC", -1)
	if err != nil {
		t.Fatal(err)
	}
	if len(orders.Buy) < 1 || len(orders.Sell) < 1 {
		t.Fatal("empty")
	}
}

func TestGetMarketOrderGroups(t *testing.T) {
	c := &Client{}

	groups, err := c.GetMarketOrderGroups(-1, []string{"LTC/BTC", "SKY/BTC"})
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) != 2 {
		t.Fatal("count of groups should be 2")
	}
}
