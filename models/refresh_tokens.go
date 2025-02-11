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

// RefreshToken is an object representing the database table.
type RefreshToken struct {
	ID           uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt    null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt    null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	UserID       uint        `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	RefreshToken string      `boil:"refresh_token" json:"refresh_token" toml:"refresh_token" yaml:"refresh_token"`
	ExpireAt     time.Time   `boil:"expire_at" json:"expire_at" toml:"expire_at" yaml:"expire_at"`
	DeviceAppID  uint        `boil:"device_app_id" json:"device_app_id" toml:"device_app_id" yaml:"device_app_id"`
	Scopes       null.String `boil:"scopes" json:"scopes,omitempty" toml:"scopes" yaml:"scopes,omitempty"`
	SessionKey   string      `boil:"session_key" json:"session_key" toml:"session_key" yaml:"session_key"`

	R *refreshTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L refreshTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RefreshTokenColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	UserID       string
	RefreshToken string
	ExpireAt     string
	DeviceAppID  string
	Scopes       string
	SessionKey   string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	UserID:       "user_id",
	RefreshToken: "refresh_token",
	ExpireAt:     "expire_at",
	DeviceAppID:  "device_app_id",
	Scopes:       "scopes",
	SessionKey:   "session_key",
}

var RefreshTokenTableColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	UserID       string
	RefreshToken string
	ExpireAt     string
	DeviceAppID  string
	Scopes       string
	SessionKey   string
}{
	ID:           "refresh_tokens.id",
	CreatedAt:    "refresh_tokens.created_at",
	UpdatedAt:    "refresh_tokens.updated_at",
	DeletedAt:    "refresh_tokens.deleted_at",
	UserID:       "refresh_tokens.user_id",
	RefreshToken: "refresh_tokens.refresh_token",
	ExpireAt:     "refresh_tokens.expire_at",
	DeviceAppID:  "refresh_tokens.device_app_id",
	Scopes:       "refresh_tokens.scopes",
	SessionKey:   "refresh_tokens.session_key",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var RefreshTokenWhere = struct {
	ID           whereHelperuint
	CreatedAt    whereHelpernull_Time
	UpdatedAt    whereHelpernull_Time
	DeletedAt    whereHelpernull_Time
	UserID       whereHelperuint
	RefreshToken whereHelperstring
	ExpireAt     whereHelpertime_Time
	DeviceAppID  whereHelperuint
	Scopes       whereHelpernull_String
	SessionKey   whereHelperstring
}{
	ID:           whereHelperuint{field: "`refresh_tokens`.`id`"},
	CreatedAt:    whereHelpernull_Time{field: "`refresh_tokens`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`refresh_tokens`.`updated_at`"},
	DeletedAt:    whereHelpernull_Time{field: "`refresh_tokens`.`deleted_at`"},
	UserID:       whereHelperuint{field: "`refresh_tokens`.`user_id`"},
	RefreshToken: whereHelperstring{field: "`refresh_tokens`.`refresh_token`"},
	ExpireAt:     whereHelpertime_Time{field: "`refresh_tokens`.`expire_at`"},
	DeviceAppID:  whereHelperuint{field: "`refresh_tokens`.`device_app_id`"},
	Scopes:       whereHelpernull_String{field: "`refresh_tokens`.`scopes`"},
	SessionKey:   whereHelperstring{field: "`refresh_tokens`.`session_key`"},
}

// RefreshTokenRels is where relationship names are stored.
var RefreshTokenRels = struct {
}{}

// refreshTokenR is where relationships are stored.
type refreshTokenR struct {
}

// NewStruct creates a new relationship struct
func (*refreshTokenR) NewStruct() *refreshTokenR {
	return &refreshTokenR{}
}

// refreshTokenL is where Load methods for each relationship are stored.
type refreshTokenL struct{}

var (
	refreshTokenAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "user_id", "refresh_token", "expire_at", "device_app_id", "scopes", "session_key"}
	refreshTokenColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "user_id", "refresh_token", "device_app_id", "scopes", "session_key"}
	refreshTokenColumnsWithDefault    = []string{"id", "expire_at"}
	refreshTokenPrimaryKeyColumns     = []string{"id"}
	refreshTokenGeneratedColumns      = []string{}
)

type (
	// RefreshTokenSlice is an alias for a slice of pointers to RefreshToken.
	// This should almost always be used instead of []RefreshToken.
	RefreshTokenSlice []*RefreshToken
	// RefreshTokenHook is the signature for custom RefreshToken hook methods
	RefreshTokenHook func(boil.Executor, *RefreshToken) error

	refreshTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	refreshTokenType                 = reflect.TypeOf(&RefreshToken{})
	refreshTokenMapping              = queries.MakeStructMapping(refreshTokenType)
	refreshTokenPrimaryKeyMapping, _ = queries.BindMapping(refreshTokenType, refreshTokenMapping, refreshTokenPrimaryKeyColumns)
	refreshTokenInsertCacheMut       sync.RWMutex
	refreshTokenInsertCache          = make(map[string]insertCache)
	refreshTokenUpdateCacheMut       sync.RWMutex
	refreshTokenUpdateCache          = make(map[string]updateCache)
	refreshTokenUpsertCacheMut       sync.RWMutex
	refreshTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var refreshTokenAfterSelectHooks []RefreshTokenHook

var refreshTokenBeforeInsertHooks []RefreshTokenHook
var refreshTokenAfterInsertHooks []RefreshTokenHook

var refreshTokenBeforeUpdateHooks []RefreshTokenHook
var refreshTokenAfterUpdateHooks []RefreshTokenHook

var refreshTokenBeforeDeleteHooks []RefreshTokenHook
var refreshTokenAfterDeleteHooks []RefreshTokenHook

var refreshTokenBeforeUpsertHooks []RefreshTokenHook
var refreshTokenAfterUpsertHooks []RefreshTokenHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RefreshToken) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RefreshToken) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RefreshToken) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RefreshToken) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RefreshToken) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RefreshToken) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RefreshToken) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RefreshToken) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RefreshToken) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range refreshTokenAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRefreshTokenHook registers your hook function for all future operations.
func AddRefreshTokenHook(hookPoint boil.HookPoint, refreshTokenHook RefreshTokenHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		refreshTokenAfterSelectHooks = append(refreshTokenAfterSelectHooks, refreshTokenHook)
	case boil.BeforeInsertHook:
		refreshTokenBeforeInsertHooks = append(refreshTokenBeforeInsertHooks, refreshTokenHook)
	case boil.AfterInsertHook:
		refreshTokenAfterInsertHooks = append(refreshTokenAfterInsertHooks, refreshTokenHook)
	case boil.BeforeUpdateHook:
		refreshTokenBeforeUpdateHooks = append(refreshTokenBeforeUpdateHooks, refreshTokenHook)
	case boil.AfterUpdateHook:
		refreshTokenAfterUpdateHooks = append(refreshTokenAfterUpdateHooks, refreshTokenHook)
	case boil.BeforeDeleteHook:
		refreshTokenBeforeDeleteHooks = append(refreshTokenBeforeDeleteHooks, refreshTokenHook)
	case boil.AfterDeleteHook:
		refreshTokenAfterDeleteHooks = append(refreshTokenAfterDeleteHooks, refreshTokenHook)
	case boil.BeforeUpsertHook:
		refreshTokenBeforeUpsertHooks = append(refreshTokenBeforeUpsertHooks, refreshTokenHook)
	case boil.AfterUpsertHook:
		refreshTokenAfterUpsertHooks = append(refreshTokenAfterUpsertHooks, refreshTokenHook)
	}
}

// One returns a single refreshToken record from the query.
func (q refreshTokenQuery) One(exec boil.Executor) (*RefreshToken, error) {
	o := &RefreshToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for refresh_tokens")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RefreshToken records from the query.
func (q refreshTokenQuery) All(exec boil.Executor) (RefreshTokenSlice, error) {
	var o []*RefreshToken

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RefreshToken slice")
	}

	if len(refreshTokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RefreshToken records in the query.
func (q refreshTokenQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count refresh_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q refreshTokenQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if refresh_tokens exists")
	}

	return count > 0, nil
}

// RefreshTokens retrieves all the records using an executor.
func RefreshTokens(mods ...qm.QueryMod) refreshTokenQuery {
	mods = append(mods, qm.From("`refresh_tokens`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`refresh_tokens`.*"})
	}

	return refreshTokenQuery{q}
}

// FindRefreshToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRefreshToken(exec boil.Executor, iD uint, selectCols ...string) (*RefreshToken, error) {
	refreshTokenObj := &RefreshToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `refresh_tokens` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, refreshTokenObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from refresh_tokens")
	}

	if err = refreshTokenObj.doAfterSelectHooks(exec); err != nil {
		return refreshTokenObj, err
	}

	return refreshTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RefreshToken) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no refresh_tokens provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(refreshTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	refreshTokenInsertCacheMut.RLock()
	cache, cached := refreshTokenInsertCache[key]
	refreshTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			refreshTokenAllColumns,
			refreshTokenColumnsWithDefault,
			refreshTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `refresh_tokens` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `refresh_tokens` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `refresh_tokens` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, refreshTokenPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into refresh_tokens")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == refreshTokenMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for refresh_tokens")
	}

CacheNoHooks:
	if !cached {
		refreshTokenInsertCacheMut.Lock()
		refreshTokenInsertCache[key] = cache
		refreshTokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the RefreshToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RefreshToken) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	refreshTokenUpdateCacheMut.RLock()
	cache, cached := refreshTokenUpdateCache[key]
	refreshTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			refreshTokenAllColumns,
			refreshTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update refresh_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `refresh_tokens` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, refreshTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, append(wl, refreshTokenPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update refresh_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for refresh_tokens")
	}

	if !cached {
		refreshTokenUpdateCacheMut.Lock()
		refreshTokenUpdateCache[key] = cache
		refreshTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q refreshTokenQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for refresh_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for refresh_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RefreshTokenSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `refresh_tokens` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, refreshTokenPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in refreshToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all refreshToken")
	}
	return rowsAff, nil
}

var mySQLRefreshTokenUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RefreshToken) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no refresh_tokens provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(refreshTokenColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRefreshTokenUniqueColumns, o)

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

	refreshTokenUpsertCacheMut.RLock()
	cache, cached := refreshTokenUpsertCache[key]
	refreshTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			refreshTokenAllColumns,
			refreshTokenColumnsWithDefault,
			refreshTokenColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			refreshTokenAllColumns,
			refreshTokenPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert refresh_tokens, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`refresh_tokens`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `refresh_tokens` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for refresh_tokens")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == refreshTokenMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(refreshTokenType, refreshTokenMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for refresh_tokens")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for refresh_tokens")
	}

CacheNoHooks:
	if !cached {
		refreshTokenUpsertCacheMut.Lock()
		refreshTokenUpsertCache[key] = cache
		refreshTokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single RefreshToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RefreshToken) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RefreshToken provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), refreshTokenPrimaryKeyMapping)
	sql := "DELETE FROM `refresh_tokens` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from refresh_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for refresh_tokens")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q refreshTokenQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no refreshTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from refresh_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for refresh_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RefreshTokenSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(refreshTokenBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `refresh_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, refreshTokenPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from refreshToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for refresh_tokens")
	}

	if len(refreshTokenAfterDeleteHooks) != 0 {
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
func (o *RefreshToken) Reload(exec boil.Executor) error {
	ret, err := FindRefreshToken(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RefreshTokenSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RefreshTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), refreshTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `refresh_tokens`.* FROM `refresh_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, refreshTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RefreshTokenSlice")
	}

	*o = slice

	return nil
}

// RefreshTokenExists checks if the RefreshToken row exists.
func RefreshTokenExists(exec boil.Executor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `refresh_tokens` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if refresh_tokens exists")
	}

	return exists, nil
}

// Exists checks if the RefreshToken row exists.
func (o *RefreshToken) Exists(exec boil.Executor) (bool, error) {
	return RefreshTokenExists(exec, o.ID)
}
