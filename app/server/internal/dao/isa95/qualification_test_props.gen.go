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

func newQualificationTestProp(db *gorm.DB, opts ...gen.DOOption) qualificationTestProp {
	_qualificationTestProp := qualificationTestProp{}

	_qualificationTestProp.qualificationTestPropDo.UseDB(db, opts...)
	_qualificationTestProp.qualificationTestPropDo.UseModel(&isa95.QualificationTestProp{})

	tableName := _qualificationTestProp.qualificationTestPropDo.TableName()
	_qualificationTestProp.ALL = field.NewAsterisk(tableName)
	_qualificationTestProp.ID = field.NewUint(tableName, "id")
	_qualificationTestProp.CreatedAt = field.NewTime(tableName, "created_at")
	_qualificationTestProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_qualificationTestProp.DeletedAt = field.NewField(tableName, "deleted_at")
	_qualificationTestProp.Description = field.NewString(tableName, "description")
	_qualificationTestProp.Date = field.NewTime(tableName, "date")
	_qualificationTestProp.Result = field.NewBool(tableName, "result")
	_qualificationTestProp.QualificationTestID = field.NewUint(tableName, "qualification_test_id")
	_qualificationTestProp.Expiration = field.NewTime(tableName, "expiration")
	_qualificationTestProp.QualificationTest = qualificationTestPropBelongsToQualificationTest{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("QualificationTest", "isa95.QualificationTest"),
	}

	_qualificationTestProp.fillFieldMap()

	return _qualificationTestProp
}

type qualificationTestProp struct {
	qualificationTestPropDo

	ALL                 field.Asterisk
	ID                  field.Uint
	CreatedAt           field.Time
	UpdatedAt           field.Time
	DeletedAt           field.Field
	Description         field.String
	Date                field.Time
	Result              field.Bool
	QualificationTestID field.Uint
	Expiration          field.Time
	QualificationTest   qualificationTestPropBelongsToQualificationTest

	fieldMap map[string]field.Expr
}

func (q qualificationTestProp) Table(newTableName string) *qualificationTestProp {
	q.qualificationTestPropDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q qualificationTestProp) As(alias string) *qualificationTestProp {
	q.qualificationTestPropDo.DO = *(q.qualificationTestPropDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *qualificationTestProp) updateTableName(table string) *qualificationTestProp {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewUint(table, "id")
	q.CreatedAt = field.NewTime(table, "created_at")
	q.UpdatedAt = field.NewTime(table, "updated_at")
	q.DeletedAt = field.NewField(table, "deleted_at")
	q.Description = field.NewString(table, "description")
	q.Date = field.NewTime(table, "date")
	q.Result = field.NewBool(table, "result")
	q.QualificationTestID = field.NewUint(table, "qualification_test_id")
	q.Expiration = field.NewTime(table, "expiration")

	q.fillFieldMap()

	return q
}

func (q *qualificationTestProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *qualificationTestProp) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 10)
	q.fieldMap["id"] = q.ID
	q.fieldMap["created_at"] = q.CreatedAt
	q.fieldMap["updated_at"] = q.UpdatedAt
	q.fieldMap["deleted_at"] = q.DeletedAt
	q.fieldMap["description"] = q.Description
	q.fieldMap["date"] = q.Date
	q.fieldMap["result"] = q.Result
	q.fieldMap["qualification_test_id"] = q.QualificationTestID
	q.fieldMap["expiration"] = q.Expiration

}

func (q qualificationTestProp) clone(db *gorm.DB) qualificationTestProp {
	q.qualificationTestPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q qualificationTestProp) replaceDB(db *gorm.DB) qualificationTestProp {
	q.qualificationTestPropDo.ReplaceDB(db)
	return q
}

type qualificationTestPropBelongsToQualificationTest struct {
	db *gorm.DB

	field.RelationField
}

func (a qualificationTestPropBelongsToQualificationTest) Where(conds ...field.Expr) *qualificationTestPropBelongsToQualificationTest {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a qualificationTestPropBelongsToQualificationTest) WithContext(ctx context.Context) *qualificationTestPropBelongsToQualificationTest {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a qualificationTestPropBelongsToQualificationTest) Session(session *gorm.Session) *qualificationTestPropBelongsToQualificationTest {
	a.db = a.db.Session(session)
	return &a
}

func (a qualificationTestPropBelongsToQualificationTest) Model(m *isa95.QualificationTestProp) *qualificationTestPropBelongsToQualificationTestTx {
	return &qualificationTestPropBelongsToQualificationTestTx{a.db.Model(m).Association(a.Name())}
}

type qualificationTestPropBelongsToQualificationTestTx struct{ tx *gorm.Association }

func (a qualificationTestPropBelongsToQualificationTestTx) Find() (result *isa95.QualificationTest, err error) {
	return result, a.tx.Find(&result)
}

func (a qualificationTestPropBelongsToQualificationTestTx) Append(values ...*isa95.QualificationTest) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a qualificationTestPropBelongsToQualificationTestTx) Replace(values ...*isa95.QualificationTest) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a qualificationTestPropBelongsToQualificationTestTx) Delete(values ...*isa95.QualificationTest) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a qualificationTestPropBelongsToQualificationTestTx) Clear() error {
	return a.tx.Clear()
}

func (a qualificationTestPropBelongsToQualificationTestTx) Count() int64 {
	return a.tx.Count()
}

type qualificationTestPropDo struct{ gen.DO }

func (q qualificationTestPropDo) Debug() *qualificationTestPropDo {
	return q.withDO(q.DO.Debug())
}

func (q qualificationTestPropDo) WithContext(ctx context.Context) *qualificationTestPropDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q qualificationTestPropDo) ReadDB() *qualificationTestPropDo {
	return q.Clauses(dbresolver.Read)
}

func (q qualificationTestPropDo) WriteDB() *qualificationTestPropDo {
	return q.Clauses(dbresolver.Write)
}

func (q qualificationTestPropDo) Session(config *gorm.Session) *qualificationTestPropDo {
	return q.withDO(q.DO.Session(config))
}

func (q qualificationTestPropDo) Clauses(conds ...clause.Expression) *qualificationTestPropDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q qualificationTestPropDo) Returning(value interface{}, columns ...string) *qualificationTestPropDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q qualificationTestPropDo) Not(conds ...gen.Condition) *qualificationTestPropDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q qualificationTestPropDo) Or(conds ...gen.Condition) *qualificationTestPropDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q qualificationTestPropDo) Select(conds ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q qualificationTestPropDo) Where(conds ...gen.Condition) *qualificationTestPropDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q qualificationTestPropDo) Order(conds ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q qualificationTestPropDo) Distinct(cols ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q qualificationTestPropDo) Omit(cols ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q qualificationTestPropDo) Join(table schema.Tabler, on ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q qualificationTestPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q qualificationTestPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q qualificationTestPropDo) Group(cols ...field.Expr) *qualificationTestPropDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q qualificationTestPropDo) Having(conds ...gen.Condition) *qualificationTestPropDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q qualificationTestPropDo) Limit(limit int) *qualificationTestPropDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q qualificationTestPropDo) Offset(offset int) *qualificationTestPropDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q qualificationTestPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *qualificationTestPropDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q qualificationTestPropDo) Unscoped() *qualificationTestPropDo {
	return q.withDO(q.DO.Unscoped())
}

func (q qualificationTestPropDo) Create(values ...*isa95.QualificationTestProp) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q qualificationTestPropDo) CreateInBatches(values []*isa95.QualificationTestProp, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q qualificationTestPropDo) Save(values ...*isa95.QualificationTestProp) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q qualificationTestPropDo) First() (*isa95.QualificationTestProp, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTestProp), nil
	}
}

func (q qualificationTestPropDo) Take() (*isa95.QualificationTestProp, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTestProp), nil
	}
}

func (q qualificationTestPropDo) Last() (*isa95.QualificationTestProp, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTestProp), nil
	}
}

func (q qualificationTestPropDo) Find() ([]*isa95.QualificationTestProp, error) {
	result, err := q.DO.Find()
	return result.([]*isa95.QualificationTestProp), err
}

func (q qualificationTestPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.QualificationTestProp, err error) {
	buf := make([]*isa95.QualificationTestProp, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q qualificationTestPropDo) FindInBatches(result *[]*isa95.QualificationTestProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q qualificationTestPropDo) Attrs(attrs ...field.AssignExpr) *qualificationTestPropDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q qualificationTestPropDo) Assign(attrs ...field.AssignExpr) *qualificationTestPropDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q qualificationTestPropDo) Joins(fields ...field.RelationField) *qualificationTestPropDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q qualificationTestPropDo) Preload(fields ...field.RelationField) *qualificationTestPropDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q qualificationTestPropDo) FirstOrInit() (*isa95.QualificationTestProp, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTestProp), nil
	}
}

func (q qualificationTestPropDo) FirstOrCreate() (*isa95.QualificationTestProp, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.QualificationTestProp), nil
	}
}

func (q qualificationTestPropDo) FindByPage(offset int, limit int) (result []*isa95.QualificationTestProp, count int64, err error) {
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

func (q qualificationTestPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q qualificationTestPropDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q qualificationTestPropDo) Delete(models ...*isa95.QualificationTestProp) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *qualificationTestPropDo) withDO(do gen.Dao) *qualificationTestPropDo {
	q.DO = *do.(*gen.DO)
	return q
}
