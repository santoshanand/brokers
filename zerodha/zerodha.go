package zerodha

import (
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/samber/lo"
	"resty.dev/v3"
)

// Instrument represents a row from the instruments.csv
type Instrument struct {
	InstrumentToken int64  `csv:"instrument_token"`
	ExchangeToken   int64  `csv:"exchange_token"`
	Tradingsymbol   string `csv:"tradingsymbol"`
	Name            string `csv:"name"`
	Expiry          string `csv:"expiry"`
	LotSize         int64  `csv:"lot_size"`
	InstrumentType  string `csv:"instrument_type"`
	Segment         string `csv:"segment"`
	Exchange        string `csv:"exchange"`
}

type Zerodha struct {
	client *resty.Client
}

const (
	instrumentURL = "https://api.kite.trade/instruments"
)

func NewZerodha(client *resty.Client) *Zerodha {
	return &Zerodha{client: client}
}

func (z *Zerodha) fetchInstruments() ([]*Instrument, error) {
	resp, err := z.client.R().Get(instrumentURL)
	if err != nil {
		return nil, err
	}
	csvData := string(resp.Bytes())
	reader := strings.NewReader(csvData)
	var instruments []*Instrument
	if err := gocsv.Unmarshal(reader, &instruments); err != nil {
		return nil, err
	}
	return instruments, nil
}

func (z *Zerodha) LoadInstrument() ([]*Instrument, error) {
	instruments, err := z.fetchInstruments()
	if err != nil {
		return nil, err
	}
	return instruments, nil
}

func (z *Zerodha) LoadOptionStockInstruments() ([]*Instrument, error) {
	instruments, err := z.fetchInstruments()
	if err != nil {
		return nil, err
	}
	finalInstruments := []*Instrument{}
	for _, inst := range instruments {
		if inst.Exchange != "NFO" {
			continue
		}
		if strings.Contains(inst.Tradingsymbol, "CE") || strings.Contains(inst.Tradingsymbol, "PE") {
			continue
		}
		if strings.Contains(inst.Exchange, "BFO") && inst.Name == "SENSEX" {
			finalInstruments = append(finalInstruments, inst)
		}
		total := len(inst.Tradingsymbol)
		if total > 8 {
			inst.Tradingsymbol = inst.Tradingsymbol[:len(inst.Tradingsymbol)-8]
		}
		finalInstruments = append(finalInstruments, inst)
	}
	uniqueItems := lo.UniqBy(finalInstruments, func(v *Instrument) string {
		return v.Tradingsymbol
	})
	return uniqueItems, nil
}
