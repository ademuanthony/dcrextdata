// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testBlocks(t *testing.T) {
	t.Parallel()

	query := Blocks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBlocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
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

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Blocks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BlockExists(ctx, tx, o.Height)
	if err != nil {
		t.Errorf("Unable to check if Block exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlockExists to return true, but got false.")
	}
}

func testBlocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	blockFound, err := FindBlock(ctx, tx, o.Height)
	if err != nil {
		t.Error(err)
	}

	if blockFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBlocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Blocks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBlocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Blocks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBlocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(blockColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
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

func testBlocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBlocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blockDBTypes = map[string]string{`Height`: `integer`, `ReceiveTime`: `bigint`, `InternalTimestamp`: `bigint`, `Hash`: `character varying`}
	_            = bytes.MinRead
)

func testBlocksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBlocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Block{}
	if err = randomize.Struct(seed, o, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blockAllColumns, blockPrimaryKeyColumns) {
		fields = blockAllColumns
	} else {
		fields = strmangle.SetComplement(
			blockAllColumns,
			blockPrimaryKeyColumns,
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

	slice := BlockSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBlocksUpsert(t *testing.T) {
	t.Parallel()

	if len(blockAllColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Block{}
	if err = randomize.Struct(seed, &o, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err := Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, blockDBTypes, false, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err = Blocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
