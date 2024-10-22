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

func newPhysicalAssetSpecificationProp(db *gorm.DB, opts ...gen.DOOption) physicalAssetSpecificationProp {
	_physicalAssetSpecificationProp := physicalAssetSpecificationProp{}

	_physicalAssetSpecificationProp.physicalAssetSpecificationPropDo.UseDB(db, opts...)
	_physicalAssetSpecificationProp.physicalAssetSpecificationPropDo.UseModel(&isa95.PhysicalAssetSpecificationProp{})

	tableName := _physicalAssetSpecificationProp.physicalAssetSpecificationPropDo.TableName()
	_physicalAssetSpecificationProp.ALL = field.NewAsterisk(tableName)
	_physicalAssetSpecificationProp.ID = field.NewUint(tableName, "id")
	_physicalAssetSpecificationProp.CreatedAt = field.NewTime(tableName, "created_at")
	_physicalAssetSpecificationProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_physicalAssetSpecificationProp.DeletedAt = field.NewField(tableName, "deleted_at")

	_physicalAssetSpecificationProp.fillFieldMap()

	return _physicalAssetSpecificationProp
}

type physicalAssetSpecificationProp struct {
	physicalAssetSpecificationPropDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (p physicalAssetSpecificationProp) Table(newTableName string) *physicalAssetSpecificationProp {
	p.physicalAssetSpecificationPropDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p physicalAssetSpecificationProp) As(alias string) *physicalAssetSpecificationProp {
	p.physicalAssetSpecificationPropDo.DO = *(p.physicalAssetSpecificationPropDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *physicalAssetSpecificationProp) updateTableName(table string) *physicalAssetSpecificationProp {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *physicalAssetSpecificationProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *physicalAssetSpecificationProp) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p physicalAssetSpecificationProp) clone(db *gorm.DB) physicalAssetSpecificationProp {
	p.physicalAssetSpecificationPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p physicalAssetSpecificationProp) replaceDB(db *gorm.DB) physicalAssetSpecificationProp {
	p.physicalAssetSpecificationPropDo.ReplaceDB(db)
	return p
}

type physicalAssetSpecificationPropDo struct{ gen.DO }

func (p physicalAssetSpecificationPropDo) Debug() *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Debug())
}

func (p physicalAssetSpecificationPropDo) WithContext(ctx context.Context) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p physicalAssetSpecificationPropDo) ReadDB() *physicalAssetSpecificationPropDo {
	return p.Clauses(dbresolver.Read)
}

func (p physicalAssetSpecificationPropDo) WriteDB() *physicalAssetSpecificationPropDo {
	return p.Clauses(dbresolver.Write)
}

func (p physicalAssetSpecificationPropDo) Session(config *gorm.Session) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Session(config))
}

func (p physicalAssetSpecificationPropDo) Clauses(conds ...clause.Expression) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p physicalAssetSpecificationPropDo) Returning(value interface{}, columns ...string) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p physicalAssetSpecificationPropDo) Not(conds ...gen.Condition) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p physicalAssetSpecificationPropDo) Or(conds ...gen.Condition) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p physicalAssetSpecificationPropDo) Select(conds ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p physicalAssetSpecificationPropDo) Where(conds ...gen.Condition) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p physicalAssetSpecificationPropDo) Order(conds ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p physicalAssetSpecificationPropDo) Distinct(cols ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p physicalAssetSpecificationPropDo) Omit(cols ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p physicalAssetSpecificationPropDo) Join(table schema.Tabler, on ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p physicalAssetSpecificationPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p physicalAssetSpecificationPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p physicalAssetSpecificationPropDo) Group(cols ...field.Expr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p physicalAssetSpecificationPropDo) Having(conds ...gen.Condition) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p physicalAssetSpecificationPropDo) Limit(limit int) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p physicalAssetSpecificationPropDo) Offset(offset int) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p physicalAssetSpecificationPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p physicalAssetSpecificationPropDo) Unscoped() *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Unscoped())
}

func (p physicalAssetSpecificationPropDo) Create(values ...*isa95.PhysicalAssetSpecificationProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p physicalAssetSpecificationPropDo) CreateInBatches(values []*isa95.PhysicalAssetSpecificationProp, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p physicalAssetSpecificationPropDo) Save(values ...*isa95.PhysicalAssetSpecificationProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p physicalAssetSpecificationPropDo) First() (*isa95.PhysicalAssetSpecificationProp, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetSpecificationProp), nil
	}
}

func (p physicalAssetSpecificationPropDo) Take() (*isa95.PhysicalAssetSpecificationProp, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetSpecificationProp), nil
	}
}

func (p physicalAssetSpecificationPropDo) Last() (*isa95.PhysicalAssetSpecificationProp, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetSpecificationProp), nil
	}
}

func (p physicalAssetSpecificationPropDo) Find() ([]*isa95.PhysicalAssetSpecificationProp, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.PhysicalAssetSpecificationProp), err
}

func (p physicalAssetSpecificationPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.PhysicalAssetSpecificationProp, err error) {
	buf := make([]*isa95.PhysicalAssetSpecificationProp, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p physicalAssetSpecificationPropDo) FindInBatches(result *[]*isa95.PhysicalAssetSpecificationProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p physicalAssetSpecificationPropDo) Attrs(attrs ...field.AssignExpr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p physicalAssetSpecificationPropDo) Assign(attrs ...field.AssignExpr) *physicalAssetSpecificationPropDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p physicalAssetSpecificationPropDo) Joins(fields ...field.RelationField) *physicalAssetSpecificationPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p physicalAssetSpecificationPropDo) Preload(fields ...field.RelationField) *physicalAssetSpecificationPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p physicalAssetSpecificationPropDo) FirstOrInit() (*isa95.PhysicalAssetSpecificationProp, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetSpecificationProp), nil
	}
}

func (p physicalAssetSpecificationPropDo) FirstOrCreate() (*isa95.PhysicalAssetSpecificationProp, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAssetSpecificationProp), nil
	}
}

func (p physicalAssetSpecificationPropDo) FindByPage(offset int, limit int) (result []*isa95.PhysicalAssetSpecificationProp, count int64, err error) {
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

func (p physicalAssetSpecificationPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p physicalAssetSpecificationPropDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p physicalAssetSpecificationPropDo) Delete(models ...*isa95.PhysicalAssetSpecificationProp) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *physicalAssetSpecificationPropDo) withDO(do gen.Dao) *physicalAssetSpecificationPropDo {
	p.DO = *do.(*gen.DO)
	return p
}
