package pkg

import "text/template"

// Make sure that the template compiles during package initialization
func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var outputTemplate = parseTemplateOrPanic(`
///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package bis

import (
	{{.DefaultImportPackage}}
)

type {{.StructName}} struct {
	Base
}

func NewModel() *{{.StructName}} {
	return new({{.StructName}})
}

func NewQueryBuilder() *{{.QueryBuilderName}} {
	return new({{.QueryBuilderName}})
}

func (t *{{.StructName}}) Create(ctx context.Context,db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errno.NewMsgLog("create err", errno.CommonErrorNo, err, ctx)
	}
	return t.Id, nil
}

type {{.QueryBuilderName}} struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *{{.QueryBuilderName}}) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *{{.QueryBuilderName}}) Updates(ctx context.Context,db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&{{.StructName}}{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errno.NewMsgLog("updates err", errno.CommonErrorNo, err, ctx)
	}
	return nil
}

func (qb *{{.QueryBuilderName}}) Delete(ctx context.Context,db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&{{.StructName}}{}).Error; err != nil {
		return errno.NewMsgLog("delete err", errno.CommonErrorNo, err, ctx)
	}
	return nil
}

func (qb *{{.QueryBuilderName}}) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&{{.StructName}}{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *{{.QueryBuilderName}}) First(db *gorm.DB) (*{{.StructName}}, error) {
	ret := &{{.StructName}}{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *{{.QueryBuilderName}}) QueryOne(db *gorm.DB) (*{{.StructName}}, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *{{.QueryBuilderName}}) QueryAll(db *gorm.DB) ([]*{{.StructName}}, error) {
	var ret []*{{.StructName}}
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *{{.QueryBuilderName}}) Limit(limit int) *{{.QueryBuilderName}} {
	qb.limit = limit
	return qb
}

func (qb *{{.QueryBuilderName}}) Offset(offset int) *{{.QueryBuilderName}} {
	qb.offset = offset
	return qb
}

{{$queryBuilderName := .QueryBuilderName}}
{{range .OptionFields}}
func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}(p mysql.Predicate, value {{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", p),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}In(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) Where{{call $.Helpers.Titelize .FieldName}}NotIn(value []{{.FieldType}}) *{{$queryBuilderName}} {
	 qb.where = append(qb.where, struct {
		prefix string
		value interface{}
	}{
		fmt.Sprintf("%v %v ?", "{{.ColumnName}}", "NOT IN"),
		value,
	})
	return qb
}

func (qb *{{$queryBuilderName}}) OrderBy{{call $.Helpers.Titelize .FieldName}}(asc bool) *{{$queryBuilderName}} {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "{{.ColumnName}} " + order)
	return qb
}
{{end}}
`)
