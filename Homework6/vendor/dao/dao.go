package dao

import (
	"bytes"
	"container/list"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 时间格式
const (
	TF = "2006-01-02 15:04:05"
)

var sqlDB *sql.DB

func Open() *sql.DB {
	if sqlDB == nil {
		db, err := sql.Open("mysql", "root:qyhfbqzh@tcp(127.0.0.1:3306)/Homework6?charset=utf8&parseTime=true")
		if err != nil {
			panic(err)
		}
		sqlDB = db
	}

	return sqlDB
}

type IDaoBase interface {
	Init()
	Save(data interface{}) error
	Update(data interface{}) error
	Find() (*list.List, error)
}

type DaoBase struct {
	EntityType    reflect.Type
	sqlDB         *sql.DB
	tableName     string            // 表名
	pk            string            // 主键
	columnToField map[string]string // 字段名:属性名
	fieldToColumn map[string]string // 属性名:字段名
}

// 初始化
func (this *DaoBase) Init() {
	this.columnToField = make(map[string]string)
	this.fieldToColumn = make(map[string]string)

	types := this.EntityType

	for i := 0; i < types.NumField(); i++ {
		proto := types.Field(i)
		tag := proto.Tag

		if len(tag) > 0 {
			column := tag.Get("column")
			name := proto.Name
			this.columnToField[column] = name
			this.fieldToColumn[name] = column

			if len(tag.Get("table")) > 0 {
				this.tableName = tag.Get("table")
				this.pk = column
			}
		}
	}
}

// 预处理插入sql
func (this *DaoBase) preInsertSQL() (fieldNames list.List, sql string) {
	names := new(bytes.Buffer)
	values := new(bytes.Buffer)

	i := 0

	for column, fieldName := range this.columnToField {

		if i != 0 {
			names.WriteString(",")
			values.WriteString(",")
		}
		fieldNames.PushBack(fieldName)
		names.WriteString(column)
		values.WriteString("?")
		i++
	}
	sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", this.tableName, names.String(), values.String())
	return
}

//增加单列
func (this *DaoBase) Save(data interface{}) error {
	columns, sql := this.preInsertSQL()

	stmt, err := Open().Prepare(sql)
	args := this.preValue(data, columns)
	fmt.Println(sql, " ", args)
	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
	return err
}

// 更新一个实体
func (this *DaoBase) Update(data interface{}) error {
	columns, sql := this.preUpdateSQL()

	stmt, err := Open().Prepare(sql)
	args := this.preValue(data, columns)

	fmt.Println(sql, " ", args)
	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
	return err
}

// 预处理更新sql
func (this *DaoBase) preUpdateSQL() (fieldNames list.List, sql string) {
	sets := new(bytes.Buffer)

	i := 0
	for column, fieldName := range this.columnToField {
		if strings.EqualFold(column, this.pk) {
			continue
		}
		if i > 0 {
			sets.WriteString(",")
		}

		fieldNames.PushBack(fieldName)
		sets.WriteString(column)
		sets.WriteString("=?")
		i++
	}
	fieldNames.PushBack(this.columnToField[this.pk])
	sql = fmt.Sprintf("UPDATE %s SET %s WHERE %s=?", this.tableName, sets.String(), this.pk)
	return
}

// 预处理占位符的数据
func (this *DaoBase) preValue(data interface{}, fieldNames list.List) []interface{} {
	values := make([]interface{}, len(this.columnToField))
	object := reflect.ValueOf(data).Elem()

	for e, i := fieldNames.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		name := e.Value.(string)
		field := object.FieldByName(name)
		values[i] = this.fieldValue(field)
	}

	return values
}

// 通过reflect.Value来获取值
func (this *DaoBase) fieldValue(v reflect.Value) interface{} {
	if !v.IsValid() {
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Ptr:
		m := v.MethodByName("Format")
		rets := m.Call([]reflect.Value{reflect.ValueOf(TF)})
		t := rets[0].String()
		return t
	default:
		return nil
	}
}

// 根据sql语句查询多条记录
func (this *DaoBase) Find() (*list.List, error) {
	var sql = "SELECT * FROM userinfo"
	rows, err := Open().Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	l := len(columns)
	// 构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	values := make([]interface{}, l)
	scanArgs := make([]interface{}, l)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	data := list.New()
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}
		obj := this.parseQuery(columns, values)
		data.PushBack(obj)
	}
	return data, err
}

// 封装一条查询结果
func (this *DaoBase) parseQuery(columns []string, values []interface{}) interface{} {

	obj := reflect.New(this.EntityType).Interface()
	proto := reflect.ValueOf(obj).Elem()

	for i, col := range values {
		if col != nil {
			name := this.columnToField[columns[i]]
			var field = proto.FieldByName(name)
			b, ok := col.([]byte)
			if ok {
				this.parseQueryColumn(field, string(b))
			} else {
				this.parseQueryColumn(field, col)
			}
		}
	}
	return obj
}

// 赋值单个属性
func (this *DaoBase) parseQueryColumn(field reflect.Value, s interface{}) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(reflect.ValueOf(s).String())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, _ := strconv.ParseUint(reflect.ValueOf(s).String(), 10, 0)
		field.SetUint(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, _ := strconv.ParseInt(reflect.ValueOf(s).String(), 10, 0)
		field.SetInt(v)
	case reflect.Float32:
		v, _ := strconv.ParseFloat(reflect.ValueOf(s).String(), 32)
		field.SetFloat(v)
	case reflect.Float64:
		v, _ := strconv.ParseFloat(reflect.ValueOf(s).String(), 64)
		field.SetFloat(v)
	case reflect.Ptr:
		values := new(bytes.Buffer)
		vs := reflect.ValueOf(s)
		m := vs.MethodByName("Format")
		rets := m.Call([]reflect.Value{reflect.ValueOf(TF)})
		t := rets[0].String()
		values.WriteString(t)
		v, _ := time.Parse(TF, values.String())
		field.Set(reflect.ValueOf(&v))
	}
}
