package chain

import "testing"

func TestChain(t *testing.T) {
	finance := NewFinanceHandler("小李")
	department := NewDepartmentHeadHandler("小强")
	department.SetSuccessor(finance)
	manager := NewManagerHandler("小明")
	manager.SetSuccessor(department)

	request1 := NewReimbursementRequest(799.00, "王五北京出差高铁报销费用")
	request2 := NewReimbursementRequest(3999.00, "王五上海出差8天住宿报销费用")
	request3 := NewReimbursementRequest(4500.00, "王五杭州出差15天住宿报销费用")

	manager.Handle(request1) //经理能处理

	manager.Handle(request2) //经理不能处理，转给部门主管处理

	manager.Handle(request3) //经理不能处理，部分主管
}
