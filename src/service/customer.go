package service

import (
	"context"
	"fmt"
	"log"

	"bsm.com/gqlgengofirestore/graph/model"
	"bsm.com/gqlgengofirestore/src/repository"

	"google.golang.org/api/iterator"
)

func FechCustomers() []*model.Customer {
	var customers []*model.Customer
	ctx := context.Background()
	client := repository.CreateFirestoreClient()
	//	defer client.Close()
	iter := client.Collection("Customer").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		customerMap := doc.Data()
		fmt.Println("customerMap: ", customerMap)
		var cust = &model.Customer{FirstName: customerMap["FIRST_NAME"].(string),
			LastName:  customerMap["LAST_NAME"].(string),
			Dob:       customerMap["DOB"].(string),
			Gender:    customerMap["GENDER"].(string),
			ContactNo: customerMap["CONTACT_NO"].(string)}
		customers = append(customers, cust)
	}

	return customers
}

func CreateCustomer(customerReq model.CustomerReq) *model.Customer {
	ctx := context.Background()
	client := repository.CreateFirestoreClient()
	//	defer client.Close()

	var cust = &model.Customer{FirstName: customerReq.FirstName,
		LastName:  customerReq.LastName,
		Dob:       customerReq.Dob,
		Gender:    customerReq.Gender,
		ContactNo: customerReq.ContactNo}

	_, _, err := client.Collection("Customer").Add(ctx, map[string]interface{}{
		"FIRST_NAME": customerReq.FirstName,
		"LAST_NAME":  customerReq.LastName,
		"GENDER":     customerReq.Gender,
		"CONTACT_NO": customerReq.ContactNo,
		"DOB":        customerReq.Dob,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	return cust
}
