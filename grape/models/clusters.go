// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Cluster is an object representing the database table.
type Cluster struct {
	ID         int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name       string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Code       string `boil:"code" json:"code" toml:"code" yaml:"code"`
	DeployType string `boil:"deploy_type" json:"deploy_type" toml:"deploy_type" yaml:"deploy_type"`
	Note       string `boil:"note" json:"note" toml:"note" yaml:"note"`

	R *clusterR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clusterL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClusterColumns = struct {
	ID         string
	Name       string
	Code       string
	DeployType string
	Note       string
}{
	ID:         "id",
	Name:       "name",
	Code:       "code",
	DeployType: "deploy_type",
	Note:       "note",
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
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ClusterWhere = struct {
	ID         whereHelperint64
	Name       whereHelperstring
	Code       whereHelperstring
	DeployType whereHelperstring
	Note       whereHelperstring
}{
	ID:         whereHelperint64{field: "\"clusters\".\"id\""},
	Name:       whereHelperstring{field: "\"clusters\".\"name\""},
	Code:       whereHelperstring{field: "\"clusters\".\"code\""},
	DeployType: whereHelperstring{field: "\"clusters\".\"deploy_type\""},
	Note:       whereHelperstring{field: "\"clusters\".\"note\""},
}

// ClusterRels is where relationship names are stored.
var ClusterRels = struct {
}{}

// clusterR is where relationships are stored.
type clusterR struct {
}

// NewStruct creates a new relationship struct
func (*clusterR) NewStruct() *clusterR {
	return &clusterR{}
}

// clusterL is where Load methods for each relationship are stored.
type clusterL struct{}

var (
	clusterAllColumns            = []string{"id", "name", "code", "deploy_type", "note"}
	clusterColumnsWithoutDefault = []string{"name", "code"}
	clusterColumnsWithDefault    = []string{"id", "deploy_type", "note"}
	clusterPrimaryKeyColumns     = []string{"id"}
)

type (
	// ClusterSlice is an alias for a slice of pointers to Cluster.
	// This should generally be used opposed to []Cluster.
	ClusterSlice []*Cluster
	// ClusterHook is the signature for custom Cluster hook methods
	ClusterHook func(context.Context, boil.ContextExecutor, *Cluster) error

	clusterQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clusterType                 = reflect.TypeOf(&Cluster{})
	clusterMapping              = queries.MakeStructMapping(clusterType)
	clusterPrimaryKeyMapping, _ = queries.BindMapping(clusterType, clusterMapping, clusterPrimaryKeyColumns)
	clusterInsertCacheMut       sync.RWMutex
	clusterInsertCache          = make(map[string]insertCache)
	clusterUpdateCacheMut       sync.RWMutex
	clusterUpdateCache          = make(map[string]updateCache)
	clusterUpsertCacheMut       sync.RWMutex
	clusterUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clusterBeforeInsertHooks []ClusterHook
var clusterBeforeUpdateHooks []ClusterHook
var clusterBeforeDeleteHooks []ClusterHook
var clusterBeforeUpsertHooks []ClusterHook

var clusterAfterInsertHooks []ClusterHook
var clusterAfterSelectHooks []ClusterHook
var clusterAfterUpdateHooks []ClusterHook
var clusterAfterDeleteHooks []ClusterHook
var clusterAfterUpsertHooks []ClusterHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cluster) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cluster) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cluster) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cluster) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cluster) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cluster) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cluster) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cluster) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cluster) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clusterAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClusterHook registers your hook function for all future operations.
func AddClusterHook(hookPoint boil.HookPoint, clusterHook ClusterHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		clusterBeforeInsertHooks = append(clusterBeforeInsertHooks, clusterHook)
	case boil.BeforeUpdateHook:
		clusterBeforeUpdateHooks = append(clusterBeforeUpdateHooks, clusterHook)
	case boil.BeforeDeleteHook:
		clusterBeforeDeleteHooks = append(clusterBeforeDeleteHooks, clusterHook)
	case boil.BeforeUpsertHook:
		clusterBeforeUpsertHooks = append(clusterBeforeUpsertHooks, clusterHook)
	case boil.AfterInsertHook:
		clusterAfterInsertHooks = append(clusterAfterInsertHooks, clusterHook)
	case boil.AfterSelectHook:
		clusterAfterSelectHooks = append(clusterAfterSelectHooks, clusterHook)
	case boil.AfterUpdateHook:
		clusterAfterUpdateHooks = append(clusterAfterUpdateHooks, clusterHook)
	case boil.AfterDeleteHook:
		clusterAfterDeleteHooks = append(clusterAfterDeleteHooks, clusterHook)
	case boil.AfterUpsertHook:
		clusterAfterUpsertHooks = append(clusterAfterUpsertHooks, clusterHook)
	}
}

// OneG returns a single cluster record from the query using the global executor.
func (q clusterQuery) OneG(ctx context.Context) (*Cluster, error) {
	return q.One(ctx, boil.GetContextDB())
}

// OneGP returns a single cluster record from the query using the global executor, and panics on error.
func (q clusterQuery) OneGP(ctx context.Context) *Cluster {
	o, err := q.One(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// OneP returns a single cluster record from the query, and panics on error.
func (q clusterQuery) OneP(ctx context.Context, exec boil.ContextExecutor) *Cluster {
	o, err := q.One(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cluster record from the query.
func (q clusterQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Cluster, error) {
	o := &Cluster{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for clusters")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Cluster records from the query using the global executor.
func (q clusterQuery) AllG(ctx context.Context) (ClusterSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// AllGP returns all Cluster records from the query using the global executor, and panics on error.
func (q clusterQuery) AllGP(ctx context.Context) ClusterSlice {
	o, err := q.All(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// AllP returns all Cluster records from the query, and panics on error.
func (q clusterQuery) AllP(ctx context.Context, exec boil.ContextExecutor) ClusterSlice {
	o, err := q.All(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Cluster records from the query.
func (q clusterQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClusterSlice, error) {
	var o []*Cluster

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Cluster slice")
	}

	if len(clusterAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Cluster records in the query, and panics on error.
func (q clusterQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// CountGP returns the count of all Cluster records in the query using the global executor, and panics on error.
func (q clusterQuery) CountGP(ctx context.Context) int64 {
	c, err := q.Count(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// CountP returns the count of all Cluster records in the query, and panics on error.
func (q clusterQuery) CountP(ctx context.Context, exec boil.ContextExecutor) int64 {
	c, err := q.Count(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Cluster records in the query.
func (q clusterQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count clusters rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q clusterQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// ExistsGP checks if the row exists in the table using the global executor, and panics on error.
func (q clusterQuery) ExistsGP(ctx context.Context) bool {
	e, err := q.Exists(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q clusterQuery) ExistsP(ctx context.Context, exec boil.ContextExecutor) bool {
	e, err := q.Exists(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q clusterQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if clusters exists")
	}

	return count > 0, nil
}

// Clusters retrieves all the records using an executor.
func Clusters(mods ...qm.QueryMod) clusterQuery {
	mods = append(mods, qm.From("\"clusters\""))
	return clusterQuery{NewQuery(mods...)}
}

// FindClusterG retrieves a single record by ID.
func FindClusterG(ctx context.Context, iD int64, selectCols ...string) (*Cluster, error) {
	return FindCluster(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindClusterP retrieves a single record by ID with an executor, and panics on error.
func FindClusterP(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) *Cluster {
	retobj, err := FindCluster(ctx, exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindClusterGP retrieves a single record by ID, and panics on error.
func FindClusterGP(ctx context.Context, iD int64, selectCols ...string) *Cluster {
	retobj, err := FindCluster(ctx, boil.GetContextDB(), iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCluster retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCluster(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Cluster, error) {
	clusterObj := &Cluster{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"clusters\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clusterObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from clusters")
	}

	return clusterObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Cluster) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Cluster) InsertP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) {
	if err := o.Insert(ctx, exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Cluster) InsertGP(ctx context.Context, columns boil.Columns) {
	if err := o.Insert(ctx, boil.GetContextDB(), columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Cluster) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clusters provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clusterColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clusterInsertCacheMut.RLock()
	cache, cached := clusterInsertCache[key]
	clusterInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clusterAllColumns,
			clusterColumnsWithDefault,
			clusterColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clusterType, clusterMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clusterType, clusterMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"clusters\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"clusters\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into clusters")
	}

	if !cached {
		clusterInsertCacheMut.Lock()
		clusterInsertCache[key] = cache
		clusterInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Cluster record using the global executor.
// See Update for more documentation.
func (o *Cluster) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// UpdateP uses an executor to update the Cluster, and panics on error.
// See Update for more documentation.
func (o *Cluster) UpdateP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) int64 {
	rowsAff, err := o.Update(ctx, exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateGP a single Cluster record using the global executor. Panics on error.
// See Update for more documentation.
func (o *Cluster) UpdateGP(ctx context.Context, columns boil.Columns) int64 {
	rowsAff, err := o.Update(ctx, boil.GetContextDB(), columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Update uses an executor to update the Cluster.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Cluster) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clusterUpdateCacheMut.RLock()
	cache, cached := clusterUpdateCache[key]
	clusterUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clusterAllColumns,
			clusterPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update clusters, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"clusters\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, clusterPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clusterType, clusterMapping, append(wl, clusterPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update clusters row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for clusters")
	}

	if !cached {
		clusterUpdateCacheMut.Lock()
		clusterUpdateCache[key] = cache
		clusterUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q clusterQuery) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := q.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAllG updates all rows with the specified column values.
func (q clusterQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q clusterQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for clusters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for clusters")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ClusterSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ClusterSlice) UpdateAllGP(ctx context.Context, cols M) int64 {
	rowsAff, err := o.UpdateAll(ctx, boil.GetContextDB(), cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ClusterSlice) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := o.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClusterSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clusterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"clusters\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, clusterPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cluster slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cluster")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Cluster) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Cluster) UpsertGP(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Cluster) UpsertP(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(ctx, exec, updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Cluster) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clusters provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(clusterColumnsWithDefault, o)

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

	clusterUpsertCacheMut.RLock()
	cache, cached := clusterUpsertCache[key]
	clusterUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clusterAllColumns,
			clusterColumnsWithDefault,
			clusterColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			clusterAllColumns,
			clusterPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert clusters, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(clusterPrimaryKeyColumns))
			copy(conflict, clusterPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"clusters\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(clusterType, clusterMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clusterType, clusterMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert clusters")
	}

	if !cached {
		clusterUpsertCacheMut.Lock()
		clusterUpsertCache[key] = cache
		clusterUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Cluster record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Cluster) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// DeleteP deletes a single Cluster record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cluster) DeleteP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.Delete(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteGP deletes a single Cluster record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cluster) DeleteGP(ctx context.Context) int64 {
	rowsAff, err := o.Delete(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Delete deletes a single Cluster record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cluster) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Cluster provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clusterPrimaryKeyMapping)
	sql := "DELETE FROM \"clusters\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from clusters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for clusters")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q clusterQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAllP deletes all rows, and panics on error.
func (q clusterQuery) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := q.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all matching rows.
func (q clusterQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clusterQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clusters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clusters")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ClusterSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ClusterSlice) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ClusterSlice) DeleteAllGP(ctx context.Context) int64 {
	rowsAff, err := o.DeleteAll(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClusterSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clusterBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clusterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"clusters\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clusterPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cluster slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clusters")
	}

	if len(clusterAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Cluster) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Cluster provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Cluster) ReloadP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.Reload(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Cluster) ReloadGP(ctx context.Context) {
	if err := o.Reload(ctx, boil.GetContextDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cluster) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCluster(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClusterSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ClusterSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ClusterSlice) ReloadAllP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.ReloadAll(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ClusterSlice) ReloadAllGP(ctx context.Context) {
	if err := o.ReloadAll(ctx, boil.GetContextDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClusterSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClusterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clusterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"clusters\".* FROM \"clusters\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, clusterPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClusterSlice")
	}

	*o = slice

	return nil
}

// ClusterExistsG checks if the Cluster row exists.
func ClusterExistsG(ctx context.Context, iD int64) (bool, error) {
	return ClusterExists(ctx, boil.GetContextDB(), iD)
}

// ClusterExistsP checks if the Cluster row exists. Panics on error.
func ClusterExistsP(ctx context.Context, exec boil.ContextExecutor, iD int64) bool {
	e, err := ClusterExists(ctx, exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ClusterExistsGP checks if the Cluster row exists. Panics on error.
func ClusterExistsGP(ctx context.Context, iD int64) bool {
	e, err := ClusterExists(ctx, boil.GetContextDB(), iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ClusterExists checks if the Cluster row exists.
func ClusterExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"clusters\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if clusters exists")
	}

	return exists, nil
}
