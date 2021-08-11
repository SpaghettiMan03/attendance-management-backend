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

	employee1 := &schema.Employee{
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

	employe2 := &schema.Employee{
		FirstName: "太郎",
		LastName: "田中",
		Birthday: &schema.Date{
			Day: 5,
			Month: 6,
			Year: 15,
		},
		Gender: schema.Employee_MALE,
		Position: schema.Employee_FULL_TIME,
	}

	employe3 := &schema.Employee{
		FirstName: "いいい",
		LastName: "内山",
		Birthday: &schema.Date{
			Day: 5,
			Month: 6,
			Year: 15,
		},
		Gender: schema.Employee_NEITHER,
		Position: schema.Employee_FULL_TIME,
	}

	employe4 := &schema.Employee{
		FirstName: "適当",
		LastName: "下田",
		Birthday: &schema.Date{
			Day: 5,
			Month: 6,
			Year: 15,
		},
		Gender: schema.Employee_NOT_ANSWER,
		Position: schema.Employee_FULL_TIME,
	}

	employees := []*schema.Employee{ employee1, employe2, employe3, employe4 }

	resp := &schema.ListResponse{
		Employees: employees,
	}

	return resp, nil
}