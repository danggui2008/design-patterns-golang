package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {

	collection := NewConcreteCollection[string]()
	collection.Add("A")
	collection.Add("B")
	collection.Add("C")

	iterator := collection.CreateIterator()
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}

}
