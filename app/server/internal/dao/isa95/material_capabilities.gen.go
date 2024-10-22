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

func newMaterialCapability(db *gorm.DB, opts ...gen.DOOption) materialCapability {
	_materialCapability := materialCapability{}

	_materialCapability.materialCapabilityDo.UseDB(db, opts...)
	_materialCapability.materialCapabilityDo.UseModel(&isa95.MaterialCapability{})

	tableName := _materialCapability.materialCapabilityDo.TableName()
	_materialCapability.ALL = field.NewAsterisk(tableName)
	_materialCapability.ID = field.NewUint(tableName, "id")
	_materialCapability.CreatedAt = field.NewTime(tableName, "created_at")
	_materialCapability.UpdatedAt = field.NewTime(tableName, "updated_at")
	_materialCapability.DeletedAt = field.NewField(tableName, "deleted_at")
	_materialCapability.MaterialClassID = field.NewUint(tableName, "material_class_id")
	_materialCapability.MaterialDefinitionID = field.NewUint(tableName, "material_definition_id")
	_materialCapability.MaterialLotID = field.NewUint(tableName, "material_lot_id")
	_materialCapability.MaterialSublotID = field.NewUint(tableName, "material_sublot_id")
	_materialCapability.Description = field.NewString(tableName, "description")
	_materialCapability.CapabilityType = field.NewString(tableName, "capability_type")
	_materialCapability.Reason = field.NewString(tableName, "reason")
	_materialCapability.ConfidenceFactor = field.NewString(tableName, "confidence_factor")
	_materialCapability.LevelID = field.NewUint(tableName, "level_id")
	_materialCapability.UseID = field.NewUint(tableName, "use_id")
	_materialCapability.StartTime = field.NewTime(tableName, "start_time")
	_materialCapability.EndTime = field.NewTime(tableName, "end_time")
	_materialCapability.Quantity = field.NewString(tableName, "quantity")
	_materialCapability.MeasurementID = field.NewUint(tableName, "measurement_id")
	_materialCapability.AssemblyTypeID = field.NewUint(tableName, "assembly_type_id")
	_materialCapability.AssemblyRelationshipID = field.NewUint(tableName, "assembly_relationship_id")

	_materialCapability.fillFieldMap()

	return _materialCapability
}

type materialCapability struct {
	materialCapabilityDo

	ALL                    field.Asterisk
	ID                     field.Uint
	CreatedAt              field.Time
	UpdatedAt              field.Time
	DeletedAt              field.Field
	MaterialClassID        field.Uint
	MaterialDefinitionID   field.Uint
	MaterialLotID          field.Uint
	MaterialSublotID       field.Uint
	Description            field.String
	CapabilityType         field.String
	Reason                 field.String
	ConfidenceFactor       field.String
	LevelID                field.Uint
	UseID                  field.Uint
	StartTime              field.Time
	EndTime                field.Time
	Quantity               field.String
	MeasurementID          field.Uint
	AssemblyTypeID         field.Uint
	AssemblyRelationshipID field.Uint

	fieldMap map[string]field.Expr
}

func (m materialCapability) Table(newTableName string) *materialCapability {
	m.materialCapabilityDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialCapability) As(alias string) *materialCapability {
	m.materialCapabilityDo.DO = *(m.materialCapabilityDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialCapability) updateTableName(table string) *materialCapability {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.MaterialClassID = field.NewUint(table, "material_class_id")
	m.MaterialDefinitionID = field.NewUint(table, "material_definition_id")
	m.MaterialLotID = field.NewUint(table, "material_lot_id")
	m.MaterialSublotID = field.NewUint(table, "material_sublot_id")
	m.Description = field.NewString(table, "description")
	m.CapabilityType = field.NewString(table, "capability_type")
	m.Reason = field.NewString(table, "reason")
	m.ConfidenceFactor = field.NewString(table, "confidence_factor")
	m.LevelID = field.NewUint(table, "level_id")
	m.UseID = field.NewUint(table, "use_id")
	m.StartTime = field.NewTime(table, "start_time")
	m.EndTime = field.NewTime(table, "end_time")
	m.Quantity = field.NewString(table, "quantity")
	m.MeasurementID = field.NewUint(table, "measurement_id")
	m.AssemblyTypeID = field.NewUint(table, "assembly_type_id")
	m.AssemblyRelationshipID = field.NewUint(table, "assembly_relationship_id")

	m.fillFieldMap()

	return m
}

func (m *materialCapability) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialCapability) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 20)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["material_class_id"] = m.MaterialClassID
	m.fieldMap["material_definition_id"] = m.MaterialDefinitionID
	m.fieldMap["material_lot_id"] = m.MaterialLotID
	m.fieldMap["material_sublot_id"] = m.MaterialSublotID
	m.fieldMap["description"] = m.Description
	m.fieldMap["capability_type"] = m.CapabilityType
	m.fieldMap["reason"] = m.Reason
	m.fieldMap["confidence_factor"] = m.ConfidenceFactor
	m.fieldMap["level_id"] = m.LevelID
	m.fieldMap["use_id"] = m.UseID
	m.fieldMap["start_time"] = m.StartTime
	m.fieldMap["end_time"] = m.EndTime
	m.fieldMap["quantity"] = m.Quantity
	m.fieldMap["measurement_id"] = m.MeasurementID
	m.fieldMap["assembly_type_id"] = m.AssemblyTypeID
	m.fieldMap["assembly_relationship_id"] = m.AssemblyRelationshipID
}

func (m materialCapability) clone(db *gorm.DB) materialCapability {
	m.materialCapabilityDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialCapability) replaceDB(db *gorm.DB) materialCapability {
	m.materialCapabilityDo.ReplaceDB(db)
	return m
}

type materialCapabilityDo struct{ gen.DO }

func (m materialCapabilityDo) Debug() *materialCapabilityDo {
	return m.withDO(m.DO.Debug())
}

func (m materialCapabilityDo) WithContext(ctx context.Context) *materialCapabilityDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialCapabilityDo) ReadDB() *materialCapabilityDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialCapabilityDo) WriteDB() *materialCapabilityDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialCapabilityDo) Session(config *gorm.Session) *materialCapabilityDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialCapabilityDo) Clauses(conds ...clause.Expression) *materialCapabilityDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialCapabilityDo) Returning(value interface{}, columns ...string) *materialCapabilityDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialCapabilityDo) Not(conds ...gen.Condition) *materialCapabilityDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialCapabilityDo) Or(conds ...gen.Condition) *materialCapabilityDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialCapabilityDo) Select(conds ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialCapabilityDo) Where(conds ...gen.Condition) *materialCapabilityDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialCapabilityDo) Order(conds ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialCapabilityDo) Distinct(cols ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialCapabilityDo) Omit(cols ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialCapabilityDo) Join(table schema.Tabler, on ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialCapabilityDo) LeftJoin(table schema.Tabler, on ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialCapabilityDo) RightJoin(table schema.Tabler, on ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialCapabilityDo) Group(cols ...field.Expr) *materialCapabilityDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialCapabilityDo) Having(conds ...gen.Condition) *materialCapabilityDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialCapabilityDo) Limit(limit int) *materialCapabilityDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialCapabilityDo) Offset(offset int) *materialCapabilityDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialCapabilityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *materialCapabilityDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialCapabilityDo) Unscoped() *materialCapabilityDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialCapabilityDo) Create(values ...*isa95.MaterialCapability) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialCapabilityDo) CreateInBatches(values []*isa95.MaterialCapability, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialCapabilityDo) Save(values ...*isa95.MaterialCapability) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialCapabilityDo) First() (*isa95.MaterialCapability, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialCapability), nil
	}
}

func (m materialCapabilityDo) Take() (*isa95.MaterialCapability, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialCapability), nil
	}
}

func (m materialCapabilityDo) Last() (*isa95.MaterialCapability, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialCapability), nil
	}
}

func (m materialCapabilityDo) Find() ([]*isa95.MaterialCapability, error) {
	result, err := m.DO.Find()
	return result.([]*isa95.MaterialCapability), err
}

func (m materialCapabilityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.MaterialCapability, err error) {
	buf := make([]*isa95.MaterialCapability, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialCapabilityDo) FindInBatches(result *[]*isa95.MaterialCapability, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialCapabilityDo) Attrs(attrs ...field.AssignExpr) *materialCapabilityDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialCapabilityDo) Assign(attrs ...field.AssignExpr) *materialCapabilityDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialCapabilityDo) Joins(fields ...field.RelationField) *materialCapabilityDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialCapabilityDo) Preload(fields ...field.RelationField) *materialCapabilityDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialCapabilityDo) FirstOrInit() (*isa95.MaterialCapability, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialCapability), nil
	}
}

func (m materialCapabilityDo) FirstOrCreate() (*isa95.MaterialCapability, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialCapability), nil
	}
}

func (m materialCapabilityDo) FindByPage(offset int, limit int) (result []*isa95.MaterialCapability, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m materialCapabilityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialCapabilityDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialCapabilityDo) Delete(models ...*isa95.MaterialCapability) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialCapabilityDo) withDO(do gen.Dao) *materialCapabilityDo {
	m.DO = *do.(*gen.DO)
	return m
}
