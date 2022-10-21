package pack

type CMD struct {
	Pre_cmd       string
	Short_options string
	i             int
	parts         []*string //指向指针(string类型)的数组
}
