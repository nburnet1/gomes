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

func newOperationSegment(db *gorm.DB, opts ...gen.DOOption) operationSegment {
	_operationSegment := operationSegment{}

	_operationSegment.operationSegmentDo.UseDB(db, opts...)
	_operationSegment.operationSegmentDo.UseModel(&isa95.OperationSegment{})

	tableName := _operationSegment.operationSegmentDo.TableName()
	_operationSegment.ALL = field.NewAsterisk(tableName)
	_operationSegment.ID = field.NewUint(tableName, "id")
	_operationSegment.CreatedAt = field.NewTime(tableName, "created_at")
	_operationSegment.UpdatedAt = field.NewTime(tableName, "updated_at")
	_operationSegment.DeletedAt = field.NewField(tableName, "deleted_at")
	_operationSegment.Code = field.NewString(tableName, "code")
	_operationSegment.Description = field.NewString(tableName, "description")
	_operationSegment.LevelID = field.NewUint(tableName, "level_id")
	_operationSegment.Duration = field.NewString(tableName, "duration")
	_operationSegment.MeasurementID = field.NewUint(tableName, "measurement_id")
	_operationSegment.ProcessSegmentID = field.NewUint(tableName, "process_segment_id")
	_operationSegment.OperationTypeID = field.NewUint(tableName, "operation_type_id")
	_operationSegment.WorkDefinitionID = field.NewUint(tableName, "work_definition_id")
	_operationSegment.Measurement = operationSegmentBelongsToMeasurement{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Measurement", "isa95.Measurement"),
	}

	_operationSegment.fillFieldMap()

	return _operationSegment
}

type operationSegment struct {
	operationSegmentDo

	ALL              field.Asterisk
	ID               field.Uint
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Code             field.String
	Description      field.String
	LevelID          field.Uint
	Duration         field.String
	MeasurementID    field.Uint
	ProcessSegmentID field.Uint
	OperationTypeID  field.Uint
	WorkDefinitionID field.Uint
	Measurement      operationSegmentBelongsToMeasurement

	fieldMap map[string]field.Expr
}

func (o operationSegment) Table(newTableName string) *operationSegment {
	o.operationSegmentDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o operationSegment) As(alias string) *operationSegment {
	o.operationSegmentDo.DO = *(o.operationSegmentDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *operationSegment) updateTableName(table string) *operationSegment {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewUint(table, "id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.Code = field.NewString(table, "code")
	o.Description = field.NewString(table, "description")
	o.LevelID = field.NewUint(table, "level_id")
	o.Duration = field.NewString(table, "duration")
	o.MeasurementID = field.NewUint(table, "measurement_id")
	o.ProcessSegmentID = field.NewUint(table, "process_segment_id")
	o.OperationTypeID = field.NewUint(table, "operation_type_id")
	o.WorkDefinitionID = field.NewUint(table, "work_definition_id")

	o.fillFieldMap()

	return o
}

func (o *operationSegment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *operationSegment) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 13)
	o.fieldMap["id"] = o.ID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["code"] = o.Code
	o.fieldMap["description"] = o.Description
	o.fieldMap["level_id"] = o.LevelID
	o.fieldMap["duration"] = o.Duration
	o.fieldMap["measurement_id"] = o.MeasurementID
	o.fieldMap["process_segment_id"] = o.ProcessSegmentID
	o.fieldMap["operation_type_id"] = o.OperationTypeID
	o.fieldMap["work_definition_id"] = o.WorkDefinitionID

}

func (o operationSegment) clone(db *gorm.DB) operationSegment {
	o.operationSegmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o operationSegment) replaceDB(db *gorm.DB) operationSegment {
	o.operationSegmentDo.ReplaceDB(db)
	return o
}

type operationSegmentBelongsToMeasurement struct {
	db *gorm.DB

	field.RelationField
}

func (a operationSegmentBelongsToMeasurement) Where(conds ...field.Expr) *operationSegmentBelongsToMeasurement {
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

func (a operationSegmentBelongsToMeasurement) WithContext(ctx context.Context) *operationSegmentBelongsToMeasurement {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a operationSegmentBelongsToMeasurement) Session(session *gorm.Session) *operationSegmentBelongsToMeasurement {
	a.db = a.db.Session(session)
	return &a
}

func (a operationSegmentBelongsToMeasurement) Model(m *isa95.OperationSegment) *operationSegmentBelongsToMeasurementTx {
	return &operationSegmentBelongsToMeasurementTx{a.db.Model(m).Association(a.Name())}
}

type operationSegmentBelongsToMeasurementTx struct{ tx *gorm.Association }

func (a operationSegmentBelongsToMeasurementTx) Find() (result *isa95.Measurement, err error) {
	return result, a.tx.Find(&result)
}

func (a operationSegmentBelongsToMeasurementTx) Append(values ...*isa95.Measurement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a operationSegmentBelongsToMeasurementTx) Replace(values ...*isa95.Measurement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a operationSegmentBelongsToMeasurementTx) Delete(values ...*isa95.Measurement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a operationSegmentBelongsToMeasurementTx) Clear() error {
	return a.tx.Clear()
}

func (a operationSegmentBelongsToMeasurementTx) Count() int64 {
	return a.tx.Count()
}

type operationSegmentDo struct{ gen.DO }

func (o operationSegmentDo) Debug() *operationSegmentDo {
	return o.withDO(o.DO.Debug())
}

func (o operationSegmentDo) WithContext(ctx context.Context) *operationSegmentDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o operationSegmentDo) ReadDB() *operationSegmentDo {
	return o.Clauses(dbresolver.Read)
}

func (o operationSegmentDo) WriteDB() *operationSegmentDo {
	return o.Clauses(dbresolver.Write)
}

func (o operationSegmentDo) Session(config *gorm.Session) *operationSegmentDo {
	return o.withDO(o.DO.Session(config))
}

func (o operationSegmentDo) Clauses(conds ...clause.Expression) *operationSegmentDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o operationSegmentDo) Returning(value interface{}, columns ...string) *operationSegmentDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o operationSegmentDo) Not(conds ...gen.Condition) *operationSegmentDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o operationSegmentDo) Or(conds ...gen.Condition) *operationSegmentDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o operationSegmentDo) Select(conds ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o operationSegmentDo) Where(conds ...gen.Condition) *operationSegmentDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o operationSegmentDo) Order(conds ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o operationSegmentDo) Distinct(cols ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o operationSegmentDo) Omit(cols ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o operationSegmentDo) Join(table schema.Tabler, on ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o operationSegmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o operationSegmentDo) RightJoin(table schema.Tabler, on ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o operationSegmentDo) Group(cols ...field.Expr) *operationSegmentDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o operationSegmentDo) Having(conds ...gen.Condition) *operationSegmentDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o operationSegmentDo) Limit(limit int) *operationSegmentDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o operationSegmentDo) Offset(offset int) *operationSegmentDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o operationSegmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *operationSegmentDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o operationSegmentDo) Unscoped() *operationSegmentDo {
	return o.withDO(o.DO.Unscoped())
}

func (o operationSegmentDo) Create(values ...*isa95.OperationSegment) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o operationSegmentDo) CreateInBatches(values []*isa95.OperationSegment, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o operationSegmentDo) Save(values ...*isa95.OperationSegment) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o operationSegmentDo) First() (*isa95.OperationSegment, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationSegment), nil
	}
}

func (o operationSegmentDo) Take() (*isa95.OperationSegment, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationSegment), nil
	}
}

func (o operationSegmentDo) Last() (*isa95.OperationSegment, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationSegment), nil
	}
}

func (o operationSegmentDo) Find() ([]*isa95.OperationSegment, error) {
	result, err := o.DO.Find()
	return result.([]*isa95.OperationSegment), err
}

func (o operationSegmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.OperationSegment, err error) {
	buf := make([]*isa95.OperationSegment, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o operationSegmentDo) FindInBatches(result *[]*isa95.OperationSegment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o operationSegmentDo) Attrs(attrs ...field.AssignExpr) *operationSegmentDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o operationSegmentDo) Assign(attrs ...field.AssignExpr) *operationSegmentDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o operationSegmentDo) Joins(fields ...field.RelationField) *operationSegmentDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o operationSegmentDo) Preload(fields ...field.RelationField) *operationSegmentDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o operationSegmentDo) FirstOrInit() (*isa95.OperationSegment, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationSegment), nil
	}
}

func (o operationSegmentDo) FirstOrCreate() (*isa95.OperationSegment, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationSegment), nil
	}
}

func (o operationSegmentDo) FindByPage(offset int, limit int) (result []*isa95.OperationSegment, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o operationSegmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o operationSegmentDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o operationSegmentDo) Delete(models ...*isa95.OperationSegment) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *operationSegmentDo) withDO(do gen.Dao) *operationSegmentDo {
	o.DO = *do.(*gen.DO)
	return o
}
