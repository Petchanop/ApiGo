/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   app.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/05 01:37:57 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/07 22:06:28 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package app

import (
	"log"
	"net/http"

	"github.com/ApiGo/banking/pkg/handlers"
	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	//define routes
	// router.HandleFunc("/", handlers.Greeting)
	router.HandleFunc("/", handlers.Greeting).Methods(http.MethodGet)
	// router.HandleFunc("/customer", handlers.GetAllCustomers)
	router.HandleFunc("/customer", handlers.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer", handlers.CreateCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/customer/{customer_id}", handlers.GetCustomers)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", handlers.GetCustomers)

	//starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
