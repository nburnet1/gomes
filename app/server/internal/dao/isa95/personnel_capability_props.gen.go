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

func newPersonnelCapabilityProp(db *gorm.DB, opts ...gen.DOOption) personnelCapabilityProp {
	_personnelCapabilityProp := personnelCapabilityProp{}

	_personnelCapabilityProp.personnelCapabilityPropDo.UseDB(db, opts...)
	_personnelCapabilityProp.personnelCapabilityPropDo.UseModel(&isa95.PersonnelCapabilityProp{})

	tableName := _personnelCapabilityProp.personnelCapabilityPropDo.TableName()
	_personnelCapabilityProp.ALL = field.NewAsterisk(tableName)
	_personnelCapabilityProp.ID = field.NewUint(tableName, "id")
	_personnelCapabilityProp.CreatedAt = field.NewTime(tableName, "created_at")
	_personnelCapabilityProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_personnelCapabilityProp.DeletedAt = field.NewField(tableName, "deleted_at")

	_personnelCapabilityProp.fillFieldMap()

	return _personnelCapabilityProp
}

type personnelCapabilityProp struct {
	personnelCapabilityPropDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (p personnelCapabilityProp) Table(newTableName string) *personnelCapabilityProp {
	p.personnelCapabilityPropDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p personnelCapabilityProp) As(alias string) *personnelCapabilityProp {
	p.personnelCapabilityPropDo.DO = *(p.personnelCapabilityPropDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *personnelCapabilityProp) updateTableName(table string) *personnelCapabilityProp {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *personnelCapabilityProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *personnelCapabilityProp) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p personnelCapabilityProp) clone(db *gorm.DB) personnelCapabilityProp {
	p.personnelCapabilityPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p personnelCapabilityProp) replaceDB(db *gorm.DB) personnelCapabilityProp {
	p.personnelCapabilityPropDo.ReplaceDB(db)
	return p
}

type personnelCapabilityPropDo struct{ gen.DO }

func (p personnelCapabilityPropDo) Debug() *personnelCapabilityPropDo {
	return p.withDO(p.DO.Debug())
}

func (p personnelCapabilityPropDo) WithContext(ctx context.Context) *personnelCapabilityPropDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personnelCapabilityPropDo) ReadDB() *personnelCapabilityPropDo {
	return p.Clauses(dbresolver.Read)
}

func (p personnelCapabilityPropDo) WriteDB() *personnelCapabilityPropDo {
	return p.Clauses(dbresolver.Write)
}

func (p personnelCapabilityPropDo) Session(config *gorm.Session) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Session(config))
}

func (p personnelCapabilityPropDo) Clauses(conds ...clause.Expression) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personnelCapabilityPropDo) Returning(value interface{}, columns ...string) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personnelCapabilityPropDo) Not(conds ...gen.Condition) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personnelCapabilityPropDo) Or(conds ...gen.Condition) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personnelCapabilityPropDo) Select(conds ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personnelCapabilityPropDo) Where(conds ...gen.Condition) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personnelCapabilityPropDo) Order(conds ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personnelCapabilityPropDo) Distinct(cols ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personnelCapabilityPropDo) Omit(cols ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personnelCapabilityPropDo) Join(table schema.Tabler, on ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personnelCapabilityPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personnelCapabilityPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personnelCapabilityPropDo) Group(cols ...field.Expr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personnelCapabilityPropDo) Having(conds ...gen.Condition) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personnelCapabilityPropDo) Limit(limit int) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personnelCapabilityPropDo) Offset(offset int) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personnelCapabilityPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personnelCapabilityPropDo) Unscoped() *personnelCapabilityPropDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personnelCapabilityPropDo) Create(values ...*isa95.PersonnelCapabilityProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personnelCapabilityPropDo) CreateInBatches(values []*isa95.PersonnelCapabilityProp, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personnelCapabilityPropDo) Save(values ...*isa95.PersonnelCapabilityProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personnelCapabilityPropDo) First() (*isa95.PersonnelCapabilityProp, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelCapabilityProp), nil
	}
}

func (p personnelCapabilityPropDo) Take() (*isa95.PersonnelCapabilityProp, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelCapabilityProp), nil
	}
}

func (p personnelCapabilityPropDo) Last() (*isa95.PersonnelCapabilityProp, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelCapabilityProp), nil
	}
}

func (p personnelCapabilityPropDo) Find() ([]*isa95.PersonnelCapabilityProp, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.PersonnelCapabilityProp), err
}

func (p personnelCapabilityPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.PersonnelCapabilityProp, err error) {
	buf := make([]*isa95.PersonnelCapabilityProp, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personnelCapabilityPropDo) FindInBatches(result *[]*isa95.PersonnelCapabilityProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personnelCapabilityPropDo) Attrs(attrs ...field.AssignExpr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personnelCapabilityPropDo) Assign(attrs ...field.AssignExpr) *personnelCapabilityPropDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personnelCapabilityPropDo) Joins(fields ...field.RelationField) *personnelCapabilityPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personnelCapabilityPropDo) Preload(fields ...field.RelationField) *personnelCapabilityPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personnelCapabilityPropDo) FirstOrInit() (*isa95.PersonnelCapabilityProp, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelCapabilityProp), nil
	}
}

func (p personnelCapabilityPropDo) FirstOrCreate() (*isa95.PersonnelCapabilityProp, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelCapabilityProp), nil
	}
}

func (p personnelCapabilityPropDo) FindByPage(offset int, limit int) (result []*isa95.PersonnelCapabilityProp, count int64, err error) {
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

func (p personnelCapabilityPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personnelCapabilityPropDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personnelCapabilityPropDo) Delete(models ...*isa95.PersonnelCapabilityProp) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personnelCapabilityPropDo) withDO(do gen.Dao) *personnelCapabilityPropDo {
	p.DO = *do.(*gen.DO)
	return p
}