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

func newEquipmentSegmentProp(db *gorm.DB, opts ...gen.DOOption) equipmentSegmentProp {
	_equipmentSegmentProp := equipmentSegmentProp{}

	_equipmentSegmentProp.equipmentSegmentPropDo.UseDB(db, opts...)
	_equipmentSegmentProp.equipmentSegmentPropDo.UseModel(&isa95.EquipmentSegmentProp{})

	tableName := _equipmentSegmentProp.equipmentSegmentPropDo.TableName()
	_equipmentSegmentProp.ALL = field.NewAsterisk(tableName)
	_equipmentSegmentProp.ID = field.NewUint(tableName, "id")
	_equipmentSegmentProp.CreatedAt = field.NewTime(tableName, "created_at")
	_equipmentSegmentProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_equipmentSegmentProp.DeletedAt = field.NewField(tableName, "deleted_at")
	_equipmentSegmentProp.Code = field.NewString(tableName, "code")
	_equipmentSegmentProp.Description = field.NewString(tableName, "description")
	_equipmentSegmentProp.Value = field.NewString(tableName, "value")
	_equipmentSegmentProp.ValueMeasurementID = field.NewUint(tableName, "value_measurement_id")
	_equipmentSegmentProp.Quantity = field.NewString(tableName, "quantity")
	_equipmentSegmentProp.QuantityMeasurementID = field.NewUint(tableName, "quantity_measurement_id")

	_equipmentSegmentProp.fillFieldMap()

	return _equipmentSegmentProp
}

type equipmentSegmentProp struct {
	equipmentSegmentPropDo

	ALL                   field.Asterisk
	ID                    field.Uint
	CreatedAt             field.Time
	UpdatedAt             field.Time
	DeletedAt             field.Field
	Code                  field.String
	Description           field.String
	Value                 field.String
	ValueMeasurementID    field.Uint
	Quantity              field.String
	QuantityMeasurementID field.Uint

	fieldMap map[string]field.Expr
}

func (e equipmentSegmentProp) Table(newTableName string) *equipmentSegmentProp {
	e.equipmentSegmentPropDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e equipmentSegmentProp) As(alias string) *equipmentSegmentProp {
	e.equipmentSegmentPropDo.DO = *(e.equipmentSegmentPropDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *equipmentSegmentProp) updateTableName(table string) *equipmentSegmentProp {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewUint(table, "id")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")
	e.DeletedAt = field.NewField(table, "deleted_at")
	e.Code = field.NewString(table, "code")
	e.Description = field.NewString(table, "description")
	e.Value = field.NewString(table, "value")
	e.ValueMeasurementID = field.NewUint(table, "value_measurement_id")
	e.Quantity = field.NewString(table, "quantity")
	e.QuantityMeasurementID = field.NewUint(table, "quantity_measurement_id")

	e.fillFieldMap()

	return e
}

func (e *equipmentSegmentProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *equipmentSegmentProp) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 10)
	e.fieldMap["id"] = e.ID
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["deleted_at"] = e.DeletedAt
	e.fieldMap["code"] = e.Code
	e.fieldMap["description"] = e.Description
	e.fieldMap["value"] = e.Value
	e.fieldMap["value_measurement_id"] = e.ValueMeasurementID
	e.fieldMap["quantity"] = e.Quantity
	e.fieldMap["quantity_measurement_id"] = e.QuantityMeasurementID
}

func (e equipmentSegmentProp) clone(db *gorm.DB) equipmentSegmentProp {
	e.equipmentSegmentPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e equipmentSegmentProp) replaceDB(db *gorm.DB) equipmentSegmentProp {
	e.equipmentSegmentPropDo.ReplaceDB(db)
	return e
}

type equipmentSegmentPropDo struct{ gen.DO }

func (e equipmentSegmentPropDo) Debug() *equipmentSegmentPropDo {
	return e.withDO(e.DO.Debug())
}

func (e equipmentSegmentPropDo) WithContext(ctx context.Context) *equipmentSegmentPropDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e equipmentSegmentPropDo) ReadDB() *equipmentSegmentPropDo {
	return e.Clauses(dbresolver.Read)
}

func (e equipmentSegmentPropDo) WriteDB() *equipmentSegmentPropDo {
	return e.Clauses(dbresolver.Write)
}

func (e equipmentSegmentPropDo) Session(config *gorm.Session) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Session(config))
}

func (e equipmentSegmentPropDo) Clauses(conds ...clause.Expression) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e equipmentSegmentPropDo) Returning(value interface{}, columns ...string) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e equipmentSegmentPropDo) Not(conds ...gen.Condition) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e equipmentSegmentPropDo) Or(conds ...gen.Condition) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e equipmentSegmentPropDo) Select(conds ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e equipmentSegmentPropDo) Where(conds ...gen.Condition) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e equipmentSegmentPropDo) Order(conds ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e equipmentSegmentPropDo) Distinct(cols ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e equipmentSegmentPropDo) Omit(cols ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e equipmentSegmentPropDo) Join(table schema.Tabler, on ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e equipmentSegmentPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e equipmentSegmentPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e equipmentSegmentPropDo) Group(cols ...field.Expr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e equipmentSegmentPropDo) Having(conds ...gen.Condition) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e equipmentSegmentPropDo) Limit(limit int) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e equipmentSegmentPropDo) Offset(offset int) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e equipmentSegmentPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e equipmentSegmentPropDo) Unscoped() *equipmentSegmentPropDo {
	return e.withDO(e.DO.Unscoped())
}

func (e equipmentSegmentPropDo) Create(values ...*isa95.EquipmentSegmentProp) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e equipmentSegmentPropDo) CreateInBatches(values []*isa95.EquipmentSegmentProp, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e equipmentSegmentPropDo) Save(values ...*isa95.EquipmentSegmentProp) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e equipmentSegmentPropDo) First() (*isa95.EquipmentSegmentProp, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentSegmentProp), nil
	}
}

func (e equipmentSegmentPropDo) Take() (*isa95.EquipmentSegmentProp, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentSegmentProp), nil
	}
}

func (e equipmentSegmentPropDo) Last() (*isa95.EquipmentSegmentProp, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentSegmentProp), nil
	}
}

func (e equipmentSegmentPropDo) Find() ([]*isa95.EquipmentSegmentProp, error) {
	result, err := e.DO.Find()
	return result.([]*isa95.EquipmentSegmentProp), err
}

func (e equipmentSegmentPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.EquipmentSegmentProp, err error) {
	buf := make([]*isa95.EquipmentSegmentProp, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e equipmentSegmentPropDo) FindInBatches(result *[]*isa95.EquipmentSegmentProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e equipmentSegmentPropDo) Attrs(attrs ...field.AssignExpr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e equipmentSegmentPropDo) Assign(attrs ...field.AssignExpr) *equipmentSegmentPropDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e equipmentSegmentPropDo) Joins(fields ...field.RelationField) *equipmentSegmentPropDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e equipmentSegmentPropDo) Preload(fields ...field.RelationField) *equipmentSegmentPropDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e equipmentSegmentPropDo) FirstOrInit() (*isa95.EquipmentSegmentProp, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentSegmentProp), nil
	}
}

func (e equipmentSegmentPropDo) FirstOrCreate() (*isa95.EquipmentSegmentProp, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentSegmentProp), nil
	}
}

func (e equipmentSegmentPropDo) FindByPage(offset int, limit int) (result []*isa95.EquipmentSegmentProp, count int64, err error) {
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

func (e equipmentSegmentPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e equipmentSegmentPropDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e equipmentSegmentPropDo) Delete(models ...*isa95.EquipmentSegmentProp) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *equipmentSegmentPropDo) withDO(do gen.Dao) *equipmentSegmentPropDo {
	e.DO = *do.(*gen.DO)
	return e
}