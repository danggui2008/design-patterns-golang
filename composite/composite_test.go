package composite

import "testing"

func TestComposite(t *testing.T) {
	file1 := File{"file1.txt"}
	file2 := File{"file2.txt"}

	subDirectory := NewDirectory("subDirectory")
	subDirectory.AddComponent(&file1)
	subDirectory.AddComponent(&file2)

	root := NewDirectory("root")
	root.AddComponent(subDirectory)
	root.DisplayInfo()

}
