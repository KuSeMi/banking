package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KuSeMi/banking/domain"
	"github.com/KuSeMi/banking/service"
	"github.com/gorilla/mux"
)

const webPort = 8000

func Start() {
	router := mux.NewRouter()

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	//routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", webPort), router))
}
