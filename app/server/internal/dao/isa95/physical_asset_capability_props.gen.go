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

func newPhysicalAssetCapabilityProp(db *gorm.DB, opts ...gen.DOOption) physicalAssetCapabilityProp {
	_physicalAssetCapabilityProp := physicalAssetCapabilityProp{}

	_physicalAssetCapabilityProp.physicalAssetCapabilityPropDo.UseDB(db, opts...)
	_physicalAssetCapabilityProp.physicalAssetCapabilityPropDo.UseModel(&isa95.PhysicalAssetCapabilityProp{})

	tableName := _physicalAssetCapabilityProp.physicalAssetCapabilityPropDo.TableName()
	_physicalAssetCapabilityProp.ALL = field.NewAsterisk(tableName)
	_physicalAssetCapabilityProp.ID = field.NewUint(tableName, "id")
	_physicalAssetCapabilityProp.CreatedAt = field.NewTime(tableName, "created_at")
	_physicalAssetCapabilityProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_physicalAssetCapabilityProp.DeletedAt = field.NewField(tableName, "deleted_at")

	_physicalAssetCapabilityProp.fillFieldMap()

	return _physicalAssetCapabilityProp
}

type physicalAssetCapabilityProp struct {
	physicalAssetCapabilityPropDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (p physicalAssetCapabilityProp) Table(newTableName string) *physicalAssetCapabilityProp {
	p.physicalAssetCapabilityPropDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p physicalAssetCapabilityProp) As(alias string) *physicalAssetCapabilityProp {
	p.physicalAssetCapabilityPropDo.DO = *(p.physicalAssetCapabilityPropDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *physicalAssetCapabilityProp) updateTableName(table string) *physicalAssetCapabilityProp {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *physicalAssetCapabilityProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *physicalAssetCapabilityProp) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p physicalAssetCapabilityProp) clone(db *gorm.DB) physicalAssetCapabilityProp {
	p.physicalAssetCapabilityPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p physicalAssetCapabilityProp) replaceDB(db *gorm.DB) physicalAssetCapabilityProp {
	p.physicalAssetCapabilityPropDo.ReplaceDB(db)
	return p
}

type physicalAssetCapabilityPropDo struct{ gen.DO }

func (p physicalAssetCapabilityPropDo) Debug() *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Debug())
}

func (p physicalAssetCapabilityPropDo) WithContext(ctx context.Context) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p physicalAssetCapabilityPropDo) ReadDB() *physicalAssetCapabilityPropDo {
	return p.Clauses(dbresolver.Read)
}

func (p physicalAssetCapabilityPropDo) WriteDB() *physicalAssetCapabilityPropDo {
	return p.Clauses(dbresolver.Write)
}

func (p physicalAssetCapabilityPropDo) Session(config *gorm.Session) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Session(config))
}

func (p physicalAssetCapabilityPropDo) Clauses(conds ...clause.Expression) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p physicalAssetCapabilityPropDo) Returning(value interface{}, columns ...string) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p physicalAssetCapabilityPropDo) Not(conds ...gen.Condition) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p physicalAssetCapabilityPropDo) Or(conds ...gen.Condition) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p physicalAssetCapabilityPropDo) Select(conds ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p physicalAssetCapabilityPropDo) Where(conds ...gen.Condition) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p physicalAssetCapabilityPropDo) Order(conds ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p physicalAssetCapabilityPropDo) Distinct(cols ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p physicalAssetCapabilityPropDo) Omit(cols ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p physicalAssetCapabilityPropDo) Join(table schema.Tabler, on ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p physicalAssetCapabilityPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p physicalAssetCapabilityPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p physicalAssetCapabilityPropDo) Group(cols ...field.Expr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p physicalAssetCapabilityPropDo) Having(conds ...gen.Condition) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p physicalAssetCapabilityPropDo) Limit(limit int) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p physicalAssetCapabilityPropDo) Offset(offset int) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p physicalAssetCapabilityPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p physicalAssetCapabilityPropDo) Unscoped() *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Unscoped())
}

func (p physicalAssetCapabilityPropDo) Create(values ...*isa95.PhysicalAssetCapabilityProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p physicalAssetCapabilityPropDo) CreateInBatches(values []*isa95.PhysicalAssetCapabilityProp, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p physicalAssetCapabilityPropDo) Save(values ...*isa95.PhysicalAssetCapabilityProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p physicalAssetCapabilityPropDo) First() (*isa95.PhysicalAssetCapabilityProp, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetCapabilityProp), nil
	}
}

func (p physicalAssetCapabilityPropDo) Take() (*isa95.PhysicalAssetCapabilityProp, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetCapabilityProp), nil
	}
}

func (p physicalAssetCapabilityPropDo) Last() (*isa95.PhysicalAssetCapabilityProp, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetCapabilityProp), nil
	}
}

func (p physicalAssetCapabilityPropDo) Find() ([]*isa95.PhysicalAssetCapabilityProp, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.PhysicalAssetCapabilityProp), err
}

func (p physicalAssetCapabilityPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.PhysicalAssetCapabilityProp, err error) {
	buf := make([]*isa95.PhysicalAssetCapabilityProp, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p physicalAssetCapabilityPropDo) FindInBatches(result *[]*isa95.PhysicalAssetCapabilityProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p physicalAssetCapabilityPropDo) Attrs(attrs ...field.AssignExpr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p physicalAssetCapabilityPropDo) Assign(attrs ...field.AssignExpr) *physicalAssetCapabilityPropDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p physicalAssetCapabilityPropDo) Joins(fields ...field.RelationField) *physicalAssetCapabilityPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p physicalAssetCapabilityPropDo) Preload(fields ...field.RelationField) *physicalAssetCapabilityPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p physicalAssetCapabilityPropDo) FirstOrInit() (*isa95.PhysicalAssetCapabilityProp, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetCapabilityProp), nil
	}
}

func (p physicalAssetCapabilityPropDo) FirstOrCreate() (*isa95.PhysicalAssetCapabilityProp, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetCapabilityProp), nil
	}
}

func (p physicalAssetCapabilityPropDo) FindByPage(offset int, limit int) (result []*isa95.PhysicalAssetCapabilityProp, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p physicalAssetCapabilityPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p physicalAssetCapabilityPropDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p physicalAssetCapabilityPropDo) Delete(models ...*isa95.PhysicalAssetCapabilityProp) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *physicalAssetCapabilityPropDo) withDO(do gen.Dao) *physicalAssetCapabilityPropDo {
	p.DO = *do.(*gen.DO)
	return p
}