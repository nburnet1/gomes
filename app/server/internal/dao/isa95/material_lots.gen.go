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

func newMaterialLot(db *gorm.DB, opts ...gen.DOOption) materialLot {
	_materialLot := materialLot{}

	_materialLot.materialLotDo.UseDB(db, opts...)
	_materialLot.materialLotDo.UseModel(&isa95.MaterialLot{})

	tableName := _materialLot.materialLotDo.TableName()
	_materialLot.ALL = field.NewAsterisk(tableName)
	_materialLot.ID = field.NewUint(tableName, "id")
	_materialLot.CreatedAt = field.NewTime(tableName, "created_at")
	_materialLot.UpdatedAt = field.NewTime(tableName, "updated_at")
	_materialLot.DeletedAt = field.NewField(tableName, "deleted_at")
	_materialLot.Code = field.NewString(tableName, "code")
	_materialLot.Description = field.NewString(tableName, "description")
	_materialLot.AssemblyTypeID = field.NewUint(tableName, "assembly_type_id")
	_materialLot.AssemblyRelationshipID = field.NewUint(tableName, "assembly_relationship_id")
	_materialLot.Quantity = field.NewString(tableName, "quantity")
	_materialLot.MeasurementID = field.NewUint(tableName, "measurement_id")
	_materialLot.LevelID = field.NewUint(tableName, "level_id")
	_materialLot.StatusID = field.NewUint(tableName, "status_id")

	_materialLot.fillFieldMap()

	return _materialLot
}

type materialLot struct {
	materialLotDo

	ALL                    field.Asterisk
	ID                     field.Uint
	CreatedAt              field.Time
	UpdatedAt              field.Time
	DeletedAt              field.Field
	Code                   field.String
	Description            field.String
	AssemblyTypeID         field.Uint
	AssemblyRelationshipID field.Uint
	Quantity               field.String
	MeasurementID          field.Uint
	LevelID                field.Uint
	StatusID               field.Uint

	fieldMap map[string]field.Expr
}

func (m materialLot) Table(newTableName string) *materialLot {
	m.materialLotDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m materialLot) As(alias string) *materialLot {
	m.materialLotDo.DO = *(m.materialLotDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *materialLot) updateTableName(table string) *materialLot {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.Code = field.NewString(table, "code")
	m.Description = field.NewString(table, "description")
	m.AssemblyTypeID = field.NewUint(table, "assembly_type_id")
	m.AssemblyRelationshipID = field.NewUint(table, "assembly_relationship_id")
	m.Quantity = field.NewString(table, "quantity")
	m.MeasurementID = field.NewUint(table, "measurement_id")
	m.LevelID = field.NewUint(table, "level_id")
	m.StatusID = field.NewUint(table, "status_id")

	m.fillFieldMap()

	return m
}

func (m *materialLot) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *materialLot) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 12)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["code"] = m.Code
	m.fieldMap["description"] = m.Description
	m.fieldMap["assembly_type_id"] = m.AssemblyTypeID
	m.fieldMap["assembly_relationship_id"] = m.AssemblyRelationshipID
	m.fieldMap["quantity"] = m.Quantity
	m.fieldMap["measurement_id"] = m.MeasurementID
	m.fieldMap["level_id"] = m.LevelID
	m.fieldMap["status_id"] = m.StatusID
}

func (m materialLot) clone(db *gorm.DB) materialLot {
	m.materialLotDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m materialLot) replaceDB(db *gorm.DB) materialLot {
	m.materialLotDo.ReplaceDB(db)
	return m
}

type materialLotDo struct{ gen.DO }

func (m materialLotDo) Debug() *materialLotDo {
	return m.withDO(m.DO.Debug())
}

func (m materialLotDo) WithContext(ctx context.Context) *materialLotDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m materialLotDo) ReadDB() *materialLotDo {
	return m.Clauses(dbresolver.Read)
}

func (m materialLotDo) WriteDB() *materialLotDo {
	return m.Clauses(dbresolver.Write)
}

func (m materialLotDo) Session(config *gorm.Session) *materialLotDo {
	return m.withDO(m.DO.Session(config))
}

func (m materialLotDo) Clauses(conds ...clause.Expression) *materialLotDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m materialLotDo) Returning(value interface{}, columns ...string) *materialLotDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m materialLotDo) Not(conds ...gen.Condition) *materialLotDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m materialLotDo) Or(conds ...gen.Condition) *materialLotDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m materialLotDo) Select(conds ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m materialLotDo) Where(conds ...gen.Condition) *materialLotDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m materialLotDo) Order(conds ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m materialLotDo) Distinct(cols ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m materialLotDo) Omit(cols ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m materialLotDo) Join(table schema.Tabler, on ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m materialLotDo) LeftJoin(table schema.Tabler, on ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m materialLotDo) RightJoin(table schema.Tabler, on ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m materialLotDo) Group(cols ...field.Expr) *materialLotDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m materialLotDo) Having(conds ...gen.Condition) *materialLotDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m materialLotDo) Limit(limit int) *materialLotDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m materialLotDo) Offset(offset int) *materialLotDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m materialLotDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *materialLotDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m materialLotDo) Unscoped() *materialLotDo {
	return m.withDO(m.DO.Unscoped())
}

func (m materialLotDo) Create(values ...*isa95.MaterialLot) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m materialLotDo) CreateInBatches(values []*isa95.MaterialLot, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m materialLotDo) Save(values ...*isa95.MaterialLot) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m materialLotDo) First() (*isa95.MaterialLot, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLot), nil
	}
}

func (m materialLotDo) Take() (*isa95.MaterialLot, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLot), nil
	}
}

func (m materialLotDo) Last() (*isa95.MaterialLot, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLot), nil
	}
}

func (m materialLotDo) Find() ([]*isa95.MaterialLot, error) {
	result, err := m.DO.Find()
	return result.([]*isa95.MaterialLot), err
}

func (m materialLotDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.MaterialLot, err error) {
	buf := make([]*isa95.MaterialLot, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m materialLotDo) FindInBatches(result *[]*isa95.MaterialLot, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m materialLotDo) Attrs(attrs ...field.AssignExpr) *materialLotDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m materialLotDo) Assign(attrs ...field.AssignExpr) *materialLotDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m materialLotDo) Joins(fields ...field.RelationField) *materialLotDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m materialLotDo) Preload(fields ...field.RelationField) *materialLotDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m materialLotDo) FirstOrInit() (*isa95.MaterialLot, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLot), nil
	}
}

func (m materialLotDo) FirstOrCreate() (*isa95.MaterialLot, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.MaterialLot), nil
	}
}

func (m materialLotDo) FindByPage(offset int, limit int) (result []*isa95.MaterialLot, count int64, err error) {
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

func (m materialLotDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m materialLotDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m materialLotDo) Delete(models ...*isa95.MaterialLot) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *materialLotDo) withDO(do gen.Dao) *materialLotDo {
	m.DO = *do.(*gen.DO)
	return m
}
