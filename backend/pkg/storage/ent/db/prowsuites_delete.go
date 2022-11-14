// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowsuites"
)

// ProwSuitesDelete is the builder for deleting a ProwSuites entity.
type ProwSuitesDelete struct {
	config
	hooks    []Hook
	mutation *ProwSuitesMutation
}

// Where appends a list predicates to the ProwSuitesDelete builder.
func (psd *ProwSuitesDelete) Where(ps ...predicate.ProwSuites) *ProwSuitesDelete {
	psd.mutation.Where(ps...)
	return psd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (psd *ProwSuitesDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(psd.hooks) == 0 {
		affected, err = psd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProwSuitesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psd.mutation = mutation
			affected, err = psd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psd.hooks) - 1; i >= 0; i-- {
			if psd.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = psd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (psd *ProwSuitesDelete) ExecX(ctx context.Context) int {
	n, err := psd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (psd *ProwSuitesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: prowsuites.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prowsuites.FieldID,
			},
		},
	}
	if ps := psd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, psd.driver, _spec)
}

// ProwSuitesDeleteOne is the builder for deleting a single ProwSuites entity.
type ProwSuitesDeleteOne struct {
	psd *ProwSuitesDelete
}

// Exec executes the deletion query.
func (psdo *ProwSuitesDeleteOne) Exec(ctx context.Context) error {
	n, err := psdo.psd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{prowsuites.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (psdo *ProwSuitesDeleteOne) ExecX(ctx context.Context) {
	psdo.psd.ExecX(ctx)
}