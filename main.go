package main

import (
	"datastar-demo/views"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	datastar "github.com/starfederation/datastar/sdk/go"
)

var accounts = []views.Accounts{
	{Name: "Operating Account", Currency: "USD"},
	{Name: "Payroll Account", Currency: "CAD"},
	{Name: "Expense Account", Currency: "GBP"},
	{Name: "Corporate Savings", Currency: "EUR"},
	{Name: "Investment Account", Currency: "JPY"},
	{Name: "Tax Reserve Fund", Currency: "AUD"},
	{Name: "Petty Cash", Currency: "INR"},
	{Name: "Vendor Payments", Currency: "CNY"},
	{Name: "Client Deposits", Currency: "CHF"},
	{Name: "Foreign Exchange Account", Currency: "BRL"},
	{Name: "Capital Expenditure", Currency: "ZAR"},
	{Name: "General Ledger Account", Currency: "RUB"},
	{Name: "Retirement Fund", Currency: "MXN"},
	{Name: "Escrow Account", Currency: "SAR"},
	{Name: "Reserve Fund", Currency: "AED"},
}

func main() {
	fmt.Println("Hello, World!")
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Country_List(accounts)).ServeHTTP(w, r)
	})

	r.Get("/modal/{account}", func(w http.ResponseWriter, r *http.Request) {
		currency_code := chi.URLParam(r, "account")

		// db query to get account details to render in modal as a form

		sse := datastar.NewSSE(w, r)

		sse.MergeFragmentTempl(views.Modal(currency_code), datastar.WithMergeMode("morph"), datastar.WithSelector("#modal-form"))
		sse.MergeFragments(`<div id="overlay" class="fixed inset-0 flex items-end justify-center bg-black/20"></div>`)
	})

	r.Get("/hide-modal", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)

		sse.MergeFragments(`<div id="modal-form"></div>`)
		sse.MergeFragments(`<div id="overlay"></div>`)
	})

	http.ListenAndServe(":8080", r)
}
