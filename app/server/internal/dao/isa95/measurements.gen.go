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

func newMeasurement(db *gorm.DB, opts ...gen.DOOption) measurement {
	_measurement := measurement{}

	_measurement.measurementDo.UseDB(db, opts...)
	_measurement.measurementDo.UseModel(&isa95.Measurement{})

	tableName := _measurement.measurementDo.TableName()
	_measurement.ALL = field.NewAsterisk(tableName)
	_measurement.ID = field.NewUint(tableName, "id")
	_measurement.CreatedAt = field.NewTime(tableName, "created_at")
	_measurement.UpdatedAt = field.NewTime(tableName, "updated_at")
	_measurement.DeletedAt = field.NewField(tableName, "deleted_at")
	_measurement.Code = field.NewString(tableName, "code")
	_measurement.Name = field.NewString(tableName, "name")
	_measurement.Description = field.NewString(tableName, "description")

	_measurement.fillFieldMap()

	return _measurement
}

type measurement struct {
	measurementDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Code        field.String
	Name        field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (m measurement) Table(newTableName string) *measurement {
	m.measurementDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m measurement) As(alias string) *measurement {
	m.measurementDo.DO = *(m.measurementDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *measurement) updateTableName(table string) *measurement {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.Code = field.NewString(table, "code")
	m.Name = field.NewString(table, "name")
	m.Description = field.NewString(table, "description")

	m.fillFieldMap()

	return m
}

func (m *measurement) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *measurement) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 7)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["code"] = m.Code
	m.fieldMap["name"] = m.Name
	m.fieldMap["description"] = m.Description
}

func (m measurement) clone(db *gorm.DB) measurement {
	m.measurementDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m measurement) replaceDB(db *gorm.DB) measurement {
	m.measurementDo.ReplaceDB(db)
	return m
}

type measurementDo struct{ gen.DO }

func (m measurementDo) Debug() *measurementDo {
	return m.withDO(m.DO.Debug())
}

func (m measurementDo) WithContext(ctx context.Context) *measurementDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m measurementDo) ReadDB() *measurementDo {
	return m.Clauses(dbresolver.Read)
}

func (m measurementDo) WriteDB() *measurementDo {
	return m.Clauses(dbresolver.Write)
}

func (m measurementDo) Session(config *gorm.Session) *measurementDo {
	return m.withDO(m.DO.Session(config))
}

func (m measurementDo) Clauses(conds ...clause.Expression) *measurementDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m measurementDo) Returning(value interface{}, columns ...string) *measurementDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m measurementDo) Not(conds ...gen.Condition) *measurementDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m measurementDo) Or(conds ...gen.Condition) *measurementDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m measurementDo) Select(conds ...field.Expr) *measurementDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m measurementDo) Where(conds ...gen.Condition) *measurementDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m measurementDo) Order(conds ...field.Expr) *measurementDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m measurementDo) Distinct(cols ...field.Expr) *measurementDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m measurementDo) Omit(cols ...field.Expr) *measurementDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m measurementDo) Join(table schema.Tabler, on ...field.Expr) *measurementDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m measurementDo) LeftJoin(table schema.Tabler, on ...field.Expr) *measurementDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m measurementDo) RightJoin(table schema.Tabler, on ...field.Expr) *measurementDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m measurementDo) Group(cols ...field.Expr) *measurementDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m measurementDo) Having(conds ...gen.Condition) *measurementDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m measurementDo) Limit(limit int) *measurementDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m measurementDo) Offset(offset int) *measurementDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m measurementDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *measurementDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m measurementDo) Unscoped() *measurementDo {
	return m.withDO(m.DO.Unscoped())
}

func (m measurementDo) Create(values ...*isa95.Measurement) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m measurementDo) CreateInBatches(values []*isa95.Measurement, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m measurementDo) Save(values ...*isa95.Measurement) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m measurementDo) First() (*isa95.Measurement, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Measurement), nil
	}
}

func (m measurementDo) Take() (*isa95.Measurement, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Measurement), nil
	}
}

func (m measurementDo) Last() (*isa95.Measurement, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Measurement), nil
	}
}

func (m measurementDo) Find() ([]*isa95.Measurement, error) {
	result, err := m.DO.Find()
	return result.([]*isa95.Measurement), err
}

func (m measurementDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.Measurement, err error) {
	buf := make([]*isa95.Measurement, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m measurementDo) FindInBatches(result *[]*isa95.Measurement, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m measurementDo) Attrs(attrs ...field.AssignExpr) *measurementDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m measurementDo) Assign(attrs ...field.AssignExpr) *measurementDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m measurementDo) Joins(fields ...field.RelationField) *measurementDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m measurementDo) Preload(fields ...field.RelationField) *measurementDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m measurementDo) FirstOrInit() (*isa95.Measurement, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Measurement), nil
	}
}

func (m measurementDo) FirstOrCreate() (*isa95.Measurement, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Measurement), nil
	}
}

func (m measurementDo) FindByPage(offset int, limit int) (result []*isa95.Measurement, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m measurementDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m measurementDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m measurementDo) Delete(models ...*isa95.Measurement) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *measurementDo) withDO(do gen.Dao) *measurementDo {
	m.DO = *do.(*gen.DO)
	return m
}
