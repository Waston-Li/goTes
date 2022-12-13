package designPattern

import "fmt"

// 迭代器模式是一种行为设计模式，让你能在不暴露集合底层表现形式 （列表、 栈和树等）的情况下遍历集合中所有的元素
// 在迭代器的帮助下， 客户端可以用一个迭代器接口以相似的方式 遍历不同集合中的元素。

// 这里需要注意的是有两个典型的迭代器接口需要分清楚；
// 一个是集合类实现的可以创建迭代器的工厂方法接口一般命名为Iterable，包含的方法类似CreateIterator；
// 另一个是迭代器本身的接口，命名为Iterator，有Next及hasMore两个主要方法；

//eg
// 一个班级类中包括一个老师和若干个学生，我们要对班级所有成员进行遍历，
// 班级中老师存储在单独的结构字段中，学生存储在另外一个slice字段中，通过迭代器，我们实现统一遍历处理；

// Teacher 老师
type Teacher struct {
	name    string // 名称
	subject string // 所教课程
}

func NewTeacher(name, subject string) *Teacher {
	return &Teacher{
		name:    name,
		subject: subject,
	}
}

// Student 学生
type Student struct {
	name     string // 姓名
	sumScore int    // 考试总分数
}

// NewStudent 创建学生对象
func NewStudent(name string, sumScore int) *Student {
	return &Student{
		name:     name,
		sumScore: sumScore,
	}
}

// Member 成员接口
type Member interface {
	Desc() string // 输出成员描述信息
}

//实现接口
func (t *Teacher) Desc() string {
	return fmt.Sprintf("%s班主任老师负责教%s", t.name, t.subject)
}
func (t *Student) Desc() string {
	return fmt.Sprintf("%s同学考试总分为%d", t.name, t.sumScore)
}

// Class 班级，包括老师和同学
type Class struct {
	name     string
	teacher  *Teacher
	students []*Student
}

// NewClass 根据班主任老师名称，授课创建班级
func NewClass(name, teacherName, teacherSubject string) *Class {
	return &Class{
		name:    name,
		teacher: NewTeacher(teacherName, teacherSubject), //函数返回老师对象

	}
}
func (c *Class) Name() string {
	return c.name
}

// AddStudent 班级添加同学
func (c *Class) AddStudent(students ...*Student) { //注意多个参数的写法
	c.students = append(c.students, students...)
}

//---------------------------------------班级成员迭代器----------------------------------
// Iterator 迭代器接口
type Iterator interface {
	Next() Member  // 迭代下一个成员
	HasMore() bool // 是否还有
}

// memberIterator 班级成员迭代器实现
type memberIterator struct {
	class *Class // 需迭代的班级
	index int    // 迭代索引
}

func (c *Class) CreateIterator() Iterator {
	return &memberIterator{
		class: c,
		index: -1, // 迭代索引初始化为-1，从老师开始迭代
	}
}

// Iterable 可迭代集合接口，实现此接口返回迭代器
type Iterable interface {
	CreateIterator() Iterator
}

//接口方法实现
func (m *memberIterator) Next() Member {
	// 迭代索引为-1时，返回老师成员，否则遍历学生slice
	if m.index == -1 {
		m.index++
		return m.class.teacher
	}
	student := m.class.students[m.index]
	m.index++
	return student
}
func (m *memberIterator) HasMore() bool {
	return m.index < len(m.class.students)
}

func IteratorMode() {
	class := NewClass("三年级一班", "王明", "数学课")
	class.AddStudent(NewStudent("张三", 389),
		NewStudent("李四", 378),
		NewStudent("王五", 347)) //New返回指针

	fmt.Printf("%s成员如下:\n", class.Name())
	classIterator := class.CreateIterator()
	for classIterator.HasMore() {
		member := classIterator.Next()
		fmt.Println(member.Desc())
	}

}
