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

// Node is an object representing the database table.
type Node struct {
	Address         string `boil:"address" json:"address" toml:"address" yaml:"address"`
	IPVersion       int    `boil:"ip_version" json:"ip_version" toml:"ip_version" yaml:"ip_version"`
	Country         string `boil:"country" json:"country" toml:"country" yaml:"country"`
	State           string `boil:"state" json:"state" toml:"state" yaml:"state"`
	City            string `boil:"city" json:"city" toml:"city" yaml:"city"`
	Locality        string `boil:"locality" json:"locality" toml:"locality" yaml:"locality"`
	LastAttempt     int64  `boil:"last_attempt" json:"last_attempt" toml:"last_attempt" yaml:"last_attempt"`
	LastSeen        int64  `boil:"last_seen" json:"last_seen" toml:"last_seen" yaml:"last_seen"`
	IsDead          bool   `boil:"is_dead" json:"is_dead" toml:"is_dead" yaml:"is_dead"`
	ConnectionTime  int64  `boil:"connection_time" json:"connection_time" toml:"connection_time" yaml:"connection_time"`
	ProtocolVersion int    `boil:"protocol_version" json:"protocol_version" toml:"protocol_version" yaml:"protocol_version"`
	UserAgent       string `boil:"user_agent" json:"user_agent" toml:"user_agent" yaml:"user_agent"`
	Services        string `boil:"services" json:"services" toml:"services" yaml:"services"`
	StartingHeight  int64  `boil:"starting_height" json:"starting_height" toml:"starting_height" yaml:"starting_height"`
	CurrentHeight   int64  `boil:"current_height" json:"current_height" toml:"current_height" yaml:"current_height"`

	R *nodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L nodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NodeColumns = struct {
	Address         string
	IPVersion       string
	Country         string
	State           string
	City            string
	Locality        string
	LastAttempt     string
	LastSeen        string
	IsDead          string
	ConnectionTime  string
	ProtocolVersion string
	UserAgent       string
	Services        string
	StartingHeight  string
	CurrentHeight   string
}{
	Address:         "address",
	IPVersion:       "ip_version",
	Country:         "country",
	State:           "state",
	City:            "city",
	Locality:        "locality",
	LastAttempt:     "last_attempt",
	LastSeen:        "last_seen",
	IsDead:          "is_dead",
	ConnectionTime:  "connection_time",
	ProtocolVersion: "protocol_version",
	UserAgent:       "user_agent",
	Services:        "services",
	StartingHeight:  "starting_height",
	CurrentHeight:   "current_height",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var NodeWhere = struct {
	Address         whereHelperstring
	IPVersion       whereHelperint
	Country         whereHelperstring
	State           whereHelperstring
	City            whereHelperstring
	Locality        whereHelperstring
	LastAttempt     whereHelperint64
	LastSeen        whereHelperint64
	IsDead          whereHelperbool
	ConnectionTime  whereHelperint64
	ProtocolVersion whereHelperint
	UserAgent       whereHelperstring
	Services        whereHelperstring
	StartingHeight  whereHelperint64
	CurrentHeight   whereHelperint64
}{
	Address:         whereHelperstring{field: "\"node\".\"address\""},
	IPVersion:       whereHelperint{field: "\"node\".\"ip_version\""},
	Country:         whereHelperstring{field: "\"node\".\"country\""},
	State:           whereHelperstring{field: "\"node\".\"state\""},
	City:            whereHelperstring{field: "\"node\".\"city\""},
	Locality:        whereHelperstring{field: "\"node\".\"locality\""},
	LastAttempt:     whereHelperint64{field: "\"node\".\"last_attempt\""},
	LastSeen:        whereHelperint64{field: "\"node\".\"last_seen\""},
	IsDead:          whereHelperbool{field: "\"node\".\"is_dead\""},
	ConnectionTime:  whereHelperint64{field: "\"node\".\"connection_time\""},
	ProtocolVersion: whereHelperint{field: "\"node\".\"protocol_version\""},
	UserAgent:       whereHelperstring{field: "\"node\".\"user_agent\""},
	Services:        whereHelperstring{field: "\"node\".\"services\""},
	StartingHeight:  whereHelperint64{field: "\"node\".\"starting_height\""},
	CurrentHeight:   whereHelperint64{field: "\"node\".\"current_height\""},
}

// NodeRels is where relationship names are stored.
var NodeRels = struct {
	Heartbeats string
}{
	Heartbeats: "Heartbeats",
}

// nodeR is where relationships are stored.
type nodeR struct {
	Heartbeats HeartbeatSlice
}

// NewStruct creates a new relationship struct
func (*nodeR) NewStruct() *nodeR {
	return &nodeR{}
}

// nodeL is where Load methods for each relationship are stored.
type nodeL struct{}

var (
	nodeAllColumns            = []string{"address", "ip_version", "country", "state", "city", "locality", "last_attempt", "last_seen", "is_dead", "connection_time", "protocol_version", "user_agent", "services", "starting_height", "current_height"}
	nodeColumnsWithoutDefault = []string{"address", "ip_version", "country", "state", "city", "locality", "last_attempt", "last_seen", "is_dead", "connection_time", "protocol_version", "user_agent", "services", "starting_height", "current_height"}
	nodeColumnsWithDefault    = []string{}
	nodePrimaryKeyColumns     = []string{"address"}
)

type (
	// NodeSlice is an alias for a slice of pointers to Node.
	// This should generally be used opposed to []Node.
	NodeSlice []*Node

	nodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	nodeType                 = reflect.TypeOf(&Node{})
	nodeMapping              = queries.MakeStructMapping(nodeType)
	nodePrimaryKeyMapping, _ = queries.BindMapping(nodeType, nodeMapping, nodePrimaryKeyColumns)
	nodeInsertCacheMut       sync.RWMutex
	nodeInsertCache          = make(map[string]insertCache)
	nodeUpdateCacheMut       sync.RWMutex
	nodeUpdateCache          = make(map[string]updateCache)
	nodeUpsertCacheMut       sync.RWMutex
	nodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single node record from the query.
func (q nodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Node, error) {
	o := &Node{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for node")
	}

	return o, nil
}

// All returns all Node records from the query.
func (q nodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (NodeSlice, error) {
	var o []*Node

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Node slice")
	}

	return o, nil
}

// Count returns the count of all Node records in the query.
func (q nodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count node rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q nodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if node exists")
	}

	return count > 0, nil
}

// Heartbeats retrieves all the heartbeat's Heartbeats with an executor.
func (o *Node) Heartbeats(mods ...qm.QueryMod) heartbeatQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"heartbeat\".\"node_id\"=?", o.Address),
	)

	query := Heartbeats(queryMods...)
	queries.SetFrom(query.Query, "\"heartbeat\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"heartbeat\".*"})
	}

	return query
}

// LoadHeartbeats allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (nodeL) LoadHeartbeats(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNode interface{}, mods queries.Applicator) error {
	var slice []*Node
	var object *Node

	if singular {
		object = maybeNode.(*Node)
	} else {
		slice = *maybeNode.(*[]*Node)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &nodeR{}
		}
		args = append(args, object.Address)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &nodeR{}
			}

			for _, a := range args {
				if a == obj.Address {
					continue Outer
				}
			}

			args = append(args, obj.Address)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`heartbeat`), qm.WhereIn(`heartbeat.node_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load heartbeat")
	}

	var resultSlice []*Heartbeat
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice heartbeat")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on heartbeat")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for heartbeat")
	}

	if singular {
		object.R.Heartbeats = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &heartbeatR{}
			}
			foreign.R.Node = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Address == foreign.NodeID {
				local.R.Heartbeats = append(local.R.Heartbeats, foreign)
				if foreign.R == nil {
					foreign.R = &heartbeatR{}
				}
				foreign.R.Node = local
				break
			}
		}
	}

	return nil
}

// AddHeartbeats adds the given related objects to the existing relationships
// of the node, optionally inserting them as new records.
// Appends related to o.R.Heartbeats.
// Sets related.R.Node appropriately.
func (o *Node) AddHeartbeats(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Heartbeat) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.NodeID = o.Address
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"heartbeat\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"node_id"}),
				strmangle.WhereClause("\"", "\"", 2, heartbeatPrimaryKeyColumns),
			)
			values := []interface{}{o.Address, rel.Timestamp, rel.NodeID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.NodeID = o.Address
		}
	}

	if o.R == nil {
		o.R = &nodeR{
			Heartbeats: related,
		}
	} else {
		o.R.Heartbeats = append(o.R.Heartbeats, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &heartbeatR{
				Node: o,
			}
		} else {
			rel.R.Node = o
		}
	}
	return nil
}

// Nodes retrieves all the records using an executor.
func Nodes(mods ...qm.QueryMod) nodeQuery {
	mods = append(mods, qm.From("\"node\""))
	return nodeQuery{NewQuery(mods...)}
}

// FindNode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNode(ctx context.Context, exec boil.ContextExecutor, address string, selectCols ...string) (*Node, error) {
	nodeObj := &Node{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"node\" where \"address\"=$1", sel,
	)

	q := queries.Raw(query, address)

	err := q.Bind(ctx, exec, nodeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from node")
	}

	return nodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Node) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no node provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(nodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	nodeInsertCacheMut.RLock()
	cache, cached := nodeInsertCache[key]
	nodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			nodeAllColumns,
			nodeColumnsWithDefault,
			nodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(nodeType, nodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(nodeType, nodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"node\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"node\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into node")
	}

	if !cached {
		nodeInsertCacheMut.Lock()
		nodeInsertCache[key] = cache
		nodeInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Node.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Node) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	nodeUpdateCacheMut.RLock()
	cache, cached := nodeUpdateCache[key]
	nodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			nodeAllColumns,
			nodePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update node, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"node\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, nodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(nodeType, nodeMapping, append(wl, nodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update node row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for node")
	}

	if !cached {
		nodeUpdateCacheMut.Lock()
		nodeUpdateCache[key] = cache
		nodeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q nodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for node")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for node")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"node\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, nodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in node slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all node")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Node) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no node provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(nodeColumnsWithDefault, o)

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

	nodeUpsertCacheMut.RLock()
	cache, cached := nodeUpsertCache[key]
	nodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			nodeAllColumns,
			nodeColumnsWithDefault,
			nodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			nodeAllColumns,
			nodePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert node, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(nodePrimaryKeyColumns))
			copy(conflict, nodePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"node\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(nodeType, nodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(nodeType, nodeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert node")
	}

	if !cached {
		nodeUpsertCacheMut.Lock()
		nodeUpsertCache[key] = cache
		nodeUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Node record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Node) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Node provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), nodePrimaryKeyMapping)
	sql := "DELETE FROM \"node\" WHERE \"address\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from node")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for node")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q nodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no nodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from node")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"node\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, nodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from node slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for node")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Node) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNode(ctx, exec, o.Address)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"node\".* FROM \"node\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, nodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NodeSlice")
	}

	*o = slice

	return nil
}

// NodeExists checks if the Node row exists.
func NodeExists(ctx context.Context, exec boil.ContextExecutor, address string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"node\" where \"address\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, address)
	}
	row := exec.QueryRowContext(ctx, sql, address)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if node exists")
	}

	return exists, nil
}
