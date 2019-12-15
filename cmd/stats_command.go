package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/currency"
)

const BankStatsApiUrl = "https://bank.hackclub.com/stats.json"

// BankStats types are used to unmarshal the Bank stats API
type BankStats struct {
	TransactionsVolume int64 `json:"transactions_volume"`
	TransactionsCount  int64 `json:"transactions_count"`
	EventsCount        int64 `json:"events_count"`
	Raised             int64 `json:"raised"`

	LastYear  BankTimedStats `json:"last_year"`
	LastQtr   BankTimedStats `json:"last_qtr"`
	LastMonth BankTimedStats `json:"last_month"`
}

type BankTimedStats struct {
	TransactionsVolume int64 `json:"transactions_volume"`
	Raised             int64 `json:"raised"`
	Revenue            int64 `json:"revenue"`
}

type BankEventStats struct {
	CreatedAt int64 `json:"created_at"`
}

// Utility used to format USD currency amounts
func RenderMoney(amount int64) string {
	amt := currency.USD.Amount(float64(amount) / 100)
	return fmt.Sprintf("%d", currency.NarrowSymbol(amt))
}

// StatsCommand defines the command to retrieve common
// Hack Club Bank statistics.
type StatsCommand struct{}

func (cmd StatsCommand) Init() {}

func (cmd StatsCommand) ShouldRespond(msg string) bool {
	return strings.Contains(strings.ToLower(msg), "stats")
}

func (cmd StatsCommand) Respond(msg string) string {
	resp, err := http.Get(BankStatsApiUrl)
	if err != nil {
		return "Oops, couldn't fetch stats from Hack Club Bank!"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Oops, couldn't read stats returned by Hack Club Bank!"
	}

	// unmarshal stats.json
	bankStats := BankStats{}
	json.Unmarshal(body, &bankStats)

	lines := []string{
		"*Hack Club Bank Stats* :bank-hackclub:",
		fmt.Sprintf("Total transactions: %s", RenderMoney(bankStats.TransactionsVolume)),
		fmt.Sprintf("Total raised: %s", RenderMoney(bankStats.Raised)),
		fmt.Sprintf("Transaction count: %d", bankStats.TransactionsCount),
		fmt.Sprintf("Events count: %d", bankStats.EventsCount),
		"---",
		fmt.Sprintf("*Last month*, Bank made %s", RenderMoney(bankStats.LastMonth.Revenue)),
		fmt.Sprintf("Transactions: %s | Raised: %s",
			RenderMoney(bankStats.LastMonth.TransactionsVolume),
			RenderMoney(bankStats.LastMonth.Raised),
		),
		"---",
		fmt.Sprintf("*Last quarter*, Bank made %s", RenderMoney(bankStats.LastQtr.Revenue)),
		fmt.Sprintf("Transactions: %s | Raised: %s",
			RenderMoney(bankStats.LastQtr.TransactionsVolume),
			RenderMoney(bankStats.LastQtr.Raised),
		),
		"---",
		fmt.Sprintf("*Last year*, Bank made %s", RenderMoney(bankStats.LastYear.Revenue)),
		fmt.Sprintf("Transactions: %s | Raised: %s",
			RenderMoney(bankStats.LastYear.TransactionsVolume),
			RenderMoney(bankStats.LastYear.Raised),
		),
	}

	return strings.Join(lines, "\n")
}
