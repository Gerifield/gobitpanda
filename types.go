package gobitpanda

import (
	"net/http"
	"sync"
	"time"
)

const (
	// APIBase points to the public Bitpanda Global Exchange API
	APIBase = "https://api.exchange.bitpanda.com/public"
)

// Currency codes
const (
	CurrencyBEST  string = "BEST"
	CurrencyBTC   string = "BTC"
	CurrencyETH   string = "ETH"
	CurrencyEUR   string = "EUR"
	CurrencyMIOTA string = "MIOTA"
	CurrencyPAN   string = "PAN"
	CurrencyUSDT  string = "USDT"
	CurrencyXRP   string = "XRP"
)

// Instrument codes
const (
	InstrumentBESTBTC  string = "BEST_BTC"
	InstrumentBESTEUR  string = "BEST_EUR"
	InstrumentBESTUSDT string = "BEST_USDT"
	InstrumentBTCEUR   string = "BTC_EUR"
	InstrumentBTCUSDT  string = "BTC_USDT"
	InstrumentETHBTC   string = "ETH_BTC"
	InstrumentETHEUR   string = "ETH_EUR"
	InstrumentMIOTABTC string = "MIOTA_BTC"
	InstrumentMIOTAEUR string = "MIOTA_EUR"
	InstrumentPANBTC   string = "PAN_BTC"
	InstrumentXRPBTC   string = "XRP_BTC"
	InstrumentXRPEUR   string = "XRP_EUR"
)

// Levels
const (
	LevelOne     int = 1
	LevelTwo     int = 2
	LevelThree   int = 3
	LevelDefault int = 3
)

// Possible value for `type` in orders
//
// https://developers.bitpanda.com/exchange/#/orders-post
const (
	OrderTypeMarket string = "MARKET"
	OrderTypeLimit  string = "LIMIT"
	OrderTypeStop   string = "STOP"
)

// Possible value for `side` in orders
const (
	OrderSideBuy  string = "BUY"
	OrderSideSell string = "SELL"
)

// Periods
const (
	PeriodOneMinute      int = 1
	PeriodFiveMinutes    int = 5
	PeriodFifteenMinutes int = 15
	PeriodThirtyMinutes  int = 30
	PeriodOneHour        int = 1
	PeriodFourHours      int = 4
	PeriodOneDay         int = 1
	PeriodOneWeek        int = 1
	PeriodOneMonth       int = 1
)

// Status
const (
	StatusOpen           string = "OPEN"
	StatusStopTriggered  string = "STOP_TRIGGERED"
	StatusFilled         string = "FILLED"
	StatusFilledFully    string = "FILLED_FULLY"
	StatusFilledClosed   string = "FILLED_CLOSED"
	StatusFilledRejected string = "FILLED_REJECTED"
	StatusRejected       string = "REJECTED"
	StatusClosed         string = "CLOSED"
	StatusFailed         string = "FAILED"
)

// Units
const (
	UnitMinutes string = "MINUTES"
	UnitHours   string = "HOURS"
	UnitDays    string = "DAYS"
	UnitWeeks   string = "WEEKS"
	UnitMonths  string = "MONTHS"
)

type (

	// Account details of a registered user's balance(s).
	Account struct {
		AccountID string    `json:"account_id"`
		Balances  []Balance `json:"balances"`
	}

	// AccountFees struct
	AccountFees struct {
		AccountID            string    `json:"account_id"`
		RunningTradingVolume string    `json:"running_trading_volume"`
		FeeGroupID           string    `json:"fee_group_id"`
		FeeTiers             []FeeTier `json:"fee_tiers"`
		ActiveFeeTier        FeeTier   `json:"active_fee_tier"`
		CollectFeesInBest    bool      `json:"collect_fees_in_best"`
		FeeDiscountRate      string    `json:"fee_discount_rate"`
		MinimumPriceValue    string    `json:"minimum_price_value"`
	}

	// Ask struct
	Ask struct {
		Value Value `json:"value"`
	}

	//Asks struct
	Asks struct {
		Price          string `json:"price,omitempty"`
		Amount         string `json:"amount,omitempty"`
		NumberOfOrders uint   `json:"number_of_orders,omitempty"`
		OrderID        string `json:"order_id,omitempty"`
	}

	// Balance Account balance for one single currency
	Balance struct {
		AccountID    string    `json:"account_id"`
		CurrencyCode string    `json:"currency_code"`
		Change       string    `json:"change"`
		Available    string    `json:"available"`
		Locked       string    `json:"locked"`
		Sequence     uint      `json:"sequence"`
		Time         time.Time `json:"time"` // RFC3339
	}

	// Bid struct
	Bid struct {
		Value Value `json:"value"`
	}

	//Bids struct
	Bids struct {
		Price          string `json:"price,omitempty"`
		Amount         string `json:"amount,omitempty"`
		NumberOfOrders uint   `json:"number_of_orders,omitempty"`
		OrderID        string `json:"order_id,omitempty"`
	}

	// Candlestick representing price action for a given period
	Candlestick struct {
		LastSequence   uint        `json:"last_sequence"`
		InstrumentCode string      `json:"instrument_code"`
		Granularity    Granularity `json:"granularity"`
		High           string      `json:"high"`
		Low            string      `json:"low"`
		Open           string      `json:"open"`
		Close          string      `json:"close"`
		Volume         string      `json:"volume"`
		Time           time.Time   `json:"time"` // RFC3339
	}

	// Client represents a Bitpanda REST API Client
	Client struct {
		sync.Mutex
		Client   *http.Client
		APIBase  string
		APIToken string
	}

	// CreateOrder struct
	CreateOrder struct {
		InstrumentCode string `json:"instrument_code"`
		Type           string `json:"type"`
		Side           string `json:"side"`
		Amount         string `json:"amount"`
		Price          string `json:"price,omitempty"`
		TriggerPrice   string `json:"trigger_price,omitempty"`
		ClientID       string `json:"client_id,omitempty"`
	}

	// Currency struct
	Currency struct {
		Code      string `json:"code"`
		Precision int    `json:"precision,omitempty"`
	}

	// CurrencyCode struct
	CurrencyCode struct {
		Code string `json:"currency"`
	}

	// DepositReturn struct
	DepositReturn struct {
		Address        string `json:"address"`
		DestinationTag string `json:"destinationTag"`
		Enabled        bool   `json:"enabled"`
		CanCreateMore  bool   `json:"can_create_more"`
		SmartContract  bool   `json:"is_smart_contract"`
	}

	// WithdrawReturn struct
	WithdrawReturn struct {
		Amount         string `json:"amount"`
		Recipient      string `json:"recipient"`
		DestinationTag string `json:"destinationTag"`
		Fee            string `json:"fee"`
	}

	// ErrorResponse struct
	ErrorResponse struct {
		Response *http.Response `json:"-"`
		Error    string         `json:"error"`
	}

	// Fee applied account balance as part of trade settlement
	Fee struct {
		FeeAmount            string `json:"fee_amount"`
		FeeCurreny           string `json:"fee_currency"`
		FeePercentage        string `json:"fee_percentage"`
		FeeGroupID           string `json:"fee_group_id"`
		FeeType              string `json:"fee_type"`
		RunningTradingVolume string `json:"running_trading_volume"`
	}

	// FeeGroup struct
	FeeGroup struct {
		FeeGroupID        string    `json:"fee_group_id"`
		DisplayText       string    `json:"display_text,omitempty"`
		FeeTiers          []FeeTier `json:"fee_tiers"`
		FeeDiscountRate   string    `json:"fee_discount_rate"`
		MinimumPriceValue string    `json:"minimum_price_value"`
	}

	// FeeMode toggle to enable or disable fee collection with BEST
	FeeMode struct {
		CollectFeesInBest bool `json:"collect_fees_in_best"`
	}

	// FeeTier struct
	FeeTier struct {
		FeeGroupID string `json:"fee_group_id"`
		Volume     string `json:"volume"`
		MakerFee   string `json:"maker_fee"`
		TakerFee   string `json:"taker_fee"`
	}

	// FiatDepositReturn struct
	FiatDepositReturn struct {
		IBAN                string `json:"iban"`
		BIC                 string `json:"bic"`
		Bank                string `json:"bank"`
		Address             string `json:"address"`
		Receiver            string `json:"receiver"`
		ReceiverAddress     string `json:"receiver_address"`
		UniquePaymentNumber string `json:"unique_payment_number"`
	}

	// Granularity struct
	Granularity struct {
		Unit   string `json:"unit"`
		Period uint   `json:"period"`
	}

	// Instrument struct
	Instrument struct {
		State           string   `json:"state"`
		Base            Currency `json:"base"`
		Quote           Currency `json:"quote"`
		AmountPrecision uint     `json:"amount_precision"`
		MarketPrecision uint     `json:"market_precision"`
		MinSize         string   `json:"min_size"`
	}

	// MarketTick struct
	MarketTick struct {
		InstrumentCode        string `json:"instrument_code"`
		Sequence              int    `json:"sequence"`
		State                 string `json:"state"`
		IsFrozen              int    `json:"is_frozen"`
		QuoteVolume           string `json:"quote_volume"`
		BaseVolume            string `json:"base_volume"`
		LastPrice             string `json:"last_price"`
		BestBid               string `json:"best_bid"`
		BestAsk               string `json:"best_ask"`
		PriceChange           string `json:"price_change"`
		PriceChangePercentage string `json:"price_change_percentage"`
		High                  string `json:"high"`
		Low                   string `json:"low"`
	}

	// Order struct
	Order struct {
		OrderID         string    `json:"order_id"`
		AccountID       string    `json:"account_id"`
		InstrumentCode  string    `json:"instrument_code"`
		Amount          string    `json:"amount"`
		FilledAmount    string    `json:"filled_amount"`
		Side            string    `json:"side"`
		Type            string    `json:"type"`
		Status          string    `json:"status"`
		Sequence        uint      `json:"sequence,omitempty"`
		Price           string    `json:"price"`
		Reason          string    `json:"reason,omitempty"`
		Time            time.Time `json:"time"`                        // RFC3339
		TimeLastUpdated time.Time `json:"time_last_updated,omitempty"` // RFC3339
		TimeTriggered   time.Time `json:"time_triggered,omitempty"`    // RFC3339
		TriggerPrice    string    `json:"trigger_price,omitempty"`
	}

	// OrderBook a snapshot of the order book state
	OrderBook struct {
		InstrumentCode string    `json:"instrument_code"`
		Time           time.Time `json:"time"` // RFC3339
		Bids           []Bids    `json:"bids"`
		Asks           []Asks    `json:"asks"`
	}

	// OrderBookLvlOne a snapshot of the order book state
	OrderBookLvlOne struct {
		InstrumentCode string    `json:"instrument_code"`
		Time           time.Time `json:"time"` // RFC3339
		Bids           Bid       `json:"bids"`
		Asks           Ask       `json:"asks"`
	}

	// OrderHistory Paginated collection of account orders
	OrderHistory struct {
		OrderHistory []OrderHistoryEntry `json:"order_history"`
		MaxPageSize  uint                `json:"max_page_size,omitempty"`
		Cursor       string              `json:"cursor,omitempty"`
	}

	// OrderHistoryEntry active or inactive order, for orders with the status FILLED, FILLED_FULLY, FILLED_CLOSED and FILLED_REJECTED, information about trades and fees is returned.
	OrderHistoryEntry struct {
		Order  Order               `json:"order"`
		Trades []TradeHistoryEntry `json:"trades"`
	}

	// PriceTick struct
	PriceTick struct {
		InstrumentCode string    `json:"instrument_code"`
		Price          string    `json:"price"`
		Amount         string    `json:"amount"`
		Volume         string    `json:"volume"`
		Sequence       int       `json:"sequence"`
		TakerSide      string    `json:"taker_side"`
		Time           time.Time `json:"time"`
		TradeTimestamp int64     `json:"trade_timestamp"`
	}

	// Recipient struct
	Recipient struct {
		Address        string `json:"address"`
		DestinationTag string `json:"destination_tag"`
	}

	// Time struct
	Time struct {
		Iso         string `json:"iso"` // RFC3339
		EpochMillis uint64 `json:"epoch_millis"`
	}

	// TimeGranularity is a length of time defined by unit and period used to identify the type of candlestick.
	// Supported resolutions are 1, 5, 15, 30 MINUTES & 1, 4 HOURS & 1 DAYS & 1 WEEKS & 1 MONTHS.
	TimeGranularity struct {
		Unit   string `json:"unit"`
		Period uint   `json:"period"`
	}

	// Trade struct
	Trade struct {
		TradeID        string    `json:"trade_id"`
		OrderID        string    `json:"order_id"`
		AccountID      string    `json:"account_id"`
		Amount         string    `json:"amount"`
		Side           string    `json:"side"`
		InstrumentCode string    `json:"instrument_code"`
		Price          string    `json:"price"`
		Time           time.Time `json:"time"` // RFC3339
		Sequence       uint      `json:"sequence,omitempty"`
	}

	// TradeHistory Paginated collection of account trades
	TradeHistory struct {
		TradeHistory []TradeHistoryEntry `json:"trade_history"`
		MaxPageSize  uint                `json:"max_page_size,omitempty"`
		Cursor       string              `json:"cursor,omitempty"`
	}

	// TradeHistoryEntry Trade recorded for exactly one order
	TradeHistoryEntry struct {
		Trade Trade `json:"trade"`
		Fee   Fee   `json:"fee"`
	}

	// TradingVolume struct
	TradingVolume struct {
		Volume string `json:"volume"`
	}

	// Value struct
	Value struct {
		Price          string `json:"price"`
		Amount         string `json:"amount"`
		NumberOfOrders uint   `json:"number_of_orders"`
	}

	//Withdraw struct
	Withdraw struct {
		Currency  string    `json:"currency"`
		Amount    string    `json:"amount"`
		Recipient Recipient `json:"recipient"`
	}
)
