// Code generated by ent, DO NOT EDIT.

package movie

import (
	"PopcornMovie/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldTitle, v))
}

// Genre applies equality check predicate on the "genre" field. It's identical to GenreEQ.
func Genre(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldGenre, v))
}

// Language applies equality check predicate on the "language" field. It's identical to LanguageEQ.
func Language(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldLanguage, v))
}

// Director applies equality check predicate on the "director" field. It's identical to DirectorEQ.
func Director(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldDirector, v))
}

// Cast applies equality check predicate on the "cast" field. It's identical to CastEQ.
func Cast(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldCast, v))
}

// Poster applies equality check predicate on the "poster" field. It's identical to PosterEQ.
func Poster(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldPoster, v))
}

// Rated applies equality check predicate on the "rated" field. It's identical to RatedEQ.
func Rated(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldRated, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldDuration, v))
}

// Trailer applies equality check predicate on the "trailer" field. It's identical to TrailerEQ.
func Trailer(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldTrailer, v))
}

// OpeningDay applies equality check predicate on the "opening_day" field. It's identical to OpeningDayEQ.
func OpeningDay(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldOpeningDay, v))
}

// Story applies equality check predicate on the "story" field. It's identical to StoryEQ.
func Story(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldStory, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldTitle, v))
}

// GenreEQ applies the EQ predicate on the "genre" field.
func GenreEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldGenre, v))
}

// GenreNEQ applies the NEQ predicate on the "genre" field.
func GenreNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldGenre, v))
}

// GenreIn applies the In predicate on the "genre" field.
func GenreIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldGenre, vs...))
}

// GenreNotIn applies the NotIn predicate on the "genre" field.
func GenreNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldGenre, vs...))
}

// GenreGT applies the GT predicate on the "genre" field.
func GenreGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldGenre, v))
}

// GenreGTE applies the GTE predicate on the "genre" field.
func GenreGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldGenre, v))
}

// GenreLT applies the LT predicate on the "genre" field.
func GenreLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldGenre, v))
}

// GenreLTE applies the LTE predicate on the "genre" field.
func GenreLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldGenre, v))
}

// GenreContains applies the Contains predicate on the "genre" field.
func GenreContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldGenre, v))
}

// GenreHasPrefix applies the HasPrefix predicate on the "genre" field.
func GenreHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldGenre, v))
}

// GenreHasSuffix applies the HasSuffix predicate on the "genre" field.
func GenreHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldGenre, v))
}

// GenreEqualFold applies the EqualFold predicate on the "genre" field.
func GenreEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldGenre, v))
}

// GenreContainsFold applies the ContainsFold predicate on the "genre" field.
func GenreContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldGenre, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldStatus, vs...))
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldLanguage, v))
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldLanguage, v))
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldLanguage, vs...))
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldLanguage, vs...))
}

// LanguageGT applies the GT predicate on the "language" field.
func LanguageGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldLanguage, v))
}

// LanguageGTE applies the GTE predicate on the "language" field.
func LanguageGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldLanguage, v))
}

// LanguageLT applies the LT predicate on the "language" field.
func LanguageLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldLanguage, v))
}

// LanguageLTE applies the LTE predicate on the "language" field.
func LanguageLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldLanguage, v))
}

// LanguageContains applies the Contains predicate on the "language" field.
func LanguageContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldLanguage, v))
}

// LanguageHasPrefix applies the HasPrefix predicate on the "language" field.
func LanguageHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldLanguage, v))
}

// LanguageHasSuffix applies the HasSuffix predicate on the "language" field.
func LanguageHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldLanguage, v))
}

// LanguageEqualFold applies the EqualFold predicate on the "language" field.
func LanguageEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldLanguage, v))
}

// LanguageContainsFold applies the ContainsFold predicate on the "language" field.
func LanguageContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldLanguage, v))
}

// DirectorEQ applies the EQ predicate on the "director" field.
func DirectorEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldDirector, v))
}

// DirectorNEQ applies the NEQ predicate on the "director" field.
func DirectorNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldDirector, v))
}

// DirectorIn applies the In predicate on the "director" field.
func DirectorIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldDirector, vs...))
}

// DirectorNotIn applies the NotIn predicate on the "director" field.
func DirectorNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldDirector, vs...))
}

// DirectorGT applies the GT predicate on the "director" field.
func DirectorGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldDirector, v))
}

// DirectorGTE applies the GTE predicate on the "director" field.
func DirectorGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldDirector, v))
}

// DirectorLT applies the LT predicate on the "director" field.
func DirectorLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldDirector, v))
}

// DirectorLTE applies the LTE predicate on the "director" field.
func DirectorLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldDirector, v))
}

// DirectorContains applies the Contains predicate on the "director" field.
func DirectorContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldDirector, v))
}

// DirectorHasPrefix applies the HasPrefix predicate on the "director" field.
func DirectorHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldDirector, v))
}

// DirectorHasSuffix applies the HasSuffix predicate on the "director" field.
func DirectorHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldDirector, v))
}

// DirectorEqualFold applies the EqualFold predicate on the "director" field.
func DirectorEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldDirector, v))
}

// DirectorContainsFold applies the ContainsFold predicate on the "director" field.
func DirectorContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldDirector, v))
}

// CastEQ applies the EQ predicate on the "cast" field.
func CastEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldCast, v))
}

// CastNEQ applies the NEQ predicate on the "cast" field.
func CastNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldCast, v))
}

// CastIn applies the In predicate on the "cast" field.
func CastIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldCast, vs...))
}

// CastNotIn applies the NotIn predicate on the "cast" field.
func CastNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldCast, vs...))
}

// CastGT applies the GT predicate on the "cast" field.
func CastGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldCast, v))
}

// CastGTE applies the GTE predicate on the "cast" field.
func CastGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldCast, v))
}

// CastLT applies the LT predicate on the "cast" field.
func CastLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldCast, v))
}

// CastLTE applies the LTE predicate on the "cast" field.
func CastLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldCast, v))
}

// CastContains applies the Contains predicate on the "cast" field.
func CastContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldCast, v))
}

// CastHasPrefix applies the HasPrefix predicate on the "cast" field.
func CastHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldCast, v))
}

// CastHasSuffix applies the HasSuffix predicate on the "cast" field.
func CastHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldCast, v))
}

// CastEqualFold applies the EqualFold predicate on the "cast" field.
func CastEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldCast, v))
}

// CastContainsFold applies the ContainsFold predicate on the "cast" field.
func CastContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldCast, v))
}

// PosterEQ applies the EQ predicate on the "poster" field.
func PosterEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldPoster, v))
}

// PosterNEQ applies the NEQ predicate on the "poster" field.
func PosterNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldPoster, v))
}

// PosterIn applies the In predicate on the "poster" field.
func PosterIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldPoster, vs...))
}

// PosterNotIn applies the NotIn predicate on the "poster" field.
func PosterNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldPoster, vs...))
}

// PosterGT applies the GT predicate on the "poster" field.
func PosterGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldPoster, v))
}

// PosterGTE applies the GTE predicate on the "poster" field.
func PosterGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldPoster, v))
}

// PosterLT applies the LT predicate on the "poster" field.
func PosterLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldPoster, v))
}

// PosterLTE applies the LTE predicate on the "poster" field.
func PosterLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldPoster, v))
}

// PosterContains applies the Contains predicate on the "poster" field.
func PosterContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldPoster, v))
}

// PosterHasPrefix applies the HasPrefix predicate on the "poster" field.
func PosterHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldPoster, v))
}

// PosterHasSuffix applies the HasSuffix predicate on the "poster" field.
func PosterHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldPoster, v))
}

// PosterEqualFold applies the EqualFold predicate on the "poster" field.
func PosterEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldPoster, v))
}

// PosterContainsFold applies the ContainsFold predicate on the "poster" field.
func PosterContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldPoster, v))
}

// RatedEQ applies the EQ predicate on the "rated" field.
func RatedEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldRated, v))
}

// RatedNEQ applies the NEQ predicate on the "rated" field.
func RatedNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldRated, v))
}

// RatedIn applies the In predicate on the "rated" field.
func RatedIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldRated, vs...))
}

// RatedNotIn applies the NotIn predicate on the "rated" field.
func RatedNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldRated, vs...))
}

// RatedGT applies the GT predicate on the "rated" field.
func RatedGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldRated, v))
}

// RatedGTE applies the GTE predicate on the "rated" field.
func RatedGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldRated, v))
}

// RatedLT applies the LT predicate on the "rated" field.
func RatedLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldRated, v))
}

// RatedLTE applies the LTE predicate on the "rated" field.
func RatedLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldRated, v))
}

// RatedContains applies the Contains predicate on the "rated" field.
func RatedContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldRated, v))
}

// RatedHasPrefix applies the HasPrefix predicate on the "rated" field.
func RatedHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldRated, v))
}

// RatedHasSuffix applies the HasSuffix predicate on the "rated" field.
func RatedHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldRated, v))
}

// RatedEqualFold applies the EqualFold predicate on the "rated" field.
func RatedEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldRated, v))
}

// RatedContainsFold applies the ContainsFold predicate on the "rated" field.
func RatedContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldRated, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldDuration, v))
}

// TrailerEQ applies the EQ predicate on the "trailer" field.
func TrailerEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldTrailer, v))
}

// TrailerNEQ applies the NEQ predicate on the "trailer" field.
func TrailerNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldTrailer, v))
}

// TrailerIn applies the In predicate on the "trailer" field.
func TrailerIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldTrailer, vs...))
}

// TrailerNotIn applies the NotIn predicate on the "trailer" field.
func TrailerNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldTrailer, vs...))
}

// TrailerGT applies the GT predicate on the "trailer" field.
func TrailerGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldTrailer, v))
}

// TrailerGTE applies the GTE predicate on the "trailer" field.
func TrailerGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldTrailer, v))
}

// TrailerLT applies the LT predicate on the "trailer" field.
func TrailerLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldTrailer, v))
}

// TrailerLTE applies the LTE predicate on the "trailer" field.
func TrailerLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldTrailer, v))
}

// TrailerContains applies the Contains predicate on the "trailer" field.
func TrailerContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldTrailer, v))
}

// TrailerHasPrefix applies the HasPrefix predicate on the "trailer" field.
func TrailerHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldTrailer, v))
}

// TrailerHasSuffix applies the HasSuffix predicate on the "trailer" field.
func TrailerHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldTrailer, v))
}

// TrailerEqualFold applies the EqualFold predicate on the "trailer" field.
func TrailerEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldTrailer, v))
}

// TrailerContainsFold applies the ContainsFold predicate on the "trailer" field.
func TrailerContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldTrailer, v))
}

// OpeningDayEQ applies the EQ predicate on the "opening_day" field.
func OpeningDayEQ(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldOpeningDay, v))
}

// OpeningDayNEQ applies the NEQ predicate on the "opening_day" field.
func OpeningDayNEQ(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldOpeningDay, v))
}

// OpeningDayIn applies the In predicate on the "opening_day" field.
func OpeningDayIn(vs ...time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldOpeningDay, vs...))
}

// OpeningDayNotIn applies the NotIn predicate on the "opening_day" field.
func OpeningDayNotIn(vs ...time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldOpeningDay, vs...))
}

// OpeningDayGT applies the GT predicate on the "opening_day" field.
func OpeningDayGT(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldOpeningDay, v))
}

// OpeningDayGTE applies the GTE predicate on the "opening_day" field.
func OpeningDayGTE(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldOpeningDay, v))
}

// OpeningDayLT applies the LT predicate on the "opening_day" field.
func OpeningDayLT(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldOpeningDay, v))
}

// OpeningDayLTE applies the LTE predicate on the "opening_day" field.
func OpeningDayLTE(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldOpeningDay, v))
}

// StoryEQ applies the EQ predicate on the "story" field.
func StoryEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldStory, v))
}

// StoryNEQ applies the NEQ predicate on the "story" field.
func StoryNEQ(v string) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldStory, v))
}

// StoryIn applies the In predicate on the "story" field.
func StoryIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldStory, vs...))
}

// StoryNotIn applies the NotIn predicate on the "story" field.
func StoryNotIn(vs ...string) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldStory, vs...))
}

// StoryGT applies the GT predicate on the "story" field.
func StoryGT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldStory, v))
}

// StoryGTE applies the GTE predicate on the "story" field.
func StoryGTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldStory, v))
}

// StoryLT applies the LT predicate on the "story" field.
func StoryLT(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldStory, v))
}

// StoryLTE applies the LTE predicate on the "story" field.
func StoryLTE(v string) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldStory, v))
}

// StoryContains applies the Contains predicate on the "story" field.
func StoryContains(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContains(FieldStory, v))
}

// StoryHasPrefix applies the HasPrefix predicate on the "story" field.
func StoryHasPrefix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasPrefix(FieldStory, v))
}

// StoryHasSuffix applies the HasSuffix predicate on the "story" field.
func StoryHasSuffix(v string) predicate.Movie {
	return predicate.Movie(sql.FieldHasSuffix(FieldStory, v))
}

// StoryEqualFold applies the EqualFold predicate on the "story" field.
func StoryEqualFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldEqualFold(FieldStory, v))
}

// StoryContainsFold applies the ContainsFold predicate on the "story" field.
func StoryContainsFold(v string) predicate.Movie {
	return predicate.Movie(sql.FieldContainsFold(FieldStory, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Movie {
	return predicate.Movie(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasShowTimes applies the HasEdge predicate on the "showTimes" edge.
func HasShowTimes() predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ShowTimesTable, ShowTimesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShowTimesWith applies the HasEdge predicate on the "showTimes" edge with a given conditions (other predicates).
func HasShowTimesWith(preds ...predicate.ShowTime) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		step := newShowTimesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Movie) predicate.Movie {
	return predicate.Movie(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Movie) predicate.Movie {
	return predicate.Movie(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Movie) predicate.Movie {
	return predicate.Movie(sql.NotPredicates(p))
}
