# Nexus-Gateway: Micro-Infraestructura de Alto Rendimiento en Go

Nexus-Gateway es un **API Gateway minimalista** diseñado para demostrar conceptos avanzados de sistemas distribuidos, redes y concurrencia. El proyecto implementa desde cero componentes críticos de infraestructura sin depender de librerías externas pesadas.

## 🚀 Características Técnicas

El core del sistema se divide en cuatro pilares de ingeniería:

### 1. Generador de IDs Distribuidos (Estilo Snowflake)
Implementación de un generador de identificadores únicos de 64 bits. 
- **Teoría:** Utiliza una combinación de *Timestamp* (milisegundos), *Worker ID* y un *Sequence Counter*.
- **Beneficio:** Garantiza colisión cero en ambientes distribuidos sin necesidad de una base de datos centralizada.

### 2. Escáner de Puertos Multihilo
Un motor de descubrimiento de servicios que utiliza el modelo de concurrencia de Go.
- **Teoría:** Implementación de `sync.WaitGroup` y `goroutines` para realizar escaneos TCP no bloqueantes.
- **Beneficio:** Identifica dinámicamente servicios activos en la red local antes de balancear tráfico.

### 3. Balanceador de Carga (Round Robin)
Distribuidor de tráfico para alta disponibilidad.
- **Teoría:** Algoritmo de rotación cíclica utilizando operaciones atómicas (`sync/atomic`) para garantizar la consistencia del índice incluso bajo alta concurrencia.
- **Beneficio:** Distribuye la carga de trabajo equitativamente entre los nodos del backend.

### 4. Limitador de Tasa (Token Bucket)
Middleware de control de flujo para protección del sistema.
- **Teoría:** Algoritmo de "Cubo de Tokens" que permite ráfagas controladas de tráfico pero mantiene un promedio constante a largo plazo.
- **Beneficio:** Protege los recursos del servidor contra picos de tráfico (Spikes) y ataques de denegación de servicio (DoS).

## 🛠️ Tecnologías Utilizadas
- **Lenguaje:** Go (Golang) 1.21+
- **Concurrencia:** Goroutines, Channels, Mutexes, Atomic operations.
- **Redes:** Net/HTTP, TCP Dialing.

## 📦 Instalación y Uso

1. Clonar el repositorio:
   ```bash
   git clone [https://github.com/tu-usuario/nexus-gateway.git](https://github.com/tu-usuario/nexus-gateway.git)

2. Ejecutar el gateway:
go run main.go

3. Probar el endpoint:
curl http://localhost:3000

🧠 ¿Por qué Go?
Elegí Go para este proyecto debido a su eficiencia en el manejo de operaciones de I/O y su modelo de memoria ligero. La capacidad de disparar miles de "hilos" (goroutines) con un consumo mínimo de RAM lo hace la herramienta perfecta para construir middleware de infraestructura como este. 

No usé frameworks porque quería implementar manualmente la lógica de un **Token Bucket** para el Rate Limiting y un generador tipo **Snowflake** para asegurar la trazabilidad de cada petición con IDs únicos. Esto me permitió profundizar en cómo se gestiona la concurrencia y la estabilidad en sistemas que escalan.
