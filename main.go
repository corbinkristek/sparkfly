package main

import "github.com/corbinkristek/sparkfly/concurrency"

func main() {
	b := concurrency.InitBarcodeScanner()
	b.ScanFile()
}
