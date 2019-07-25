package container

import (
	"fmt"
	"github.com/firmeve/firmeve/testdata"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
	"time"
)

type message string

func TestContainer_Resolve_String_Alias_Type(t *testing.T) {
	i := NewInstance()

	i.Bind("message", message("message bar"))
	i.Bind("string", "string value")

	//assert.IsType(t,message(""),f.Get("message"))
	z := i.Get("message").(message)
	assert.Equal(t, "message bar", fmt.Sprintf("%s", z))
	assert.Equal(t, "string value", i.Get("string").(string))
}

type IntAlias int32

func TestContainer_Resolve_Number(t *testing.T) {
	var num int32
	num = 255
	var fnum float32
	fnum = 255.02
	var IntAliasNum IntAlias
	IntAliasNum = 80

	f := NewInstance()
	f.Bind(t.Name()+"num", num, )
	f.Bind(t.Name()+"fnum", fnum, )
	f.Bind(t.Name()+"IntAliasNum", IntAliasNum, )

	assert.Equal(t, num, f.Get(t.Name() + "num").(int32))
	assert.Equal(t, fnum, f.Get(t.Name() + "fnum").(float32))
	assert.Equal(t, IntAliasNum, f.Get(t.Name() + "IntAliasNum").(IntAlias))
}

func TestContainer_Resolve_Bool(t *testing.T) {
	f := NewInstance()
	f.Bind("bool", (false))

	assert.Equal(t, false, f.Get("bool").(bool))
}

func TestContainer_Resolve_Struct_Prt(t *testing.T) {
	t1 := testdata.NewT1()

	f := NewInstance()
	f.Bind(t.Name()+"t1", (t1))
	assert.Equal(t, t1, f.Get(t.Name()+"t1"))
}

type FuncType func() int64

func TestContainer_Resolve_Func_Simple(t *testing.T) {
	f := NewInstance()

	var z FuncType
	z = func() int64 {
		return time.Now().UnixNano()
	}

	fmt.Println(z())
	fmt.Println(z())
	fmt.Println(z())
	//t1 := testdata.NewT1()
	f.Bind(t.Name()+"z", (z))
	//assert.IsType(t, z, f.Get(t.Name()+"z"))

	result := f.Get(t.Name() + "z").(int64)
	log.Println(result)
}

//
func TestContainer_Resolve_Cover(t *testing.T) {
	t1 := testdata.NewT1()
	//t3 := testdata.NewT1Sturct()

	f := NewInstance()
	f.Bind(t.Name()+"t1", (t1), )
	assert.Equal(t, t1, f.Get(t.Name()+"t1"))

	f.Bind(t.Name()+"t1", (t1), WithBindCover(true))
	assert.Equal(t, t1, f.Get(t.Name()+"t1"))

	assert.Panics(t, func() {
		f.Bind(t.Name()+"t1", (t1), WithBindCover(false))
	}, "binding alias type already exists")
}

// 测试单例
func TestContainer_Singleton(t *testing.T) {
	//t1 := testdata.NewT1
	//t3 := testdata.NewT1Sturct()
	f := NewInstance()
	f.Bind(t.Name()+"t1.singleton", (testdata.NewT1), WithBindShare(true))
	assert.Equal(t, fmt.Sprintf("%p", f.Get(t.Name()+"t1.singleton")), fmt.Sprintf("%p", f.Get(t.Name()+"t1.singleton")))

	f.Bind(t.Name()+"t2.prototype", (testdata.NewT1), WithBindShare(false))
	assert.NotEqual(t, fmt.Sprintf("%p", f.Get(t.Name()+"t2.prototype")), fmt.Sprintf("%p", f.Get(t.Name()+"t2.prototype")))
}

func TestReflectType(t *testing.T) {
	// 字符串
	s1 := "abc"
	s2 := "def"
	s3 := "abc"
	fmt.Println(reflect.TypeOf(s1) == reflect.TypeOf(s2))
	fmt.Println(reflect.ValueOf(s1) == reflect.ValueOf(s3))
	// number
	n1 := 20
	n2 := 40
	var n3 int64
	n3 = 50
	fmt.Println(reflect.TypeOf(n1) == reflect.TypeOf(n2))
	fmt.Println(reflect.TypeOf(n1) == reflect.TypeOf(n3))
	// bool
	b1 := false
	b2 := true
	fmt.Println(reflect.TypeOf(b1) == reflect.TypeOf(b2))
	// struct
	t1 := testdata.NewT1Sturct()
	t2 := testdata.NewT1Sturct()
	t2.Name = "def"
	fmt.Println(reflect.TypeOf(t1) == reflect.TypeOf(t2))
	// *struct
	t3 := testdata.NewT1()
	t4 := testdata.NewT1()
	t5 := testdata.NewT22()
	fmt.Println(reflect.TypeOf(t3) == reflect.TypeOf(t4))
	fmt.Println(reflect.TypeOf(t3) == reflect.TypeOf(t5))
	// func
}

// 测试struct字段反射
func TestContainer_Resolve_Struct_Field(t *testing.T) {
	f := NewInstance()
	t1 := testdata.NewT1()
	f.Bind("t1", (t1))

	fmt.Printf("%#v\n", f.Resolve(testdata.NewT2))

	t1struct := testdata.NewT1Sturct()
	f.Bind("t1struct", (t1struct))
	fmt.Printf("%#v", f.Resolve(testdata.NewTStruct))
}

// 测试非单例注入
func TestContainer_Resolve_Prototype(t *testing.T) {
	f := NewInstance()
	f.Bind("t1", (testdata.NewT1), WithBindCover(true))
	log.Printf("%#v\n", f.Resolve(testdata.NewT2))
	assert.IsType(t, testdata.NewT2(f.Get("t1").(*testdata.T1)), f.Resolve(testdata.NewT2))
}

////
func TestContainer_Bind_Struct_Prt2(t *testing.T) {
	t1 := testdata.NewT1()
	//t2 := testdata.NewT2(t1)

	f := NewInstance()
	f.Bind(t.Name()+"t1", (t1))

	t2 := new(testdata.T2)
	//fmt.Printf("%#v\n", t2)
	result := f.Resolve(t2)
	fmt.Printf("%#v\n", result)
	//result.(*testdata.T2).Age = 10
	//fmt.Printf("%#v\n", t2)

	//t4 := testdata.T2{}
	//result2 := f.Get(t4)
	//fmt.Printf("%#v\n", result2.(testdata.T2))
}

func TestGetInstance(t *testing.T) {
	f := NewInstance()
	f1 := GetInstance()
	assert.Equal(t, f, f1)
}

func TestInstance_GetError(t *testing.T) {
	assert.Panics(t, func() {
		GetInstance()
	},`instance not exists`)
}

func TestContainer_Remove(t *testing.T) {
	t1 := testdata.NewT1()
	//t2 := testdata.NewT2(t1)

	f := NewInstance()
	f.Bind(t.Name()+"t1", t1)
	f.Remove(t.Name() + "t1")

	assert.Equal(t, false, f.Has(t.Name()+"t1"))
}



//
//func TestContainer_Bind(t *testing.T) {
//	//c:=b
//	//println(c())
//	//
//	//
//
//	//testReject1 := &testReject{"simon",30}
//
//	firmeve := NewFirmeve()
//	//t1 := NewT1()
//	//firmeve.Bind(t1)
//
//	//firmeve.Bind(WithBindShare(true),(func() (string,int) {
//	//	return `abc`,10
//	//}),WithBindName("abc"))
//
//	//firmeve.Bind(WithBindShare(true),(testReject))
//	//firmeve.Bind(WithBindShare(true),(testReject{"simon",30}))
//	z := []string{"a", "b"}
//	firmeve.Bind(WithBindName("abcd"), (z))
//
//	//firmeve.Bind(func() interface{} {
//	//	return NewT1()
//	//},false)
//
//	//fmt.Printf("%#v",firmeve.Resolve(demo.NewT2).(T2))
//}

//func TestContainer_Register(t *testing.T) {
// mock对象
//f := NewFirmeve(basePath)
//config := config2.NewConfig(strings.Join([]string{f.GetBasePath(),`testdata/conf`},`/`))
//f.Register(config)
//}

//func TestContainer_multi_level_resolution(t *testing.T) {
//	f := NewFirmeve(".")
//	z := new(FirmeveServiceProvider)
//	f.Bind("firmeve.provider", z,WithBindShare(true))
//	//fmt.Printf("%#v",f.Resolve(z).(*firmeve.FirmeveServiceProvider).Firmeve)
//	serviceProvider := new(cache.CacheServiceProvider)
//	//serviceProvider := cache.CacheServiceProvider{}
//	//f.Bind("service.provider2",serviceProvider)
//	//zs := func(cs *cache.CacheServiceProvider) *cache.CacheServiceProvider {
//	//	return cs
//	//}
//
//	//fmt.Println("#############")
//	fmt.Printf("%#v\n", f.Resolve(serviceProvider).(*cache.CacheServiceProvider).Provider)
//	//fmt.Printf("%#v\n",serviceProvider.Provider.Firmeve)
//	//f.Register(`cache`, new(cache.CacheServiceProvider))
//}

//func TestRun(t *testing.T) {
//	//f := NewFirmeve(basePath)
//	handlers := make(map[string]func(ctx context.Context) interface{},0)
//	handlers[`def`] = func(ctx context.Context) interface{} {
//		fmt.Println("==============")
//		log.Printf("%#v",ctx.Value("params"))
//		fmt.Println("==============")
//		return "abc"
//	}
//	gin := gin2.Default()
//	//ctx := context.Background()
//	//ctx := context.Background()
//	//fmt.Printf("%#p\n",ctx)
//	gin.GET(`/def`, func(context2 *gin2.Context) {
//		//fmt.Printf("%#p\n",context2)
//		//subctx,cancel := context.WithDeadline(context2)
//		//fmt.Printf("%#p",subctx)
//		//subctx = context.WithValue(subctx,"params",context2.Request.RequestURI)
//		result := handlers["def"](context2)
//		context2.String(200,result.(string))
//	})
//	gin.Run(":28088")
//}