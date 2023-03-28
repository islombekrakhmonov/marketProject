package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"products/model"	
)

var products []string
var collection string
var summarize int
var readCompany = ReadCompany()
var readProduct = ReaderProduct()
var readUser = ReaderUser()
var userInput = GetUserInput()


func ReaderProduct() []model.Product{
	data, err := ioutil.ReadFile("../data/productJson.json")
	if err != nil{
		fmt.Println(err)
		return nil
	}
	var something []model.Product
	err = json.Unmarshal(data, &something)
		if err != nil{
			fmt.Println(err)
		}
	return something
}

func ReadCompany() []model.Company{
	data, err := ioutil.ReadFile("../data/companyJson.json")
	if err != nil{
		fmt.Println(err)
		return nil
	}
	var something []model.Company
	err = json.Unmarshal(data, &something)
		if err != nil{
			fmt.Println(err)
		}
	return something
}
func ReaderUser() []model.User{
	data, err := ioutil.ReadFile("../data/userJson.json")
	if err != nil{
		fmt.Println(err)
		return nil
	}
	var something []model.User
	err = json.Unmarshal(data, &something)
		if err != nil{
			fmt.Println(err)
		}
	return something
}

func GetUserInput()string{
	var userText string
	readCompany := ReadCompany()
	fmt.Println("Welcome to ", readCompany[0].Name)
	fmt.Println("Please enter your name")
	fmt.Scan(&userText)
	return userText
}

func CheckUser(userName string) (model.User,error) {
	for i := range readUser{
		if userName == readUser[i].Name{
			return readUser[i], nil
		} 
	} 
	return model.User{}, errors.New("You entered wrong name or your name is not in our system")
}

func Smth(){
	user, err := CheckUser(userInput)
    if err != nil {
		panic(err)
	} else {
		fmt.Println("Please choose your products")
		for i := range readProduct{
			fmt.Println("-----",readProduct[i].Name, "-----", readProduct[i].Price,"-----")
		}
	}

	for {
		fmt.Scan(&collection)
		products = append(products, collection)
		if collection == "stop" {
			fmt.Println("Stopped")
			fmt.Println("You have selected", products[:len(products)-1])
			break
		}
	}

	for k:= range products{
		for t := range readProduct{
			if products[k] == readProduct[t].Name{
				summarize += readProduct[t].Price
			}
		}	
		}
		fmt.Printf("Your total is %v\n",summarize)

	for p := range readUser{
		if user == readUser[p]{
			readUser[p].Balance -= summarize
			UserJsonWrite()
		}
	}
	CompanyJsonWrite() 
}

func UserJsonWrite(){
	newUserJson, err := json.Marshal(readUser)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("../data/userJson.json", newUserJson, 0644)
	if err != nil {
		panic(err)
	}
}

func CompanyJsonWrite(){
	readCompany[0].Balance += summarize;
    newCompanyBalance, err := json.Marshal(readCompany)
    if err != nil{
	   panic(err)
    } 
    err = ioutil.WriteFile("../data/companyJson.json", newCompanyBalance, 0644)
    if err != nil{
	panic(err)
    }
}