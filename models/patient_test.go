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

func testPatients(t *testing.T) {
	t.Parallel()

	query := Patients()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPatientsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
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

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Patients().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatientSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatientsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PatientExists(ctx, tx, o.Userid)
	if err != nil {
		t.Errorf("Unable to check if Patient exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PatientExists to return true, but got false.")
	}
}

func testPatientsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	patientFound, err := FindPatient(ctx, tx, o.Userid)
	if err != nil {
		t.Error(err)
	}

	if patientFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPatientsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Patients().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPatientsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Patients().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPatientsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	patientOne := &Patient{}
	patientTwo := &Patient{}
	if err = randomize.Struct(seed, patientOne, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}
	if err = randomize.Struct(seed, patientTwo, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPatientsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	patientOne := &Patient{}
	patientTwo := &Patient{}
	if err = randomize.Struct(seed, patientOne, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}
	if err = randomize.Struct(seed, patientTwo, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func patientBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func patientAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Patient) error {
	*o = Patient{}
	return nil
}

func testPatientsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Patient{}
	o := &Patient{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, patientDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Patient object: %s", err)
	}

	AddPatientHook(boil.BeforeInsertHook, patientBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	patientBeforeInsertHooks = []PatientHook{}

	AddPatientHook(boil.AfterInsertHook, patientAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	patientAfterInsertHooks = []PatientHook{}

	AddPatientHook(boil.AfterSelectHook, patientAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	patientAfterSelectHooks = []PatientHook{}

	AddPatientHook(boil.BeforeUpdateHook, patientBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	patientBeforeUpdateHooks = []PatientHook{}

	AddPatientHook(boil.AfterUpdateHook, patientAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	patientAfterUpdateHooks = []PatientHook{}

	AddPatientHook(boil.BeforeDeleteHook, patientBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	patientBeforeDeleteHooks = []PatientHook{}

	AddPatientHook(boil.AfterDeleteHook, patientAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	patientAfterDeleteHooks = []PatientHook{}

	AddPatientHook(boil.BeforeUpsertHook, patientBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	patientBeforeUpsertHooks = []PatientHook{}

	AddPatientHook(boil.AfterUpsertHook, patientAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	patientAfterUpsertHooks = []PatientHook{}
}

func testPatientsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatientsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(patientColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatientToManyInsuranceidInsurances(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patient
	var b, c Insurance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, insuranceDBTypes, false, insuranceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, insuranceDBTypes, false, insuranceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"insurancesonpatient\" (\"userid\", \"insuranceid\") values ($1, $2)", a.Userid, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"insurancesonpatient\" (\"userid\", \"insuranceid\") values ($1, $2)", a.Userid, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.InsuranceidInsurances().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PatientSlice{&a}
	if err = a.L.LoadInsuranceidInsurances(ctx, tx, false, (*[]*Patient)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.InsuranceidInsurances); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.InsuranceidInsurances = nil
	if err = a.L.LoadInsuranceidInsurances(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.InsuranceidInsurances); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPatientToManyAddOpInsuranceidInsurances(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patient
	var b, c, d, e Insurance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientDBTypes, false, strmangle.SetComplement(patientPrimaryKeyColumns, patientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Insurance{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, insuranceDBTypes, false, strmangle.SetComplement(insurancePrimaryKeyColumns, insuranceColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Insurance{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddInsuranceidInsurances(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.UseridPatients[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.UseridPatients[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.InsuranceidInsurances[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.InsuranceidInsurances[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.InsuranceidInsurances().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPatientToManySetOpInsuranceidInsurances(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patient
	var b, c, d, e Insurance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientDBTypes, false, strmangle.SetComplement(patientPrimaryKeyColumns, patientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Insurance{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, insuranceDBTypes, false, strmangle.SetComplement(insurancePrimaryKeyColumns, insuranceColumnsWithoutDefault)...); err != nil {
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

	err = a.SetInsuranceidInsurances(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.InsuranceidInsurances().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetInsuranceidInsurances(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.InsuranceidInsurances().Count(ctx, tx)
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
	// if len(b.R.UseridPatients) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.UseridPatients) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.UseridPatients[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.UseridPatients[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.InsuranceidInsurances[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.InsuranceidInsurances[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPatientToManyRemoveOpInsuranceidInsurances(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patient
	var b, c, d, e Insurance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientDBTypes, false, strmangle.SetComplement(patientPrimaryKeyColumns, patientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Insurance{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, insuranceDBTypes, false, strmangle.SetComplement(insurancePrimaryKeyColumns, insuranceColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddInsuranceidInsurances(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.InsuranceidInsurances().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveInsuranceidInsurances(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.InsuranceidInsurances().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.UseridPatients) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.UseridPatients) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.UseridPatients[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.UseridPatients[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.InsuranceidInsurances) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.InsuranceidInsurances[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.InsuranceidInsurances[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPatientToOneGenericuserUsingUseridGenericuser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Patient
	var foreign Genericuser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, patientDBTypes, false, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, genericuserDBTypes, false, genericuserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Genericuser struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.Userid = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.UseridGenericuser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddGenericuserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Genericuser) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := PatientSlice{&local}
	if err = local.L.LoadUseridGenericuser(ctx, tx, false, (*[]*Patient)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UseridGenericuser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.UseridGenericuser = nil
	if err = local.L.LoadUseridGenericuser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.UseridGenericuser == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testPatientToOneSetOpGenericuserUsingUseridGenericuser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Patient
	var b, c Genericuser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patientDBTypes, false, strmangle.SetComplement(patientPrimaryKeyColumns, patientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, genericuserDBTypes, false, strmangle.SetComplement(genericuserPrimaryKeyColumns, genericuserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, genericuserDBTypes, false, strmangle.SetComplement(genericuserPrimaryKeyColumns, genericuserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Genericuser{&b, &c} {
		err = a.SetUseridGenericuser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.UseridGenericuser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UseridPatient != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Userid != x.ID {
			t.Error("foreign key was wrong value", a.Userid)
		}

		if exists, err := PatientExists(ctx, tx, a.Userid); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testPatientsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
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

func testPatientsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatientSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPatientsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	patientDBTypes = map[string]string{`Userid`: `integer`}
	_              = bytes.MinRead
)

func testPatientsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(patientAllColumns) == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patientDBTypes, true, patientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPatientsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(patientAllColumns) == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Patient{}
	if err = randomize.Struct(seed, o, patientDBTypes, true, patientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patientDBTypes, true, patientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(patientAllColumns, patientPrimaryKeyColumns) {
		fields = patientAllColumns
	} else {
		fields = strmangle.SetComplement(
			patientAllColumns,
			patientPrimaryKeyColumns,
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

	slice := PatientSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPatientsUpsert(t *testing.T) {
	t.Parallel()

	if len(patientAllColumns) == len(patientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Patient{}
	if err = randomize.Struct(seed, &o, patientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Patient: %s", err)
	}

	count, err := Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, patientDBTypes, false, patientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Patient struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Patient: %s", err)
	}

	count, err = Patients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
