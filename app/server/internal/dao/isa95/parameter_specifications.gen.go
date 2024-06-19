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

func newParameterSpecification(db *gorm.DB, opts ...gen.DOOption) parameterSpecification {
	_parameterSpecification := parameterSpecification{}

	_parameterSpecification.parameterSpecificationDo.UseDB(db, opts...)
	_parameterSpecification.parameterSpecificationDo.UseModel(&isa95.ParameterSpecification{})

	tableName := _parameterSpecification.parameterSpecificationDo.TableName()
	_parameterSpecification.ALL = field.NewAsterisk(tableName)
	_parameterSpecification.ID = field.NewUint(tableName, "id")
	_parameterSpecification.CreatedAt = field.NewTime(tableName, "created_at")
	_parameterSpecification.UpdatedAt = field.NewTime(tableName, "updated_at")
	_parameterSpecification.DeletedAt = field.NewField(tableName, "deleted_at")
	_parameterSpecification.Code = field.NewString(tableName, "code")
	_parameterSpecification.Description = field.NewString(tableName, "description")
	_parameterSpecification.Value = field.NewString(tableName, "value")
	_parameterSpecification.MeasurementID = field.NewUint(tableName, "measurement_id")

	_parameterSpecification.fillFieldMap()

	return _parameterSpecification
}

type parameterSpecification struct {
	parameterSpecificationDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	Code          field.String
	Description   field.String
	Value         field.String
	MeasurementID field.Uint

	fieldMap map[string]field.Expr
}

func (p parameterSpecification) Table(newTableName string) *parameterSpecification {
	p.parameterSpecificationDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p parameterSpecification) As(alias string) *parameterSpecification {
	p.parameterSpecificationDo.DO = *(p.parameterSpecificationDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *parameterSpecification) updateTableName(table string) *parameterSpecification {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Code = field.NewString(table, "code")
	p.Description = field.NewString(table, "description")
	p.Value = field.NewString(table, "value")
	p.MeasurementID = field.NewUint(table, "measurement_id")

	p.fillFieldMap()

	return p
}

func (p *parameterSpecification) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *parameterSpecification) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["code"] = p.Code
	p.fieldMap["description"] = p.Description
	p.fieldMap["value"] = p.Value
	p.fieldMap["measurement_id"] = p.MeasurementID
}

func (p parameterSpecification) clone(db *gorm.DB) parameterSpecification {
	p.parameterSpecificationDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p parameterSpecification) replaceDB(db *gorm.DB) parameterSpecification {
	p.parameterSpecificationDo.ReplaceDB(db)
	return p
}

type parameterSpecificationDo struct{ gen.DO }

func (p parameterSpecificationDo) Debug() *parameterSpecificationDo {
	return p.withDO(p.DO.Debug())
}

func (p parameterSpecificationDo) WithContext(ctx context.Context) *parameterSpecificationDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p parameterSpecificationDo) ReadDB() *parameterSpecificationDo {
	return p.Clauses(dbresolver.Read)
}

func (p parameterSpecificationDo) WriteDB() *parameterSpecificationDo {
	return p.Clauses(dbresolver.Write)
}

func (p parameterSpecificationDo) Session(config *gorm.Session) *parameterSpecificationDo {
	return p.withDO(p.DO.Session(config))
}

func (p parameterSpecificationDo) Clauses(conds ...clause.Expression) *parameterSpecificationDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p parameterSpecificationDo) Returning(value interface{}, columns ...string) *parameterSpecificationDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p parameterSpecificationDo) Not(conds ...gen.Condition) *parameterSpecificationDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p parameterSpecificationDo) Or(conds ...gen.Condition) *parameterSpecificationDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p parameterSpecificationDo) Select(conds ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p parameterSpecificationDo) Where(conds ...gen.Condition) *parameterSpecificationDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p parameterSpecificationDo) Order(conds ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p parameterSpecificationDo) Distinct(cols ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p parameterSpecificationDo) Omit(cols ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p parameterSpecificationDo) Join(table schema.Tabler, on ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p parameterSpecificationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p parameterSpecificationDo) RightJoin(table schema.Tabler, on ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p parameterSpecificationDo) Group(cols ...field.Expr) *parameterSpecificationDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p parameterSpecificationDo) Having(conds ...gen.Condition) *parameterSpecificationDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p parameterSpecificationDo) Limit(limit int) *parameterSpecificationDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p parameterSpecificationDo) Offset(offset int) *parameterSpecificationDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p parameterSpecificationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *parameterSpecificationDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p parameterSpecificationDo) Unscoped() *parameterSpecificationDo {
	return p.withDO(p.DO.Unscoped())
}

func (p parameterSpecificationDo) Create(values ...*isa95.ParameterSpecification) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p parameterSpecificationDo) CreateInBatches(values []*isa95.ParameterSpecification, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p parameterSpecificationDo) Save(values ...*isa95.ParameterSpecification) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p parameterSpecificationDo) First() (*isa95.ParameterSpecification, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.ParameterSpecification), nil
	}
}

func (p parameterSpecificationDo) Take() (*isa95.ParameterSpecification, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.ParameterSpecification), nil
	}
}

func (p parameterSpecificationDo) Last() (*isa95.ParameterSpecification, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.ParameterSpecification), nil
	}
}

func (p parameterSpecificationDo) Find() ([]*isa95.ParameterSpecification, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.ParameterSpecification), err
}

func (p parameterSpecificationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.ParameterSpecification, err error) {
	buf := make([]*isa95.ParameterSpecification, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p parameterSpecificationDo) FindInBatches(result *[]*isa95.ParameterSpecification, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p parameterSpecificationDo) Attrs(attrs ...field.AssignExpr) *parameterSpecificationDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p parameterSpecificationDo) Assign(attrs ...field.AssignExpr) *parameterSpecificationDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p parameterSpecificationDo) Joins(fields ...field.RelationField) *parameterSpecificationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p parameterSpecificationDo) Preload(fields ...field.RelationField) *parameterSpecificationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p parameterSpecificationDo) FirstOrInit() (*isa95.ParameterSpecification, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.ParameterSpecification), nil
	}
}

func (p parameterSpecificationDo) FirstOrCreate() (*isa95.ParameterSpecification, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.ParameterSpecification), nil
	}
}

func (p parameterSpecificationDo) FindByPage(offset int, limit int) (result []*isa95.ParameterSpecification, count int64, err error) {
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

func (p parameterSpecificationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p parameterSpecificationDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p parameterSpecificationDo) Delete(models ...*isa95.ParameterSpecification) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *parameterSpecificationDo) withDO(do gen.Dao) *parameterSpecificationDo {
	p.DO = *do.(*gen.DO)
	return p
}