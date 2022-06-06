package concurrency

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type barcodeScanner struct {
	m  map[string]int
	mu sync.Mutex
	wg sync.WaitGroup
}

func InitBarcodeScanner() *barcodeScanner {
	return &barcodeScanner{m: make(map[string]int)}
}

func (b *barcodeScanner) ScanFile() {

	for i := 0; i < 5; i++ {
		b.wg.Add(1)

		go func(j int) {
			defer b.wg.Done()
			path := fmt.Sprintf("./concurrency/testdata/TestProcess2_%s.csv", strconv.Itoa(j))

			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			csvReader := csv.NewReader(f)
			data, err := csvReader.ReadAll()
			if err != nil {
				log.Fatal(err)
			}
			for i, line := range data {
				if i > 0 {
					var rec string
					for j, field := range line {
						if j == 1 {
							rec = field
						}
					}

					b.mu.Lock()
					b.m[rec]++
					count := b.m[rec]
					if count > 1 {
						//stops once duplicate is found
						log.Fatal("barcode already exists: ", rec)
					}

					b.mu.Unlock()
				}
			}
		}(i)

	}
	b.wg.Wait()
	log.Println("no duplicates found")
}
