package main

import (
	"errors"
	"sync"
)

type EmployeeStore struct {
	mu        sync.RWMutex
	employees map[int]Employee
	nextID    int
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]Employee),
		nextID:    1,
	}
}

func (s *EmployeeStore) CreateEmployee(e Employee) Employee {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.ID = s.nextID
	s.employees[s.nextID] = e
	s.nextID++
	return e
}

func (s *EmployeeStore) GetEmployeeByID(id int) (Employee, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, exists := s.employees[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}
	return e, nil
}

func (s *EmployeeStore) UpdateEmployee(id int, e Employee) (Employee, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.employees[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}
	e.ID = id
	s.employees[id] = e
	return e, nil
}

func (s *EmployeeStore) DeleteEmployee(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.employees[id]
	if !exists {
		return errors.New("employee not found")
	}
	delete(s.employees, id)
	return nil
}
