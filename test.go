


package main
import (
	"fmt"
	"reflect" // 这里引入reflect模块
)


//import "github.com/therecipe/qt/widgets"

type DrawDlg2 struct {
//	widgets.QWidget

	_ func() `constructor:"init"`
	_ func() `slot:"animate"`
}


type User struct {
	Name   string "user name" //这引号里面的就是tag
	Passwd string "user passsword"
}

func main2() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}


	type S struct{
		F string `species:"gopher" color:"blue"`
	}

	s2 := DrawDlg2{}
	st := reflect.TypeOf(s2)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}