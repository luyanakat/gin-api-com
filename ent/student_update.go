// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gin-api/ent/predicate"
	"gin-api/ent/student"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// SetAge sets the "age" field.
func (su *StudentUpdate) SetAge(s string) *StudentUpdate {
	su.mutation.SetAge(s)
	return su
}

// SetSchool sets the "school" field.
func (su *StudentUpdate) SetSchool(s string) *StudentUpdate {
	su.mutation.SetSchool(s)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *StudentUpdate) SetCreatedAt(t time.Time) *StudentUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *StudentUpdate) SetNillableCreatedAt(t *time.Time) *StudentUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StudentUpdate) SetUpdatedAt(t time.Time) *StudentUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks[int, StudentMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StudentUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeString, value)
	}
	if value, ok := su.mutation.School(); ok {
		_spec.SetField(student.FieldSchool, field.TypeString, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(student.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetAge sets the "age" field.
func (suo *StudentUpdateOne) SetAge(s string) *StudentUpdateOne {
	suo.mutation.SetAge(s)
	return suo
}

// SetSchool sets the "school" field.
func (suo *StudentUpdateOne) SetSchool(s string) *StudentUpdateOne {
	suo.mutation.SetSchool(s)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *StudentUpdateOne) SetCreatedAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableCreatedAt(t *time.Time) *StudentUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StudentUpdateOne) SetUpdatedAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	suo.defaults()
	return withHooks[*Student, StudentMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StudentUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeString, value)
	}
	if value, ok := suo.mutation.School(); ok {
		_spec.SetField(student.FieldSchool, field.TypeString, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(student.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
