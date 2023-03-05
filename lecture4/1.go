package main

import "fmt"

type List interface{
	Push(value int)
	Get(ind int)int
	Delete(ind int)
	Contains(element int) bool
	size() int
}

type ArrayList struct{
	list []int
}

type Vector struct{
	listv []int
}

func(a *ArrayList) Push(value int){
	a.list = append(a.list,value)
	fmt.Println("Okay, element is added,my array: ",a.list)
}
func(v *Vector) Push(value int){
	v.listv = append(v.listv,value)
	fmt.Println("Okay, element is added,my vector:  ",v.listv)
}
func(v *Vector) Contains(e int) bool{ 
    for _,v := range v.listv{ 
        if v == e{ 
            fmt.Println("Vector contains this item") 
            return true 
        } 
    } 
    return false 
} 

func (v *Vector) size() int{
	size := len(v.listv)
	return size
}

func(v *Vector) Get(ind int){
	element:=v.listv[ind]
	fmt.Println("Value of the element under the given index:",element)
}
func(a *ArrayList) Get(ind int){
	element:=a.list[ind]
	fmt.Println("Value of the element under the given index:",element)
}
func(a *ArrayList) Delete(ind int){
	element:=a.list[ind]
	a.list=append(a.list[:ind],a.list[ind+1:]...)
	fmt.Println("Value of the deleted element in ArrayList: ",element)
}
func(v *Vector) Delete(ind int){
	element:=v.listv[ind]
	v.listv=append(v.listv[:ind],v.listv[ind+1:]...)
	fmt.Println("Value of the deleted element in Vector: ",element)
}

func (v *Vector) Front() int {
    if len(v.listv) == 0 { 
        return 0 
    } 
    return v.listv[0]
}

func (v *Vector) Back() int {
    if len(v.listv) == 0 { 
        return 0 
    } 
    return v.listv[len(v.listv) - 1]
}

func main(){
 	vec := Vector{}
 	vec.Push(9)
	vec.Push(8)
	vec.Delete(1)
	ar:= ArrayList{}
	ar.Push(8)
	ar.Push(2)
	ar.Delete(0)
	ar.Get(0)
	fmt.Println(vec.Front())
}