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

func newSegmentParameter(db *gorm.DB, opts ...gen.DOOption) segmentParameter {
	_segmentParameter := segmentParameter{}

	_segmentParameter.segmentParameterDo.UseDB(db, opts...)
	_segmentParameter.segmentParameterDo.UseModel(&isa95.SegmentParameter{})

	tableName := _segmentParameter.segmentParameterDo.TableName()
	_segmentParameter.ALL = field.NewAsterisk(tableName)
	_segmentParameter.ID = field.NewUint(tableName, "id")
	_segmentParameter.CreatedAt = field.NewTime(tableName, "created_at")
	_segmentParameter.UpdatedAt = field.NewTime(tableName, "updated_at")
	_segmentParameter.DeletedAt = field.NewField(tableName, "deleted_at")
	_segmentParameter.Code = field.NewString(tableName, "code")
	_segmentParameter.Description = field.NewString(tableName, "description")
	_segmentParameter.Value = field.NewString(tableName, "value")
	_segmentParameter.MeasurementID = field.NewUint(tableName, "measurement_id")

	_segmentParameter.fillFieldMap()

	return _segmentParameter
}

type segmentParameter struct {
	segmentParameterDo

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

func (s segmentParameter) Table(newTableName string) *segmentParameter {
	s.segmentParameterDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s segmentParameter) As(alias string) *segmentParameter {
	s.segmentParameterDo.DO = *(s.segmentParameterDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *segmentParameter) updateTableName(table string) *segmentParameter {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Code = field.NewString(table, "code")
	s.Description = field.NewString(table, "description")
	s.Value = field.NewString(table, "value")
	s.MeasurementID = field.NewUint(table, "measurement_id")

	s.fillFieldMap()

	return s
}

func (s *segmentParameter) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *segmentParameter) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["code"] = s.Code
	s.fieldMap["description"] = s.Description
	s.fieldMap["value"] = s.Value
	s.fieldMap["measurement_id"] = s.MeasurementID
}

func (s segmentParameter) clone(db *gorm.DB) segmentParameter {
	s.segmentParameterDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s segmentParameter) replaceDB(db *gorm.DB) segmentParameter {
	s.segmentParameterDo.ReplaceDB(db)
	return s
}

type segmentParameterDo struct{ gen.DO }

func (s segmentParameterDo) Debug() *segmentParameterDo {
	return s.withDO(s.DO.Debug())
}

func (s segmentParameterDo) WithContext(ctx context.Context) *segmentParameterDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s segmentParameterDo) ReadDB() *segmentParameterDo {
	return s.Clauses(dbresolver.Read)
}

func (s segmentParameterDo) WriteDB() *segmentParameterDo {
	return s.Clauses(dbresolver.Write)
}

func (s segmentParameterDo) Session(config *gorm.Session) *segmentParameterDo {
	return s.withDO(s.DO.Session(config))
}

func (s segmentParameterDo) Clauses(conds ...clause.Expression) *segmentParameterDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s segmentParameterDo) Returning(value interface{}, columns ...string) *segmentParameterDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s segmentParameterDo) Not(conds ...gen.Condition) *segmentParameterDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s segmentParameterDo) Or(conds ...gen.Condition) *segmentParameterDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s segmentParameterDo) Select(conds ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s segmentParameterDo) Where(conds ...gen.Condition) *segmentParameterDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s segmentParameterDo) Order(conds ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s segmentParameterDo) Distinct(cols ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s segmentParameterDo) Omit(cols ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s segmentParameterDo) Join(table schema.Tabler, on ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s segmentParameterDo) LeftJoin(table schema.Tabler, on ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s segmentParameterDo) RightJoin(table schema.Tabler, on ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s segmentParameterDo) Group(cols ...field.Expr) *segmentParameterDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s segmentParameterDo) Having(conds ...gen.Condition) *segmentParameterDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s segmentParameterDo) Limit(limit int) *segmentParameterDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s segmentParameterDo) Offset(offset int) *segmentParameterDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s segmentParameterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *segmentParameterDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s segmentParameterDo) Unscoped() *segmentParameterDo {
	return s.withDO(s.DO.Unscoped())
}

func (s segmentParameterDo) Create(values ...*isa95.SegmentParameter) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s segmentParameterDo) CreateInBatches(values []*isa95.SegmentParameter, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s segmentParameterDo) Save(values ...*isa95.SegmentParameter) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s segmentParameterDo) First() (*isa95.SegmentParameter, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentParameter), nil
	}
}

func (s segmentParameterDo) Take() (*isa95.SegmentParameter, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentParameter), nil
	}
}

func (s segmentParameterDo) Last() (*isa95.SegmentParameter, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentParameter), nil
	}
}

func (s segmentParameterDo) Find() ([]*isa95.SegmentParameter, error) {
	result, err := s.DO.Find()
	return result.([]*isa95.SegmentParameter), err
}

func (s segmentParameterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.SegmentParameter, err error) {
	buf := make([]*isa95.SegmentParameter, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s segmentParameterDo) FindInBatches(result *[]*isa95.SegmentParameter, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s segmentParameterDo) Attrs(attrs ...field.AssignExpr) *segmentParameterDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s segmentParameterDo) Assign(attrs ...field.AssignExpr) *segmentParameterDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s segmentParameterDo) Joins(fields ...field.RelationField) *segmentParameterDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s segmentParameterDo) Preload(fields ...field.RelationField) *segmentParameterDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s segmentParameterDo) FirstOrInit() (*isa95.SegmentParameter, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentParameter), nil
	}
}

func (s segmentParameterDo) FirstOrCreate() (*isa95.SegmentParameter, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentParameter), nil
	}
}

func (s segmentParameterDo) FindByPage(offset int, limit int) (result []*isa95.SegmentParameter, count int64, err error) {
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

func (s segmentParameterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s segmentParameterDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s segmentParameterDo) Delete(models ...*isa95.SegmentParameter) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *segmentParameterDo) withDO(do gen.Dao) *segmentParameterDo {
	s.DO = *do.(*gen.DO)
	return s
}