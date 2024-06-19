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

func newMaterialActual(db *gorm.DB, opts ...gen.DOOption) materialActual {
	_materialActual := materialActual{}

	_materialActual.materialActualDo.UseDB(db, opts...)
	_materialActual.materialActualDo.UseModel(&isa95.MaterialActual{})

	tableName := _materialActual.materialActualDo.TableName()
	_materialActual.ALL = field.NewAsterisk(tableName)
	_materialActual.ID = field.NewUint(tableName, "id")
	_materialActual.CreatedAt = field.NewTime(tableName, "created_at")
	_materialActual.UpdatedAt = field.NewTime(tableName, "updated_at")
	_materialActual.DeletedAt = field.NewField(tableName, "deleted_at")
	_materialActual.MaterialClassID = field.NewUint(tableName, "material_class_id")
	_materialActual.MaterialDefinitionID = field.NewUint(tableName, "material_definition_id")
	_materialActual.MaterialLotID = field.NewUint(tableName, "material_lot_id")
	_materialActual.MaterialSublotID = field.NewUint(tableName, "material_sublot_id")
	_materialActual.Description = field.NewString(tableName, "description")
	_materialActual.UseID = field.NewUint(tableName, "use_id")
	_materialActual.LevelID = field.NewUint(tableName, "level_id")
	_materialActual.Quantity = field.NewString(tableName, "quantity")
	_materialActual.MeasurementID = field.NewUint(tableName, "measurement_id")
	_materialActual.AssemblyTypeID = field.NewUint(tableName, "assembly_type_id")
	_materialActual.AssemblyRelationshipID = field.NewUint(tableName, "assembly_relationship_id")

	_materialActual.fillFieldMap()

	return _materialActual
}

type materialActual struct {
	materialActualDo

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
	UseID                  field.Uint
	LevelID                field.Uint
	Quantity               field.String
	MeasurementID          field.Uint
	AssemblyTypeID         field.Uint
	AssemblyRelationshipID field.Uint

	fieldMap map[string]field.Expr
}

func (m materialActual) Table(newTableName string) *materialActual {
	m.materialActualDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialActual) As(alias string) *materialActual {
	m.materialActualDo.DO = *(m.materialActualDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialActual) updateTableName(table string) *materialActual {
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
	m.UseID = field.NewUint(table, "use_id")
	m.LevelID = field.NewUint(table, "level_id")
	m.Quantity = field.NewString(table, "quantity")
	m.MeasurementID = field.NewUint(table, "measurement_id")
	m.AssemblyTypeID = field.NewUint(table, "assembly_type_id")
	m.AssemblyRelationshipID = field.NewUint(table, "assembly_relationship_id")

	m.fillFieldMap()

	return m
}

func (m *materialActual) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialActual) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 15)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["material_class_id"] = m.MaterialClassID
	m.fieldMap["material_definition_id"] = m.MaterialDefinitionID
	m.fieldMap["material_lot_id"] = m.MaterialLotID
	m.fieldMap["material_sublot_id"] = m.MaterialSublotID
	m.fieldMap["description"] = m.Description
	m.fieldMap["use_id"] = m.UseID
	m.fieldMap["level_id"] = m.LevelID
	m.fieldMap["quantity"] = m.Quantity
	m.fieldMap["measurement_id"] = m.MeasurementID
	m.fieldMap["assembly_type_id"] = m.AssemblyTypeID
	m.fieldMap["assembly_relationship_id"] = m.AssemblyRelationshipID
}

func (m materialActual) clone(db *gorm.DB) materialActual {
	m.materialActualDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialActual) replaceDB(db *gorm.DB) materialActual {
	m.materialActualDo.ReplaceDB(db)
	return m
}

type materialActualDo struct{ gen.DO }

func (m materialActualDo) Debug() *materialActualDo {
	return m.withDO(m.DO.Debug())
}

func (m materialActualDo) WithContext(ctx context.Context) *materialActualDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialActualDo) ReadDB() *materialActualDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialActualDo) WriteDB() *materialActualDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialActualDo) Session(config *gorm.Session) *materialActualDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialActualDo) Clauses(conds ...clause.Expression) *materialActualDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialActualDo) Returning(value interface{}, columns ...string) *materialActualDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialActualDo) Not(conds ...gen.Condition) *materialActualDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialActualDo) Or(conds ...gen.Condition) *materialActualDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialActualDo) Select(conds ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialActualDo) Where(conds ...gen.Condition) *materialActualDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialActualDo) Order(conds ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialActualDo) Distinct(cols ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialActualDo) Omit(cols ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialActualDo) Join(table schema.Tabler, on ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialActualDo) LeftJoin(table schema.Tabler, on ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialActualDo) RightJoin(table schema.Tabler, on ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialActualDo) Group(cols ...field.Expr) *materialActualDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialActualDo) Having(conds ...gen.Condition) *materialActualDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialActualDo) Limit(limit int) *materialActualDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialActualDo) Offset(offset int) *materialActualDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialActualDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *materialActualDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialActualDo) Unscoped() *materialActualDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialActualDo) Create(values ...*isa95.MaterialActual) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialActualDo) CreateInBatches(values []*isa95.MaterialActual, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialActualDo) Save(values ...*isa95.MaterialActual) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialActualDo) First() (*isa95.MaterialActual, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialActual), nil
	}
}

func (m materialActualDo) Take() (*isa95.MaterialActual, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialActual), nil
	}
}

func (m materialActualDo) Last() (*isa95.MaterialActual, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialActual), nil
	}
}

func (m materialActualDo) Find() ([]*isa95.MaterialActual, error) {
	result, err := m.DO.Find()
	return result.([]*isa95.MaterialActual), err
}

func (m materialActualDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.MaterialActual, err error) {
	buf := make([]*isa95.MaterialActual, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialActualDo) FindInBatches(result *[]*isa95.MaterialActual, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialActualDo) Attrs(attrs ...field.AssignExpr) *materialActualDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialActualDo) Assign(attrs ...field.AssignExpr) *materialActualDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialActualDo) Joins(fields ...field.RelationField) *materialActualDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialActualDo) Preload(fields ...field.RelationField) *materialActualDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialActualDo) FirstOrInit() (*isa95.MaterialActual, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialActual), nil
	}
}

func (m materialActualDo) FirstOrCreate() (*isa95.MaterialActual, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialActual), nil
	}
}

func (m materialActualDo) FindByPage(offset int, limit int) (result []*isa95.MaterialActual, count int64, err error) {
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

func (m materialActualDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialActualDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialActualDo) Delete(models ...*isa95.MaterialActual) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialActualDo) withDO(do gen.Dao) *materialActualDo {
	m.DO = *do.(*gen.DO)
	return m
}
