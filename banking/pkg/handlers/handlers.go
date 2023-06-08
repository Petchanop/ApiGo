/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handlers.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/05 01:35:23 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/08 15:12:18 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ApiGo/banking/pkg/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

func Greeting(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello world.")
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, req *http.Request) {
	// customers := []domain.Customer{
	// 	{Name: "Rati", City: "Bangkok", Zipcode: "10000"},
	// 	{Name: "Sivaluck", City: "Chiangmai", Zipcode: "50000"},
	// }
	customers, _ := ch.Service.GetAllCustomer()
	if req.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func GetCustomers(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCustomer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Post request received.")
}
