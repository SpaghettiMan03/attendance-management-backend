package handler

import (
	"context"
	"fmt"

	schema "attendance-management-backend/pkg/schema/gen/server"
)

type EmployeeHandler struct {}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{}
}

func (h *EmployeeHandler) List(
	ctx context.Context,
	req *schema.ListRequest,
	) (*schema.ListResponse, error) {
	fmt.Println(ctx)
	fmt.Println(req)
	fmt.Println("お試し")

	employee := &schema.Employee{
		FirstName: "名前",
		LastName: "苗字",
		Birthday: &schema.Date{
			Day: 4,
			Month: 12,
			Year: 12,
		},
		Gender: schema.Employee_FEMALE,
		Position: schema.Employee_PART_TIME,
	}

	employees := []*schema.Employee{ employee }

	resp := &schema.ListResponse{
		Employees: employees,
	}

	return resp, nil
}