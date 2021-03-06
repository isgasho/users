package schema

import (
	"time"

	"github.com/facebook/ent/schema/edge"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Workspace holds the schema definition for the Workspace entity.
type Workspace struct {
	ent.Schema
}

// Annotations of the Workspace.
func (Workspace) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "workspacesx"},
	}
}

// Fields of the Workspace.
func (Workspace) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("plan").NotEmpty(),
		field.String("description").NotEmpty().Optional(),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Workspace.
func (Workspace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("workspace").Unique().Required(),
		edge.To("workspace_roles", WorkspaceRole.Type),
		edge.To("groups", Group.Type),
	}
}
