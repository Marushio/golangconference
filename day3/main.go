package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	//Ferramenta para para fazer profile de teste de performance
	_ "net/http/pprof"

	//O _ significa que nao vai usar explicitamente a li
	_ "github.com/mattn/go-sqlite3"

	
)

type user struct {
	ID    int    `json:"id"` //A tag json entre `` muda o valor ID para id quando convertido para json
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	//Cria o servidor http multiplex
	mux := http.NewServeMux()

	//Cria a rota apontando para a funcao listUsersHandler
	mux.HandleFunc("/users", listUsersHandler)
	mux.HandleFunc("/users/{id}", getUserHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("/cpu", CPUIntensiveEndpoint)

	//Sobe o listen apontando para o servidor mux
	go http.ListenAndServe(":8081", mux)

	http.ListenAndServe(":6060", nil)

}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	//Cria a conecxao com o banco
	db, err := sql.Open("sqlite3", "users.db")

	//Tratamento de erro caso algum problema com o banco de dados
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Fecha a conexao com o banco depois que todos os codigos forem executados
	defer db.Close()

	//Executa o select na base
	rows, err := db.Query("SELECT * FROM users")
	//Tratamento de erro caso algum problema no select
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//Declaracao de um slice de users (como se fosse um vetor dinamico)
	users := []*user{}

	//Iteracao em cima do retorno do select
	for rows.Next() {
		//Cria variavel user seguindo a struct
		var u user

		//Scan da linha para pegar um user com tratamento caso algo de errado
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Adiciona o usuario da linha ao slice de users
		users = append(users, &u)
	}

	//Converte usuaros para json e joga no response, com o tratamento em caso de erro
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	//Cria a conecxao com o banco
	db, err := sql.Open("sqlite3", "users.db")
	//Tratamento de erro caso algum problema com o banco de dados
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Recebe um json e converte para a struc user
	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	//Insere na base de dados
	if _, err := db.Exec("INSERT INTO users (name, email) VALUES (?,?)", u.Name, u.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Retorna o status created
	w.WriteHeader(http.StatusCreated)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	//Cria a conecxao com o banco
	db, err := sql.Open("sqlite3", "users.db")
	//Tratamento de erro caso algum problema com o banco de dados
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Fecha a conexao com o banco depois que todos os codigos forem executados
	defer db.Close()

	//Pega o id do usuario da url jeito antigo
	//id := r.URL.Query().Get("id")

	//Pega o id do usuario da url jeito depois da versao 1.22 do go
	id := r.PathValue("id")

	//Executa o select na base
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	//Cria variavel user seguindo a struct
	var u user

	//Scan da linha para pegar um user com tratamento caso algo de errado
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Converte usuaros para json e joga no response, com o tratamento em caso de erro
	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

// Teste de performance com go
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func CPUIntensiveEndpoint(w http.ResponseWriter, r *http.Request) {
	// Simulate CPU intensive task
	result := fib(40)
	w.Write([]byte(strconv.Itoa(result)))
}
