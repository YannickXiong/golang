/**
* @Author  : Yannick
* @File    : learn_anonymous_func.go
* @Date    : 2017-11-11
* @Desc    : As we know that there is no set data type in go lang,
*            this project implements a simple set with golang.
* @Feature :
*     Basic: Provide below basic features.
*  --------: --------------------------------------------------
*     New(): Create a new set.
*     Add(): Add an element into a set.
*  Delete(): Delete an element from a set.
*   Clear(): Clear a set, delete all elements from a set.
* IsEmpty(): Determin whether a set is empty.
*Contains()：Determin whether a set contains an element.
*     Len(): Get the number of all element in a set.
*    Same(): Determin whether a set is the same to another set.
*Snapshot(): Returns a snapshot of a set.
*  String(): Implements String().
* Destroy(): Destroy a set.
*
*   Advance: Provide below advance features.
*  --------: --------------------------------------------------
*IsSupper(): Determin whether a set is a supper set(A ∈ B) of another one.
*   Union(): Union operation for set A and set B(A∪B = B∪A).
Intersect(): Intersect operation for set A and set B(A∩B = B∩A).
*    Diff(): Diff operation for set A to set B(A-B != B-A)
SymmeDiff(): Symmetric diff operation for A and B(A△B = B△A = (A∪B)—(A∩B))
*/

package set

import (
	"bytes"
	"fmt"
	"sync"
	"unsafe"
)

// HashSet : key of map is an empyt interface, can store any date type.
type HashSet struct {
	// First letter capital will not lead a error when parse into json.
	// In go, first letter captill means public, or means private.
	Element      map[interface{}]bool
	sync.RWMutex // thread level security
}

/*
** New    : Create a new set
** @Para  :
** @Rece  :
** Return : the pointer which refers to a HashSet.
 */
func New() *HashSet {
	return &HashSet{
		Element: map[interface{}]bool{},
	}
}

/*
** Add    : Add an element into a set.
** @Para  : interface{}, support any data type.
** @Rece  : *HashSet
** Return : true  - Add successfully.
**        : false - Add failed.
** Addition: Here is no limit that the type of each element in a set must be the same.
**           Which means set.Add(1) & set.Add("xyang") both are OK.
 */
func (set *HashSet) Add(e interface{}) bool {
	set.Lock()
	defer set.Unlock()

	if _, ok := set.Element[e]; !ok { // e does not exist in set
		set.Element[e] = true

		return true
	}

	return false
}

/*
** Delete : Delete an element from a set.
** @Para  : interface{}, support any data type.
** @Rece  : *HashSet
** Return : true  - Delete successfully.
**        : false - Delete failed.
 */
func (set *HashSet) Delete(e interface{}) bool {
	set.Lock()
	defer set.Unlock()

	if _, ok := set.Element[e]; ok { // e does exist in set
		delete(set.Element, e)

		return true
	}

	return false
}

/*
** Clear :  Clear a set, delete all elements from a set.
** @Para  :
** @Rece  : *HashSet
** Return :
 */
func (set *HashSet) Clear() {
	set.Lock()
	set.Unlock()

	// set.Element = map[interface{}]bool{}
	// fmt.Printf("before => %p, %p\n", s, s.Element)
	// s.Clear()
	// fmt.Printf("after => %p, %p\n", s, s.Element)
	// addr of s will not change, while addr of s.Element changed.
	// So this is not a good solution.
	for key := range set.Element { // is key, not value.
		set.Delete(key)
	}
}

/*
** IsEmpty: Determin whether a set is empty.
** @Para  :
** @Rece  : *HashSet
** Return : true  - empty
**          false - not empty
 */
func (set *HashSet) IsEmpty() bool {

	if set.Len() == 0 {
		return true
	}

	return false
}

/*
** Contains：Determin whether a set contains an element.
** @Para   :
** @Rece   : *HashSet
** Return  : true  - empty
**          false - not empty
 */
func (set *HashSet) Contains(e interface{}) bool {
	set.RLock()
	defer set.RUnlock()

	if _, ok := set.Element[e]; ok {
		return true
	}

	return false
}

/*
** Len     ：Get the number of all element in a set.
** @Para   :
** @Rece   : *HashSet
** Return  : int
**
 */
func (set *HashSet) Len() int {
	return len(set.Element)
}

/*
** Same    : Determin whether a set is the same to another set.
** @Para   : other *HashSet
** @Rece   : *HashSet
** Return  : true  - set is the same to set other
**           false - set is different to set other
**
 */
func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}

	if len(set.Element) != len(other.Element) {
		return false
	}

	set.RLock()
	defer set.RUnlock()

	for key, value := range set.Element {
		if !other.Contains(key) { // has same key
			return false
		}

		// the same key should have the same value.
		// actually it's no need to check value, because HashSet.Element[key]=true.
		if other.Element[key] != value {
			return false
		}

	}

	return true

}

/*
** Snapshot：Returns a snapshot of a set.
** @Para   :
** @Rece   : *HashSet
** Return  : a slice, type is interface{}
 */
func (set *HashSet) Snapshot() []interface{} {

	set.Lock()
	defer set.Unlock()

	originLen := set.Len()

	snap := make([]interface{}, originLen)
	// elements can be added into or deleted from a set, so during snapshot,
	// len(snap) is changable.
	actualLen := 0

	// for key, value = range ... is no need, because all value is true.
	for key := range set.Element {
		if actualLen <= originLen { // some elements have been deleted from set.
			snap[actualLen] = key
		} else { //some elements have been added into set. This situation the rest of snap will be nil.
			snap = append(snap, key)
		}
		actualLen++
	}

	// delete nil elements from sanp.
	if actualLen < originLen {
		snap = snap[:actualLen]
	}

	return snap
}

/*
** String  ：Implements String().
** @Para   :
** @Rece   : *HashSet
** Return  : a formated string.
 */
// If String () is implemented, then fmt calls String () by default
func (set *HashSet) String() string {
	var buf bytes.Buffer

	set.RLock()
	defer set.RUnlock()

	buf.WriteString("HashSet Info: {")

	first := true
	for key := range set.Element {
		// control to print ","
		if first {
			first = false
		} else {
			buf.WriteString(", ")
		}

		buf.WriteString(fmt.Sprintf("%v", key))
	}

	buf.WriteString("}")

	return buf.String()
}

/*
** Destroy ：Destroy a set.
** @Para   : *HashSet
** @Rece   : *HashSet
** Return  :
 */
func Destroy(self *HashSet) *HashSet {
	self.Lock()
	defer self.Unlock()

	// s1 = {1:ture, "xyang":ture, 3.23:true}
	// in func Destroy, self is nil, while after set.Destroy(s1)
	// s1 is not nil, and s1 is still Still {1:ture, "xyang":ture, 3.23:true},
	// really don't understand.

	fmt.Printf("addr of self => %p, %p, %p, %d, %d\n ", &self.Element, self, &self, unsafe.Sizeof(self), unsafe.Sizeof(&self))
	// output: self is the address refers to struct, the same to. &self is the addr of pointer itself.
	// addr of s1 => 0xc42000a080, 0xc42000a080, 0xc42000c030, 8, 8
	// addr of self => 0xc42000a080, 0xc42000a080, 0xc42000c038, 8, 8
	// Conclusion: Destroy(self *HashSet), self is a copy of s1, but self and s1 both point the same data.
	// this is why we can change data by modify self.Element, but cannot cannot reset s1 to nil after reset self = nil.
	// More important is that, both value or pointer as a func argument, they are passed by value, not preferenced.
	// If value type, copy a value; if pointer, copy a pointer.

	// More details see https://studygolang.com/topics/4089.
	self = nil

	return self
	// * HashSet
	// fmt.Printf("type => %v\n", reflect.TypeOf(self))
	// // nil
	// fmt.Printf("value => %v\n", reflect.ValueOf(self))
}
