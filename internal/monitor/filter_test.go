package monitor

import (
	"reflect"
	"testing"
)

func TestFilterRestartedPods(t *testing.T) {
	input := []PodRestartInfo{
		{"default", "pod1", "container1", 2},
		{"default", "pod2", "container2", 0},
		{"kube-system", "pod3", "container3", 5},
	}

	expected := []PodRestartInfo{
		{"default", "pod1", "container1", 2},
		{"kube-system", "pod3", "container3", 5},
	}

	result := FilterRestartedPods(input, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

/*
Is It Best Practice to Place Unit Tests in the Same Folder in Go? Yes — this is idiomatic Go.
In Go, test files go in the same package and directory.
*/
