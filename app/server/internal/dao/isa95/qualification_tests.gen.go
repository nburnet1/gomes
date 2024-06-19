// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package isa95

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gomes/server/internal/model/isa95"
)

func newQualificationTest(db *gorm.DB, opts ...gen.DOOption) qualificationTest {
	_qualificationTest := qualificationTest{}

	_qualificationTest.qualificationTestDo.UseDB(db, opts...)
	_qualificationTest.qualificationTestDo.UseModel(&isa95.QualificationTest{})

	tableName := _qualificationTest.qualificationTestDo.TableName()
	_qualificationTest.ALL = field.NewAsterisk(tableName)
	_qualificationTest.ID = field.NewUint(tableName, "id")
	_qualificationTest.CreatedAt = field.NewTime(tableName, "created_at")
	_qualificationTest.UpdatedAt = field.NewTime(tableName, "updated_at")
	_qualificationTest.DeletedAt = field.NewField(tableName, "deleted_at")
	_qualificationTest.Description = field.NewString(tableName, "description")
	_qualificationTest.Version = field.NewString(tableName, "version")

	_qualificationTest.fillFieldMap()

	return _qualificationTest
}

type qualificationTest struct {
	qualificationTestDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Description field.String
	Version     field.String

	fieldMap map[string]field.Expr
}

func (q qualificationTest) Table(newTableName string) *qualificationTest {
	q.qualificationTestDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qualificationTest) As(alias string) *qualificationTest {
	q.qualificationTestDo.DO = *(q.qualificationTestDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qualificationTest) updateTableName(table string) *qualificationTest {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewUint(table, "id")
	q.CreatedAt = field.NewTime(table, "created_at")
	q.UpdatedAt = field.NewTime(table, "updated_at")
	q.DeletedAt = field.NewField(table, "deleted_at")
	q.Description = field.NewString(table, "description")
	q.Version = field.NewString(table, "version")

	q.fillFieldMap()

	return q
}

func (q *qualificationTest) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qualificationTest) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 6)
	q.fieldMap["id"] = q.ID
	q.fieldMap["created_at"] = q.CreatedAt
	q.fieldMap["updated_at"] = q.UpdatedAt
	q.fieldMap["deleted_at"] = q.DeletedAt
	q.fieldMap["description"] = q.Description
	q.fieldMap["version"] = q.Version
}

func (q qualificationTest) clone(db *gorm.DB) qualificationTest {
	q.qualificationTestDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qualificationTest) replaceDB(db *gorm.DB) qualificationTest {
	q.qualificationTestDo.ReplaceDB(db)
	return q
}

type qualificationTestDo struct{ gen.DO }

func (q qualificationTestDo) Debug() *qualificationTestDo {
	return q.withDO(q.DO.Debug())
}

func (q qualificationTestDo) WithContext(ctx context.Context) *qualificationTestDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qualificationTestDo) ReadDB() *qualificationTestDo {
	return q.Clauses(dbresolver.Read)
}

func (q qualificationTestDo) WriteDB() *qualificationTestDo {
	return q.Clauses(dbresolver.Write)
}

func (q qualificationTestDo) Session(config *gorm.Session) *qualificationTestDo {
	return q.withDO(q.DO.Session(config))
}

func (q qualificationTestDo) Clauses(conds ...clause.Expression) *qualificationTestDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qualificationTestDo) Returning(value interface{}, columns ...string) *qualificationTestDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qualificationTestDo) Not(conds ...gen.Condition) *qualificationTestDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qualificationTestDo) Or(conds ...gen.Condition) *qualificationTestDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qualificationTestDo) Select(conds ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qualificationTestDo) Where(conds ...gen.Condition) *qualificationTestDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qualificationTestDo) Order(conds ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qualificationTestDo) Distinct(cols ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qualificationTestDo) Omit(cols ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qualificationTestDo) Join(table schema.Tabler, on ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qualificationTestDo) LeftJoin(table schema.Tabler, on ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qualificationTestDo) RightJoin(table schema.Tabler, on ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qualificationTestDo) Group(cols ...field.Expr) *qualificationTestDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qualificationTestDo) Having(conds ...gen.Condition) *qualificationTestDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qualificationTestDo) Limit(limit int) *qualificationTestDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qualificationTestDo) Offset(offset int) *qualificationTestDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qualificationTestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *qualificationTestDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qualificationTestDo) Unscoped() *qualificationTestDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qualificationTestDo) Create(values ...*isa95.QualificationTest) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qualificationTestDo) CreateInBatches(values []*isa95.QualificationTest, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qualificationTestDo) Save(values ...*isa95.QualificationTest) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qualificationTestDo) First() (*isa95.QualificationTest, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTest), nil
	}
}

func (q qualificationTestDo) Take() (*isa95.QualificationTest, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTest), nil
	}
}

func (q qualificationTestDo) Last() (*isa95.QualificationTest, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTest), nil
	}
}

func (q qualificationTestDo) Find() ([]*isa95.QualificationTest, error) {
	result, err := q.DO.Find()
	return result.([]*isa95.QualificationTest), err
}

func (q qualificationTestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.QualificationTest, err error) {
	buf := make([]*isa95.QualificationTest, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qualificationTestDo) FindInBatches(result *[]*isa95.QualificationTest, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qualificationTestDo) Attrs(attrs ...field.AssignExpr) *qualificationTestDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qualificationTestDo) Assign(attrs ...field.AssignExpr) *qualificationTestDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qualificationTestDo) Joins(fields ...field.RelationField) *qualificationTestDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qualificationTestDo) Preload(fields ...field.RelationField) *qualificationTestDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qualificationTestDo) FirstOrInit() (*isa95.QualificationTest, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTest), nil
	}
}

func (q qualificationTestDo) FirstOrCreate() (*isa95.QualificationTest, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTest), nil
	}
}

func (q qualificationTestDo) FindByPage(offset int, limit int) (result []*isa95.QualificationTest, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q qualificationTestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qualificationTestDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qualificationTestDo) Delete(models ...*isa95.QualificationTest) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qualificationTestDo) withDO(do gen.Dao) *qualificationTestDo {
	q.DO = *do.(*gen.DO)
	return q
}