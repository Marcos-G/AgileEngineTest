package pkg

import "errors"

type StorageClient interface {
	Get(id string) ([]byte, error)
	GetBy(field, value string) ([][]byte, error)
	Save(id string,value []byte) error
}

type AccountsClient struct {
	storage map[string][]byte
}

func (client *AccountsClient) Get(id string) ([]byte, error) {
	return client.storage[id],nil
}

func  (client *AccountsClient)Save(id string,value []byte) error {
	if client.storage == nil{
		client.storage = map[string][]byte{}
	}

	client.storage[id] = value
	return nil
}

func (client *AccountsClient) GetBy(field,value string) ([][]byte, error) {
	return [][]byte{},errors.New("invalid query")
}

type TransactionsClient struct {
	storage map[string][]byte
	storageByAccount map[string][][]byte
}

func (client *TransactionsClient) Get(id string) ([]byte, error) {
	return client.storage[id],nil
}

func (client *TransactionsClient) GetBy(field,value string) ([][]byte, error) {
	if field != "account"{
		return [][]byte{},errors.New("invalid query")
	}
	return client.storageByAccount[value],nil
}

func  (client *TransactionsClient)Save(id string,value []byte) error {
	if client.storage == nil{
		client.storage = map[string][]byte{}
	}
	if client.storageByAccount == nil{
		client.storageByAccount =  map[string][][]byte{}
	}

	client.storage[id] = value
	client.storageByAccount["uniqueAccount"] = append(client.storageByAccount["unique_account"],value)
	return nil
}