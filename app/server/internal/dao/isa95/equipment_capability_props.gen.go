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

func newEquipmentCapabilityProp(db *gorm.DB, opts ...gen.DOOption) equipmentCapabilityProp {
	_equipmentCapabilityProp := equipmentCapabilityProp{}

	_equipmentCapabilityProp.equipmentCapabilityPropDo.UseDB(db, opts...)
	_equipmentCapabilityProp.equipmentCapabilityPropDo.UseModel(&isa95.EquipmentCapabilityProp{})

	tableName := _equipmentCapabilityProp.equipmentCapabilityPropDo.TableName()
	_equipmentCapabilityProp.ALL = field.NewAsterisk(tableName)
	_equipmentCapabilityProp.ID = field.NewUint(tableName, "id")
	_equipmentCapabilityProp.CreatedAt = field.NewTime(tableName, "created_at")
	_equipmentCapabilityProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_equipmentCapabilityProp.DeletedAt = field.NewField(tableName, "deleted_at")

	_equipmentCapabilityProp.fillFieldMap()

	return _equipmentCapabilityProp
}

type equipmentCapabilityProp struct {
	equipmentCapabilityPropDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (e equipmentCapabilityProp) Table(newTableName string) *equipmentCapabilityProp {
	e.equipmentCapabilityPropDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e equipmentCapabilityProp) As(alias string) *equipmentCapabilityProp {
	e.equipmentCapabilityPropDo.DO = *(e.equipmentCapabilityPropDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *equipmentCapabilityProp) updateTableName(table string) *equipmentCapabilityProp {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewUint(table, "id")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")
	e.DeletedAt = field.NewField(table, "deleted_at")

	e.fillFieldMap()

	return e
}

func (e *equipmentCapabilityProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *equipmentCapabilityProp) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 4)
	e.fieldMap["id"] = e.ID
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["deleted_at"] = e.DeletedAt
}

func (e equipmentCapabilityProp) clone(db *gorm.DB) equipmentCapabilityProp {
	e.equipmentCapabilityPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e equipmentCapabilityProp) replaceDB(db *gorm.DB) equipmentCapabilityProp {
	e.equipmentCapabilityPropDo.ReplaceDB(db)
	return e
}

type equipmentCapabilityPropDo struct{ gen.DO }

func (e equipmentCapabilityPropDo) Debug() *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Debug())
}

func (e equipmentCapabilityPropDo) WithContext(ctx context.Context) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e equipmentCapabilityPropDo) ReadDB() *equipmentCapabilityPropDo {
	return e.Clauses(dbresolver.Read)
}

func (e equipmentCapabilityPropDo) WriteDB() *equipmentCapabilityPropDo {
	return e.Clauses(dbresolver.Write)
}

func (e equipmentCapabilityPropDo) Session(config *gorm.Session) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Session(config))
}

func (e equipmentCapabilityPropDo) Clauses(conds ...clause.Expression) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e equipmentCapabilityPropDo) Returning(value interface{}, columns ...string) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e equipmentCapabilityPropDo) Not(conds ...gen.Condition) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e equipmentCapabilityPropDo) Or(conds ...gen.Condition) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e equipmentCapabilityPropDo) Select(conds ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e equipmentCapabilityPropDo) Where(conds ...gen.Condition) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e equipmentCapabilityPropDo) Order(conds ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e equipmentCapabilityPropDo) Distinct(cols ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e equipmentCapabilityPropDo) Omit(cols ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e equipmentCapabilityPropDo) Join(table schema.Tabler, on ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e equipmentCapabilityPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e equipmentCapabilityPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e equipmentCapabilityPropDo) Group(cols ...field.Expr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e equipmentCapabilityPropDo) Having(conds ...gen.Condition) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e equipmentCapabilityPropDo) Limit(limit int) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e equipmentCapabilityPropDo) Offset(offset int) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e equipmentCapabilityPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e equipmentCapabilityPropDo) Unscoped() *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Unscoped())
}

func (e equipmentCapabilityPropDo) Create(values ...*isa95.EquipmentCapabilityProp) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e equipmentCapabilityPropDo) CreateInBatches(values []*isa95.EquipmentCapabilityProp, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e equipmentCapabilityPropDo) Save(values ...*isa95.EquipmentCapabilityProp) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e equipmentCapabilityPropDo) First() (*isa95.EquipmentCapabilityProp, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentCapabilityProp), nil
	}
}

func (e equipmentCapabilityPropDo) Take() (*isa95.EquipmentCapabilityProp, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentCapabilityProp), nil
	}
}

func (e equipmentCapabilityPropDo) Last() (*isa95.EquipmentCapabilityProp, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentCapabilityProp), nil
	}
}

func (e equipmentCapabilityPropDo) Find() ([]*isa95.EquipmentCapabilityProp, error) {
	result, err := e.DO.Find()
	return result.([]*isa95.EquipmentCapabilityProp), err
}

func (e equipmentCapabilityPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.EquipmentCapabilityProp, err error) {
	buf := make([]*isa95.EquipmentCapabilityProp, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e equipmentCapabilityPropDo) FindInBatches(result *[]*isa95.EquipmentCapabilityProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e equipmentCapabilityPropDo) Attrs(attrs ...field.AssignExpr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e equipmentCapabilityPropDo) Assign(attrs ...field.AssignExpr) *equipmentCapabilityPropDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e equipmentCapabilityPropDo) Joins(fields ...field.RelationField) *equipmentCapabilityPropDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e equipmentCapabilityPropDo) Preload(fields ...field.RelationField) *equipmentCapabilityPropDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e equipmentCapabilityPropDo) FirstOrInit() (*isa95.EquipmentCapabilityProp, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentCapabilityProp), nil
	}
}

func (e equipmentCapabilityPropDo) FirstOrCreate() (*isa95.EquipmentCapabilityProp, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentCapabilityProp), nil
	}
}

func (e equipmentCapabilityPropDo) FindByPage(offset int, limit int) (result []*isa95.EquipmentCapabilityProp, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e equipmentCapabilityPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e equipmentCapabilityPropDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e equipmentCapabilityPropDo) Delete(models ...*isa95.EquipmentCapabilityProp) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *equipmentCapabilityPropDo) withDO(do gen.Dao) *equipmentCapabilityPropDo {
	e.DO = *do.(*gen.DO)
	return e
}
