package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique(),
		field.String("name"),
		field.String("age"),
		field.String("school"),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
