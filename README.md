# Employee Management API

This Go application provides a RESTful API for managing employees with CRUD operations and pagination.

## Usage

1. **Clone Repository**:

   ```sh
   git clone https://github.com/yourusername/employee-management.git
   cd employee-management
   ```

2. **Run Application**:

   ```sh
   go run .
   ```

3. **API Endpoints**:

   - Create Employee: `POST /employees`
   - Get Employee by ID: `GET /employees/{id}`
   - Update Employee: `PUT /employees/{id}`
   - Delete Employee: `DELETE /employees/{id}`
   - List Employees with Pagination: `GET /employees?page={page}&limit={limit}`

## Testing

Run unit tests with:

```sh
go test ./...
```

## Author

- Naman Agarwal
- agarwalnaman427@gmail.com


### Note

- Replace `{id}` with the employee ID.
- Adjust pagination parameters `{page}` and `{limit}` as needed.
