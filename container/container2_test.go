package container

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/firmeve/firmeve/testdata/structs"
)

type Account struct {
	Id   uint32
	Name string
	//Nested *struct {
	//	Age uint8
	//}
}

func ParseOne(main *structs.Main, maps map[string]string, slices []string, account *Account) interface{} {
	//x := 1
	//y := 2
	//return x, y
	return nil
}

type Bs map[string]string

func TestBaseContainer_Resolve_Func2(t *testing.T) {
	x := 1
	y := "s"

	fmt.Println(reflect.TypeOf(&x) == reflect.TypeOf(&y))
	fmt.Println("#########################")
	main := new(structs.Main)
	fmt.Println(main)
	fmt.Println("1111111111111111111111111")
	z := map[string]string{"a": "a"}
	var bs Bs
	fmt.Println(reflect.TypeOf(z) == reflect.TypeOf(bs))

	c := New()
	parseFunc := ParseOne
	c.resolveFunc(reflect.TypeOf(parseFunc), reflect.ValueOf(parseFunc))

	//m := new(structs.Main)
	//c := New()
	//c.resolveStruct2(reflect.TypeOf(m), reflect.New(reflect.TypeOf(m).Elem()))
}

func TestBaseContainer_Resolve_struct(t *testing.T) {
	c := New()
	c.Bind("sub", &structs.Sub{
		SubPublicKey: "SubPublicKey2...###",
	})

	//accout := new(Account)
	//fmt.Println(c.resolveStruct2(reflect.TypeOf(accout)))
	//
	main := new(structs.Main)
	//sub := new(structs.Sub)
	m := c.resolveStruct2(reflect.TypeOf(main)).(*structs.Main)
	fmt.Printf("%#v\n", m)
	//fmt.Println("==================")
	//fmt.Println(m.PrtSub.SubPublicKey)
	//fmt.Println("==================")
	//c.resolveStruct2(reflect.TypeOf(sub), reflect.ValueOf(sub))
}

func TestBaseContainer_Bind(t *testing.T) {
	// 不支持静态类型的绑定
	// types应该更改为resolveTypes，表示已解析的类型

	// Bind
	// bind("string","string") -> share
	// bind("bool",bool) -> share
	// bind("any",reflect.Kind < 动态类型) -> share
	// bind("struct",struct object) -> share
	// bind("ptr", ptr object) -> share
	// bind("slice", slice) -> share
	// bind("array", array) -> share array 是值
	// bind("func" ,function(xxx,xxx)) -> no share
	// 只要是非函数类型即可绑定为单例

	// Resolve
	// 必须要支持新对象的创建
	// 必须要支持递归解析
	// resolve("string") -> 表示直接读取容器key Get()
	// 01 resolve(new(struct)|ptr) // ptr类型的struct
	// 02 resolve(func) // func 是有多种类型的参数
	// 03 resolve(slice|array|struct|map) 创建一个新的slice|array|struct
	// 如果 反射类型在container中存在并且是singleton那么则返回已存在的类型

	v1 := [2]int{0, 1}
	v2 := [3]int{0, 1}
	fmt.Println(reflect.TypeOf(v1) == reflect.TypeOf(v2))
	fmt.Println("!!!!!!!!!!!!!!!!!!!")

	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(d))
	fmt.Println("====")
	fmt.Println(reflect.TypeOf("abc") == reflect.TypeOf("b"))

	fmt.Println("########")
	z := []int{0, 1}
	fmt.Println(reflect.TypeOf(z).Elem().Kind())
	//type Account struct {
	//	Id     uint32
	//	Name   string
	//	Nested struct {
	//		Age uint8
	//	}
	//}
	//account := &Account{
	//	Id: 10, Name: "jim",
	//	Nested: struct{ Age uint8 }{Age: 20},
	//}
	str := "abc"
	fmt.Println(reflect.TypeOf(str).Kind(), reflect.ValueOf(str).Kind())
	fmt.Println("########")
	m := struct {
		id        int
		substruct *struct {
			title string
		}
	}{
		id: 1,
		substruct: &struct {
			title string
		}{
			title: "abc",
		},
	}
	fmt.Println(reflect.TypeOf(m).Field(1))

	//var b, c interface{}
	//var d int
	//if b == nil {
	//	fmt.Println("111")
	//}
	//if d == nil {
	//	fmt.Println("111")
	//}
	//fmt.Println(b == c)
}

func a(a string) string {
	return a
}
func c(c string) string {
	return c
}
func d(c string, d int) string {
	return c
}
func b(b int) int {
	return b
}

//func Test(p1,p2,p3)  {
//}
//// 普通调用
//p1 := ...
//p2 := ...
//p3 := ...
//Test(p1,p2,p3)
//
////ioc
//x.Resolve(Test)
//
//x.Resolve(Test) == Test(p1,p2,p3)
