/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   customerService.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/08 14:49:33 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/08 14:59:01 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package service

import "github.com/ApiGo/banking/pkg/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
