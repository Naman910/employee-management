package main

import (
	"testing"
)

func TestEmployeeStore(t *testing.T) {
	store := NewEmployeeStore()

	// Test CreateEmployee
	emp := Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmp := store.CreateEmployee(emp)
	if createdEmp.ID == 0 {
		t.Errorf("Expected non-zero employee ID")
	}

	// Test GetEmployeeByID
	retrievedEmp, err := store.GetEmployeeByID(createdEmp.ID)
	if err != nil {
		t.Errorf("Expected to retrieve employee, got error: %v", err)
	}
	if retrievedEmp.Name != emp.Name {
		t.Errorf("Expected employee name %v, got %v", emp.Name, retrievedEmp.Name)
	}

	// Test UpdateEmployee
	updatedEmp := Employee{Name: "John Smith", Position: "Senior Developer", Salary: 80000}
	updatedEmp, err = store.UpdateEmployee(createdEmp.ID, updatedEmp)
	if err != nil {
		t.Errorf("Expected to update employee, got error: %v", err)
	}
	if updatedEmp.Name != "John Smith" {
		t.Errorf("Expected employee name to be updated to John Smith, got %v", updatedEmp.Name)
	}

	// Test DeleteEmployee
	err = store.DeleteEmployee(createdEmp.ID)
	if err != nil {
		t.Errorf("Expected to delete employee, got error: %v", err)
	}
	_, err = store.GetEmployeeByID(createdEmp.ID)
	if err == nil {
		t.Errorf("Expected error when retrieving deleted employee")
	}
}
