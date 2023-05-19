package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("content").Annotations(entsql.Annotation{
			Size: 255,
		}),
		field.String("authorId").Annotations(entsql.Annotation{
			Size: 255,
		}),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
