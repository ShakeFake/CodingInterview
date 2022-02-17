package list

import "container/list"

// 实现一个栈
type List struct {
	item *list.List
}

// 获得一个新 List
func (l *List) NewList() *List {
	return l.Init()
}

// 从 List 尾部返回一个值
func (l *List) Back() interface{} {
	return l.item.Back()
}

// 从 list 头部返回一个值
func (l *List) Front() interface{} {
	return l.item.Front()
}

func (l *List) Init() *List {
	l.item.Init()
	return l
}

// 返回 list 的长度
func (l *List) Len() int {
	return l.item.Len()
}

// 在确定位置后插入一个值
func (l *List) InsertAfter() {
	// TODO
}

// 在确定位置前插入一个值
func (l *List) InsertBefore() {
	// TODO
}

func (l *List) MoveAfter() {
	// todo
}

func (l *List) MoveBefore() {
	// todo
}

func (l *List) MoveToBack() {
	// todo
}

func (l *List) MoveToBefore() {
	// todo
}

// 从尾部放入一个值
func (l *List) PushBack(v interface{}) interface{} {
	return l.item.PushBack(v).Value
}

// 在头部放入一个值
func (l *List) PushFront(v interface{}) interface{} {
	return l.item.PushFront(v).Value
}

// 从头部去除一个值
func (l *List) RemoveFront() interface{} {
	return l.item.Remove(l.item.Front())
}

// 从尾部去除一个值
func (l *List) RemoveBack() interface{} {
	return l.item.Remove(l.item.Back())
}
