package composite

import "testing"

func TestComposite(t *testing.T) {
	file1 := &File{name: "AAA"}
	file2 := &File{name: "BBB"}
	file3 := &File{name: "AAAA1"}

	folder1 := &Fold{
		name: "ziliao",
	}

	folder1.Append(file1)

	folder2 := &Fold{
		name: "shiti",
	}
	folder2.Append(file2)
	folder2.Append(file3)
	folder2.Append(folder1)

	folder2.Search("AAA")
}
