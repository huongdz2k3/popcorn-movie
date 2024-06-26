// Code generated by ent, DO NOT EDIT.

package theater

import (
	"PopcornMovie/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Theater {
	return predicate.Theater(sql.FieldLTE(FieldID, id))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldAddress, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldName, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldPhoneNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldUpdatedAt, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Theater {
	return predicate.Theater(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Theater {
	return predicate.Theater(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Theater {
	return predicate.Theater(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Theater {
	return predicate.Theater(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Theater {
	return predicate.Theater(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Theater {
	return predicate.Theater(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Theater {
	return predicate.Theater(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Theater {
	return predicate.Theater(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Theater {
	return predicate.Theater(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Theater {
	return predicate.Theater(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Theater {
	return predicate.Theater(sql.FieldContainsFold(FieldAddress, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Theater {
	return predicate.Theater(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Theater {
	return predicate.Theater(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Theater {
	return predicate.Theater(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Theater {
	return predicate.Theater(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Theater {
	return predicate.Theater(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Theater {
	return predicate.Theater(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Theater {
	return predicate.Theater(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Theater {
	return predicate.Theater(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Theater {
	return predicate.Theater(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Theater {
	return predicate.Theater(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Theater {
	return predicate.Theater(sql.FieldContainsFold(FieldName, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Theater {
	return predicate.Theater(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Theater {
	return predicate.Theater(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Theater {
	return predicate.Theater(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Theater {
	return predicate.Theater(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Theater {
	return predicate.Theater(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Theater {
	return predicate.Theater(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Theater {
	return predicate.Theater(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Theater {
	return predicate.Theater(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Theater {
	return predicate.Theater(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Theater {
	return predicate.Theater(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Theater {
	return predicate.Theater(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Theater {
	return predicate.Theater(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Theater {
	return predicate.Theater(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasRooms applies the HasEdge predicate on the "rooms" edge.
func HasRooms() predicate.Theater {
	return predicate.Theater(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomsTable, RoomsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomsWith applies the HasEdge predicate on the "rooms" edge with a given conditions (other predicates).
func HasRoomsWith(preds ...predicate.Room) predicate.Theater {
	return predicate.Theater(func(s *sql.Selector) {
		step := newRoomsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Theater) predicate.Theater {
	return predicate.Theater(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Theater) predicate.Theater {
	return predicate.Theater(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Theater) predicate.Theater {
	return predicate.Theater(sql.NotPredicates(p))
}
