package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KuSeMi/banking/domain"
	"github.com/KuSeMi/banking/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	requiredEnv := []string{"HOST", "WEB_PORT", "MYSQL_HOST",
		"MYSQL_USER", "MYSQL_PASSWORD", "MYSQL_DATABASE"}
	for _, env := range requiredEnv {
		if os.Getenv(env) == "" {
			log.Fatalf("Environment variable %s not defined...", env)
		}
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	// database
	dbClient := getDbClient()

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	//routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	host := os.Getenv("HOST")
	webPort := os.Getenv("WEB_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, webPort), router))
}
