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

func newPersonnelRequirementProp(db *gorm.DB, opts ...gen.DOOption) personnelRequirementProp {
	_personnelRequirementProp := personnelRequirementProp{}

	_personnelRequirementProp.personnelRequirementPropDo.UseDB(db, opts...)
	_personnelRequirementProp.personnelRequirementPropDo.UseModel(&isa95.PersonnelRequirementProp{})

	tableName := _personnelRequirementProp.personnelRequirementPropDo.TableName()
	_personnelRequirementProp.ALL = field.NewAsterisk(tableName)
	_personnelRequirementProp.ID = field.NewUint(tableName, "id")
	_personnelRequirementProp.CreatedAt = field.NewTime(tableName, "created_at")
	_personnelRequirementProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_personnelRequirementProp.DeletedAt = field.NewField(tableName, "deleted_at")

	_personnelRequirementProp.fillFieldMap()

	return _personnelRequirementProp
}

type personnelRequirementProp struct {
	personnelRequirementPropDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (p personnelRequirementProp) Table(newTableName string) *personnelRequirementProp {
	p.personnelRequirementPropDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p personnelRequirementProp) As(alias string) *personnelRequirementProp {
	p.personnelRequirementPropDo.DO = *(p.personnelRequirementPropDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *personnelRequirementProp) updateTableName(table string) *personnelRequirementProp {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *personnelRequirementProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *personnelRequirementProp) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p personnelRequirementProp) clone(db *gorm.DB) personnelRequirementProp {
	p.personnelRequirementPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p personnelRequirementProp) replaceDB(db *gorm.DB) personnelRequirementProp {
	p.personnelRequirementPropDo.ReplaceDB(db)
	return p
}

type personnelRequirementPropDo struct{ gen.DO }

func (p personnelRequirementPropDo) Debug() *personnelRequirementPropDo {
	return p.withDO(p.DO.Debug())
}

func (p personnelRequirementPropDo) WithContext(ctx context.Context) *personnelRequirementPropDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personnelRequirementPropDo) ReadDB() *personnelRequirementPropDo {
	return p.Clauses(dbresolver.Read)
}

func (p personnelRequirementPropDo) WriteDB() *personnelRequirementPropDo {
	return p.Clauses(dbresolver.Write)
}

func (p personnelRequirementPropDo) Session(config *gorm.Session) *personnelRequirementPropDo {
	return p.withDO(p.DO.Session(config))
}

func (p personnelRequirementPropDo) Clauses(conds ...clause.Expression) *personnelRequirementPropDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personnelRequirementPropDo) Returning(value interface{}, columns ...string) *personnelRequirementPropDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personnelRequirementPropDo) Not(conds ...gen.Condition) *personnelRequirementPropDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personnelRequirementPropDo) Or(conds ...gen.Condition) *personnelRequirementPropDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personnelRequirementPropDo) Select(conds ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personnelRequirementPropDo) Where(conds ...gen.Condition) *personnelRequirementPropDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personnelRequirementPropDo) Order(conds ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personnelRequirementPropDo) Distinct(cols ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personnelRequirementPropDo) Omit(cols ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personnelRequirementPropDo) Join(table schema.Tabler, on ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personnelRequirementPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personnelRequirementPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personnelRequirementPropDo) Group(cols ...field.Expr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personnelRequirementPropDo) Having(conds ...gen.Condition) *personnelRequirementPropDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personnelRequirementPropDo) Limit(limit int) *personnelRequirementPropDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personnelRequirementPropDo) Offset(offset int) *personnelRequirementPropDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personnelRequirementPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personnelRequirementPropDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personnelRequirementPropDo) Unscoped() *personnelRequirementPropDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personnelRequirementPropDo) Create(values ...*isa95.PersonnelRequirementProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personnelRequirementPropDo) CreateInBatches(values []*isa95.PersonnelRequirementProp, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personnelRequirementPropDo) Save(values ...*isa95.PersonnelRequirementProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personnelRequirementPropDo) First() (*isa95.PersonnelRequirementProp, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelRequirementProp), nil
	}
}

func (p personnelRequirementPropDo) Take() (*isa95.PersonnelRequirementProp, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelRequirementProp), nil
	}
}

func (p personnelRequirementPropDo) Last() (*isa95.PersonnelRequirementProp, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelRequirementProp), nil
	}
}

func (p personnelRequirementPropDo) Find() ([]*isa95.PersonnelRequirementProp, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.PersonnelRequirementProp), err
}

func (p personnelRequirementPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.PersonnelRequirementProp, err error) {
	buf := make([]*isa95.PersonnelRequirementProp, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personnelRequirementPropDo) FindInBatches(result *[]*isa95.PersonnelRequirementProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personnelRequirementPropDo) Attrs(attrs ...field.AssignExpr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personnelRequirementPropDo) Assign(attrs ...field.AssignExpr) *personnelRequirementPropDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personnelRequirementPropDo) Joins(fields ...field.RelationField) *personnelRequirementPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personnelRequirementPropDo) Preload(fields ...field.RelationField) *personnelRequirementPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personnelRequirementPropDo) FirstOrInit() (*isa95.PersonnelRequirementProp, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelRequirementProp), nil
	}
}

func (p personnelRequirementPropDo) FirstOrCreate() (*isa95.PersonnelRequirementProp, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelRequirementProp), nil
	}
}

func (p personnelRequirementPropDo) FindByPage(offset int, limit int) (result []*isa95.PersonnelRequirementProp, count int64, err error) {
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

func (p personnelRequirementPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personnelRequirementPropDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personnelRequirementPropDo) Delete(models ...*isa95.PersonnelRequirementProp) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personnelRequirementPropDo) withDO(do gen.Dao) *personnelRequirementPropDo {
	p.DO = *do.(*gen.DO)
	return p
}