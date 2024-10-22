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

func newPersonnelActualProp(db *gorm.DB, opts ...gen.DOOption) personnelActualProp {
	_personnelActualProp := personnelActualProp{}

	_personnelActualProp.personnelActualPropDo.UseDB(db, opts...)
	_personnelActualProp.personnelActualPropDo.UseModel(&isa95.PersonnelActualProp{})

	tableName := _personnelActualProp.personnelActualPropDo.TableName()
	_personnelActualProp.ALL = field.NewAsterisk(tableName)
	_personnelActualProp.ID = field.NewUint(tableName, "id")
	_personnelActualProp.CreatedAt = field.NewTime(tableName, "created_at")
	_personnelActualProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_personnelActualProp.DeletedAt = field.NewField(tableName, "deleted_at")

	_personnelActualProp.fillFieldMap()

	return _personnelActualProp
}

type personnelActualProp struct {
	personnelActualPropDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (p personnelActualProp) Table(newTableName string) *personnelActualProp {
	p.personnelActualPropDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p personnelActualProp) As(alias string) *personnelActualProp {
	p.personnelActualPropDo.DO = *(p.personnelActualPropDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *personnelActualProp) updateTableName(table string) *personnelActualProp {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *personnelActualProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *personnelActualProp) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p personnelActualProp) clone(db *gorm.DB) personnelActualProp {
	p.personnelActualPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p personnelActualProp) replaceDB(db *gorm.DB) personnelActualProp {
	p.personnelActualPropDo.ReplaceDB(db)
	return p
}

type personnelActualPropDo struct{ gen.DO }

func (p personnelActualPropDo) Debug() *personnelActualPropDo {
	return p.withDO(p.DO.Debug())
}

func (p personnelActualPropDo) WithContext(ctx context.Context) *personnelActualPropDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personnelActualPropDo) ReadDB() *personnelActualPropDo {
	return p.Clauses(dbresolver.Read)
}

func (p personnelActualPropDo) WriteDB() *personnelActualPropDo {
	return p.Clauses(dbresolver.Write)
}

func (p personnelActualPropDo) Session(config *gorm.Session) *personnelActualPropDo {
	return p.withDO(p.DO.Session(config))
}

func (p personnelActualPropDo) Clauses(conds ...clause.Expression) *personnelActualPropDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personnelActualPropDo) Returning(value interface{}, columns ...string) *personnelActualPropDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personnelActualPropDo) Not(conds ...gen.Condition) *personnelActualPropDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personnelActualPropDo) Or(conds ...gen.Condition) *personnelActualPropDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personnelActualPropDo) Select(conds ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personnelActualPropDo) Where(conds ...gen.Condition) *personnelActualPropDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personnelActualPropDo) Order(conds ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personnelActualPropDo) Distinct(cols ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personnelActualPropDo) Omit(cols ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personnelActualPropDo) Join(table schema.Tabler, on ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personnelActualPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personnelActualPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personnelActualPropDo) Group(cols ...field.Expr) *personnelActualPropDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personnelActualPropDo) Having(conds ...gen.Condition) *personnelActualPropDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personnelActualPropDo) Limit(limit int) *personnelActualPropDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personnelActualPropDo) Offset(offset int) *personnelActualPropDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personnelActualPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personnelActualPropDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personnelActualPropDo) Unscoped() *personnelActualPropDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personnelActualPropDo) Create(values ...*isa95.PersonnelActualProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personnelActualPropDo) CreateInBatches(values []*isa95.PersonnelActualProp, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personnelActualPropDo) Save(values ...*isa95.PersonnelActualProp) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personnelActualPropDo) First() (*isa95.PersonnelActualProp, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelActualProp), nil
	}
}

func (p personnelActualPropDo) Take() (*isa95.PersonnelActualProp, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelActualProp), nil
	}
}

func (p personnelActualPropDo) Last() (*isa95.PersonnelActualProp, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelActualProp), nil
	}
}

func (p personnelActualPropDo) Find() ([]*isa95.PersonnelActualProp, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.PersonnelActualProp), err
}

func (p personnelActualPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.PersonnelActualProp, err error) {
	buf := make([]*isa95.PersonnelActualProp, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personnelActualPropDo) FindInBatches(result *[]*isa95.PersonnelActualProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personnelActualPropDo) Attrs(attrs ...field.AssignExpr) *personnelActualPropDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personnelActualPropDo) Assign(attrs ...field.AssignExpr) *personnelActualPropDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personnelActualPropDo) Joins(fields ...field.RelationField) *personnelActualPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personnelActualPropDo) Preload(fields ...field.RelationField) *personnelActualPropDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personnelActualPropDo) FirstOrInit() (*isa95.PersonnelActualProp, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelActualProp), nil
	}
}

func (p personnelActualPropDo) FirstOrCreate() (*isa95.PersonnelActualProp, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PersonnelActualProp), nil
	}
}

func (p personnelActualPropDo) FindByPage(offset int, limit int) (result []*isa95.PersonnelActualProp, count int64, err error) {
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

func (p personnelActualPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personnelActualPropDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personnelActualPropDo) Delete(models ...*isa95.PersonnelActualProp) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personnelActualPropDo) withDO(do gen.Dao) *personnelActualPropDo {
	p.DO = *do.(*gen.DO)
	return p
}
