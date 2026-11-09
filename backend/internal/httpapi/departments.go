package httpapi

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"amerta/backend/internal/store"
)

type deptInput struct {
	Name string `json:"name"`
}

type empView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type deptView struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	EmployeeCount int       `json:"employee_count"`
	Employees     []empView `json:"employees"`
}

func toDept(d store.Department) deptView {
	emps := make([]empView, 0, len(d.Employees))
	for _, e := range d.Employees {
		emps = append(emps, empView{ID: e.ID, Name: e.Name})
	}
	return deptView{
		ID:            d.ID,
		Name:          d.Name,
		EmployeeCount: d.EmployeeCount,
		Employees:     emps,
	}
}

func deptName(w http.ResponseWriter, r *http.Request) (string, bool) {
	var in deptInput
	if err := decode(w, r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "Permintaan tidak valid.")
		return "", false
	}
	name := strings.ToUpper(strings.TrimSpace(in.Name))
	if name == "" || utf8.RuneCountInString(name) > 100 {
		writeError(w, http.StatusUnprocessableEntity, "Nama departemen wajib diisi (maksimal 100 karakter).")
		return "", false
	}
	return name, true
}

func (s *server) listDepartments(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	depts, err := s.store.Departments(ctx)
	if err != nil {
		s.logger.Error("list departments", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal memuat departemen.")
		return
	}

	list := make([]deptView, 0, len(depts))
	for _, d := range depts {
		list = append(list, toDept(d))
	}
	writeJSON(w, http.StatusOK, map[string]any{"departments": list})
}

func (s *server) createDepartment(w http.ResponseWriter, r *http.Request) {
	name, ok := deptName(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	d, err := s.store.CreateDepartment(ctx, name)
	if err != nil {
		if isUnique(err) {
			writeError(w, http.StatusConflict, "Nama departemen sudah digunakan.")
			return
		}
		s.logger.Error("create department", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal menambah departemen.")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"department": toDept(d)})
}

func (s *server) updateDepartment(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "ID departemen tidak valid.")
		return
	}
	name, ok := deptName(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	d, err := s.store.UpdateDepartment(ctx, id, name)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusNotFound, "Departemen tidak ditemukan.")
			return
		}
		if isUnique(err) {
			writeError(w, http.StatusConflict, "Nama departemen sudah digunakan.")
			return
		}
		s.logger.Error("update department", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal mengubah departemen.")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"department": toDept(d)})
}

func (s *server) deleteDepartment(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "ID departemen tidak valid.")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := s.store.DeleteDepartment(ctx, id); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusNotFound, "Departemen tidak ditemukan.")
			return
		}
		if errors.Is(err, store.ErrHasEmployees) {
			writeError(w, http.StatusConflict, "Departemen masih memiliki karyawan.")
			return
		}
		s.logger.Error("delete department", "error", err)
		writeError(w, http.StatusInternalServerError, "Gagal menghapus departemen.")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
