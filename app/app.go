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

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", webPort), router))
}
