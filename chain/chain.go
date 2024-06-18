package chain

import "fmt"

/*
	责任链模式
*/

//报销请求
type ReimbursementRequest struct {
	amount float64
	desc   string
}

func NewReimbursementRequest(amount float64, desc string) *ReimbursementRequest {
	return &ReimbursementRequest{amount: amount, desc: desc}
}

//ReimbursementHandler【抽象处理角色（Handler）】
type ReimbursementHandler interface {
	Handle(request *ReimbursementRequest)
}

//ManagerHandler经理【具体处理者角色（ConcreteHandler）】
type ManagerHandler struct {
	name      string
	successor ReimbursementHandler
}

func NewManagerHandler(name string) *ManagerHandler {
	return &ManagerHandler{name: name}
}
func (m *ManagerHandler) SetSuccessor(successor ReimbursementHandler) {
	m.successor = successor
}

//经理处理报销:800以下可以处理
func (m *ManagerHandler) Handle(request *ReimbursementRequest) {
	if request.amount < 800 {
		fmt.Printf("经理:%s处理报销：%0.2f元，报销情况：%s\n", m.name, request.amount, request.desc)
	} else {
		if m.successor != nil {
			m.successor.Handle(request)
		}
	}
}

//DepartmentHeadHandler部门负责人【具体处理者角色（ConcreteHandler）】
type DepartmentHeadHandler struct {
	name      string
	successor ReimbursementHandler
}

func NewDepartmentHeadHandler(name string) *DepartmentHeadHandler {
	return &DepartmentHeadHandler{name: name}
}
func (d *DepartmentHeadHandler) SetSuccessor(successor ReimbursementHandler) {
	d.successor = successor
}

//部门处理报销:4000以下可以处理
func (d *DepartmentHeadHandler) Handle(request *ReimbursementRequest) {
	if request.amount < 4000 {
		fmt.Printf("部门主管:%s处理报销：%0.2f元，报销情况：%s\n", d.name, request.amount, request.desc)
	} else {
		if d.successor != nil {
			d.successor.Handle(request)
		}
	}
}

//FinanceHandler财务处理人【具体处理者角色（ConcreteHandler）】
type FinanceHandler struct {
	name string
}

func NewFinanceHandler(name string) *FinanceHandler {
	return &FinanceHandler{name: name}
}

//财务处理人处理报销
func (f *FinanceHandler) Handle(request *ReimbursementRequest) {
	fmt.Printf("财务:%s处理报销：%0.2f元，报销情况：%s\n", f.name, request.amount, request.desc)
}
