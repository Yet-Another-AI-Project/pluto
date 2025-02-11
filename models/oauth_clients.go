// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OauthClient is an object representing the database table.
type OauthClient struct {
	ID          uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt   null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt   null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Key         string    `boil:"key" json:"key" toml:"key" yaml:"key"`
	Secret      string    `boil:"secret" json:"secret" toml:"secret" yaml:"secret"`
	Status      string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	UserID      uint      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	RedirectURI string    `boil:"redirect_uri" json:"redirect_uri" toml:"redirect_uri" yaml:"redirect_uri"`

	R *oauthClientR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L oauthClientL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OauthClientColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	Key         string
	Secret      string
	Status      string
	UserID      string
	RedirectURI string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Key:         "key",
	Secret:      "secret",
	Status:      "status",
	UserID:      "user_id",
	RedirectURI: "redirect_uri",
}

var OauthClientTableColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	Key         string
	Secret      string
	Status      string
	UserID      string
	RedirectURI string
}{
	ID:          "oauth_clients.id",
	CreatedAt:   "oauth_clients.created_at",
	UpdatedAt:   "oauth_clients.updated_at",
	DeletedAt:   "oauth_clients.deleted_at",
	Key:         "oauth_clients.key",
	Secret:      "oauth_clients.secret",
	Status:      "oauth_clients.status",
	UserID:      "oauth_clients.user_id",
	RedirectURI: "oauth_clients.redirect_uri",
}

// Generated where

var OauthClientWhere = struct {
	ID          whereHelperuint
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
	DeletedAt   whereHelpernull_Time
	Key         whereHelperstring
	Secret      whereHelperstring
	Status      whereHelperstring
	UserID      whereHelperuint
	RedirectURI whereHelperstring
}{
	ID:          whereHelperuint{field: "`oauth_clients`.`id`"},
	CreatedAt:   whereHelpernull_Time{field: "`oauth_clients`.`created_at`"},
	UpdatedAt:   whereHelpernull_Time{field: "`oauth_clients`.`updated_at`"},
	DeletedAt:   whereHelpernull_Time{field: "`oauth_clients`.`deleted_at`"},
	Key:         whereHelperstring{field: "`oauth_clients`.`key`"},
	Secret:      whereHelperstring{field: "`oauth_clients`.`secret`"},
	Status:      whereHelperstring{field: "`oauth_clients`.`status`"},
	UserID:      whereHelperuint{field: "`oauth_clients`.`user_id`"},
	RedirectURI: whereHelperstring{field: "`oauth_clients`.`redirect_uri`"},
}

// OauthClientRels is where relationship names are stored.
var OauthClientRels = struct {
}{}

// oauthClientR is where relationships are stored.
type oauthClientR struct {
}

// NewStruct creates a new relationship struct
func (*oauthClientR) NewStruct() *oauthClientR {
	return &oauthClientR{}
}

// oauthClientL is where Load methods for each relationship are stored.
type oauthClientL struct{}

var (
	oauthClientAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "key", "secret", "status", "user_id", "redirect_uri"}
	oauthClientColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "key", "secret", "status", "user_id", "redirect_uri"}
	oauthClientColumnsWithDefault    = []string{"id"}
	oauthClientPrimaryKeyColumns     = []string{"id"}
	oauthClientGeneratedColumns      = []string{}
)

type (
	// OauthClientSlice is an alias for a slice of pointers to OauthClient.
	// This should almost always be used instead of []OauthClient.
	OauthClientSlice []*OauthClient
	// OauthClientHook is the signature for custom OauthClient hook methods
	OauthClientHook func(boil.Executor, *OauthClient) error

	oauthClientQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	oauthClientType                 = reflect.TypeOf(&OauthClient{})
	oauthClientMapping              = queries.MakeStructMapping(oauthClientType)
	oauthClientPrimaryKeyMapping, _ = queries.BindMapping(oauthClientType, oauthClientMapping, oauthClientPrimaryKeyColumns)
	oauthClientInsertCacheMut       sync.RWMutex
	oauthClientInsertCache          = make(map[string]insertCache)
	oauthClientUpdateCacheMut       sync.RWMutex
	oauthClientUpdateCache          = make(map[string]updateCache)
	oauthClientUpsertCacheMut       sync.RWMutex
	oauthClientUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var oauthClientAfterSelectHooks []OauthClientHook

var oauthClientBeforeInsertHooks []OauthClientHook
var oauthClientAfterInsertHooks []OauthClientHook

var oauthClientBeforeUpdateHooks []OauthClientHook
var oauthClientAfterUpdateHooks []OauthClientHook

var oauthClientBeforeDeleteHooks []OauthClientHook
var oauthClientAfterDeleteHooks []OauthClientHook

var oauthClientBeforeUpsertHooks []OauthClientHook
var oauthClientAfterUpsertHooks []OauthClientHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OauthClient) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OauthClient) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OauthClient) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OauthClient) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OauthClient) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OauthClient) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OauthClient) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OauthClient) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OauthClient) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range oauthClientAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOauthClientHook registers your hook function for all future operations.
func AddOauthClientHook(hookPoint boil.HookPoint, oauthClientHook OauthClientHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		oauthClientAfterSelectHooks = append(oauthClientAfterSelectHooks, oauthClientHook)
	case boil.BeforeInsertHook:
		oauthClientBeforeInsertHooks = append(oauthClientBeforeInsertHooks, oauthClientHook)
	case boil.AfterInsertHook:
		oauthClientAfterInsertHooks = append(oauthClientAfterInsertHooks, oauthClientHook)
	case boil.BeforeUpdateHook:
		oauthClientBeforeUpdateHooks = append(oauthClientBeforeUpdateHooks, oauthClientHook)
	case boil.AfterUpdateHook:
		oauthClientAfterUpdateHooks = append(oauthClientAfterUpdateHooks, oauthClientHook)
	case boil.BeforeDeleteHook:
		oauthClientBeforeDeleteHooks = append(oauthClientBeforeDeleteHooks, oauthClientHook)
	case boil.AfterDeleteHook:
		oauthClientAfterDeleteHooks = append(oauthClientAfterDeleteHooks, oauthClientHook)
	case boil.BeforeUpsertHook:
		oauthClientBeforeUpsertHooks = append(oauthClientBeforeUpsertHooks, oauthClientHook)
	case boil.AfterUpsertHook:
		oauthClientAfterUpsertHooks = append(oauthClientAfterUpsertHooks, oauthClientHook)
	}
}

// One returns a single oauthClient record from the query.
func (q oauthClientQuery) One(exec boil.Executor) (*OauthClient, error) {
	o := &OauthClient{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for oauth_clients")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OauthClient records from the query.
func (q oauthClientQuery) All(exec boil.Executor) (OauthClientSlice, error) {
	var o []*OauthClient

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OauthClient slice")
	}

	if len(oauthClientAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OauthClient records in the query.
func (q oauthClientQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count oauth_clients rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q oauthClientQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if oauth_clients exists")
	}

	return count > 0, nil
}

// OauthClients retrieves all the records using an executor.
func OauthClients(mods ...qm.QueryMod) oauthClientQuery {
	mods = append(mods, qm.From("`oauth_clients`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`oauth_clients`.*"})
	}

	return oauthClientQuery{q}
}

// FindOauthClient retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOauthClient(exec boil.Executor, iD uint, selectCols ...string) (*OauthClient, error) {
	oauthClientObj := &OauthClient{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `oauth_clients` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, oauthClientObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from oauth_clients")
	}

	if err = oauthClientObj.doAfterSelectHooks(exec); err != nil {
		return oauthClientObj, err
	}

	return oauthClientObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OauthClient) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no oauth_clients provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(oauthClientColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	oauthClientInsertCacheMut.RLock()
	cache, cached := oauthClientInsertCache[key]
	oauthClientInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			oauthClientAllColumns,
			oauthClientColumnsWithDefault,
			oauthClientColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(oauthClientType, oauthClientMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(oauthClientType, oauthClientMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `oauth_clients` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `oauth_clients` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `oauth_clients` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, oauthClientPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into oauth_clients")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == oauthClientMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for oauth_clients")
	}

CacheNoHooks:
	if !cached {
		oauthClientInsertCacheMut.Lock()
		oauthClientInsertCache[key] = cache
		oauthClientInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the OauthClient.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OauthClient) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	oauthClientUpdateCacheMut.RLock()
	cache, cached := oauthClientUpdateCache[key]
	oauthClientUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			oauthClientAllColumns,
			oauthClientPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update oauth_clients, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `oauth_clients` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, oauthClientPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(oauthClientType, oauthClientMapping, append(wl, oauthClientPrimaryKeyColumns...))
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
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update oauth_clients row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for oauth_clients")
	}

	if !cached {
		oauthClientUpdateCacheMut.Lock()
		oauthClientUpdateCache[key] = cache
		oauthClientUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q oauthClientQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for oauth_clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for oauth_clients")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OauthClientSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthClientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `oauth_clients` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthClientPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in oauthClient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all oauthClient")
	}
	return rowsAff, nil
}

var mySQLOauthClientUniqueColumns = []string{
	"id",
	"key",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OauthClient) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no oauth_clients provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(oauthClientColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOauthClientUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	oauthClientUpsertCacheMut.RLock()
	cache, cached := oauthClientUpsertCache[key]
	oauthClientUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			oauthClientAllColumns,
			oauthClientColumnsWithDefault,
			oauthClientColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			oauthClientAllColumns,
			oauthClientPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert oauth_clients, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`oauth_clients`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `oauth_clients` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(oauthClientType, oauthClientMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(oauthClientType, oauthClientMapping, ret)
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
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for oauth_clients")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == oauthClientMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(oauthClientType, oauthClientMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for oauth_clients")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for oauth_clients")
	}

CacheNoHooks:
	if !cached {
		oauthClientUpsertCacheMut.Lock()
		oauthClientUpsertCache[key] = cache
		oauthClientUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single OauthClient record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OauthClient) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OauthClient provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), oauthClientPrimaryKeyMapping)
	sql := "DELETE FROM `oauth_clients` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from oauth_clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for oauth_clients")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q oauthClientQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no oauthClientQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oauth_clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for oauth_clients")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OauthClientSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(oauthClientBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthClientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `oauth_clients` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthClientPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oauthClient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for oauth_clients")
	}

	if len(oauthClientAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OauthClient) Reload(exec boil.Executor) error {
	ret, err := FindOauthClient(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OauthClientSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OauthClientSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthClientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `oauth_clients`.* FROM `oauth_clients` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oauthClientPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OauthClientSlice")
	}

	*o = slice

	return nil
}

// OauthClientExists checks if the OauthClient row exists.
func OauthClientExists(exec boil.Executor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `oauth_clients` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if oauth_clients exists")
	}

	return exists, nil
}

// Exists checks if the OauthClient row exists.
func (o *OauthClient) Exists(exec boil.Executor) (bool, error) {
	return OauthClientExists(exec, o.ID)
}
