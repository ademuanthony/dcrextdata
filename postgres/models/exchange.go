// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Exchange is an object representing the database table.
type Exchange struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`
	URL  string `boil:"url" json:"url" toml:"url" yaml:"url"`

	R *exchangeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L exchangeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExchangeColumns = struct {
	ID   string
	Name string
	URL  string
}{
	ID:   "id",
	Name: "name",
	URL:  "url",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var ExchangeWhere = struct {
	ID   whereHelperint
	Name whereHelperstring
	URL  whereHelperstring
}{
	ID:   whereHelperint{field: "\"exchange\".\"id\""},
	Name: whereHelperstring{field: "\"exchange\".\"name\""},
	URL:  whereHelperstring{field: "\"exchange\".\"url\""},
}

// ExchangeRels is where relationship names are stored.
var ExchangeRels = struct {
	ExchangeTicks string
}{
	ExchangeTicks: "ExchangeTicks",
}

// exchangeR is where relationships are stored.
type exchangeR struct {
	ExchangeTicks ExchangeTickSlice
}

// NewStruct creates a new relationship struct
func (*exchangeR) NewStruct() *exchangeR {
	return &exchangeR{}
}

// exchangeL is where Load methods for each relationship are stored.
type exchangeL struct{}

var (
	exchangeAllColumns            = []string{"id", "name", "url"}
	exchangeColumnsWithoutDefault = []string{"name", "url"}
	exchangeColumnsWithDefault    = []string{"id"}
	exchangePrimaryKeyColumns     = []string{"id"}
)

type (
	// ExchangeSlice is an alias for a slice of pointers to Exchange.
	// This should generally be used opposed to []Exchange.
	ExchangeSlice []*Exchange

	exchangeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	exchangeType                 = reflect.TypeOf(&Exchange{})
	exchangeMapping              = queries.MakeStructMapping(exchangeType)
	exchangePrimaryKeyMapping, _ = queries.BindMapping(exchangeType, exchangeMapping, exchangePrimaryKeyColumns)
	exchangeInsertCacheMut       sync.RWMutex
	exchangeInsertCache          = make(map[string]insertCache)
	exchangeUpdateCacheMut       sync.RWMutex
	exchangeUpdateCache          = make(map[string]updateCache)
	exchangeUpsertCacheMut       sync.RWMutex
	exchangeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single exchange record from the query.
func (q exchangeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Exchange, error) {
	o := &Exchange{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for exchange")
	}

	return o, nil
}

// All returns all Exchange records from the query.
func (q exchangeQuery) All(ctx context.Context, exec boil.ContextExecutor) (ExchangeSlice, error) {
	var o []*Exchange

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Exchange slice")
	}

	return o, nil
}

// Count returns the count of all Exchange records in the query.
func (q exchangeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count exchange rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q exchangeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if exchange exists")
	}

	return count > 0, nil
}

// ExchangeTicks retrieves all the exchange_tick's ExchangeTicks with an executor.
func (o *Exchange) ExchangeTicks(mods ...qm.QueryMod) exchangeTickQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"exchange_tick\".\"exchange_id\"=?", o.ID),
	)

	query := ExchangeTicks(queryMods...)
	queries.SetFrom(query.Query, "\"exchange_tick\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"exchange_tick\".*"})
	}

	return query
}

// LoadExchangeTicks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (exchangeL) LoadExchangeTicks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeExchange interface{}, mods queries.Applicator) error {
	var slice []*Exchange
	var object *Exchange

	if singular {
		object = maybeExchange.(*Exchange)
	} else {
		slice = *maybeExchange.(*[]*Exchange)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &exchangeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &exchangeR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`exchange_tick`), qm.WhereIn(`exchange_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load exchange_tick")
	}

	var resultSlice []*ExchangeTick
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice exchange_tick")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on exchange_tick")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for exchange_tick")
	}

	if singular {
		object.R.ExchangeTicks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &exchangeTickR{}
			}
			foreign.R.Exchange = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ExchangeID {
				local.R.ExchangeTicks = append(local.R.ExchangeTicks, foreign)
				if foreign.R == nil {
					foreign.R = &exchangeTickR{}
				}
				foreign.R.Exchange = local
				break
			}
		}
	}

	return nil
}

// AddExchangeTicks adds the given related objects to the existing relationships
// of the exchange, optionally inserting them as new records.
// Appends related to o.R.ExchangeTicks.
// Sets related.R.Exchange appropriately.
func (o *Exchange) AddExchangeTicks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ExchangeTick) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ExchangeID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"exchange_tick\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"exchange_id"}),
				strmangle.WhereClause("\"", "\"", 2, exchangeTickPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ExchangeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &exchangeR{
			ExchangeTicks: related,
		}
	} else {
		o.R.ExchangeTicks = append(o.R.ExchangeTicks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &exchangeTickR{
				Exchange: o,
			}
		} else {
			rel.R.Exchange = o
		}
	}
	return nil
}

// Exchanges retrieves all the records using an executor.
func Exchanges(mods ...qm.QueryMod) exchangeQuery {
	mods = append(mods, qm.From("\"exchange\""))
	return exchangeQuery{NewQuery(mods...)}
}

// FindExchange retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExchange(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Exchange, error) {
	exchangeObj := &Exchange{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"exchange\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, exchangeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from exchange")
	}

	return exchangeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Exchange) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no exchange provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(exchangeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	exchangeInsertCacheMut.RLock()
	cache, cached := exchangeInsertCache[key]
	exchangeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			exchangeAllColumns,
			exchangeColumnsWithDefault,
			exchangeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(exchangeType, exchangeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"exchange\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"exchange\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into exchange")
	}

	if !cached {
		exchangeInsertCacheMut.Lock()
		exchangeInsertCache[key] = cache
		exchangeInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Exchange.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Exchange) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	exchangeUpdateCacheMut.RLock()
	cache, cached := exchangeUpdateCache[key]
	exchangeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			exchangeAllColumns,
			exchangePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update exchange, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"exchange\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, exchangePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, append(wl, exchangePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update exchange row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for exchange")
	}

	if !cached {
		exchangeUpdateCacheMut.Lock()
		exchangeUpdateCache[key] = cache
		exchangeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q exchangeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for exchange")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for exchange")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExchangeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"exchange\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, exchangePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in exchange slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all exchange")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Exchange) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no exchange provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(exchangeColumnsWithDefault, o)

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

	exchangeUpsertCacheMut.RLock()
	cache, cached := exchangeUpsertCache[key]
	exchangeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			exchangeAllColumns,
			exchangeColumnsWithDefault,
			exchangeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			exchangeAllColumns,
			exchangePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert exchange, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(exchangePrimaryKeyColumns))
			copy(conflict, exchangePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"exchange\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(exchangeType, exchangeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(exchangeType, exchangeMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
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
		return errors.Wrap(err, "models: unable to upsert exchange")
	}

	if !cached {
		exchangeUpsertCacheMut.Lock()
		exchangeUpsertCache[key] = cache
		exchangeUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Exchange record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Exchange) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Exchange provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), exchangePrimaryKeyMapping)
	sql := "DELETE FROM \"exchange\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from exchange")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for exchange")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q exchangeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no exchangeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from exchange")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for exchange")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExchangeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"exchange\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, exchangePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from exchange slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for exchange")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Exchange) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindExchange(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExchangeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ExchangeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exchangePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"exchange\".* FROM \"exchange\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, exchangePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ExchangeSlice")
	}

	*o = slice

	return nil
}

// ExchangeExists checks if the Exchange row exists.
func ExchangeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"exchange\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if exchange exists")
	}

	return exists, nil
}
