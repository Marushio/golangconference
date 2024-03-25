# golangconference
Repository for Go Lang conference scripts

to run
```go
go run main.go
```

to build for linux
```go
go build main.go
```

to build for windows (.exe)
```go 
GOOS=windows go build main.go
```

//go-wrk programa que faz teste de carga 
//-c forma concorrente
//-d 15 segundos
go-wrk -c 10 -d 15 http://localhost:8081/cpu

//Ferramenta de profile endpoint para status de Processamento
go tool pprof -seconds 5 http://localhost:6060/debug/pprof/profile
//Ferramenta de profile endpoint para status de Memoria
go tool pprof -seconds 5 http://localhost:6060/debug/pprof/allocs
comandos pprof:
```top 20``` 20 top processos mais custosos
```list CPUIntensiveEndpoint``` monstra em qual linha da func  CPUIntensiveEndpoint esta demorando mais para executar
```web``` Abre um grafico no navegador com o passo a passo das alocacoes de recursos


go test examples:
marcio_d@MarcioD-Gaming:~/repos/golangconference/day3$ go test -bench=. -benchmem
PASS
ok      golangconference        0.003s
marcio_d@MarcioD-Gaming:~/repos/golangconference/day3$ go test -bench=. -benchmem -count 10 > 1.bench
marcio_d@MarcioD-Gaming:~/repos/golangconference/day3$ go test -bench=. -benchmem -count 10 > 2.
bench
marcio_d@MarcioD-Gaming:~/repos/golangconference/day3$ benchstat 1.bench 2.bench 
bash: benchstat: command not found
marcio_d@MarcioD-Gaming:~/repos/golangconference/day3$ benchstat 1.bench 2.bench

->Metricas
	Otel collector https://opentelemetry.io/docs/collector/
MELT
	- Metrics
		RED Metrics
		Rate
		Error
		Duration
		P50,P75,P99
	- Events
		- Uma acao discreta que acontece em um momento no tempo.
		- Tem releavancia para o negocio.
	- Logs
	- Traces
	  - Loki

Operators
https://operatorhub.io/


Links
GO com bigtable (google)
https://github.com/Marushio/golangconference_bigtable_example

Tsuro - Platform as a service
https://tsuru.io/

