// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testNetworkSnapshots(t *testing.T) {
	t.Parallel()

	query := NetworkSnapshots()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNetworkSnapshotsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
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

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkSnapshotsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NetworkSnapshots().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkSnapshotsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NetworkSnapshotSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkSnapshotsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NetworkSnapshotExists(ctx, tx, o.Timestamp)
	if err != nil {
		t.Errorf("Unable to check if NetworkSnapshot exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NetworkSnapshotExists to return true, but got false.")
	}
}

func testNetworkSnapshotsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	networkSnapshotFound, err := FindNetworkSnapshot(ctx, tx, o.Timestamp)
	if err != nil {
		t.Error(err)
	}

	if networkSnapshotFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNetworkSnapshotsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NetworkSnapshots().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNetworkSnapshotsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NetworkSnapshots().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNetworkSnapshotsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	networkSnapshotOne := &NetworkSnapshot{}
	networkSnapshotTwo := &NetworkSnapshot{}
	if err = randomize.Struct(seed, networkSnapshotOne, networkSnapshotDBTypes, false, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}
	if err = randomize.Struct(seed, networkSnapshotTwo, networkSnapshotDBTypes, false, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = networkSnapshotOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = networkSnapshotTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NetworkSnapshots().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNetworkSnapshotsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	networkSnapshotOne := &NetworkSnapshot{}
	networkSnapshotTwo := &NetworkSnapshot{}
	if err = randomize.Struct(seed, networkSnapshotOne, networkSnapshotDBTypes, false, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}
	if err = randomize.Struct(seed, networkSnapshotTwo, networkSnapshotDBTypes, false, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = networkSnapshotOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = networkSnapshotTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testNetworkSnapshotsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNetworkSnapshotsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(networkSnapshotColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNetworkSnapshotsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
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

func testNetworkSnapshotsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NetworkSnapshotSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNetworkSnapshotsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NetworkSnapshots().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	networkSnapshotDBTypes = map[string]string{`Timestamp`: `bigint`, `Height`: `bigint`, `NodeCount`: `integer`}
	_                      = bytes.MinRead
)

func testNetworkSnapshotsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(networkSnapshotPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(networkSnapshotAllColumns) == len(networkSnapshotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNetworkSnapshotsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(networkSnapshotAllColumns) == len(networkSnapshotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NetworkSnapshot{}
	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, networkSnapshotDBTypes, true, networkSnapshotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(networkSnapshotAllColumns, networkSnapshotPrimaryKeyColumns) {
		fields = networkSnapshotAllColumns
	} else {
		fields = strmangle.SetComplement(
			networkSnapshotAllColumns,
			networkSnapshotPrimaryKeyColumns,
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

	slice := NetworkSnapshotSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNetworkSnapshotsUpsert(t *testing.T) {
	t.Parallel()

	if len(networkSnapshotAllColumns) == len(networkSnapshotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NetworkSnapshot{}
	if err = randomize.Struct(seed, &o, networkSnapshotDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NetworkSnapshot: %s", err)
	}

	count, err := NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, networkSnapshotDBTypes, false, networkSnapshotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkSnapshot struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NetworkSnapshot: %s", err)
	}

	count, err = NetworkSnapshots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
