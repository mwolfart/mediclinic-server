// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSpecialties(t *testing.T) {
	t.Parallel()

	query := Specialties()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpecialtiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpecialtiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Specialties().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpecialtiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpecialtySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpecialtiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpecialtyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Specialty exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpecialtyExists to return true, but got false.")
	}
}

func testSpecialtiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	specialtyFound, err := FindSpecialty(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if specialtyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpecialtiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Specialties().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpecialtiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Specialties().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpecialtiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	specialtyOne := &Specialty{}
	specialtyTwo := &Specialty{}
	if err = randomize.Struct(seed, specialtyOne, specialtyDBTypes, false, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}
	if err = randomize.Struct(seed, specialtyTwo, specialtyDBTypes, false, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = specialtyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = specialtyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Specialties().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpecialtiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	specialtyOne := &Specialty{}
	specialtyTwo := &Specialty{}
	if err = randomize.Struct(seed, specialtyOne, specialtyDBTypes, false, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}
	if err = randomize.Struct(seed, specialtyTwo, specialtyDBTypes, false, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = specialtyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = specialtyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func specialtyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func specialtyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Specialty) error {
	*o = Specialty{}
	return nil
}

func testSpecialtiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Specialty{}
	o := &Specialty{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, specialtyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Specialty object: %s", err)
	}

	AddSpecialtyHook(boil.BeforeInsertHook, specialtyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	specialtyBeforeInsertHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.AfterInsertHook, specialtyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	specialtyAfterInsertHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.AfterSelectHook, specialtyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	specialtyAfterSelectHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.BeforeUpdateHook, specialtyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	specialtyBeforeUpdateHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.AfterUpdateHook, specialtyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	specialtyAfterUpdateHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.BeforeDeleteHook, specialtyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	specialtyBeforeDeleteHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.AfterDeleteHook, specialtyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	specialtyAfterDeleteHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.BeforeUpsertHook, specialtyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	specialtyBeforeUpsertHooks = []SpecialtyHook{}

	AddSpecialtyHook(boil.AfterUpsertHook, specialtyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	specialtyAfterUpsertHooks = []SpecialtyHook{}
}

func testSpecialtiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpecialtiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(specialtyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpecialtyToManyUseridDoctors(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Specialty
	var b, c Doctor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, doctorDBTypes, false, doctorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, doctorDBTypes, false, doctorColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"specialtiesondoctor\" (\"specialtyid\", \"userid\") values ($1, $2)", a.ID, b.Userid)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"specialtiesondoctor\" (\"specialtyid\", \"userid\") values ($1, $2)", a.ID, c.Userid)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.UseridDoctors().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.Userid == b.Userid {
			bFound = true
		}
		if v.Userid == c.Userid {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SpecialtySlice{&a}
	if err = a.L.LoadUseridDoctors(ctx, tx, false, (*[]*Specialty)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UseridDoctors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UseridDoctors = nil
	if err = a.L.LoadUseridDoctors(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UseridDoctors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSpecialtyToManyAddOpUseridDoctors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Specialty
	var b, c, d, e Doctor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, specialtyDBTypes, false, strmangle.SetComplement(specialtyPrimaryKeyColumns, specialtyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Doctor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, doctorDBTypes, false, strmangle.SetComplement(doctorPrimaryKeyColumns, doctorColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Doctor{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUseridDoctors(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.SpecialtyidSpecialties[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.SpecialtyidSpecialties[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.UseridDoctors[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UseridDoctors[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UseridDoctors().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testSpecialtyToManySetOpUseridDoctors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Specialty
	var b, c, d, e Doctor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, specialtyDBTypes, false, strmangle.SetComplement(specialtyPrimaryKeyColumns, specialtyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Doctor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, doctorDBTypes, false, strmangle.SetComplement(doctorPrimaryKeyColumns, doctorColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUseridDoctors(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UseridDoctors().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUseridDoctors(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UseridDoctors().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.SpecialtyidSpecialties) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.SpecialtyidSpecialties) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.SpecialtyidSpecialties[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.SpecialtyidSpecialties[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.UseridDoctors[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.UseridDoctors[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testSpecialtyToManyRemoveOpUseridDoctors(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Specialty
	var b, c, d, e Doctor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, specialtyDBTypes, false, strmangle.SetComplement(specialtyPrimaryKeyColumns, specialtyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Doctor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, doctorDBTypes, false, strmangle.SetComplement(doctorPrimaryKeyColumns, doctorColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUseridDoctors(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UseridDoctors().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUseridDoctors(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UseridDoctors().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.SpecialtyidSpecialties) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.SpecialtyidSpecialties) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.SpecialtyidSpecialties[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.SpecialtyidSpecialties[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.UseridDoctors) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.UseridDoctors[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.UseridDoctors[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testSpecialtiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpecialtiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpecialtySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpecialtiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Specialties().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	specialtyDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`}
	_                = bytes.MinRead
)

func testSpecialtiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(specialtyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(specialtyAllColumns) == len(specialtyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpecialtiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(specialtyAllColumns) == len(specialtyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Specialty{}
	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, specialtyDBTypes, true, specialtyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(specialtyAllColumns, specialtyPrimaryKeyColumns) {
		fields = specialtyAllColumns
	} else {
		fields = strmangle.SetComplement(
			specialtyAllColumns,
			specialtyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := SpecialtySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpecialtiesUpsert(t *testing.T) {
	t.Parallel()

	if len(specialtyAllColumns) == len(specialtyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Specialty{}
	if err = randomize.Struct(seed, &o, specialtyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Specialty: %s", err)
	}

	count, err := Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, specialtyDBTypes, false, specialtyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Specialty struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Specialty: %s", err)
	}

	count, err = Specialties().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
