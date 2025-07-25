// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/quailyquaily/gzk9000/core"
)

func newFact(db *gorm.DB, opts ...gen.DOOption) fact {
	_fact := fact{}

	_fact.factDo.UseDB(db, opts...)
	_fact.factDo.UseModel(&core.Fact{})

	tableName := _fact.factDo.TableName()
	_fact.ALL = field.NewAsterisk(tableName)
	_fact.ID = field.NewUint64(tableName, "id")
	_fact.AgentID = field.NewUint64(tableName, "agent_id")
	_fact.Content = field.NewString(tableName, "content")
	_fact.Sentiment = field.NewFloat64(tableName, "sentiment")
	_fact.CreatedAt = field.NewTime(tableName, "created_at")
	_fact.UpdatedAt = field.NewTime(tableName, "updated_at")

	_fact.fillFieldMap()

	return _fact
}

type fact struct {
	factDo

	ALL       field.Asterisk
	ID        field.Uint64
	AgentID   field.Uint64
	Content   field.String
	Sentiment field.Float64
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (f fact) Table(newTableName string) *fact {
	f.factDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fact) As(alias string) *fact {
	f.factDo.DO = *(f.factDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fact) updateTableName(table string) *fact {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewUint64(table, "id")
	f.AgentID = field.NewUint64(table, "agent_id")
	f.Content = field.NewString(table, "content")
	f.Sentiment = field.NewFloat64(table, "sentiment")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fact) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fact) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 6)
	f.fieldMap["id"] = f.ID
	f.fieldMap["agent_id"] = f.AgentID
	f.fieldMap["content"] = f.Content
	f.fieldMap["sentiment"] = f.Sentiment
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fact) clone(db *gorm.DB) fact {
	f.factDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fact) replaceDB(db *gorm.DB) fact {
	f.factDo.ReplaceDB(db)
	return f
}

type factDo struct{ gen.DO }

type IFactDo interface {
	gen.SubQuery
	Debug() IFactDo
	WithContext(ctx context.Context) IFactDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFactDo
	WriteDB() IFactDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFactDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFactDo
	Not(conds ...gen.Condition) IFactDo
	Or(conds ...gen.Condition) IFactDo
	Select(conds ...field.Expr) IFactDo
	Where(conds ...gen.Condition) IFactDo
	Order(conds ...field.Expr) IFactDo
	Distinct(cols ...field.Expr) IFactDo
	Omit(cols ...field.Expr) IFactDo
	Join(table schema.Tabler, on ...field.Expr) IFactDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFactDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFactDo
	Group(cols ...field.Expr) IFactDo
	Having(conds ...gen.Condition) IFactDo
	Limit(limit int) IFactDo
	Offset(offset int) IFactDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFactDo
	Unscoped() IFactDo
	Create(values ...*core.Fact) error
	CreateInBatches(values []*core.Fact, batchSize int) error
	Save(values ...*core.Fact) error
	First() (*core.Fact, error)
	Take() (*core.Fact, error)
	Last() (*core.Fact, error)
	Find() ([]*core.Fact, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*core.Fact, err error)
	FindInBatches(result *[]*core.Fact, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*core.Fact) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFactDo
	Assign(attrs ...field.AssignExpr) IFactDo
	Joins(fields ...field.RelationField) IFactDo
	Preload(fields ...field.RelationField) IFactDo
	FirstOrInit() (*core.Fact, error)
	FirstOrCreate() (*core.Fact, error)
	FindByPage(offset int, limit int) (result []*core.Fact, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFactDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	CreateFact(ctx context.Context, fact *core.Fact) (err error)
	GetFactByID(ctx context.Context, id uint64) (result *core.Fact, err error)
	GetFactsByIDs(ctx context.Context, ids []uint64) (result []*core.Fact, err error)
}

// INSERT INTO @@table (
//
//	agent_id, content, sentiment,
//	created_at, updated_at
//
// ) VALUES (
//
//	@fact.AgentID, @fact.Content, @fact.Sentiment,
//	NOW(), NOW()
//
// ) RETURNING id;
func (f factDo) CreateFact(ctx context.Context, fact *core.Fact) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, fact.AgentID)
	params = append(params, fact.Content)
	params = append(params, fact.Sentiment)
	generateSQL.WriteString("INSERT INTO facts ( agent_id, content, sentiment, created_at, updated_at ) VALUES ( ?, ?, ?, NOW(), NOW() ) RETURNING id; ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// WHERE id = @id;
func (f factDo) GetFactByID(ctx context.Context, id uint64) (result *core.Fact, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM facts WHERE id = ?; ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// WHERE id IN (@ids);
func (f factDo) GetFactsByIDs(ctx context.Context, ids []uint64) (result []*core.Fact, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("SELECT * FROM facts WHERE id IN (?); ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (f factDo) Debug() IFactDo {
	return f.withDO(f.DO.Debug())
}

func (f factDo) WithContext(ctx context.Context) IFactDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f factDo) ReadDB() IFactDo {
	return f.Clauses(dbresolver.Read)
}

func (f factDo) WriteDB() IFactDo {
	return f.Clauses(dbresolver.Write)
}

func (f factDo) Session(config *gorm.Session) IFactDo {
	return f.withDO(f.DO.Session(config))
}

func (f factDo) Clauses(conds ...clause.Expression) IFactDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f factDo) Returning(value interface{}, columns ...string) IFactDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f factDo) Not(conds ...gen.Condition) IFactDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f factDo) Or(conds ...gen.Condition) IFactDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f factDo) Select(conds ...field.Expr) IFactDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f factDo) Where(conds ...gen.Condition) IFactDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f factDo) Order(conds ...field.Expr) IFactDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f factDo) Distinct(cols ...field.Expr) IFactDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f factDo) Omit(cols ...field.Expr) IFactDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f factDo) Join(table schema.Tabler, on ...field.Expr) IFactDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f factDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFactDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f factDo) RightJoin(table schema.Tabler, on ...field.Expr) IFactDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f factDo) Group(cols ...field.Expr) IFactDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f factDo) Having(conds ...gen.Condition) IFactDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f factDo) Limit(limit int) IFactDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f factDo) Offset(offset int) IFactDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f factDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFactDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f factDo) Unscoped() IFactDo {
	return f.withDO(f.DO.Unscoped())
}

func (f factDo) Create(values ...*core.Fact) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f factDo) CreateInBatches(values []*core.Fact, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f factDo) Save(values ...*core.Fact) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f factDo) First() (*core.Fact, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*core.Fact), nil
	}
}

func (f factDo) Take() (*core.Fact, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*core.Fact), nil
	}
}

func (f factDo) Last() (*core.Fact, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*core.Fact), nil
	}
}

func (f factDo) Find() ([]*core.Fact, error) {
	result, err := f.DO.Find()
	return result.([]*core.Fact), err
}

func (f factDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*core.Fact, err error) {
	buf := make([]*core.Fact, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f factDo) FindInBatches(result *[]*core.Fact, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f factDo) Attrs(attrs ...field.AssignExpr) IFactDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f factDo) Assign(attrs ...field.AssignExpr) IFactDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f factDo) Joins(fields ...field.RelationField) IFactDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f factDo) Preload(fields ...field.RelationField) IFactDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f factDo) FirstOrInit() (*core.Fact, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*core.Fact), nil
	}
}

func (f factDo) FirstOrCreate() (*core.Fact, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*core.Fact), nil
	}
}

func (f factDo) FindByPage(offset int, limit int) (result []*core.Fact, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f factDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f factDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f factDo) Delete(models ...*core.Fact) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *factDo) withDO(do gen.Dao) *factDo {
	f.DO = *do.(*gen.DO)
	return f
}
