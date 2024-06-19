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

func newSegmentData(db *gorm.DB, opts ...gen.DOOption) segmentData {
	_segmentData := segmentData{}

	_segmentData.segmentDataDo.UseDB(db, opts...)
	_segmentData.segmentDataDo.UseModel(&isa95.SegmentData{})

	tableName := _segmentData.segmentDataDo.TableName()
	_segmentData.ALL = field.NewAsterisk(tableName)
	_segmentData.ID = field.NewUint(tableName, "id")
	_segmentData.CreatedAt = field.NewTime(tableName, "created_at")
	_segmentData.UpdatedAt = field.NewTime(tableName, "updated_at")
	_segmentData.DeletedAt = field.NewField(tableName, "deleted_at")
	_segmentData.Code = field.NewString(tableName, "code")
	_segmentData.Description = field.NewString(tableName, "description")
	_segmentData.Value = field.NewString(tableName, "value")
	_segmentData.MeasurementID = field.NewUint(tableName, "measurement_id")

	_segmentData.fillFieldMap()

	return _segmentData
}

type segmentData struct {
	segmentDataDo

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

func (s segmentData) Table(newTableName string) *segmentData {
	s.segmentDataDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s segmentData) As(alias string) *segmentData {
	s.segmentDataDo.DO = *(s.segmentDataDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *segmentData) updateTableName(table string) *segmentData {
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

func (s *segmentData) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *segmentData) fillFieldMap() {
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

func (s segmentData) clone(db *gorm.DB) segmentData {
	s.segmentDataDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s segmentData) replaceDB(db *gorm.DB) segmentData {
	s.segmentDataDo.ReplaceDB(db)
	return s
}

type segmentDataDo struct{ gen.DO }

func (s segmentDataDo) Debug() *segmentDataDo {
	return s.withDO(s.DO.Debug())
}

func (s segmentDataDo) WithContext(ctx context.Context) *segmentDataDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s segmentDataDo) ReadDB() *segmentDataDo {
	return s.Clauses(dbresolver.Read)
}

func (s segmentDataDo) WriteDB() *segmentDataDo {
	return s.Clauses(dbresolver.Write)
}

func (s segmentDataDo) Session(config *gorm.Session) *segmentDataDo {
	return s.withDO(s.DO.Session(config))
}

func (s segmentDataDo) Clauses(conds ...clause.Expression) *segmentDataDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s segmentDataDo) Returning(value interface{}, columns ...string) *segmentDataDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s segmentDataDo) Not(conds ...gen.Condition) *segmentDataDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s segmentDataDo) Or(conds ...gen.Condition) *segmentDataDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s segmentDataDo) Select(conds ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s segmentDataDo) Where(conds ...gen.Condition) *segmentDataDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s segmentDataDo) Order(conds ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s segmentDataDo) Distinct(cols ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s segmentDataDo) Omit(cols ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s segmentDataDo) Join(table schema.Tabler, on ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s segmentDataDo) LeftJoin(table schema.Tabler, on ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s segmentDataDo) RightJoin(table schema.Tabler, on ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s segmentDataDo) Group(cols ...field.Expr) *segmentDataDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s segmentDataDo) Having(conds ...gen.Condition) *segmentDataDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s segmentDataDo) Limit(limit int) *segmentDataDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s segmentDataDo) Offset(offset int) *segmentDataDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s segmentDataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *segmentDataDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s segmentDataDo) Unscoped() *segmentDataDo {
	return s.withDO(s.DO.Unscoped())
}

func (s segmentDataDo) Create(values ...*isa95.SegmentData) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s segmentDataDo) CreateInBatches(values []*isa95.SegmentData, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s segmentDataDo) Save(values ...*isa95.SegmentData) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s segmentDataDo) First() (*isa95.SegmentData, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentData), nil
	}
}

func (s segmentDataDo) Take() (*isa95.SegmentData, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentData), nil
	}
}

func (s segmentDataDo) Last() (*isa95.SegmentData, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentData), nil
	}
}

func (s segmentDataDo) Find() ([]*isa95.SegmentData, error) {
	result, err := s.DO.Find()
	return result.([]*isa95.SegmentData), err
}

func (s segmentDataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.SegmentData, err error) {
	buf := make([]*isa95.SegmentData, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s segmentDataDo) FindInBatches(result *[]*isa95.SegmentData, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s segmentDataDo) Attrs(attrs ...field.AssignExpr) *segmentDataDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s segmentDataDo) Assign(attrs ...field.AssignExpr) *segmentDataDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s segmentDataDo) Joins(fields ...field.RelationField) *segmentDataDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s segmentDataDo) Preload(fields ...field.RelationField) *segmentDataDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s segmentDataDo) FirstOrInit() (*isa95.SegmentData, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentData), nil
	}
}

func (s segmentDataDo) FirstOrCreate() (*isa95.SegmentData, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.SegmentData), nil
	}
}

func (s segmentDataDo) FindByPage(offset int, limit int) (result []*isa95.SegmentData, count int64, err error) {
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

func (s segmentDataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s segmentDataDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s segmentDataDo) Delete(models ...*isa95.SegmentData) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *segmentDataDo) withDO(do gen.Dao) *segmentDataDo {
	s.DO = *do.(*gen.DO)
	return s
}
