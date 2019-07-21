package sql2proto

import (
	"github.com/Yimismi/sql2go"
	"github.com/go-xorm/core"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"reflect"
	"time"
)

func FromSql(sql string, args *ProtoArgs) (string, error) {
	parse := sql2go.ParseSql
	return generateCode(sql, parse, args)
}

func FromFile(fileName string, args *ProtoArgs) (string, error) {
	parse := sql2go.ParseSqlFile
	return generateCode(fileName, parse, args)
}

func generateCode(src string, parse func(string) ([]*core.Table, error), args *ProtoArgs) (string, error) {
	tables, err := parse(src)
	if err != nil {
		return "", err
	}
	p := getPrinter()
	pfb := builder.NewFile(args.PackageName + ".proto")
	pfb.SetPackageName(args.PackageName)
	pfb.SetProto3(args.Syntax3)
	for _, table := range tables {
		msgb := getMsgProtoBuilder(table)
		pfb.AddMessage(msgb)
	}
	pf, err := pfb.Build()
	if err != nil {
		return "", err
	}
	s, err := p.PrintProtoToString(pf)
	if err != nil {
		return "", err
	}
	return s, nil
}

func getMsgProtoBuilder(table *core.Table) *builder.MessageBuilder {
	msgb := builder.NewMessage(table.Name)
	if table.Comment != "" {
		msgb.SetComments(builder.Comments{
			LeadingComment: table.Comment,
		})
	}
	for _, col := range table.Columns() {
		f := getFiledBuilder(col)
		msgb.AddField(f)
	}
	return msgb
}

func getFiledBuilder(col *core.Column) *builder.FieldBuilder {
	f := builder.NewField(col.Name, builder.FieldTypeScalar(kindMapper(col)))
	if col.Default != "" {
		f.SetDefaultValue(col.Default)
	}
	if col.Comment != "" {
		f.SetComments(builder.Comments{
			TrailingComment: col.Comment,
		})
	}
	return f
}

func kindMapper(col *core.Column) descriptor.FieldDescriptorProto_Type {
	st := col.SQLType
	t := core.SQLType2Type(st)
	switch t.String() {
	case reflect.TypeOf(1).String():
		return descriptor.FieldDescriptorProto_TYPE_INT32
	case reflect.TypeOf(int64(1)).String():
		return descriptor.FieldDescriptorProto_TYPE_INT64
	case reflect.TypeOf(float32(1)).String():
		return descriptor.FieldDescriptorProto_TYPE_FLOAT
	case reflect.TypeOf(float64(1)).String():
		return descriptor.FieldDescriptorProto_TYPE_DOUBLE
	case reflect.TypeOf("").String():
		return descriptor.FieldDescriptorProto_TYPE_STRING
	case reflect.TypeOf([]byte{}).String():
		return descriptor.FieldDescriptorProto_TYPE_BYTES
	case reflect.TypeOf(true).String():
		return descriptor.FieldDescriptorProto_TYPE_DOUBLE
	case reflect.TypeOf(time.Now()).String():
		return descriptor.FieldDescriptorProto_TYPE_STRING
	default:
		return descriptor.FieldDescriptorProto_TYPE_STRING
	}
}

func getPrinter() *protoprint.Printer {
	p := new(protoprint.Printer)
	p.SortElements = true
	p.Indent = "    "
	return p
}

type ProtoArgs struct {
	PackageName string
	Syntax3     bool
}
