// Code generated by SQLBoiler 4.19.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testBenefits(t *testing.T) {
	t.Parallel()

	query := Benefits()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBenefitsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
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

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBenefitsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Benefits().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBenefitsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BenefitSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBenefitsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BenefitExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Benefit exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BenefitExists to return true, but got false.")
	}
}

func testBenefitsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	benefitFound, err := FindBenefit(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if benefitFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBenefitsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Benefits().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBenefitsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Benefits().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBenefitsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	benefitOne := &Benefit{}
	benefitTwo := &Benefit{}
	if err = randomize.Struct(seed, benefitOne, benefitDBTypes, false, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}
	if err = randomize.Struct(seed, benefitTwo, benefitDBTypes, false, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = benefitOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = benefitTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Benefits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBenefitsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	benefitOne := &Benefit{}
	benefitTwo := &Benefit{}
	if err = randomize.Struct(seed, benefitOne, benefitDBTypes, false, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}
	if err = randomize.Struct(seed, benefitTwo, benefitDBTypes, false, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = benefitOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = benefitTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func benefitBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func benefitAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Benefit) error {
	*o = Benefit{}
	return nil
}

func testBenefitsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Benefit{}
	o := &Benefit{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, benefitDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Benefit object: %s", err)
	}

	AddBenefitHook(boil.BeforeInsertHook, benefitBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	benefitBeforeInsertHooks = []BenefitHook{}

	AddBenefitHook(boil.AfterInsertHook, benefitAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	benefitAfterInsertHooks = []BenefitHook{}

	AddBenefitHook(boil.AfterSelectHook, benefitAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	benefitAfterSelectHooks = []BenefitHook{}

	AddBenefitHook(boil.BeforeUpdateHook, benefitBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	benefitBeforeUpdateHooks = []BenefitHook{}

	AddBenefitHook(boil.AfterUpdateHook, benefitAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	benefitAfterUpdateHooks = []BenefitHook{}

	AddBenefitHook(boil.BeforeDeleteHook, benefitBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	benefitBeforeDeleteHooks = []BenefitHook{}

	AddBenefitHook(boil.AfterDeleteHook, benefitAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	benefitAfterDeleteHooks = []BenefitHook{}

	AddBenefitHook(boil.BeforeUpsertHook, benefitBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	benefitBeforeUpsertHooks = []BenefitHook{}

	AddBenefitHook(boil.AfterUpsertHook, benefitAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	benefitAfterUpsertHooks = []BenefitHook{}
}

func testBenefitsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBenefitsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(strmangle.SetMerge(benefitPrimaryKeyColumns, benefitColumnsWithoutDefault)...)); err != nil {
		t.Error(err)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBenefitsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
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

func testBenefitsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BenefitSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBenefitsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Benefits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	benefitDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`, `Description`: `text`, `Price`: `double precision`, `Image`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testBenefitsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(benefitPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(benefitAllColumns) == len(benefitPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBenefitsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(benefitAllColumns) == len(benefitPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Benefit{}
	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, benefitDBTypes, true, benefitPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(benefitAllColumns, benefitPrimaryKeyColumns) {
		fields = benefitAllColumns
	} else {
		fields = strmangle.SetComplement(
			benefitAllColumns,
			benefitPrimaryKeyColumns,
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

	slice := BenefitSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBenefitsUpsert(t *testing.T) {
	t.Parallel()

	if len(benefitAllColumns) == len(benefitPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Benefit{}
	if err = randomize.Struct(seed, &o, benefitDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Benefit: %s", err)
	}

	count, err := Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, benefitDBTypes, false, benefitPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Benefit struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Benefit: %s", err)
	}

	count, err = Benefits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
