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

func newStatus(db *gorm.DB, opts ...gen.DOOption) status {
	_status := status{}

	_status.statusDo.UseDB(db, opts...)
	_status.statusDo.UseModel(&isa95.Status{})

	tableName := _status.statusDo.TableName()
	_status.ALL = field.NewAsterisk(tableName)
	_status.ID = field.NewUint(tableName, "id")
	_status.CreatedAt = field.NewTime(tableName, "created_at")
	_status.UpdatedAt = field.NewTime(tableName, "updated_at")
	_status.DeletedAt = field.NewField(tableName, "deleted_at")
	_status.Name = field.NewString(tableName, "name")
	_status.Description = field.NewString(tableName, "description")
	_status.StatusCategoryID = field.NewUint(tableName, "status_category_id")

	_status.fillFieldMap()

	return _status
}

type status struct {
	statusDo

	ALL              field.Asterisk
	ID               field.Uint
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Name             field.String
	Description      field.String
	StatusCategoryID field.Uint

	fieldMap map[string]field.Expr
}

func (s status) Table(newTableName string) *status {
	s.statusDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s status) As(alias string) *status {
	s.statusDo.DO = *(s.statusDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *status) updateTableName(table string) *status {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Name = field.NewString(table, "name")
	s.Description = field.NewString(table, "description")
	s.StatusCategoryID = field.NewUint(table, "status_category_id")

	s.fillFieldMap()

	return s
}

func (s *status) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *status) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["name"] = s.Name
	s.fieldMap["description"] = s.Description
	s.fieldMap["status_category_id"] = s.StatusCategoryID
}

func (s status) clone(db *gorm.DB) status {
	s.statusDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s status) replaceDB(db *gorm.DB) status {
	s.statusDo.ReplaceDB(db)
	return s
}

type statusDo struct{ gen.DO }

func (s statusDo) Debug() *statusDo {
	return s.withDO(s.DO.Debug())
}

func (s statusDo) WithContext(ctx context.Context) *statusDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s statusDo) ReadDB() *statusDo {
	return s.Clauses(dbresolver.Read)
}

func (s statusDo) WriteDB() *statusDo {
	return s.Clauses(dbresolver.Write)
}

func (s statusDo) Session(config *gorm.Session) *statusDo {
	return s.withDO(s.DO.Session(config))
}

func (s statusDo) Clauses(conds ...clause.Expression) *statusDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s statusDo) Returning(value interface{}, columns ...string) *statusDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s statusDo) Not(conds ...gen.Condition) *statusDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s statusDo) Or(conds ...gen.Condition) *statusDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s statusDo) Select(conds ...field.Expr) *statusDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s statusDo) Where(conds ...gen.Condition) *statusDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s statusDo) Order(conds ...field.Expr) *statusDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s statusDo) Distinct(cols ...field.Expr) *statusDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s statusDo) Omit(cols ...field.Expr) *statusDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s statusDo) Join(table schema.Tabler, on ...field.Expr) *statusDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s statusDo) LeftJoin(table schema.Tabler, on ...field.Expr) *statusDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s statusDo) RightJoin(table schema.Tabler, on ...field.Expr) *statusDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s statusDo) Group(cols ...field.Expr) *statusDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s statusDo) Having(conds ...gen.Condition) *statusDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s statusDo) Limit(limit int) *statusDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s statusDo) Offset(offset int) *statusDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s statusDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *statusDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s statusDo) Unscoped() *statusDo {
	return s.withDO(s.DO.Unscoped())
}

func (s statusDo) Create(values ...*isa95.Status) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s statusDo) CreateInBatches(values []*isa95.Status, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s statusDo) Save(values ...*isa95.Status) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s statusDo) First() (*isa95.Status, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Status), nil
	}
}

func (s statusDo) Take() (*isa95.Status, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Status), nil
	}
}

func (s statusDo) Last() (*isa95.Status, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Status), nil
	}
}

func (s statusDo) Find() ([]*isa95.Status, error) {
	result, err := s.DO.Find()
	return result.([]*isa95.Status), err
}

func (s statusDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.Status, err error) {
	buf := make([]*isa95.Status, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s statusDo) FindInBatches(result *[]*isa95.Status, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s statusDo) Attrs(attrs ...field.AssignExpr) *statusDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s statusDo) Assign(attrs ...field.AssignExpr) *statusDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s statusDo) Joins(fields ...field.RelationField) *statusDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s statusDo) Preload(fields ...field.RelationField) *statusDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s statusDo) FirstOrInit() (*isa95.Status, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Status), nil
	}
}

func (s statusDo) FirstOrCreate() (*isa95.Status, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Status), nil
	}
}

func (s statusDo) FindByPage(offset int, limit int) (result []*isa95.Status, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s statusDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s statusDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s statusDo) Delete(models ...*isa95.Status) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *statusDo) withDO(do gen.Dao) *statusDo {
	s.DO = *do.(*gen.DO)
	return s
}