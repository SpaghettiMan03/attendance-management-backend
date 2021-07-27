package handler

import (
	"context"
	"fmt"

	schema "attendance-management-backend/schema/gen/server"
)

type EmployeeHandler struct {}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{}
}

func (h *EmployeeHandler) List(
	ctx context.Context,
	req *schema.ListRequest,
	) (*schema.ListResponse, error) {
	fmt.Println("お試し")
	return nil, nil
}