package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Definindo os parâmetros da CLI
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()

	// Verificando se os parâmetros foram fornecidos
	if *url == "" || *totalRequests == 0 || *concurrency == 0 {
		fmt.Println("Parâmetros obrigatórios: --url, --requests, --concurrency")
		return
	}

	// Iniciando o teste de carga
	fmt.Println("Iniciando o teste de carga...")
	start := time.Now()
	var wg sync.WaitGroup
	requests := make(chan struct{}, *concurrency)

	// Variáveis para armazenar os resultados
	var totalTime time.Duration
	var successCount, failureCount int
	statusCodes := make(map[int]int)
	var mu sync.Mutex

	for i := 0; i < *totalRequests; i++ {
		requests <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			startRequest := time.Now()
			resp, err := http.Get(*url)
			requestTime := time.Since(startRequest)
			mu.Lock()
			totalTime += requestTime
			if err != nil {
				failureCount++
			} else {
				statusCodes[resp.StatusCode]++
				if resp.StatusCode == 200 {
					successCount++
				} else {
					failureCount++
				}
				resp.Body.Close()
			}
			mu.Unlock()
			<-requests
		}()
	}

	wg.Wait()
	totalTimeElapsed := time.Since(start)

	// Gerando o relatório
	fmt.Println("Teste de carga concluído.")
	fmt.Printf("Tempo total gasto: %s\n", totalTimeElapsed)
	fmt.Printf("Quantidade total de requests: %d\n", *totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successCount)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCodes {
		fmt.Printf("HTTP %d: %d\n", code, count)
	}
}
