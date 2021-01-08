package transactionsrepo

import (
	"encoding/json"
	"example.com/test_task/core/domain"
)

func (repo *repository)Get(id string) (*domain.Transaction,error){
	data,err :=repo.storageClient.Get(id)
	if err != nil{
		return nil,err
	}

	if len(data) == 0{
		return nil,nil
	}

	transaction := domain.Transaction{}
	err = json.Unmarshal(data,&transaction)
	if err != nil{
		return nil,err
	}

	return &transaction,nil
}
func (repo *repository)GetByAccountID(accountId string) ([]domain.Transaction,error){
	data,err :=repo.storageClient.GetBy("account",accountId)
	if err != nil{
		return nil,err
	}
	transactionList := []domain.Transaction{}

	for _,d := range data{
		transaction :=domain.Transaction{}
		err = json.Unmarshal(d,&transaction)
		if err != nil{
			return nil,err
		}
		transactionList = append(transactionList, transaction)
	}


	return transactionList,nil
}
func (repo *repository)Save(transaction *domain.Transaction) error{
	dataToSave,err := json.Marshal(transaction)
	if err != nil{
		return err
	}
	err = repo.storageClient.Save(transaction.ID,dataToSave)
	if err != nil{
		return err
	}
	return nil
}