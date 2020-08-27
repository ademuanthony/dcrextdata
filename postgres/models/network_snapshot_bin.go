// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// NetworkSnapshotBin is an object representing the database table.
type NetworkSnapshotBin struct {
	Timestamp      int64  `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Height         int64  `boil:"height" json:"height" toml:"height" yaml:"height"`
	NodeCount      int    `boil:"node_count" json:"node_count" toml:"node_count" yaml:"node_count"`
	ReachableNodes int    `boil:"reachable_nodes" json:"reachable_nodes" toml:"reachable_nodes" yaml:"reachable_nodes"`
	Bin            string `boil:"bin" json:"bin" toml:"bin" yaml:"bin"`

	R *networkSnapshotBinR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L networkSnapshotBinL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NetworkSnapshotBinColumns = struct {
	Timestamp      string
	Height         string
	NodeCount      string
	ReachableNodes string
	Bin            string
}{
	Timestamp:      "timestamp",
	Height:         "height",
	NodeCount:      "node_count",
	ReachableNodes: "reachable_nodes",
	Bin:            "bin",
}

// Generated where

var NetworkSnapshotBinWhere = struct {
	Timestamp      whereHelperint64
	Height         whereHelperint64
	NodeCount      whereHelperint
	ReachableNodes whereHelperint
	Bin            whereHelperstring
}{
	Timestamp:      whereHelperint64{field: "\"network_snapshot_bin\".\"timestamp\""},
	Height:         whereHelperint64{field: "\"network_snapshot_bin\".\"height\""},
	NodeCount:      whereHelperint{field: "\"network_snapshot_bin\".\"node_count\""},
	ReachableNodes: whereHelperint{field: "\"network_snapshot_bin\".\"reachable_nodes\""},
	Bin:            whereHelperstring{field: "\"network_snapshot_bin\".\"bin\""},
}

// NetworkSnapshotBinRels is where relationship names are stored.
var NetworkSnapshotBinRels = struct {
}{}

// networkSnapshotBinR is where relationships are stored.
type networkSnapshotBinR struct {
}

// NewStruct creates a new relationship struct
func (*networkSnapshotBinR) NewStruct() *networkSnapshotBinR {
	return &networkSnapshotBinR{}
}

// networkSnapshotBinL is where Load methods for each relationship are stored.
type networkSnapshotBinL struct{}

var (
	networkSnapshotBinAllColumns            = []string{"timestamp", "height", "node_count", "reachable_nodes", "bin"}
	networkSnapshotBinColumnsWithoutDefault = []string{"timestamp", "height", "node_count", "reachable_nodes"}
	networkSnapshotBinColumnsWithDefault    = []string{"bin"}
	networkSnapshotBinPrimaryKeyColumns     = []string{"timestamp", "bin"}
)

type (
	// NetworkSnapshotBinSlice is an alias for a slice of pointers to NetworkSnapshotBin.
	// This should generally be used opposed to []NetworkSnapshotBin.
	NetworkSnapshotBinSlice []*NetworkSnapshotBin

	networkSnapshotBinQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	networkSnapshotBinType                 = reflect.TypeOf(&NetworkSnapshotBin{})
	networkSnapshotBinMapping              = queries.MakeStructMapping(networkSnapshotBinType)
	networkSnapshotBinPrimaryKeyMapping, _ = queries.BindMapping(networkSnapshotBinType, networkSnapshotBinMapping, networkSnapshotBinPrimaryKeyColumns)
	networkSnapshotBinInsertCacheMut       sync.RWMutex
	networkSnapshotBinInsertCache          = make(map[string]insertCache)
	networkSnapshotBinUpdateCacheMut       sync.RWMutex
	networkSnapshotBinUpdateCache          = make(map[string]updateCache)
	networkSnapshotBinUpsertCacheMut       sync.RWMutex
	networkSnapshotBinUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single networkSnapshotBin record from the query.
func (q networkSnapshotBinQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NetworkSnapshotBin, error) {
	o := &NetworkSnapshotBin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for network_snapshot_bin")
	}

	return o, nil
}

// All returns all NetworkSnapshotBin records from the query.
func (q networkSnapshotBinQuery) All(ctx context.Context, exec boil.ContextExecutor) (NetworkSnapshotBinSlice, error) {
	var o []*NetworkSnapshotBin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NetworkSnapshotBin slice")
	}

	return o, nil
}

// Count returns the count of all NetworkSnapshotBin records in the query.
func (q networkSnapshotBinQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count network_snapshot_bin rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q networkSnapshotBinQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if network_snapshot_bin exists")
	}

	return count > 0, nil
}

// NetworkSnapshotBins retrieves all the records using an executor.
func NetworkSnapshotBins(mods ...qm.QueryMod) networkSnapshotBinQuery {
	mods = append(mods, qm.From("\"network_snapshot_bin\""))
	return networkSnapshotBinQuery{NewQuery(mods...)}
}

// FindNetworkSnapshotBin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNetworkSnapshotBin(ctx context.Context, exec boil.ContextExecutor, timestamp int64, bin string, selectCols ...string) (*NetworkSnapshotBin, error) {
	networkSnapshotBinObj := &NetworkSnapshotBin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"network_snapshot_bin\" where \"timestamp\"=$1 AND \"bin\"=$2", sel,
	)

	q := queries.Raw(query, timestamp, bin)

	err := q.Bind(ctx, exec, networkSnapshotBinObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from network_snapshot_bin")
	}

	return networkSnapshotBinObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NetworkSnapshotBin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_snapshot_bin provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(networkSnapshotBinColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	networkSnapshotBinInsertCacheMut.RLock()
	cache, cached := networkSnapshotBinInsertCache[key]
	networkSnapshotBinInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			networkSnapshotBinAllColumns,
			networkSnapshotBinColumnsWithDefault,
			networkSnapshotBinColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(networkSnapshotBinType, networkSnapshotBinMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(networkSnapshotBinType, networkSnapshotBinMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"network_snapshot_bin\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"network_snapshot_bin\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into network_snapshot_bin")
	}

	if !cached {
		networkSnapshotBinInsertCacheMut.Lock()
		networkSnapshotBinInsertCache[key] = cache
		networkSnapshotBinInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the NetworkSnapshotBin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NetworkSnapshotBin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	networkSnapshotBinUpdateCacheMut.RLock()
	cache, cached := networkSnapshotBinUpdateCache[key]
	networkSnapshotBinUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			networkSnapshotBinAllColumns,
			networkSnapshotBinPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update network_snapshot_bin, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"network_snapshot_bin\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, networkSnapshotBinPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(networkSnapshotBinType, networkSnapshotBinMapping, append(wl, networkSnapshotBinPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update network_snapshot_bin row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for network_snapshot_bin")
	}

	if !cached {
		networkSnapshotBinUpdateCacheMut.Lock()
		networkSnapshotBinUpdateCache[key] = cache
		networkSnapshotBinUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q networkSnapshotBinQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for network_snapshot_bin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for network_snapshot_bin")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NetworkSnapshotBinSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkSnapshotBinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"network_snapshot_bin\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, networkSnapshotBinPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in networkSnapshotBin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all networkSnapshotBin")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NetworkSnapshotBin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_snapshot_bin provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(networkSnapshotBinColumnsWithDefault, o)

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

	networkSnapshotBinUpsertCacheMut.RLock()
	cache, cached := networkSnapshotBinUpsertCache[key]
	networkSnapshotBinUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			networkSnapshotBinAllColumns,
			networkSnapshotBinColumnsWithDefault,
			networkSnapshotBinColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			networkSnapshotBinAllColumns,
			networkSnapshotBinPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert network_snapshot_bin, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(networkSnapshotBinPrimaryKeyColumns))
			copy(conflict, networkSnapshotBinPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"network_snapshot_bin\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(networkSnapshotBinType, networkSnapshotBinMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(networkSnapshotBinType, networkSnapshotBinMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert network_snapshot_bin")
	}

	if !cached {
		networkSnapshotBinUpsertCacheMut.Lock()
		networkSnapshotBinUpsertCache[key] = cache
		networkSnapshotBinUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single NetworkSnapshotBin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NetworkSnapshotBin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NetworkSnapshotBin provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), networkSnapshotBinPrimaryKeyMapping)
	sql := "DELETE FROM \"network_snapshot_bin\" WHERE \"timestamp\"=$1 AND \"bin\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from network_snapshot_bin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for network_snapshot_bin")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q networkSnapshotBinQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no networkSnapshotBinQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from network_snapshot_bin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_snapshot_bin")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NetworkSnapshotBinSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkSnapshotBinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"network_snapshot_bin\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkSnapshotBinPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from networkSnapshotBin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_snapshot_bin")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NetworkSnapshotBin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNetworkSnapshotBin(ctx, exec, o.Timestamp, o.Bin)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NetworkSnapshotBinSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NetworkSnapshotBinSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkSnapshotBinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"network_snapshot_bin\".* FROM \"network_snapshot_bin\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkSnapshotBinPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NetworkSnapshotBinSlice")
	}

	*o = slice

	return nil
}

// NetworkSnapshotBinExists checks if the NetworkSnapshotBin row exists.
func NetworkSnapshotBinExists(ctx context.Context, exec boil.ContextExecutor, timestamp int64, bin string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"network_snapshot_bin\" where \"timestamp\"=$1 AND \"bin\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, timestamp, bin)
	}
	row := exec.QueryRowContext(ctx, sql, timestamp, bin)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if network_snapshot_bin exists")
	}

	return exists, nil
}