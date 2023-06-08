/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   app.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/05 01:37:57 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/08 15:12:31 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package app

import (
	"log"
	"net/http"

	"github.com/ApiGo/banking/pkg/domain"
	"github.com/ApiGo/banking/pkg/handlers"
	"github.com/ApiGo/banking/pkg/service"
	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	//define routes
	// router.HandleFunc("/", handlers.Greeting)
	// router.HandleFunc("/", handlers.Greeting).Methods(http.MethodGet)
	// router.HandleFunc("/customer", handlers.GetAllCustomers)
	//wiring handler from adaptor to service
	ch := handlers.CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customer", ch.GetAllCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customer", handlers.CreateCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/customer/{customer_id}", handlers.GetCustomers)
	// router.HandleFunc("/customer/{customer_id:[0-9]+}", handlers.GetCustomers)
	//starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
