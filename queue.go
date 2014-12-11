package main

import (
    "fmt"
    "container/list"
)

type Queue struct {
    v interface{}
    list *list.List
}

func Create() *Queue {
    return &Queue{
        v : "",
        list : list.New(),
    }
}

func (q Queue) RPush(v string) {
    q.list.PushBack(v)
}

func (q Queue) LPush(v string) {
    q.list.PushFront(v)
}

func (q Queue) LPop() interface{} {
    if q.list.Len == 0 {
        return
    }
    ele := q.list.Front()
    q.v = q.list.Remove(ele)
    return q.v
}

func (q Queue) RPop() interface{} {
    if q.list.Len == 0 {
        return
    }
    ele := q.list.Back()
    q.v = q.list.Remove(ele)
    return q.v
}

func (q Queue) RemoveAll() {
    q.list = list.New()
}

func (q Queue) Len() int {
    return q.list.Len()
}

func (q Queue) LastVal() interface{} {
    return q.v
}

func main () {
    q := Create()
    fmt.Println(q.v)
    q.RPush("demo")
    q.RPush("s")
    fmt.Println(q.Len())
    fmt.Println(q.LPop())
    fmt.Println(q.Len())
    fmt.Println(q.LPop())
    fmt.Println(q.Len())
    fmt.Println(q.LPop())
}
