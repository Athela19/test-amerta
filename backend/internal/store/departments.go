package store

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

var ErrHasEmployees = errors.New("department has employees")

type Employee struct {
	ID   int64
	Name string
}

type Department struct {
	ID            int64
	Name          string
	EmployeeCount int
	Employees     []Employee
}

const deptMatch = `
	$1 = ''
	OR position($1 in lower(d.name)) > 0
	OR position($1 in d.id::text) > 0
	OR EXISTS (
		SELECT 1 FROM employees em
		WHERE em.department_id = d.id
		  AND position($1 in lower(em.name)) > 0
	)`

func (s *Store) Departments(ctx context.Context, query string, limit, offset int) ([]Department, int, error) {
	var total int
	err := s.pool.QueryRow(ctx, `
		SELECT COUNT(*)::int
		FROM departments d
		WHERE `+deptMatch, query).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := s.pool.Query(ctx, `
		SELECT d.id, d.name, COUNT(e.id)::int
		FROM departments d
		LEFT JOIN employees e ON e.department_id = d.id
		WHERE `+deptMatch+`
		GROUP BY d.id
		ORDER BY d.id
		LIMIT $2 OFFSET $3
	`, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	depts := make([]Department, 0)
	index := make(map[int64]int)
	for rows.Next() {
		var d Department
		if err := rows.Scan(&d.ID, &d.Name, &d.EmployeeCount); err != nil {
			return nil, 0, err
		}
		d.Employees = []Employee{}
		index[d.ID] = len(depts)
		depts = append(depts, d)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	if len(depts) == 0 {
		return depts, total, nil
	}

	args := make([]any, len(depts))
	holders := make([]string, len(depts))
	for i, d := range depts {
		args[i] = d.ID
		holders[i] = fmt.Sprintf("$%d", i+1)
	}
	emps, err := s.pool.Query(ctx, `
		SELECT id, name, department_id
		FROM employees
		WHERE department_id IN (`+strings.Join(holders, ",")+`)
		ORDER BY id
	`, args...)
	if err != nil {
		return nil, 0, err
	}
	defer emps.Close()

	for emps.Next() {
		var e Employee
		var deptID int64
		if err := emps.Scan(&e.ID, &e.Name, &deptID); err != nil {
			return nil, 0, err
		}
		if i, ok := index[deptID]; ok {
			depts[i].Employees = append(depts[i].Employees, e)
		}
	}
	return depts, total, emps.Err()
}

func (s *Store) CreateDepartment(ctx context.Context, name string) (Department, error) {
	var d Department
	err := s.pool.QueryRow(ctx, `
		INSERT INTO departments (name)
		VALUES ($1)
		RETURNING id, name
	`, name).Scan(&d.ID, &d.Name)
	if err != nil {
		return Department{}, err
	}
	d.Employees = []Employee{}
	return d, nil
}

func (s *Store) UpdateDepartment(ctx context.Context, id int64, name string) (Department, error) {
	var d Department
	err := s.pool.QueryRow(ctx, `
		UPDATE departments
		SET name = $2
		WHERE id = $1
		RETURNING id, name
	`, id, name).Scan(&d.ID, &d.Name)
	if errors.Is(err, pgx.ErrNoRows) {
		return Department{}, ErrNotFound
	}
	if err != nil {
		return Department{}, err
	}
	d.Employees = []Employee{}
	return d, nil
}

func (s *Store) DeleteDepartment(ctx context.Context, id int64) error {
	var n int
	if err := s.pool.QueryRow(ctx, `
		SELECT COUNT(*)::int FROM employees WHERE department_id = $1
	`, id).Scan(&n); err != nil {
		return err
	}
	if n > 0 {
		return ErrHasEmployees
	}

	tag, err := s.pool.Exec(ctx, `DELETE FROM departments WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
