package test

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"reflect"
	"sql2proto"
	"testing"
	"time"
)

func TestKind(t *testing.T) {
	t.Log(reflect.TypeOf(1).String(),
		reflect.TypeOf(int64(1)).String(),
		reflect.TypeOf(float32(1)).String(),
		reflect.TypeOf(float64(1)).String(),
		reflect.TypeOf("").String(),
		reflect.TypeOf([]byte{}).String(),
		reflect.TypeOf(true).String(),
		reflect.TypeOf(time.Now()).String())
}

/*
func TestSql2Proto(t *testing.T) {
	sql, err := ioutil.ReadFile("a.sql")
	if err != nil {
		t.Error(err)
		return
	}
	args := sql2proto.NewConvertArgs().SetColPrefix("t_").SetTablePrefix("t_")
	bs, err := sql2proto.FromSql(string(sql), args)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v\n", string(bs))

}*/

func TestPb(t *testing.T) {
	msg, err := builder.NewMessage("MyMessage").
		AddField(builder.NewField("foo", builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_INT64)).
			SetDefaultValue("123")).
		AddField(builder.NewField("baz", builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_INT64)).
			SetLabel(descriptor.FieldDescriptorProto_LABEL_REPEATED).
			SetOptions(&descriptor.FieldOptions{Packed: proto.Bool(true)})).
		Build()
	if err != nil {
		t.Error(err)
		return
	}
	//pp := new(protoprint.Printer)
	builder.NewFile("")
	t.Log(new(protoprint.Printer).PrintProtoToString(msg))
}

func TestParse(t *testing.T) {
	s, e := sql2proto.FromFile("a.sql", &sql2proto.ProtoArgs{
		"test",
		false,
	})
	if e != nil {
		t.Error(e)
		return
	}
	t.Log(s)
}
