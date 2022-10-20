package service

import (
	"customerManage/model"
	"fmt"
)

// CustomerService  完成对Customer的操作,包括增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

// NewCustomerService 编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// List 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// Add 添加客户到customers切片
func (this *CustomerService) Add(customer model.Customer) bool {

	//我们确定一个分配id的规则,就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// Delete 根据id删除客户(从切片中删除)
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	//如果index == -1, 说明没有这个客户
	if index == -1 {
		return false
	}
	//如何从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

// Update 根据id修改客户
func (this *CustomerService) Update(id int) bool {
	index := this.FindById(id)
	//如果index == -1, 说明没有这个客户
	if index == -1 {
		return false
	}
	//修改用户
	fmt.Printf("姓名(%v)：", this.customers[index].Name)
	name := ""
	fmt.Scanln(&name)
	if name == "" {
	} else {
		this.customers[index].Name = name
	}
	fmt.Printf("性别(%v)：", this.customers[index].Gender)
	gender := ""
	fmt.Scanln(&gender)
	if gender == "" {
	} else {
		this.customers[index].Gender = gender
	}
	fmt.Printf("年龄(%v)：", this.customers[index].Age)
	age := 0
	fmt.Scanln(&age)
	if age == 0 {
	} else {
		this.customers[index].Age = age
	}
	fmt.Printf("电话(%v)：", this.customers[index].Phone)
	phone := ""
	fmt.Scanln(&phone)
	if phone == "" {
	} else {
		this.customers[index].Phone = phone
	}
	fmt.Printf("电邮(%v)：", this.customers[index].Email)
	email := ""
	fmt.Scanln(&email)
	if email == "" {
	} else {
		this.customers[index].Email = email
	}
	return true
}

// FindById 根据id查找客户在切片中对应下标,如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//找到
			index = i
		}
	}
	return index
}
