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

func newMaterialLotProp(db *gorm.DB, opts ...gen.DOOption) materialLotProp {
	_materialLotProp := materialLotProp{}

	_materialLotProp.materialLotPropDo.UseDB(db, opts...)
	_materialLotProp.materialLotPropDo.UseModel(&isa95.MaterialLotProp{})

	tableName := _materialLotProp.materialLotPropDo.TableName()
	_materialLotProp.ALL = field.NewAsterisk(tableName)
	_materialLotProp.ID = field.NewUint(tableName, "id")
	_materialLotProp.CreatedAt = field.NewTime(tableName, "created_at")
	_materialLotProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_materialLotProp.DeletedAt = field.NewField(tableName, "deleted_at")
	_materialLotProp.Code = field.NewString(tableName, "code")
	_materialLotProp.Description = field.NewString(tableName, "description")
	_materialLotProp.Value = field.NewString(tableName, "value")
	_materialLotProp.MeasurementID = field.NewUint(tableName, "measurement_id")

	_materialLotProp.fillFieldMap()

	return _materialLotProp
}

type materialLotProp struct {
	materialLotPropDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	Code          field.String
	Description   field.String
	Value         field.String
	MeasurementID field.Uint

	fieldMap map[string]field.Expr
}

func (m materialLotProp) Table(newTableName string) *materialLotProp {
	m.materialLotPropDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialLotProp) As(alias string) *materialLotProp {
	m.materialLotPropDo.DO = *(m.materialLotPropDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialLotProp) updateTableName(table string) *materialLotProp {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.Code = field.NewString(table, "code")
	m.Description = field.NewString(table, "description")
	m.Value = field.NewString(table, "value")
	m.MeasurementID = field.NewUint(table, "measurement_id")

	m.fillFieldMap()

	return m
}

func (m *materialLotProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialLotProp) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 8)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["code"] = m.Code
	m.fieldMap["description"] = m.Description
	m.fieldMap["value"] = m.Value
	m.fieldMap["measurement_id"] = m.MeasurementID
}

func (m materialLotProp) clone(db *gorm.DB) materialLotProp {
	m.materialLotPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialLotProp) replaceDB(db *gorm.DB) materialLotProp {
	m.materialLotPropDo.ReplaceDB(db)
	return m
}

type materialLotPropDo struct{ gen.DO }

func (m materialLotPropDo) Debug() *materialLotPropDo {
	return m.withDO(m.DO.Debug())
}

func (m materialLotPropDo) WithContext(ctx context.Context) *materialLotPropDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialLotPropDo) ReadDB() *materialLotPropDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialLotPropDo) WriteDB() *materialLotPropDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialLotPropDo) Session(config *gorm.Session) *materialLotPropDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialLotPropDo) Clauses(conds ...clause.Expression) *materialLotPropDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialLotPropDo) Returning(value interface{}, columns ...string) *materialLotPropDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialLotPropDo) Not(conds ...gen.Condition) *materialLotPropDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialLotPropDo) Or(conds ...gen.Condition) *materialLotPropDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialLotPropDo) Select(conds ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialLotPropDo) Where(conds ...gen.Condition) *materialLotPropDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialLotPropDo) Order(conds ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialLotPropDo) Distinct(cols ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialLotPropDo) Omit(cols ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialLotPropDo) Join(table schema.Tabler, on ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialLotPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialLotPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialLotPropDo) Group(cols ...field.Expr) *materialLotPropDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialLotPropDo) Having(conds ...gen.Condition) *materialLotPropDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialLotPropDo) Limit(limit int) *materialLotPropDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialLotPropDo) Offset(offset int) *materialLotPropDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialLotPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *materialLotPropDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialLotPropDo) Unscoped() *materialLotPropDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialLotPropDo) Create(values ...*isa95.MaterialLotProp) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialLotPropDo) CreateInBatches(values []*isa95.MaterialLotProp, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialLotPropDo) Save(values ...*isa95.MaterialLotProp) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialLotPropDo) First() (*isa95.MaterialLotProp, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLotProp), nil
	}
}

func (m materialLotPropDo) Take() (*isa95.MaterialLotProp, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLotProp), nil
	}
}

func (m materialLotPropDo) Last() (*isa95.MaterialLotProp, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLotProp), nil
	}
}

func (m materialLotPropDo) Find() ([]*isa95.MaterialLotProp, error) {
	result, err := m.DO.Find()
	return result.([]*isa95.MaterialLotProp), err
}

func (m materialLotPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.MaterialLotProp, err error) {
	buf := make([]*isa95.MaterialLotProp, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialLotPropDo) FindInBatches(result *[]*isa95.MaterialLotProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialLotPropDo) Attrs(attrs ...field.AssignExpr) *materialLotPropDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialLotPropDo) Assign(attrs ...field.AssignExpr) *materialLotPropDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialLotPropDo) Joins(fields ...field.RelationField) *materialLotPropDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialLotPropDo) Preload(fields ...field.RelationField) *materialLotPropDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialLotPropDo) FirstOrInit() (*isa95.MaterialLotProp, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLotProp), nil
	}
}

func (m materialLotPropDo) FirstOrCreate() (*isa95.MaterialLotProp, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLotProp), nil
	}
}

func (m materialLotPropDo) FindByPage(offset int, limit int) (result []*isa95.MaterialLotProp, count int64, err error) {
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

func (m materialLotPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialLotPropDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialLotPropDo) Delete(models ...*isa95.MaterialLotProp) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialLotPropDo) withDO(do gen.Dao) *materialLotPropDo {
	m.DO = *do.(*gen.DO)
	return m
}