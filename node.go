package node

import (
	"fmt"
	"log"
	"math"
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

type Seed struct {
	Timestamp      string
	Type           string
	Origin         string
	Destination    string
	AmountSent     float64
	Fee            float64
	AmountReceived float64
	Hash           string
}

/*
Node
*/

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

/*
Seed
*/

// SendC2M returns a Seed{} for a Customer-to-Merchant transaction; Sent to a Receiving Merchant client
func (c *Customer) Send(amount float64) Seed {
	now := Now(time.Now())
	fee := 1.00
	origin := c.ID
	hash := Hash(now + fmt.Sprintf("%v", c) + fmt.Sprintf("%f", amount) + fmt.Sprintf("%f", fee))

	return Seed{
		Timestamp:  now,
		Type:       "C2M",
		Origin:     origin,
		AmountSent: amount,
		Fee:        fee,
		Hash:       hash,
	}
}

// Accepts any Seed{}; Returns a Seed{} with opposite parameters; When stored together in a K/V pair in a database, this will represent a double accounted transaction
func (n *Merchant) Receive(s Seed) Seed {
	r := Seed{
		Timestamp:      Now(time.Now()),
		Type:           s.Type,
		Origin:         s.Origin,
		Destination:    n.ID,
		AmountSent:     s.AmountSent,
		Fee:            s.Fee,
		AmountReceived: math.Round(((s.AmountSent - s.Fee) * 100)) / 100,
		Hash:           s.Hash,
	}

	if Auth(s, r) {
		return r
	}

	return Seed{}
}
