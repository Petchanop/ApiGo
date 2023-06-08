/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   CustomerRepository.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/08 14:41:39 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/08 14:54:55 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package domain

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Thanapol", City: "Bangkok", Zipcode: "10000", DateofBirth: "2010-01-01", Status: "1"},
		{Id: "1002", Name: "Rujirat", City: "Thongdang", Zipcode: "20000", DateofBirth: "2007-01-01", Status: "2"},
	}
	return CustomerRepositoryStub{customers}
}
