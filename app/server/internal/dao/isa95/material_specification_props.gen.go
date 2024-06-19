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

func newMaterialSpecificationProp(db *gorm.DB, opts ...gen.DOOption) materialSpecificationProp {
	_materialSpecificationProp := materialSpecificationProp{}

	_materialSpecificationProp.materialSpecificationPropDo.UseDB(db, opts...)
	_materialSpecificationProp.materialSpecificationPropDo.UseModel(&isa95.MaterialSpecificationProp{})

	tableName := _materialSpecificationProp.materialSpecificationPropDo.TableName()
	_materialSpecificationProp.ALL = field.NewAsterisk(tableName)
	_materialSpecificationProp.ID = field.NewUint(tableName, "id")
	_materialSpecificationProp.CreatedAt = field.NewTime(tableName, "created_at")
	_materialSpecificationProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_materialSpecificationProp.DeletedAt = field.NewField(tableName, "deleted_at")

	_materialSpecificationProp.fillFieldMap()

	return _materialSpecificationProp
}

type materialSpecificationProp struct {
	materialSpecificationPropDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (m materialSpecificationProp) Table(newTableName string) *materialSpecificationProp {
	m.materialSpecificationPropDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialSpecificationProp) As(alias string) *materialSpecificationProp {
	m.materialSpecificationPropDo.DO = *(m.materialSpecificationPropDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialSpecificationProp) updateTableName(table string) *materialSpecificationProp {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *materialSpecificationProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialSpecificationProp) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
}

func (m materialSpecificationProp) clone(db *gorm.DB) materialSpecificationProp {
	m.materialSpecificationPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialSpecificationProp) replaceDB(db *gorm.DB) materialSpecificationProp {
	m.materialSpecificationPropDo.ReplaceDB(db)
	return m
}

type materialSpecificationPropDo struct{ gen.DO }

func (m materialSpecificationPropDo) Debug() *materialSpecificationPropDo {
	return m.withDO(m.DO.Debug())
}

func (m materialSpecificationPropDo) WithContext(ctx context.Context) *materialSpecificationPropDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialSpecificationPropDo) ReadDB() *materialSpecificationPropDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialSpecificationPropDo) WriteDB() *materialSpecificationPropDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialSpecificationPropDo) Session(config *gorm.Session) *materialSpecificationPropDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialSpecificationPropDo) Clauses(conds ...clause.Expression) *materialSpecificationPropDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialSpecificationPropDo) Returning(value interface{}, columns ...string) *materialSpecificationPropDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialSpecificationPropDo) Not(conds ...gen.Condition) *materialSpecificationPropDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialSpecificationPropDo) Or(conds ...gen.Condition) *materialSpecificationPropDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialSpecificationPropDo) Select(conds ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialSpecificationPropDo) Where(conds ...gen.Condition) *materialSpecificationPropDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialSpecificationPropDo) Order(conds ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialSpecificationPropDo) Distinct(cols ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialSpecificationPropDo) Omit(cols ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialSpecificationPropDo) Join(table schema.Tabler, on ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialSpecificationPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialSpecificationPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialSpecificationPropDo) Group(cols ...field.Expr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialSpecificationPropDo) Having(conds ...gen.Condition) *materialSpecificationPropDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialSpecificationPropDo) Limit(limit int) *materialSpecificationPropDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialSpecificationPropDo) Offset(offset int) *materialSpecificationPropDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialSpecificationPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *materialSpecificationPropDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialSpecificationPropDo) Unscoped() *materialSpecificationPropDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialSpecificationPropDo) Create(values ...*isa95.MaterialSpecificationProp) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialSpecificationPropDo) CreateInBatches(values []*isa95.MaterialSpecificationProp, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialSpecificationPropDo) Save(values ...*isa95.MaterialSpecificationProp) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialSpecificationPropDo) First() (*isa95.MaterialSpecificationProp, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialSpecificationProp), nil
	}
}

func (m materialSpecificationPropDo) Take() (*isa95.MaterialSpecificationProp, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialSpecificationProp), nil
	}
}

func (m materialSpecificationPropDo) Last() (*isa95.MaterialSpecificationProp, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialSpecificationProp), nil
	}
}

func (m materialSpecificationPropDo) Find() ([]*isa95.MaterialSpecificationProp, error) {
	result, err := m.DO.Find()
	return result.([]*isa95.MaterialSpecificationProp), err
}

func (m materialSpecificationPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.MaterialSpecificationProp, err error) {
	buf := make([]*isa95.MaterialSpecificationProp, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialSpecificationPropDo) FindInBatches(result *[]*isa95.MaterialSpecificationProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialSpecificationPropDo) Attrs(attrs ...field.AssignExpr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialSpecificationPropDo) Assign(attrs ...field.AssignExpr) *materialSpecificationPropDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialSpecificationPropDo) Joins(fields ...field.RelationField) *materialSpecificationPropDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialSpecificationPropDo) Preload(fields ...field.RelationField) *materialSpecificationPropDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialSpecificationPropDo) FirstOrInit() (*isa95.MaterialSpecificationProp, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialSpecificationProp), nil
	}
}

func (m materialSpecificationPropDo) FirstOrCreate() (*isa95.MaterialSpecificationProp, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialSpecificationProp), nil
	}
}

func (m materialSpecificationPropDo) FindByPage(offset int, limit int) (result []*isa95.MaterialSpecificationProp, count int64, err error) {
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

func (m materialSpecificationPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialSpecificationPropDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialSpecificationPropDo) Delete(models ...*isa95.MaterialSpecificationProp) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialSpecificationPropDo) withDO(do gen.Dao) *materialSpecificationPropDo {
	m.DO = *do.(*gen.DO)
	return m
}
