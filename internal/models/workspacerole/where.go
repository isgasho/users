// Code generated (@generated) by entc, DO NOT EDIT.

package workspacerole

import (
	"time"

	"github.com/adnaan/users/internal/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their identifier.
func ID(id uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// WorkspaceID applies equality check predicate on the "workspace_id" field. It's identical to WorkspaceIDEQ.
func WorkspaceID(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkspaceID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// WorkspaceIDEQ applies the EQ predicate on the "workspace_id" field.
func WorkspaceIDEQ(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDNEQ applies the NEQ predicate on the "workspace_id" field.
func WorkspaceIDNEQ(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDIn applies the In predicate on the "workspace_id" field.
func WorkspaceIDIn(vs ...uuid.UUID) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkspaceID), v...))
	})
}

// WorkspaceIDNotIn applies the NotIn predicate on the "workspace_id" field.
func WorkspaceIDNotIn(vs ...uuid.UUID) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkspaceID), v...))
	})
}

// WorkspaceIDGT applies the GT predicate on the "workspace_id" field.
func WorkspaceIDGT(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDGTE applies the GTE predicate on the "workspace_id" field.
func WorkspaceIDGTE(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDLT applies the LT predicate on the "workspace_id" field.
func WorkspaceIDLT(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkspaceID), v))
	})
}

// WorkspaceIDLTE applies the LTE predicate on the "workspace_id" field.
func WorkspaceIDLTE(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkspaceID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMetadata)))
	})
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMetadata)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WorkspaceRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasWorkspaces applies the HasEdge predicate on the "workspaces" edge.
func HasWorkspaces() predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspacesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkspacesTable, WorkspacesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkspacesWith applies the HasEdge predicate on the "workspaces" edge with a given conditions (other predicates).
func HasWorkspacesWith(preds ...predicate.Workspace) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkspacesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkspacesTable, WorkspacesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.WorkspaceRole) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.WorkspaceRole) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkspaceRole) predicate.WorkspaceRole {
	return predicate.WorkspaceRole(func(s *sql.Selector) {
		p(s.Not())
	})
}
