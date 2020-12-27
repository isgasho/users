// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/adnaan/users/internal/models/userrole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// UserRoleDelete is the builder for deleting a UserRole entity.
type UserRoleDelete struct {
	config
	hooks    []Hook
	mutation *UserRoleMutation
}

// Where adds a new predicate to the delete builder.
func (urd *UserRoleDelete) Where(ps ...predicate.UserRole) *UserRoleDelete {
	urd.mutation.predicates = append(urd.mutation.predicates, ps...)
	return urd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urd *UserRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(urd.hooks) == 0 {
		affected, err = urd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			urd.mutation = mutation
			affected, err = urd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(urd.hooks) - 1; i >= 0; i-- {
			mut = urd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, urd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (urd *UserRoleDelete) ExecX(ctx context.Context) int {
	n, err := urd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urd *UserRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userrole.FieldID,
			},
		},
	}
	if ps := urd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, urd.driver, _spec)
}

// UserRoleDeleteOne is the builder for deleting a single UserRole entity.
type UserRoleDeleteOne struct {
	urd *UserRoleDelete
}

// Exec executes the deletion query.
func (urdo *UserRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := urdo.urd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urdo *UserRoleDeleteOne) ExecX(ctx context.Context) {
	urdo.urd.ExecX(ctx)
}
