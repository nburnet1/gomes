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

func newOperationType(db *gorm.DB, opts ...gen.DOOption) operationType {
	_operationType := operationType{}

	_operationType.operationTypeDo.UseDB(db, opts...)
	_operationType.operationTypeDo.UseModel(&isa95.OperationType{})

	tableName := _operationType.operationTypeDo.TableName()
	_operationType.ALL = field.NewAsterisk(tableName)
	_operationType.ID = field.NewUint(tableName, "id")
	_operationType.CreatedAt = field.NewTime(tableName, "created_at")
	_operationType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_operationType.DeletedAt = field.NewField(tableName, "deleted_at")
	_operationType.Code = field.NewString(tableName, "code")
	_operationType.Description = field.NewString(tableName, "description")

	_operationType.fillFieldMap()

	return _operationType
}

type operationType struct {
	operationTypeDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Code        field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (o operationType) Table(newTableName string) *operationType {
	o.operationTypeDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o operationType) As(alias string) *operationType {
	o.operationTypeDo.DO = *(o.operationTypeDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *operationType) updateTableName(table string) *operationType {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewUint(table, "id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")
	o.Code = field.NewString(table, "code")
	o.Description = field.NewString(table, "description")

	o.fillFieldMap()

	return o
}

func (o *operationType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *operationType) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["code"] = o.Code
	o.fieldMap["description"] = o.Description
}

func (o operationType) clone(db *gorm.DB) operationType {
	o.operationTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o operationType) replaceDB(db *gorm.DB) operationType {
	o.operationTypeDo.ReplaceDB(db)
	return o
}

type operationTypeDo struct{ gen.DO }

func (o operationTypeDo) Debug() *operationTypeDo {
	return o.withDO(o.DO.Debug())
}

func (o operationTypeDo) WithContext(ctx context.Context) *operationTypeDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o operationTypeDo) ReadDB() *operationTypeDo {
	return o.Clauses(dbresolver.Read)
}

func (o operationTypeDo) WriteDB() *operationTypeDo {
	return o.Clauses(dbresolver.Write)
}

func (o operationTypeDo) Session(config *gorm.Session) *operationTypeDo {
	return o.withDO(o.DO.Session(config))
}

func (o operationTypeDo) Clauses(conds ...clause.Expression) *operationTypeDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o operationTypeDo) Returning(value interface{}, columns ...string) *operationTypeDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o operationTypeDo) Not(conds ...gen.Condition) *operationTypeDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o operationTypeDo) Or(conds ...gen.Condition) *operationTypeDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o operationTypeDo) Select(conds ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o operationTypeDo) Where(conds ...gen.Condition) *operationTypeDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o operationTypeDo) Order(conds ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o operationTypeDo) Distinct(cols ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o operationTypeDo) Omit(cols ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o operationTypeDo) Join(table schema.Tabler, on ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o operationTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o operationTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o operationTypeDo) Group(cols ...field.Expr) *operationTypeDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o operationTypeDo) Having(conds ...gen.Condition) *operationTypeDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o operationTypeDo) Limit(limit int) *operationTypeDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o operationTypeDo) Offset(offset int) *operationTypeDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o operationTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *operationTypeDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o operationTypeDo) Unscoped() *operationTypeDo {
	return o.withDO(o.DO.Unscoped())
}

func (o operationTypeDo) Create(values ...*isa95.OperationType) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o operationTypeDo) CreateInBatches(values []*isa95.OperationType, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o operationTypeDo) Save(values ...*isa95.OperationType) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o operationTypeDo) First() (*isa95.OperationType, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationType), nil
	}
}

func (o operationTypeDo) Take() (*isa95.OperationType, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationType), nil
	}
}

func (o operationTypeDo) Last() (*isa95.OperationType, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationType), nil
	}
}

func (o operationTypeDo) Find() ([]*isa95.OperationType, error) {
	result, err := o.DO.Find()
	return result.([]*isa95.OperationType), err
}

func (o operationTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.OperationType, err error) {
	buf := make([]*isa95.OperationType, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o operationTypeDo) FindInBatches(result *[]*isa95.OperationType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o operationTypeDo) Attrs(attrs ...field.AssignExpr) *operationTypeDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o operationTypeDo) Assign(attrs ...field.AssignExpr) *operationTypeDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o operationTypeDo) Joins(fields ...field.RelationField) *operationTypeDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o operationTypeDo) Preload(fields ...field.RelationField) *operationTypeDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o operationTypeDo) FirstOrInit() (*isa95.OperationType, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationType), nil
	}
}

func (o operationTypeDo) FirstOrCreate() (*isa95.OperationType, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.OperationType), nil
	}
}

func (o operationTypeDo) FindByPage(offset int, limit int) (result []*isa95.OperationType, count int64, err error) {
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

func (o operationTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o operationTypeDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o operationTypeDo) Delete(models ...*isa95.OperationType) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *operationTypeDo) withDO(do gen.Dao) *operationTypeDo {
	o.DO = *do.(*gen.DO)
	return o
}