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

func newStudygoal(db *gorm.DB, opts ...gen.DOOption) studygoal {
	_studygoal := studygoal{}

	_studygoal.studygoalDo.UseDB(db, opts...)
	_studygoal.studygoalDo.UseModel(&core.Studygoal{})

	tableName := _studygoal.studygoalDo.TableName()
	_studygoal.ALL = field.NewAsterisk(tableName)
	_studygoal.ID = field.NewUint64(tableName, "id")
	_studygoal.AgentID = field.NewUint64(tableName, "agent_id")
	_studygoal.Content = field.NewString(tableName, "content")
	_studygoal.Iteration = field.NewInt(tableName, "iteration")
	_studygoal.PriorityScore = field.NewFloat64(tableName, "priority_score")
	_studygoal.Status = field.NewInt(tableName, "status")
	_studygoal.CreatedAt = field.NewTime(tableName, "created_at")
	_studygoal.UpdatedAt = field.NewTime(tableName, "updated_at")

	_studygoal.fillFieldMap()

	return _studygoal
}

type studygoal struct {
	studygoalDo

	ALL           field.Asterisk
	ID            field.Uint64
	AgentID       field.Uint64
	Content       field.String
	Iteration     field.Int
	PriorityScore field.Float64
	Status        field.Int
	CreatedAt     field.Time
	UpdatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (s studygoal) Table(newTableName string) *studygoal {
	s.studygoalDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s studygoal) As(alias string) *studygoal {
	s.studygoalDo.DO = *(s.studygoalDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *studygoal) updateTableName(table string) *studygoal {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.AgentID = field.NewUint64(table, "agent_id")
	s.Content = field.NewString(table, "content")
	s.Iteration = field.NewInt(table, "iteration")
	s.PriorityScore = field.NewFloat64(table, "priority_score")
	s.Status = field.NewInt(table, "status")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *studygoal) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *studygoal) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["agent_id"] = s.AgentID
	s.fieldMap["content"] = s.Content
	s.fieldMap["iteration"] = s.Iteration
	s.fieldMap["priority_score"] = s.PriorityScore
	s.fieldMap["status"] = s.Status
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s studygoal) clone(db *gorm.DB) studygoal {
	s.studygoalDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s studygoal) replaceDB(db *gorm.DB) studygoal {
	s.studygoalDo.ReplaceDB(db)
	return s
}

type studygoalDo struct{ gen.DO }

type IStudygoalDo interface {
	gen.SubQuery
	Debug() IStudygoalDo
	WithContext(ctx context.Context) IStudygoalDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStudygoalDo
	WriteDB() IStudygoalDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStudygoalDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStudygoalDo
	Not(conds ...gen.Condition) IStudygoalDo
	Or(conds ...gen.Condition) IStudygoalDo
	Select(conds ...field.Expr) IStudygoalDo
	Where(conds ...gen.Condition) IStudygoalDo
	Order(conds ...field.Expr) IStudygoalDo
	Distinct(cols ...field.Expr) IStudygoalDo
	Omit(cols ...field.Expr) IStudygoalDo
	Join(table schema.Tabler, on ...field.Expr) IStudygoalDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStudygoalDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStudygoalDo
	Group(cols ...field.Expr) IStudygoalDo
	Having(conds ...gen.Condition) IStudygoalDo
	Limit(limit int) IStudygoalDo
	Offset(offset int) IStudygoalDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStudygoalDo
	Unscoped() IStudygoalDo
	Create(values ...*core.Studygoal) error
	CreateInBatches(values []*core.Studygoal, batchSize int) error
	Save(values ...*core.Studygoal) error
	First() (*core.Studygoal, error)
	Take() (*core.Studygoal, error)
	Last() (*core.Studygoal, error)
	Find() ([]*core.Studygoal, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*core.Studygoal, err error)
	FindInBatches(result *[]*core.Studygoal, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*core.Studygoal) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStudygoalDo
	Assign(attrs ...field.AssignExpr) IStudygoalDo
	Joins(fields ...field.RelationField) IStudygoalDo
	Preload(fields ...field.RelationField) IStudygoalDo
	FirstOrInit() (*core.Studygoal, error)
	FirstOrCreate() (*core.Studygoal, error)
	FindByPage(offset int, limit int) (result []*core.Studygoal, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStudygoalDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	CreateStudygoal(ctx context.Context, goal *core.Studygoal) (err error)
	GetStudygoalByID(ctx context.Context, id uint64) (result *core.Studygoal, err error)
	GetStudygoalsByAgentID(ctx context.Context, agentID uint64) (result []*core.Studygoal, err error)
	GetActiveStudygoalsByAgentID(ctx context.Context, agentID uint64) (result []*core.Studygoal, err error)
}

// INSERT INTO @@table (
//
//	agent_id, content, iteration, status,
//	created_at, updated_at
//
// ) VALUES (
//
//	@goal.AgentID, @goal.Content, 0, @goal.Status,
//	NOW(), NOW()
//
// ) RETURNING id;
func (s studygoalDo) CreateStudygoal(ctx context.Context, goal *core.Studygoal) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, goal.AgentID)
	params = append(params, goal.Content)
	params = append(params, goal.Status)
	generateSQL.WriteString("INSERT INTO studygoals ( agent_id, content, iteration, status, created_at, updated_at ) VALUES ( ?, ?, 0, ?, NOW(), NOW() ) RETURNING id; ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// WHERE id = @id;
func (s studygoalDo) GetStudygoalByID(ctx context.Context, id uint64) (result *core.Studygoal, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM studygoals WHERE id = ?; ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// WHERE agent_id = @agentID;
func (s studygoalDo) GetStudygoalsByAgentID(ctx context.Context, agentID uint64) (result []*core.Studygoal, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, agentID)
	generateSQL.WriteString("SELECT * FROM studygoals WHERE agent_id = ?; ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
// WHERE agent_id = @agentID AND status = 1;
func (s studygoalDo) GetActiveStudygoalsByAgentID(ctx context.Context, agentID uint64) (result []*core.Studygoal, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, agentID)
	generateSQL.WriteString("SELECT * FROM studygoals WHERE agent_id = ? AND status = 1; ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s studygoalDo) Debug() IStudygoalDo {
	return s.withDO(s.DO.Debug())
}

func (s studygoalDo) WithContext(ctx context.Context) IStudygoalDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s studygoalDo) ReadDB() IStudygoalDo {
	return s.Clauses(dbresolver.Read)
}

func (s studygoalDo) WriteDB() IStudygoalDo {
	return s.Clauses(dbresolver.Write)
}

func (s studygoalDo) Session(config *gorm.Session) IStudygoalDo {
	return s.withDO(s.DO.Session(config))
}

func (s studygoalDo) Clauses(conds ...clause.Expression) IStudygoalDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s studygoalDo) Returning(value interface{}, columns ...string) IStudygoalDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s studygoalDo) Not(conds ...gen.Condition) IStudygoalDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s studygoalDo) Or(conds ...gen.Condition) IStudygoalDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s studygoalDo) Select(conds ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s studygoalDo) Where(conds ...gen.Condition) IStudygoalDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s studygoalDo) Order(conds ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s studygoalDo) Distinct(cols ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s studygoalDo) Omit(cols ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s studygoalDo) Join(table schema.Tabler, on ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s studygoalDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s studygoalDo) RightJoin(table schema.Tabler, on ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s studygoalDo) Group(cols ...field.Expr) IStudygoalDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s studygoalDo) Having(conds ...gen.Condition) IStudygoalDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s studygoalDo) Limit(limit int) IStudygoalDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s studygoalDo) Offset(offset int) IStudygoalDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s studygoalDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStudygoalDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s studygoalDo) Unscoped() IStudygoalDo {
	return s.withDO(s.DO.Unscoped())
}

func (s studygoalDo) Create(values ...*core.Studygoal) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s studygoalDo) CreateInBatches(values []*core.Studygoal, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s studygoalDo) Save(values ...*core.Studygoal) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s studygoalDo) First() (*core.Studygoal, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*core.Studygoal), nil
	}
}

func (s studygoalDo) Take() (*core.Studygoal, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*core.Studygoal), nil
	}
}

func (s studygoalDo) Last() (*core.Studygoal, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*core.Studygoal), nil
	}
}

func (s studygoalDo) Find() ([]*core.Studygoal, error) {
	result, err := s.DO.Find()
	return result.([]*core.Studygoal), err
}

func (s studygoalDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*core.Studygoal, err error) {
	buf := make([]*core.Studygoal, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s studygoalDo) FindInBatches(result *[]*core.Studygoal, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s studygoalDo) Attrs(attrs ...field.AssignExpr) IStudygoalDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s studygoalDo) Assign(attrs ...field.AssignExpr) IStudygoalDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s studygoalDo) Joins(fields ...field.RelationField) IStudygoalDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s studygoalDo) Preload(fields ...field.RelationField) IStudygoalDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s studygoalDo) FirstOrInit() (*core.Studygoal, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*core.Studygoal), nil
	}
}

func (s studygoalDo) FirstOrCreate() (*core.Studygoal, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*core.Studygoal), nil
	}
}

func (s studygoalDo) FindByPage(offset int, limit int) (result []*core.Studygoal, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s studygoalDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s studygoalDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s studygoalDo) Delete(models ...*core.Studygoal) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *studygoalDo) withDO(do gen.Dao) *studygoalDo {
	s.DO = *do.(*gen.DO)
	return s
}
