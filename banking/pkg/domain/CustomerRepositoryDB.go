/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   CustomerRepositoryDB.go                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/12 15:18:57 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/12 17:20:21 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d *CustomerRepositoryDB) FindAll() ([]Customer, error) {

	// findAllPsql := "select customer_id, name, city, zipcode, age, salary"
	rows, err := d.client.Query(`SELECT * FROM bankcustomer`)
	if err != nil {
		panic(err)
	}

	customer := make([]Customer, 1000)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			fmt.Println(err)
		}
		customer = append(customer, c)
	}
	defer rows.Close()
	return customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "user=petchanop dbname=apigo"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	return CustomerRepositoryDB{db}
}
