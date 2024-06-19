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

func newEquipmentRequirementProperty(db *gorm.DB, opts ...gen.DOOption) equipmentRequirementProperty {
	_equipmentRequirementProperty := equipmentRequirementProperty{}

	_equipmentRequirementProperty.equipmentRequirementPropertyDo.UseDB(db, opts...)
	_equipmentRequirementProperty.equipmentRequirementPropertyDo.UseModel(&isa95.EquipmentRequirementProperty{})

	tableName := _equipmentRequirementProperty.equipmentRequirementPropertyDo.TableName()
	_equipmentRequirementProperty.ALL = field.NewAsterisk(tableName)
	_equipmentRequirementProperty.ID = field.NewUint(tableName, "id")
	_equipmentRequirementProperty.CreatedAt = field.NewTime(tableName, "created_at")
	_equipmentRequirementProperty.UpdatedAt = field.NewTime(tableName, "updated_at")
	_equipmentRequirementProperty.DeletedAt = field.NewField(tableName, "deleted_at")
	_equipmentRequirementProperty.Code = field.NewString(tableName, "code")
	_equipmentRequirementProperty.Description = field.NewString(tableName, "description")
	_equipmentRequirementProperty.Value = field.NewString(tableName, "value")
	_equipmentRequirementProperty.ValueMeasurementID = field.NewUint(tableName, "value_measurement_id")
	_equipmentRequirementProperty.Quantity = field.NewString(tableName, "quantity")
	_equipmentRequirementProperty.QuantityMeasurementID = field.NewUint(tableName, "quantity_measurement_id")

	_equipmentRequirementProperty.fillFieldMap()

	return _equipmentRequirementProperty
}

type equipmentRequirementProperty struct {
	equipmentRequirementPropertyDo

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

func (e equipmentRequirementProperty) Table(newTableName string) *equipmentRequirementProperty {
	e.equipmentRequirementPropertyDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e equipmentRequirementProperty) As(alias string) *equipmentRequirementProperty {
	e.equipmentRequirementPropertyDo.DO = *(e.equipmentRequirementPropertyDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *equipmentRequirementProperty) updateTableName(table string) *equipmentRequirementProperty {
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

func (e *equipmentRequirementProperty) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *equipmentRequirementProperty) fillFieldMap() {
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

func (e equipmentRequirementProperty) clone(db *gorm.DB) equipmentRequirementProperty {
	e.equipmentRequirementPropertyDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e equipmentRequirementProperty) replaceDB(db *gorm.DB) equipmentRequirementProperty {
	e.equipmentRequirementPropertyDo.ReplaceDB(db)
	return e
}

type equipmentRequirementPropertyDo struct{ gen.DO }

func (e equipmentRequirementPropertyDo) Debug() *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Debug())
}

func (e equipmentRequirementPropertyDo) WithContext(ctx context.Context) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e equipmentRequirementPropertyDo) ReadDB() *equipmentRequirementPropertyDo {
	return e.Clauses(dbresolver.Read)
}

func (e equipmentRequirementPropertyDo) WriteDB() *equipmentRequirementPropertyDo {
	return e.Clauses(dbresolver.Write)
}

func (e equipmentRequirementPropertyDo) Session(config *gorm.Session) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Session(config))
}

func (e equipmentRequirementPropertyDo) Clauses(conds ...clause.Expression) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e equipmentRequirementPropertyDo) Returning(value interface{}, columns ...string) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e equipmentRequirementPropertyDo) Not(conds ...gen.Condition) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e equipmentRequirementPropertyDo) Or(conds ...gen.Condition) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e equipmentRequirementPropertyDo) Select(conds ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e equipmentRequirementPropertyDo) Where(conds ...gen.Condition) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e equipmentRequirementPropertyDo) Order(conds ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e equipmentRequirementPropertyDo) Distinct(cols ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e equipmentRequirementPropertyDo) Omit(cols ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e equipmentRequirementPropertyDo) Join(table schema.Tabler, on ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e equipmentRequirementPropertyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e equipmentRequirementPropertyDo) RightJoin(table schema.Tabler, on ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e equipmentRequirementPropertyDo) Group(cols ...field.Expr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e equipmentRequirementPropertyDo) Having(conds ...gen.Condition) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e equipmentRequirementPropertyDo) Limit(limit int) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e equipmentRequirementPropertyDo) Offset(offset int) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e equipmentRequirementPropertyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e equipmentRequirementPropertyDo) Unscoped() *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Unscoped())
}

func (e equipmentRequirementPropertyDo) Create(values ...*isa95.EquipmentRequirementProperty) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e equipmentRequirementPropertyDo) CreateInBatches(values []*isa95.EquipmentRequirementProperty, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e equipmentRequirementPropertyDo) Save(values ...*isa95.EquipmentRequirementProperty) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e equipmentRequirementPropertyDo) First() (*isa95.EquipmentRequirementProperty, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentRequirementProperty), nil
	}
}

func (e equipmentRequirementPropertyDo) Take() (*isa95.EquipmentRequirementProperty, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentRequirementProperty), nil
	}
}

func (e equipmentRequirementPropertyDo) Last() (*isa95.EquipmentRequirementProperty, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentRequirementProperty), nil
	}
}

func (e equipmentRequirementPropertyDo) Find() ([]*isa95.EquipmentRequirementProperty, error) {
	result, err := e.DO.Find()
	return result.([]*isa95.EquipmentRequirementProperty), err
}

func (e equipmentRequirementPropertyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.EquipmentRequirementProperty, err error) {
	buf := make([]*isa95.EquipmentRequirementProperty, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e equipmentRequirementPropertyDo) FindInBatches(result *[]*isa95.EquipmentRequirementProperty, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e equipmentRequirementPropertyDo) Attrs(attrs ...field.AssignExpr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e equipmentRequirementPropertyDo) Assign(attrs ...field.AssignExpr) *equipmentRequirementPropertyDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e equipmentRequirementPropertyDo) Joins(fields ...field.RelationField) *equipmentRequirementPropertyDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e equipmentRequirementPropertyDo) Preload(fields ...field.RelationField) *equipmentRequirementPropertyDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e equipmentRequirementPropertyDo) FirstOrInit() (*isa95.EquipmentRequirementProperty, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentRequirementProperty), nil
	}
}

func (e equipmentRequirementPropertyDo) FirstOrCreate() (*isa95.EquipmentRequirementProperty, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentRequirementProperty), nil
	}
}

func (e equipmentRequirementPropertyDo) FindByPage(offset int, limit int) (result []*isa95.EquipmentRequirementProperty, count int64, err error) {
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

func (e equipmentRequirementPropertyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e equipmentRequirementPropertyDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e equipmentRequirementPropertyDo) Delete(models ...*isa95.EquipmentRequirementProperty) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *equipmentRequirementPropertyDo) withDO(do gen.Dao) *equipmentRequirementPropertyDo {
	e.DO = *do.(*gen.DO)
	return e
}
