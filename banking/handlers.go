/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handlers.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/05 01:35:23 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/05 01:38:44 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello world.")
}

func getAllCustomers(w http.ResponseWriter, req *http.Request) {
	customers := []Customer{
		{Name: "Rati", City: "Bangkok", Zipcode: "10000"},
		{Name: "Sivaluck", City: "Chiangmai", Zipcode: "50000"},
	}
	if req.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
