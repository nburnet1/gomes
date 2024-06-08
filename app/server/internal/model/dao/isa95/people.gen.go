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

	"gomes/pkg/model/isa95"
)

func newPerson(db *gorm.DB, opts ...gen.DOOption) person {
	_person := person{}

	_person.personDo.UseDB(db, opts...)
	_person.personDo.UseModel(&isa95.Person{})

	tableName := _person.personDo.TableName()
	_person.ALL = field.NewAsterisk(tableName)
	_person.ID = field.NewUint(tableName, "id")
	_person.CreatedAt = field.NewTime(tableName, "created_at")
	_person.UpdatedAt = field.NewTime(tableName, "updated_at")
	_person.DeletedAt = field.NewField(tableName, "deleted_at")
	_person.Description = field.NewString(tableName, "description")
	_person.Name = field.NewString(tableName, "name")
	_person.PersonnelClassID = field.NewUint(tableName, "personnel_class_id")
	_person.PersonnelClass = personBelongsToPersonnelClass{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("PersonnelClass", "isa95.PersonnelClass"),
	}

	_person.fillFieldMap()

	return _person
}

type person struct {
	personDo

	ALL              field.Asterisk
	ID               field.Uint
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Description      field.String
	Name             field.String
	PersonnelClassID field.Uint
	PersonnelClass   personBelongsToPersonnelClass

	fieldMap map[string]field.Expr
}

func (p person) Table(newTableName string) *person {
	p.personDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p person) As(alias string) *person {
	p.personDo.DO = *(p.personDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *person) updateTableName(table string) *person {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Description = field.NewString(table, "description")
	p.Name = field.NewString(table, "name")
	p.PersonnelClassID = field.NewUint(table, "personnel_class_id")

	p.fillFieldMap()

	return p
}

func (p *person) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *person) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["description"] = p.Description
	p.fieldMap["name"] = p.Name
	p.fieldMap["personnel_class_id"] = p.PersonnelClassID

}

func (p person) clone(db *gorm.DB) person {
	p.personDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p person) replaceDB(db *gorm.DB) person {
	p.personDo.ReplaceDB(db)
	return p
}

type personBelongsToPersonnelClass struct {
	db *gorm.DB

	field.RelationField
}

func (a personBelongsToPersonnelClass) Where(conds ...field.Expr) *personBelongsToPersonnelClass {
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

func (a personBelongsToPersonnelClass) WithContext(ctx context.Context) *personBelongsToPersonnelClass {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a personBelongsToPersonnelClass) Session(session *gorm.Session) *personBelongsToPersonnelClass {
	a.db = a.db.Session(session)
	return &a
}

func (a personBelongsToPersonnelClass) Model(m *isa95.Person) *personBelongsToPersonnelClassTx {
	return &personBelongsToPersonnelClassTx{a.db.Model(m).Association(a.Name())}
}

type personBelongsToPersonnelClassTx struct{ tx *gorm.Association }

func (a personBelongsToPersonnelClassTx) Find() (result *isa95.PersonnelClass, err error) {
	return result, a.tx.Find(&result)
}

func (a personBelongsToPersonnelClassTx) Append(values ...*isa95.PersonnelClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a personBelongsToPersonnelClassTx) Replace(values ...*isa95.PersonnelClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a personBelongsToPersonnelClassTx) Delete(values ...*isa95.PersonnelClass) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a personBelongsToPersonnelClassTx) Clear() error {
	return a.tx.Clear()
}

func (a personBelongsToPersonnelClassTx) Count() int64 {
	return a.tx.Count()
}

type personDo struct{ gen.DO }

func (p personDo) Debug() *personDo {
	return p.withDO(p.DO.Debug())
}

func (p personDo) WithContext(ctx context.Context) *personDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personDo) ReadDB() *personDo {
	return p.Clauses(dbresolver.Read)
}

func (p personDo) WriteDB() *personDo {
	return p.Clauses(dbresolver.Write)
}

func (p personDo) Session(config *gorm.Session) *personDo {
	return p.withDO(p.DO.Session(config))
}

func (p personDo) Clauses(conds ...clause.Expression) *personDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personDo) Returning(value interface{}, columns ...string) *personDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personDo) Not(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personDo) Or(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personDo) Select(conds ...field.Expr) *personDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personDo) Where(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personDo) Order(conds ...field.Expr) *personDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personDo) Distinct(cols ...field.Expr) *personDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personDo) Omit(cols ...field.Expr) *personDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personDo) Join(table schema.Tabler, on ...field.Expr) *personDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personDo) RightJoin(table schema.Tabler, on ...field.Expr) *personDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personDo) Group(cols ...field.Expr) *personDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personDo) Having(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personDo) Limit(limit int) *personDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personDo) Offset(offset int) *personDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personDo) Unscoped() *personDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personDo) Create(values ...*isa95.Person) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personDo) CreateInBatches(values []*isa95.Person, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personDo) Save(values ...*isa95.Person) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personDo) First() (*isa95.Person, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Person), nil
	}
}

func (p personDo) Take() (*isa95.Person, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Person), nil
	}
}

func (p personDo) Last() (*isa95.Person, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Person), nil
	}
}

func (p personDo) Find() ([]*isa95.Person, error) {
	result, err := p.DO.Find()
	return result.([]*isa95.Person), err
}

func (p personDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*isa95.Person, err error) {
	buf := make([]*isa95.Person, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personDo) FindInBatches(result *[]*isa95.Person, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personDo) Attrs(attrs ...field.AssignExpr) *personDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personDo) Assign(attrs ...field.AssignExpr) *personDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personDo) Joins(fields ...field.RelationField) *personDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personDo) Preload(fields ...field.RelationField) *personDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personDo) FirstOrInit() (*isa95.Person, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Person), nil
	}
}

func (p personDo) FirstOrCreate() (*isa95.Person, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*isa95.Person), nil
	}
}

func (p personDo) FindByPage(offset int, limit int) (result []*isa95.Person, count int64, err error) {
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

func (p personDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personDo) Delete(models ...*isa95.Person) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personDo) withDO(do gen.Dao) *personDo {
	p.DO = *do.(*gen.DO)
	return p
}
