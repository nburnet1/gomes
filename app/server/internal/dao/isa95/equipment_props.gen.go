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

func newEquipmentProp(db *gorm.DB, opts ...gen.DOOption) equipmentProp {
	_equipmentProp := equipmentProp{}

	_equipmentProp.equipmentPropDo.UseDB(db, opts...)
	_equipmentProp.equipmentPropDo.UseModel(&isa95.EquipmentProp{})

	tableName := _equipmentProp.equipmentPropDo.TableName()
	_equipmentProp.ALL = field.NewAsterisk(tableName)
	_equipmentProp.ID = field.NewUint(tableName, "id")
	_equipmentProp.CreatedAt = field.NewTime(tableName, "created_at")
	_equipmentProp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_equipmentProp.DeletedAt = field.NewField(tableName, "deleted_at")
	_equipmentProp.EquipmentID = field.NewUint(tableName, "equipment_id")
	_equipmentProp.Description = field.NewString(tableName, "description")
	_equipmentProp.Value = field.NewString(tableName, "value")
	_equipmentProp.MeasurementID = field.NewUint(tableName, "measurement_id")
	_equipmentProp.Equipment = equipmentPropBelongsToEquipment{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Equipment", "isa95.Equipment"),
		Level: struct {
			field.RelationField
			LevelLookup struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Equipment.Level", "isa95.Level"),
			LevelLookup: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Equipment.Level.LevelLookup", "isa95.LevelLookup"),
			},
		},
	}

	_equipmentProp.Measurement = equipmentPropBelongsToMeasurement{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Measurement", "isa95.Measurement"),
	}

	_equipmentProp.fillFieldMap()

	return _equipmentProp
}

type equipmentProp struct {
	equipmentPropDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	EquipmentID   field.Uint
	Description   field.String
	Value         field.String
	MeasurementID field.Uint
	Equipment     equipmentPropBelongsToEquipment

	Measurement equipmentPropBelongsToMeasurement

	fieldMap map[string]field.Expr
}

func (e equipmentProp) Table(newTableName string) *equipmentProp {
	e.equipmentPropDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e equipmentProp) As(alias string) *equipmentProp {
	e.equipmentPropDo.DO = *(e.equipmentPropDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *equipmentProp) updateTableName(table string) *equipmentProp {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewUint(table, "id")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")
	e.DeletedAt = field.NewField(table, "deleted_at")
	e.EquipmentID = field.NewUint(table, "equipment_id")
	e.Description = field.NewString(table, "description")
	e.Value = field.NewString(table, "value")
	e.MeasurementID = field.NewUint(table, "measurement_id")

	e.fillFieldMap()

	return e
}

func (e *equipmentProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *equipmentProp) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 10)
	e.fieldMap["id"] = e.ID
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["deleted_at"] = e.DeletedAt
	e.fieldMap["equipment_id"] = e.EquipmentID
	e.fieldMap["description"] = e.Description
	e.fieldMap["value"] = e.Value
	e.fieldMap["measurement_id"] = e.MeasurementID

}

func (e equipmentProp) clone(db *gorm.DB) equipmentProp {
	e.equipmentPropDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e equipmentProp) replaceDB(db *gorm.DB) equipmentProp {
	e.equipmentPropDo.ReplaceDB(db)
	return e
}

type equipmentPropBelongsToEquipment struct {
	db *gorm.DB

	field.RelationField

	Level struct {
		field.RelationField
		LevelLookup struct {
			field.RelationField
		}
	}
}

func (a equipmentPropBelongsToEquipment) Where(conds ...field.Expr) *equipmentPropBelongsToEquipment {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a equipmentPropBelongsToEquipment) WithContext(ctx context.Context) *equipmentPropBelongsToEquipment {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a equipmentPropBelongsToEquipment) Session(session *gorm.Session) *equipmentPropBelongsToEquipment {
	a.db = a.db.Session(session)
	return &a
}

func (a equipmentPropBelongsToEquipment) Model(m *isa95.EquipmentProp) *equipmentPropBelongsToEquipmentTx {
	return &equipmentPropBelongsToEquipmentTx{a.db.Model(m).Association(a.Name())}
}

type equipmentPropBelongsToEquipmentTx struct{ tx *gorm.Association }

func (a equipmentPropBelongsToEquipmentTx) Find() (result *isa95.Equipment, err error) {
	return result, a.tx.Find(&result)
}

func (a equipmentPropBelongsToEquipmentTx) Append(values ...*isa95.Equipment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a equipmentPropBelongsToEquipmentTx) Replace(values ...*isa95.Equipment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a equipmentPropBelongsToEquipmentTx) Delete(values ...*isa95.Equipment) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a equipmentPropBelongsToEquipmentTx) Clear() error {
	return a.tx.Clear()
}

func (a equipmentPropBelongsToEquipmentTx) Count() int64 {
	return a.tx.Count()
}

type equipmentPropBelongsToMeasurement struct {
	db *gorm.DB

	field.RelationField
}

func (a equipmentPropBelongsToMeasurement) Where(conds ...field.Expr) *equipmentPropBelongsToMeasurement {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a equipmentPropBelongsToMeasurement) WithContext(ctx context.Context) *equipmentPropBelongsToMeasurement {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a equipmentPropBelongsToMeasurement) Session(session *gorm.Session) *equipmentPropBelongsToMeasurement {
	a.db = a.db.Session(session)
	return &a
}

func (a equipmentPropBelongsToMeasurement) Model(m *isa95.EquipmentProp) *equipmentPropBelongsToMeasurementTx {
	return &equipmentPropBelongsToMeasurementTx{a.db.Model(m).Association(a.Name())}
}

type equipmentPropBelongsToMeasurementTx struct{ tx *gorm.Association }

func (a equipmentPropBelongsToMeasurementTx) Find() (result *isa95.Measurement, err error) {
	return result, a.tx.Find(&result)
}

func (a equipmentPropBelongsToMeasurementTx) Append(values ...*isa95.Measurement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a equipmentPropBelongsToMeasurementTx) Replace(values ...*isa95.Measurement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a equipmentPropBelongsToMeasurementTx) Delete(values ...*isa95.Measurement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a equipmentPropBelongsToMeasurementTx) Clear() error {
	return a.tx.Clear()
}

func (a equipmentPropBelongsToMeasurementTx) Count() int64 {
	return a.tx.Count()
}

type equipmentPropDo struct{ gen.DO }

func (e equipmentPropDo) Debug() *equipmentPropDo {
	return e.withDO(e.DO.Debug())
}

func (e equipmentPropDo) WithContext(ctx context.Context) *equipmentPropDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e equipmentPropDo) ReadDB() *equipmentPropDo {
	return e.Clauses(dbresolver.Read)
}

func (e equipmentPropDo) WriteDB() *equipmentPropDo {
	return e.Clauses(dbresolver.Write)
}

func (e equipmentPropDo) Session(config *gorm.Session) *equipmentPropDo {
	return e.withDO(e.DO.Session(config))
}

func (e equipmentPropDo) Clauses(conds ...clause.Expression) *equipmentPropDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e equipmentPropDo) Returning(value interface{}, columns ...string) *equipmentPropDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e equipmentPropDo) Not(conds ...gen.Condition) *equipmentPropDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e equipmentPropDo) Or(conds ...gen.Condition) *equipmentPropDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e equipmentPropDo) Select(conds ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e equipmentPropDo) Where(conds ...gen.Condition) *equipmentPropDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e equipmentPropDo) Order(conds ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e equipmentPropDo) Distinct(cols ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e equipmentPropDo) Omit(cols ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e equipmentPropDo) Join(table schema.Tabler, on ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e equipmentPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e equipmentPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e equipmentPropDo) Group(cols ...field.Expr) *equipmentPropDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e equipmentPropDo) Having(conds ...gen.Condition) *equipmentPropDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e equipmentPropDo) Limit(limit int) *equipmentPropDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e equipmentPropDo) Offset(offset int) *equipmentPropDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e equipmentPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *equipmentPropDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e equipmentPropDo) Unscoped() *equipmentPropDo {
	return e.withDO(e.DO.Unscoped())
}

func (e equipmentPropDo) Create(values ...*isa95.EquipmentProp) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e equipmentPropDo) CreateInBatches(values []*isa95.EquipmentProp, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e equipmentPropDo) Save(values ...*isa95.EquipmentProp) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e equipmentPropDo) First() (*isa95.EquipmentProp, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentProp), nil
	}
}

func (e equipmentPropDo) Take() (*isa95.EquipmentProp, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentProp), nil
	}
}

func (e equipmentPropDo) Last() (*isa95.EquipmentProp, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentProp), nil
	}
}

func (e equipmentPropDo) Find() ([]*isa95.EquipmentProp, error) {
	result, err := e.DO.Find()
	return result.([]*isa95.EquipmentProp), err
}

func (e equipmentPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.EquipmentProp, err error) {
	buf := make([]*isa95.EquipmentProp, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e equipmentPropDo) FindInBatches(result *[]*isa95.EquipmentProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e equipmentPropDo) Attrs(attrs ...field.AssignExpr) *equipmentPropDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e equipmentPropDo) Assign(attrs ...field.AssignExpr) *equipmentPropDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e equipmentPropDo) Joins(fields ...field.RelationField) *equipmentPropDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e equipmentPropDo) Preload(fields ...field.RelationField) *equipmentPropDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e equipmentPropDo) FirstOrInit() (*isa95.EquipmentProp, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentProp), nil
	}
}

func (e equipmentPropDo) FirstOrCreate() (*isa95.EquipmentProp, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.EquipmentProp), nil
	}
}

func (e equipmentPropDo) FindByPage(offset int, limit int) (result []*isa95.EquipmentProp, count int64, err error) {
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

func (e equipmentPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e equipmentPropDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e equipmentPropDo) Delete(models ...*isa95.EquipmentProp) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *equipmentPropDo) withDO(do gen.Dao) *equipmentPropDo {
	e.DO = *do.(*gen.DO)
	return e
}