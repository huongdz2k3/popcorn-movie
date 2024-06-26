// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/food"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Food is the model entity for the Food schema.
type Food struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FoodQuery when eager-loading is set.
	Edges        FoodEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FoodEdges holds the relations/edges for other nodes in the graph.
type FoodEdges struct {
	// FoodOrderLines holds the value of the food_order_lines edge.
	FoodOrderLines []*FoodOrderLine `json:"food_order_lines,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FoodOrderLinesOrErr returns the FoodOrderLines value or an error if the edge
// was not loaded in eager-loading.
func (e FoodEdges) FoodOrderLinesOrErr() ([]*FoodOrderLine, error) {
	if e.loadedTypes[0] {
		return e.FoodOrderLines, nil
	}
	return nil, &NotLoadedError{edge: "food_order_lines"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Food) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case food.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case food.FieldName, food.FieldImage:
			values[i] = new(sql.NullString)
		case food.FieldCreatedAt, food.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case food.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Food fields.
func (f *Food) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case food.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case food.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case food.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				f.Price = value.Float64
			}
		case food.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				f.Image = value.String
			}
		case food.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case food.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Food.
// This includes values selected through modifiers, order, etc.
func (f *Food) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryFoodOrderLines queries the "food_order_lines" edge of the Food entity.
func (f *Food) QueryFoodOrderLines() *FoodOrderLineQuery {
	return NewFoodClient(f.config).QueryFoodOrderLines(f)
}

// Update returns a builder for updating this Food.
// Note that you need to call Food.Unwrap() before calling this method if this Food
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Food) Update() *FoodUpdateOne {
	return NewFoodClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Food entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Food) Unwrap() *Food {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Food is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Food) String() string {
	var builder strings.Builder
	builder.WriteString("Food(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", f.Price))
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(f.Image)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Foods is a parsable slice of Food.
type Foods []*Food
