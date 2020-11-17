package public

import (
	"encoding/json"
	"fmt"
	"strings"
)

// HashSet holds elements in go's native map
type HashSet struct {
	items map[interface{}]struct{}
}

var itemExists = struct{}{}

// NewHashSet instantiates a new empty set and adds the passed values, if any, to the set
func NewHashSet(values ...interface{}) *HashSet {
	set := &HashSet{items: make(map[interface{}]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *HashSet) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *HashSet) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *HashSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *HashSet) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *HashSet) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *HashSet) Clear() {
	set.items = make(map[interface{}]struct{})
}

// Values returns all items in the set.
func (set *HashSet) Values() []interface{} {
	values := make([]interface{}, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *HashSet) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *HashSet) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *HashSet) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}
