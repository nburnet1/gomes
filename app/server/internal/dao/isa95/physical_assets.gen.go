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

func newPhysicalAsset(db *gorm.DB, opts ...gen.DOOption) physicalAsset {
	_physicalAsset := physicalAsset{}

	_physicalAsset.physicalAssetDo.UseDB(db, opts...)
	_physicalAsset.physicalAssetDo.UseModel(&isa95.PhysicalAsset{})

	tableName := _physicalAsset.physicalAssetDo.TableName()
	_physicalAsset.ALL = field.NewAsterisk(tableName)
	_physicalAsset.ID = field.NewUint(tableName, "id")
	_physicalAsset.CreatedAt = field.NewTime(tableName, "created_at")
	_physicalAsset.UpdatedAt = field.NewTime(tableName, "updated_at")
	_physicalAsset.DeletedAt = field.NewField(tableName, "deleted_at")
	_physicalAsset.Description = field.NewString(tableName, "description")
	_physicalAsset.Location = field.NewString(tableName, "location")
	_physicalAsset.FixedAssetID = field.NewUint(tableName, "fixed_asset_id")
	_physicalAsset.VendorID = field.NewString(tableName, "vendor_id")
	_physicalAsset.PhysicalAssetClassID = field.NewUint(tableName, "physical_asset_class_id")
	_physicalAsset.PhysicalAssetClass = physicalAssetBelongsToPhysicalAssetClass{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PhysicalAssetClass", "isa95.PhysicalAssetClass"),
	}

	_physicalAsset.fillFieldMap()

	return _physicalAsset
}

type physicalAsset struct {
	physicalAssetDo

	ALL                  field.Asterisk
	ID                   field.Uint
	CreatedAt            field.Time
	UpdatedAt            field.Time
	DeletedAt            field.Field
	Description          field.String
	Location             field.String
	FixedAssetID         field.Uint
	VendorID             field.String
	PhysicalAssetClassID field.Uint
	PhysicalAssetClass   physicalAssetBelongsToPhysicalAssetClass

	fieldMap map[string]field.Expr
}

func (p physicalAsset) Table(newTableName string) *physicalAsset {
	p.physicalAssetDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p physicalAsset) As(alias string) *physicalAsset {
	p.physicalAssetDo.DO = *(p.physicalAssetDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *physicalAsset) updateTableName(table string) *physicalAsset {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Description = field.NewString(table, "description")
	p.Location = field.NewString(table, "location")
	p.FixedAssetID = field.NewUint(table, "fixed_asset_id")
	p.VendorID = field.NewString(table, "vendor_id")
	p.PhysicalAssetClassID = field.NewUint(table, "physical_asset_class_id")

	p.fillFieldMap()

	return p
}

func (p *physicalAsset) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *physicalAsset) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["description"] = p.Description
	p.fieldMap["location"] = p.Location
	p.fieldMap["fixed_asset_id"] = p.FixedAssetID
	p.fieldMap["vendor_id"] = p.VendorID
	p.fieldMap["physical_asset_class_id"] = p.PhysicalAssetClassID

}

func (p physicalAsset) clone(db *gorm.DB) physicalAsset {
	p.physicalAssetDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p physicalAsset) replaceDB(db *gorm.DB) physicalAsset {
	p.physicalAssetDo.ReplaceDB(db)
	return p
}

type physicalAssetBelongsToPhysicalAssetClass struct {
	db *gorm.DB

	field.RelationField
}

func (a physicalAssetBelongsToPhysicalAssetClass) Where(conds ...field.Expr) *physicalAssetBelongsToPhysicalAssetClass {
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

func (a physicalAssetBelongsToPhysicalAssetClass) WithContext(ctx context.Context) *physicalAssetBelongsToPhysicalAssetClass {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a physicalAssetBelongsToPhysicalAssetClass) Session(session *gorm.Session) *physicalAssetBelongsToPhysicalAssetClass {
	a.db = a.db.Session(session)
	return &a
}

func (a physicalAssetBelongsToPhysicalAssetClass) Model(m *isa95.PhysicalAsset) *physicalAssetBelongsToPhysicalAssetClassTx {
	return &physicalAssetBelongsToPhysicalAssetClassTx{a.db.Model(m).Association(a.Name())}
}

type physicalAssetBelongsToPhysicalAssetClassTx struct{ tx *gorm.Association }

func (a physicalAssetBelongsToPhysicalAssetClassTx) Find() (result *isa95.PhysicalAssetClass, err error) {
	return result, a.tx.Find(&result)
}

func (a physicalAssetBelongsToPhysicalAssetClassTx) Append(values ...*isa95.PhysicalAssetClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a physicalAssetBelongsToPhysicalAssetClassTx) Replace(values ...*isa95.PhysicalAssetClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a physicalAssetBelongsToPhysicalAssetClassTx) Delete(values ...*isa95.PhysicalAssetClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a physicalAssetBelongsToPhysicalAssetClassTx) Clear() error {
	return a.tx.Clear()
}

func (a physicalAssetBelongsToPhysicalAssetClassTx) Count() int64 {
	return a.tx.Count()
}

type physicalAssetDo struct{ gen.DO }

func (p physicalAssetDo) Debug() *physicalAssetDo {
	return p.withDO(p.DO.Debug())
}

func (p physicalAssetDo) WithContext(ctx context.Context) *physicalAssetDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p physicalAssetDo) ReadDB() *physicalAssetDo {
	return p.Clauses(dbresolver.Read)
}

func (p physicalAssetDo) WriteDB() *physicalAssetDo {
	return p.Clauses(dbresolver.Write)
}

func (p physicalAssetDo) Session(config *gorm.Session) *physicalAssetDo {
	return p.withDO(p.DO.Session(config))
}

func (p physicalAssetDo) Clauses(conds ...clause.Expression) *physicalAssetDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p physicalAssetDo) Returning(value interface{}, columns ...string) *physicalAssetDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p physicalAssetDo) Not(conds ...gen.Condition) *physicalAssetDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p physicalAssetDo) Or(conds ...gen.Condition) *physicalAssetDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p physicalAssetDo) Select(conds ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p physicalAssetDo) Where(conds ...gen.Condition) *physicalAssetDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p physicalAssetDo) Order(conds ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p physicalAssetDo) Distinct(cols ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p physicalAssetDo) Omit(cols ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p physicalAssetDo) Join(table schema.Tabler, on ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p physicalAssetDo) LeftJoin(table schema.Tabler, on ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p physicalAssetDo) RightJoin(table schema.Tabler, on ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p physicalAssetDo) Group(cols ...field.Expr) *physicalAssetDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p physicalAssetDo) Having(conds ...gen.Condition) *physicalAssetDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p physicalAssetDo) Limit(limit int) *physicalAssetDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p physicalAssetDo) Offset(offset int) *physicalAssetDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p physicalAssetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *physicalAssetDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p physicalAssetDo) Unscoped() *physicalAssetDo {
	return p.withDO(p.DO.Unscoped())
}

func (p physicalAssetDo) Create(values ...*isa95.PhysicalAsset) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p physicalAssetDo) CreateInBatches(values []*isa95.PhysicalAsset, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p physicalAssetDo) Save(values ...*isa95.PhysicalAsset) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p physicalAssetDo) First() (*isa95.PhysicalAsset, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAsset), nil
	}
}

func (p physicalAssetDo) Take() (*isa95.PhysicalAsset, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAsset), nil
	}
}

func (p physicalAssetDo) Last() (*isa95.PhysicalAsset, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAsset), nil
	}
}

func (p physicalAssetDo) Find() ([]*isa95.PhysicalAsset, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.PhysicalAsset), err
}

func (p physicalAssetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.PhysicalAsset, err error) {
	buf := make([]*isa95.PhysicalAsset, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p physicalAssetDo) FindInBatches(result *[]*isa95.PhysicalAsset, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p physicalAssetDo) Attrs(attrs ...field.AssignExpr) *physicalAssetDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p physicalAssetDo) Assign(attrs ...field.AssignExpr) *physicalAssetDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p physicalAssetDo) Joins(fields ...field.RelationField) *physicalAssetDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p physicalAssetDo) Preload(fields ...field.RelationField) *physicalAssetDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p physicalAssetDo) FirstOrInit() (*isa95.PhysicalAsset, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAsset), nil
	}
}

func (p physicalAssetDo) FirstOrCreate() (*isa95.PhysicalAsset, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.PhysicalAsset), nil
	}
}

func (p physicalAssetDo) FindByPage(offset int, limit int) (result []*isa95.PhysicalAsset, count int64, err error) {
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

func (p physicalAssetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p physicalAssetDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p physicalAssetDo) Delete(models ...*isa95.PhysicalAsset) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *physicalAssetDo) withDO(do gen.Dao) *physicalAssetDo {
	p.DO = *do.(*gen.DO)
	return p
}
