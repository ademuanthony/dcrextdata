// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// NetworkPeer is an object representing the database table.
type NetworkPeer struct {
	Timestamp       int64  `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Address         string `boil:"address" json:"address" toml:"address" yaml:"address"`
	LastSeen        int64  `boil:"last_seen" json:"last_seen" toml:"last_seen" yaml:"last_seen"`
	ConnectionTime  int64  `boil:"connection_time" json:"connection_time" toml:"connection_time" yaml:"connection_time"`
	ProtocolVersion int    `boil:"protocol_version" json:"protocol_version" toml:"protocol_version" yaml:"protocol_version"`
	UserAgent       string `boil:"user_agent" json:"user_agent" toml:"user_agent" yaml:"user_agent"`
	StartingHeight  int64  `boil:"starting_height" json:"starting_height" toml:"starting_height" yaml:"starting_height"`
	CurrentHeight   int64  `boil:"current_height" json:"current_height" toml:"current_height" yaml:"current_height"`

	R *networkPeerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L networkPeerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NetworkPeerColumns = struct {
	Timestamp       string
	Address         string
	LastSeen        string
	ConnectionTime  string
	ProtocolVersion string
	UserAgent       string
	StartingHeight  string
	CurrentHeight   string
}{
	Timestamp:       "timestamp",
	Address:         "address",
	LastSeen:        "last_seen",
	ConnectionTime:  "connection_time",
	ProtocolVersion: "protocol_version",
	UserAgent:       "user_agent",
	StartingHeight:  "starting_height",
	CurrentHeight:   "current_height",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

var NetworkPeerWhere = struct {
	Timestamp       whereHelperint64
	Address         whereHelperstring
	LastSeen        whereHelperint64
	ConnectionTime  whereHelperint64
	ProtocolVersion whereHelperint
	UserAgent       whereHelperstring
	StartingHeight  whereHelperint64
	CurrentHeight   whereHelperint64
}{
	Timestamp:       whereHelperint64{field: "\"network_peer\".\"timestamp\""},
	Address:         whereHelperstring{field: "\"network_peer\".\"address\""},
	LastSeen:        whereHelperint64{field: "\"network_peer\".\"last_seen\""},
	ConnectionTime:  whereHelperint64{field: "\"network_peer\".\"connection_time\""},
	ProtocolVersion: whereHelperint{field: "\"network_peer\".\"protocol_version\""},
	UserAgent:       whereHelperstring{field: "\"network_peer\".\"user_agent\""},
	StartingHeight:  whereHelperint64{field: "\"network_peer\".\"starting_height\""},
	CurrentHeight:   whereHelperint64{field: "\"network_peer\".\"current_height\""},
}

// NetworkPeerRels is where relationship names are stored.
var NetworkPeerRels = struct {
}{}

// networkPeerR is where relationships are stored.
type networkPeerR struct {
}

// NewStruct creates a new relationship struct
func (*networkPeerR) NewStruct() *networkPeerR {
	return &networkPeerR{}
}

// networkPeerL is where Load methods for each relationship are stored.
type networkPeerL struct{}

var (
	networkPeerAllColumns            = []string{"timestamp", "address", "last_seen", "connection_time", "protocol_version", "user_agent", "starting_height", "current_height"}
	networkPeerColumnsWithoutDefault = []string{"timestamp", "address", "last_seen", "connection_time", "protocol_version", "user_agent", "starting_height", "current_height"}
	networkPeerColumnsWithDefault    = []string{}
	networkPeerPrimaryKeyColumns     = []string{"timestamp", "address"}
)

type (
	// NetworkPeerSlice is an alias for a slice of pointers to NetworkPeer.
	// This should generally be used opposed to []NetworkPeer.
	NetworkPeerSlice []*NetworkPeer

	networkPeerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	networkPeerType                 = reflect.TypeOf(&NetworkPeer{})
	networkPeerMapping              = queries.MakeStructMapping(networkPeerType)
	networkPeerPrimaryKeyMapping, _ = queries.BindMapping(networkPeerType, networkPeerMapping, networkPeerPrimaryKeyColumns)
	networkPeerInsertCacheMut       sync.RWMutex
	networkPeerInsertCache          = make(map[string]insertCache)
	networkPeerUpdateCacheMut       sync.RWMutex
	networkPeerUpdateCache          = make(map[string]updateCache)
	networkPeerUpsertCacheMut       sync.RWMutex
	networkPeerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single networkPeer record from the query.
func (q networkPeerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NetworkPeer, error) {
	o := &NetworkPeer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for network_peer")
	}

	return o, nil
}

// All returns all NetworkPeer records from the query.
func (q networkPeerQuery) All(ctx context.Context, exec boil.ContextExecutor) (NetworkPeerSlice, error) {
	var o []*NetworkPeer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NetworkPeer slice")
	}

	return o, nil
}

// Count returns the count of all NetworkPeer records in the query.
func (q networkPeerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count network_peer rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q networkPeerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if network_peer exists")
	}

	return count > 0, nil
}

// NetworkPeers retrieves all the records using an executor.
func NetworkPeers(mods ...qm.QueryMod) networkPeerQuery {
	mods = append(mods, qm.From("\"network_peer\""))
	return networkPeerQuery{NewQuery(mods...)}
}

// FindNetworkPeer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNetworkPeer(ctx context.Context, exec boil.ContextExecutor, timestamp int64, address string, selectCols ...string) (*NetworkPeer, error) {
	networkPeerObj := &NetworkPeer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"network_peer\" where \"timestamp\"=$1 AND \"address\"=$2", sel,
	)

	q := queries.Raw(query, timestamp, address)

	err := q.Bind(ctx, exec, networkPeerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from network_peer")
	}

	return networkPeerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NetworkPeer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_peer provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(networkPeerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	networkPeerInsertCacheMut.RLock()
	cache, cached := networkPeerInsertCache[key]
	networkPeerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			networkPeerAllColumns,
			networkPeerColumnsWithDefault,
			networkPeerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(networkPeerType, networkPeerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(networkPeerType, networkPeerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"network_peer\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"network_peer\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into network_peer")
	}

	if !cached {
		networkPeerInsertCacheMut.Lock()
		networkPeerInsertCache[key] = cache
		networkPeerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the NetworkPeer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NetworkPeer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	networkPeerUpdateCacheMut.RLock()
	cache, cached := networkPeerUpdateCache[key]
	networkPeerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			networkPeerAllColumns,
			networkPeerPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update network_peer, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"network_peer\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, networkPeerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(networkPeerType, networkPeerMapping, append(wl, networkPeerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update network_peer row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for network_peer")
	}

	if !cached {
		networkPeerUpdateCacheMut.Lock()
		networkPeerUpdateCache[key] = cache
		networkPeerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q networkPeerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for network_peer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for network_peer")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NetworkPeerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkPeerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"network_peer\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, networkPeerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in networkPeer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all networkPeer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NetworkPeer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no network_peer provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(networkPeerColumnsWithDefault, o)

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

	networkPeerUpsertCacheMut.RLock()
	cache, cached := networkPeerUpsertCache[key]
	networkPeerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			networkPeerAllColumns,
			networkPeerColumnsWithDefault,
			networkPeerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			networkPeerAllColumns,
			networkPeerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert network_peer, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(networkPeerPrimaryKeyColumns))
			copy(conflict, networkPeerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"network_peer\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(networkPeerType, networkPeerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(networkPeerType, networkPeerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert network_peer")
	}

	if !cached {
		networkPeerUpsertCacheMut.Lock()
		networkPeerUpsertCache[key] = cache
		networkPeerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single NetworkPeer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NetworkPeer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NetworkPeer provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), networkPeerPrimaryKeyMapping)
	sql := "DELETE FROM \"network_peer\" WHERE \"timestamp\"=$1 AND \"address\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from network_peer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for network_peer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q networkPeerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no networkPeerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from network_peer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_peer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NetworkPeerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkPeerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"network_peer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkPeerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from networkPeer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for network_peer")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NetworkPeer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNetworkPeer(ctx, exec, o.Timestamp, o.Address)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NetworkPeerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NetworkPeerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), networkPeerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"network_peer\".* FROM \"network_peer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, networkPeerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NetworkPeerSlice")
	}

	*o = slice

	return nil
}

// NetworkPeerExists checks if the NetworkPeer row exists.
func NetworkPeerExists(ctx context.Context, exec boil.ContextExecutor, timestamp int64, address string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"network_peer\" where \"timestamp\"=$1 AND \"address\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, timestamp, address)
	}
	row := exec.QueryRowContext(ctx, sql, timestamp, address)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if network_peer exists")
	}

	return exists, nil
}
