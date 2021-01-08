package accountsrepo

import (
	"encoding/json"
	"example.com/test_task/core/domain"
)

func (repo *repository)Get(id string) (*domain.Account,error){
	data,err :=repo.storageClient.Get(id)
	if err != nil{
		return nil,err
	}

	if len(data) == 0{
		return nil,nil
	}

	account := domain.Account{}
	err = json.Unmarshal(data,&account)
	if err != nil{
		return nil,err
	}

	return &account,nil
}

func (repo *repository)Save(account *domain.Account) error{
	dataToSave,err := json.Marshal(account)
	if err != nil{
		return err
	}
	err = repo.storageClient.Save(account.ID,dataToSave)
	if err != nil{
		return err
	}
	return nil
}