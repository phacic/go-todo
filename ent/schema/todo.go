package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("task"),
		field.Bool("completed").
			Default(false).StructTag(`json:"completed"`),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}

func (Todo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task"),
	}
}
