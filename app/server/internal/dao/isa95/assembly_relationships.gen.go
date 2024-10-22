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

func newAssemblyRelationship(db *gorm.DB, opts ...gen.DOOption) assemblyRelationship {
	_assemblyRelationship := assemblyRelationship{}

	_assemblyRelationship.assemblyRelationshipDo.UseDB(db, opts...)
	_assemblyRelationship.assemblyRelationshipDo.UseModel(&isa95.AssemblyRelationship{})

	tableName := _assemblyRelationship.assemblyRelationshipDo.TableName()
	_assemblyRelationship.ALL = field.NewAsterisk(tableName)
	_assemblyRelationship.ID = field.NewUint(tableName, "id")
	_assemblyRelationship.CreatedAt = field.NewTime(tableName, "created_at")
	_assemblyRelationship.UpdatedAt = field.NewTime(tableName, "updated_at")
	_assemblyRelationship.DeletedAt = field.NewField(tableName, "deleted_at")
	_assemblyRelationship.Code = field.NewUint(tableName, "code")
	_assemblyRelationship.Name = field.NewString(tableName, "name")

	_assemblyRelationship.fillFieldMap()

	return _assemblyRelationship
}

type assemblyRelationship struct {
	assemblyRelationshipDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Code      field.Uint
	Name      field.String

	fieldMap map[string]field.Expr
}

func (a assemblyRelationship) Table(newTableName string) *assemblyRelationship {
	a.assemblyRelationshipDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a assemblyRelationship) As(alias string) *assemblyRelationship {
	a.assemblyRelationshipDo.DO = *(a.assemblyRelationshipDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *assemblyRelationship) updateTableName(table string) *assemblyRelationship {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Code = field.NewUint(table, "code")
	a.Name = field.NewString(table, "name")

	a.fillFieldMap()

	return a
}

func (a *assemblyRelationship) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *assemblyRelationship) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["code"] = a.Code
	a.fieldMap["name"] = a.Name
}

func (a assemblyRelationship) clone(db *gorm.DB) assemblyRelationship {
	a.assemblyRelationshipDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a assemblyRelationship) replaceDB(db *gorm.DB) assemblyRelationship {
	a.assemblyRelationshipDo.ReplaceDB(db)
	return a
}

type assemblyRelationshipDo struct{ gen.DO }

func (a assemblyRelationshipDo) Debug() *assemblyRelationshipDo {
	return a.withDO(a.DO.Debug())
}

func (a assemblyRelationshipDo) WithContext(ctx context.Context) *assemblyRelationshipDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a assemblyRelationshipDo) ReadDB() *assemblyRelationshipDo {
	return a.Clauses(dbresolver.Read)
}

func (a assemblyRelationshipDo) WriteDB() *assemblyRelationshipDo {
	return a.Clauses(dbresolver.Write)
}

func (a assemblyRelationshipDo) Session(config *gorm.Session) *assemblyRelationshipDo {
	return a.withDO(a.DO.Session(config))
}

func (a assemblyRelationshipDo) Clauses(conds ...clause.Expression) *assemblyRelationshipDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a assemblyRelationshipDo) Returning(value interface{}, columns ...string) *assemblyRelationshipDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a assemblyRelationshipDo) Not(conds ...gen.Condition) *assemblyRelationshipDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a assemblyRelationshipDo) Or(conds ...gen.Condition) *assemblyRelationshipDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a assemblyRelationshipDo) Select(conds ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a assemblyRelationshipDo) Where(conds ...gen.Condition) *assemblyRelationshipDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a assemblyRelationshipDo) Order(conds ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a assemblyRelationshipDo) Distinct(cols ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a assemblyRelationshipDo) Omit(cols ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a assemblyRelationshipDo) Join(table schema.Tabler, on ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a assemblyRelationshipDo) LeftJoin(table schema.Tabler, on ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a assemblyRelationshipDo) RightJoin(table schema.Tabler, on ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a assemblyRelationshipDo) Group(cols ...field.Expr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a assemblyRelationshipDo) Having(conds ...gen.Condition) *assemblyRelationshipDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a assemblyRelationshipDo) Limit(limit int) *assemblyRelationshipDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a assemblyRelationshipDo) Offset(offset int) *assemblyRelationshipDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a assemblyRelationshipDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *assemblyRelationshipDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a assemblyRelationshipDo) Unscoped() *assemblyRelationshipDo {
	return a.withDO(a.DO.Unscoped())
}

func (a assemblyRelationshipDo) Create(values ...*isa95.AssemblyRelationship) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a assemblyRelationshipDo) CreateInBatches(values []*isa95.AssemblyRelationship, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a assemblyRelationshipDo) Save(values ...*isa95.AssemblyRelationship) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a assemblyRelationshipDo) First() (*isa95.AssemblyRelationship, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.AssemblyRelationship), nil
	}
}

func (a assemblyRelationshipDo) Take() (*isa95.AssemblyRelationship, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.AssemblyRelationship), nil
	}
}

func (a assemblyRelationshipDo) Last() (*isa95.AssemblyRelationship, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.AssemblyRelationship), nil
	}
}

func (a assemblyRelationshipDo) Find() ([]*isa95.AssemblyRelationship, error) {
	result, err := a.DO.Find()
	return result.([]*isa95.AssemblyRelationship), err
}

func (a assemblyRelationshipDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.AssemblyRelationship, err error) {
	buf := make([]*isa95.AssemblyRelationship, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a assemblyRelationshipDo) FindInBatches(result *[]*isa95.AssemblyRelationship, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a assemblyRelationshipDo) Attrs(attrs ...field.AssignExpr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a assemblyRelationshipDo) Assign(attrs ...field.AssignExpr) *assemblyRelationshipDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a assemblyRelationshipDo) Joins(fields ...field.RelationField) *assemblyRelationshipDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a assemblyRelationshipDo) Preload(fields ...field.RelationField) *assemblyRelationshipDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a assemblyRelationshipDo) FirstOrInit() (*isa95.AssemblyRelationship, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.AssemblyRelationship), nil
	}
}

func (a assemblyRelationshipDo) FirstOrCreate() (*isa95.AssemblyRelationship, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.AssemblyRelationship), nil
	}
}

func (a assemblyRelationshipDo) FindByPage(offset int, limit int) (result []*isa95.AssemblyRelationship, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a assemblyRelationshipDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a assemblyRelationshipDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a assemblyRelationshipDo) Delete(models ...*isa95.AssemblyRelationship) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *assemblyRelationshipDo) withDO(do gen.Dao) *assemblyRelationshipDo {
	a.DO = *do.(*gen.DO)
	return a
}
