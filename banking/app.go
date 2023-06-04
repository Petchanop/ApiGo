/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   app.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/05 01:37:57 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/05 01:39:37 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/", greeting)
	http.HandleFunc("/customer", getAllCustomers)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
