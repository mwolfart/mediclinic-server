// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Specialty is an object representing the database table.
type Specialty struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *specialtyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L specialtyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SpecialtyColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var SpecialtyTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "specialties.id",
	Name: "specialties.name",
}

// Generated where

var SpecialtyWhere = struct {
	ID   whereHelperint
	Name whereHelperstring
}{
	ID:   whereHelperint{field: "\"specialties\".\"id\""},
	Name: whereHelperstring{field: "\"specialties\".\"name\""},
}

// SpecialtyRels is where relationship names are stored.
var SpecialtyRels = struct {
	UseridDoctors string
}{
	UseridDoctors: "UseridDoctors",
}

// specialtyR is where relationships are stored.
type specialtyR struct {
	UseridDoctors DoctorSlice `boil:"UseridDoctors" json:"UseridDoctors" toml:"UseridDoctors" yaml:"UseridDoctors"`
}

// NewStruct creates a new relationship struct
func (*specialtyR) NewStruct() *specialtyR {
	return &specialtyR{}
}

func (r *specialtyR) GetUseridDoctors() DoctorSlice {
	if r == nil {
		return nil
	}
	return r.UseridDoctors
}

// specialtyL is where Load methods for each relationship are stored.
type specialtyL struct{}

var (
	specialtyAllColumns            = []string{"id", "name"}
	specialtyColumnsWithoutDefault = []string{"name"}
	specialtyColumnsWithDefault    = []string{"id"}
	specialtyPrimaryKeyColumns     = []string{"id"}
	specialtyGeneratedColumns      = []string{}
)

type (
	// SpecialtySlice is an alias for a slice of pointers to Specialty.
	// This should almost always be used instead of []Specialty.
	SpecialtySlice []*Specialty
	// SpecialtyHook is the signature for custom Specialty hook methods
	SpecialtyHook func(context.Context, boil.ContextExecutor, *Specialty) error

	specialtyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	specialtyType                 = reflect.TypeOf(&Specialty{})
	specialtyMapping              = queries.MakeStructMapping(specialtyType)
	specialtyPrimaryKeyMapping, _ = queries.BindMapping(specialtyType, specialtyMapping, specialtyPrimaryKeyColumns)
	specialtyInsertCacheMut       sync.RWMutex
	specialtyInsertCache          = make(map[string]insertCache)
	specialtyUpdateCacheMut       sync.RWMutex
	specialtyUpdateCache          = make(map[string]updateCache)
	specialtyUpsertCacheMut       sync.RWMutex
	specialtyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var specialtyAfterSelectMu sync.Mutex
var specialtyAfterSelectHooks []SpecialtyHook

var specialtyBeforeInsertMu sync.Mutex
var specialtyBeforeInsertHooks []SpecialtyHook
var specialtyAfterInsertMu sync.Mutex
var specialtyAfterInsertHooks []SpecialtyHook

var specialtyBeforeUpdateMu sync.Mutex
var specialtyBeforeUpdateHooks []SpecialtyHook
var specialtyAfterUpdateMu sync.Mutex
var specialtyAfterUpdateHooks []SpecialtyHook

var specialtyBeforeDeleteMu sync.Mutex
var specialtyBeforeDeleteHooks []SpecialtyHook
var specialtyAfterDeleteMu sync.Mutex
var specialtyAfterDeleteHooks []SpecialtyHook

var specialtyBeforeUpsertMu sync.Mutex
var specialtyBeforeUpsertHooks []SpecialtyHook
var specialtyAfterUpsertMu sync.Mutex
var specialtyAfterUpsertHooks []SpecialtyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Specialty) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Specialty) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Specialty) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Specialty) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Specialty) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Specialty) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Specialty) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Specialty) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Specialty) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range specialtyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSpecialtyHook registers your hook function for all future operations.
func AddSpecialtyHook(hookPoint boil.HookPoint, specialtyHook SpecialtyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		specialtyAfterSelectMu.Lock()
		specialtyAfterSelectHooks = append(specialtyAfterSelectHooks, specialtyHook)
		specialtyAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		specialtyBeforeInsertMu.Lock()
		specialtyBeforeInsertHooks = append(specialtyBeforeInsertHooks, specialtyHook)
		specialtyBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		specialtyAfterInsertMu.Lock()
		specialtyAfterInsertHooks = append(specialtyAfterInsertHooks, specialtyHook)
		specialtyAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		specialtyBeforeUpdateMu.Lock()
		specialtyBeforeUpdateHooks = append(specialtyBeforeUpdateHooks, specialtyHook)
		specialtyBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		specialtyAfterUpdateMu.Lock()
		specialtyAfterUpdateHooks = append(specialtyAfterUpdateHooks, specialtyHook)
		specialtyAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		specialtyBeforeDeleteMu.Lock()
		specialtyBeforeDeleteHooks = append(specialtyBeforeDeleteHooks, specialtyHook)
		specialtyBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		specialtyAfterDeleteMu.Lock()
		specialtyAfterDeleteHooks = append(specialtyAfterDeleteHooks, specialtyHook)
		specialtyAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		specialtyBeforeUpsertMu.Lock()
		specialtyBeforeUpsertHooks = append(specialtyBeforeUpsertHooks, specialtyHook)
		specialtyBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		specialtyAfterUpsertMu.Lock()
		specialtyAfterUpsertHooks = append(specialtyAfterUpsertHooks, specialtyHook)
		specialtyAfterUpsertMu.Unlock()
	}
}

// One returns a single specialty record from the query.
func (q specialtyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Specialty, error) {
	o := &Specialty{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for specialties")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Specialty records from the query.
func (q specialtyQuery) All(ctx context.Context, exec boil.ContextExecutor) (SpecialtySlice, error) {
	var o []*Specialty

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Specialty slice")
	}

	if len(specialtyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Specialty records in the query.
func (q specialtyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count specialties rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q specialtyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if specialties exists")
	}

	return count > 0, nil
}

// UseridDoctors retrieves all the doctor's Doctors with an executor via userid column.
func (o *Specialty) UseridDoctors(mods ...qm.QueryMod) doctorQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"specialtiesondoctor\" on \"doctor\".\"userid\" = \"specialtiesondoctor\".\"userid\""),
		qm.Where("\"specialtiesondoctor\".\"specialtyid\"=?", o.ID),
	)

	return Doctors(queryMods...)
}

// LoadUseridDoctors allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (specialtyL) LoadUseridDoctors(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpecialty interface{}, mods queries.Applicator) error {
	var slice []*Specialty
	var object *Specialty

	if singular {
		var ok bool
		object, ok = maybeSpecialty.(*Specialty)
		if !ok {
			object = new(Specialty)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSpecialty)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSpecialty))
			}
		}
	} else {
		s, ok := maybeSpecialty.(*[]*Specialty)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSpecialty)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSpecialty))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &specialtyR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &specialtyR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.Select("\"doctor\".\"userid\", \"a\".\"specialtyid\""),
		qm.From("\"doctor\""),
		qm.InnerJoin("\"specialtiesondoctor\" as \"a\" on \"doctor\".\"userid\" = \"a\".\"userid\""),
		qm.WhereIn("\"a\".\"specialtyid\" in ?", argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load doctor")
	}

	var resultSlice []*Doctor

	var localJoinCols []int
	for results.Next() {
		one := new(Doctor)
		var localJoinCol int

		err = results.Scan(&one.Userid, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for doctor")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice doctor")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on doctor")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for doctor")
	}

	if len(doctorAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UseridDoctors = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &doctorR{}
			}
			foreign.R.SpecialtyidSpecialties = append(foreign.R.SpecialtyidSpecialties, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.UseridDoctors = append(local.R.UseridDoctors, foreign)
				if foreign.R == nil {
					foreign.R = &doctorR{}
				}
				foreign.R.SpecialtyidSpecialties = append(foreign.R.SpecialtyidSpecialties, local)
				break
			}
		}
	}

	return nil
}

// AddUseridDoctors adds the given related objects to the existing relationships
// of the specialty, optionally inserting them as new records.
// Appends related to o.R.UseridDoctors.
// Sets related.R.SpecialtyidSpecialties appropriately.
func (o *Specialty) AddUseridDoctors(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Doctor) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"specialtiesondoctor\" (\"specialtyid\", \"userid\") values ($1, $2)"
		values := []interface{}{o.ID, rel.Userid}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &specialtyR{
			UseridDoctors: related,
		}
	} else {
		o.R.UseridDoctors = append(o.R.UseridDoctors, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &doctorR{
				SpecialtyidSpecialties: SpecialtySlice{o},
			}
		} else {
			rel.R.SpecialtyidSpecialties = append(rel.R.SpecialtyidSpecialties, o)
		}
	}
	return nil
}

// SetUseridDoctors removes all previously related items of the
// specialty replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.SpecialtyidSpecialties's UseridDoctors accordingly.
// Replaces o.R.UseridDoctors with related.
// Sets related.R.SpecialtyidSpecialties's UseridDoctors accordingly.
func (o *Specialty) SetUseridDoctors(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Doctor) error {
	query := "delete from \"specialtiesondoctor\" where \"specialtyid\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeUseridDoctorsFromSpecialtyidSpecialtiesSlice(o, related)
	if o.R != nil {
		o.R.UseridDoctors = nil
	}

	return o.AddUseridDoctors(ctx, exec, insert, related...)
}

// RemoveUseridDoctors relationships from objects passed in.
// Removes related items from R.UseridDoctors (uses pointer comparison, removal does not keep order)
// Sets related.R.SpecialtyidSpecialties.
func (o *Specialty) RemoveUseridDoctors(ctx context.Context, exec boil.ContextExecutor, related ...*Doctor) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"specialtiesondoctor\" where \"specialtyid\" = $1 and \"userid\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.Userid)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeUseridDoctorsFromSpecialtyidSpecialtiesSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.UseridDoctors {
			if rel != ri {
				continue
			}

			ln := len(o.R.UseridDoctors)
			if ln > 1 && i < ln-1 {
				o.R.UseridDoctors[i] = o.R.UseridDoctors[ln-1]
			}
			o.R.UseridDoctors = o.R.UseridDoctors[:ln-1]
			break
		}
	}

	return nil
}

func removeUseridDoctorsFromSpecialtyidSpecialtiesSlice(o *Specialty, related []*Doctor) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.SpecialtyidSpecialties {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.SpecialtyidSpecialties)
			if ln > 1 && i < ln-1 {
				rel.R.SpecialtyidSpecialties[i] = rel.R.SpecialtyidSpecialties[ln-1]
			}
			rel.R.SpecialtyidSpecialties = rel.R.SpecialtyidSpecialties[:ln-1]
			break
		}
	}
}

// Specialties retrieves all the records using an executor.
func Specialties(mods ...qm.QueryMod) specialtyQuery {
	mods = append(mods, qm.From("\"specialties\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"specialties\".*"})
	}

	return specialtyQuery{q}
}

// FindSpecialty retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSpecialty(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Specialty, error) {
	specialtyObj := &Specialty{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"specialties\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, specialtyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from specialties")
	}

	if err = specialtyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return specialtyObj, err
	}

	return specialtyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Specialty) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no specialties provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(specialtyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	specialtyInsertCacheMut.RLock()
	cache, cached := specialtyInsertCache[key]
	specialtyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			specialtyAllColumns,
			specialtyColumnsWithDefault,
			specialtyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(specialtyType, specialtyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(specialtyType, specialtyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"specialties\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"specialties\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into specialties")
	}

	if !cached {
		specialtyInsertCacheMut.Lock()
		specialtyInsertCache[key] = cache
		specialtyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Specialty.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Specialty) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	specialtyUpdateCacheMut.RLock()
	cache, cached := specialtyUpdateCache[key]
	specialtyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			specialtyAllColumns,
			specialtyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update specialties, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"specialties\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, specialtyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(specialtyType, specialtyMapping, append(wl, specialtyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update specialties row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for specialties")
	}

	if !cached {
		specialtyUpdateCacheMut.Lock()
		specialtyUpdateCache[key] = cache
		specialtyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q specialtyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for specialties")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for specialties")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SpecialtySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), specialtyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"specialties\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, specialtyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in specialty slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all specialty")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Specialty) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no specialties provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(specialtyColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	specialtyUpsertCacheMut.RLock()
	cache, cached := specialtyUpsertCache[key]
	specialtyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			specialtyAllColumns,
			specialtyColumnsWithDefault,
			specialtyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			specialtyAllColumns,
			specialtyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert specialties, could not build update column list")
		}

		ret := strmangle.SetComplement(specialtyAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(specialtyPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert specialties, could not build conflict column list")
			}

			conflict = make([]string, len(specialtyPrimaryKeyColumns))
			copy(conflict, specialtyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"specialties\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(specialtyType, specialtyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(specialtyType, specialtyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert specialties")
	}

	if !cached {
		specialtyUpsertCacheMut.Lock()
		specialtyUpsertCache[key] = cache
		specialtyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Specialty record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Specialty) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Specialty provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), specialtyPrimaryKeyMapping)
	sql := "DELETE FROM \"specialties\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from specialties")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for specialties")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q specialtyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no specialtyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from specialties")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for specialties")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SpecialtySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(specialtyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), specialtyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"specialties\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, specialtyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from specialty slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for specialties")
	}

	if len(specialtyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Specialty) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSpecialty(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SpecialtySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SpecialtySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), specialtyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"specialties\".* FROM \"specialties\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, specialtyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SpecialtySlice")
	}

	*o = slice

	return nil
}

// SpecialtyExists checks if the Specialty row exists.
func SpecialtyExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"specialties\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if specialties exists")
	}

	return exists, nil
}

// Exists checks if the Specialty row exists.
func (o *Specialty) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SpecialtyExists(ctx, exec, o.ID)
}
