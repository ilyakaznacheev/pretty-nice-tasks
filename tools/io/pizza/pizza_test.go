package pizza_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/ilyakaznacheev/pretty-nice-tasks/tools/io/pizza"
)

func TestInput(t *testing.T) {

	want := pizza.InputData{
		Slices:       17,
		Types:        4,
		SliceNumbers: []int{2, 5, 6, 8},
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), fmt.Sprintf("*.in"))
	if err != nil {
		t.Fatal("cannot create temporary file:", err)
	}
	defer os.Remove(tmpFile.Name())

	text := []byte("17 4\n2 5 6 8")
	if _, err = tmpFile.Write(text); err != nil {
		t.Fatal("failed to write to temporary file:", err)
	}

	got, err := pizza.Input(tmpFile.Name())
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}
	if !reflect.DeepEqual(got, &want) {
		t.Errorf("unexpected result %v, want %v", got, want)
	}
}

func TestOutput(t *testing.T) {
	testFileName := path.Join(os.TempDir(), fmt.Sprintf("testfile%d.out", rand.Int()*1000))
	out := pizza.OutputData{
		Types:    3,
		Ordering: []int{0, 2, 3},
	}
	want := "3\n0 2 3"

	err := pizza.Output(testFileName, out)
	defer os.Remove(testFileName)
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	raw, err := ioutil.ReadFile(testFileName)
	if err != nil {
		t.Fatal(err)
		return
	}

	if got := string(raw); got != want {
		t.Errorf("unexpected result %s want %s", got, want)
		return
	}

	type args struct {
		fileName string
		out      pizza.OutputData
	}
}
