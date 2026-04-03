package main

import (
	"fmt"
	"net/http"
	"nexus-gateway/internal/balancer"
	"nexus-gateway/internal/idgen"
	"nexus-gateway/internal/limiter"
	"nexus-gateway/internal/scanner"
	"time"
)

func main() {
	// 1. Escaneo previo de puertos (Simulación de discovery)
	fmt.Println("🔍 Escaneando servicios locales...")
	activePorts := scanner.ScanLocalPorts("localhost", 8080, 8085)
	fmt.Printf("✅ Servidores encontrados en puertos: %v\n", activePorts)

	// 2. Setup de componentes
	nodeID := idgen.NewSnowflake(1)
	bucket := limiter.NewTokenBucket(5, 2*time.Second) // 5 peticiones máx, recupera 1 cada 2 seg
	
	targets := []string{"Server-A", "Server-B", "Server-C"}
	lb := balancer.NewRoundRobin(targets)

	// 3. Servidor de entrada
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !bucket.Allow() {
			http.Error(w, "Too Many Requests (Rate Limit)", http.StatusTooManyRequests)
			return
		}

		requestID := nodeID.NextID()
		targetServer := lb.Next()

		fmt.Fprintf(w, "Nexus-Gateway\n")
		fmt.Fprintf(w, "ID Unico (Snowflake): %s\n", requestID)
		fmt.Fprintf(w, "Redirigiendo a: %s\n", targetServer)
		fmt.Printf("LOG: [%s] -> %s\n", requestID, targetServer)
	})

	fmt.Println("🚀 Nexus-Gateway corriendo en http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}