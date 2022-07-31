package node

import (
	"log"
	"time"
)

type Customer struct {
	ID    string
	First string
	Last  string
}

type Merchant struct {
	ID      string
	Name    string
	Address string
	Phone   string
}

type Supplier struct {
	ID      string
	Name    string
	Address string
	Phone   string
}

// Create a new Customer{}
func NewCustomer(first, last string) Customer {
	if first == "" || last == "" {
		log.Fatalf("Empty first or last name.")
	}

	return Customer{
		ID:    Hash(first + last + Now(time.Now())),
		First: first,
		Last:  last,
	}
}

// Create a new Merchant{}
func NewMerchant(name, address, phone string) Merchant {
	if name == "" {
		log.Fatalf("Empty name.")
	} else if address == "" {
		log.Fatalf("Empty address.")
	} else if phone == "" {
		log.Fatalf("Empty phone number.")
	}

	return Merchant{
		ID:      Hash(name + address + phone + Now(time.Now())),
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

// Create a new Supplier{}
func NewSupplier(name, address, phone string) Supplier {
	if name == "" {
		log.Fatalf("Empty name.")
	} else if address == "" {
		log.Fatalf("Empty address.")
	} else if phone == "" {
		log.Fatalf("Empty phone number.")
	}

	return Supplier{
		ID:      Hash(name + address + phone + Now(time.Now())),
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}
