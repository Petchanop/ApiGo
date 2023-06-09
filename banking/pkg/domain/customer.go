/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   customer.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/06 15:22:48 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/08 14:54:37 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package domain

type Customer struct {
	// Name    string `json:"full_name" xml:"name"`
	// City    string `json:"city" xml:"city"`
	// Zipcode string `json:"zip_code" xml:"zipcode"`
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}
