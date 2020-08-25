// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testNodes(t *testing.T) {
	t.Parallel()

	query := Nodes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNodesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
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

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Nodes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NodeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NodeExists(ctx, tx, o.Address)
	if err != nil {
		t.Errorf("Unable to check if Node exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NodeExists to return true, but got false.")
	}
}

func testNodesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	nodeFound, err := FindNode(ctx, tx, o.Address)
	if err != nil {
		t.Error(err)
	}

	if nodeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNodesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Nodes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNodesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Nodes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNodesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	nodeOne := &Node{}
	nodeTwo := &Node{}
	if err = randomize.Struct(seed, nodeOne, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}
	if err = randomize.Struct(seed, nodeTwo, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Nodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNodesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	nodeOne := &Node{}
	nodeTwo := &Node{}
	if err = randomize.Struct(seed, nodeOne, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}
	if err = randomize.Struct(seed, nodeTwo, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testNodesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNodesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(nodeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNodeToManyHeartbeats(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Node
	var b, c Heartbeat

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, heartbeatDBTypes, false, heartbeatColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, heartbeatDBTypes, false, heartbeatColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.NodeID = a.Address
	c.NodeID = a.Address

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Heartbeats().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.NodeID == b.NodeID {
			bFound = true
		}
		if v.NodeID == c.NodeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := NodeSlice{&a}
	if err = a.L.LoadHeartbeats(ctx, tx, false, (*[]*Node)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Heartbeats); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Heartbeats = nil
	if err = a.L.LoadHeartbeats(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Heartbeats); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testNodeToManyAddOpHeartbeats(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Node
	var b, c, d, e Heartbeat

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, nodeDBTypes, false, strmangle.SetComplement(nodePrimaryKeyColumns, nodeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Heartbeat{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, heartbeatDBTypes, false, strmangle.SetComplement(heartbeatPrimaryKeyColumns, heartbeatColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Heartbeat{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddHeartbeats(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Address != first.NodeID {
			t.Error("foreign key was wrong value", a.Address, first.NodeID)
		}
		if a.Address != second.NodeID {
			t.Error("foreign key was wrong value", a.Address, second.NodeID)
		}

		if first.R.Node != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Node != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Heartbeats[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Heartbeats[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Heartbeats().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testNodesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
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

func testNodesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NodeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNodesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Nodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	nodeDBTypes = map[string]string{`Address`: `character varying`, `IPVersion`: `integer`, `Country`: `character varying`, `Region`: `character varying`, `City`: `character varying`, `Zip`: `character varying`, `LastAttempt`: `bigint`, `LastSeen`: `bigint`, `LastSuccess`: `bigint`, `IsDead`: `boolean`, `ConnectionTime`: `bigint`, `ProtocolVersion`: `integer`, `UserAgent`: `character varying`, `Services`: `character varying`, `StartingHeight`: `bigint`, `CurrentHeight`: `bigint`, `FailureCount`: `integer`}
	_           = bytes.MinRead
)

func testNodesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNodesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(nodeAllColumns, nodePrimaryKeyColumns) {
		fields = nodeAllColumns
	} else {
		fields = strmangle.SetComplement(
			nodeAllColumns,
			nodePrimaryKeyColumns,
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

	slice := NodeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNodesUpsert(t *testing.T) {
	t.Parallel()

	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Node{}
	if err = randomize.Struct(seed, &o, nodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Node: %s", err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, nodeDBTypes, false, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Node: %s", err)
	}

	count, err = Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
