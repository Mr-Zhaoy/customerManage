package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// NewCustomer 使用工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// NewCustomer2 第二种创建Customer实例方法，不带id
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// GetInfo 返回用户的信息,格式化的字符串
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id,
		this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info

}
