package gitlab_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPathSplit(t *testing.T) {
	path := "Group"
	strs := strings.SplitN(path, string(os.PathSeparator), 4)
	fmt.Println(strs)
	path = "Group/"
	path = strings.TrimRight(path, "/")
	strs = strings.SplitN(path, string(os.PathSeparator), 4)
	fmt.Println(strs)
	path = "Group/Project"
	strs = strings.SplitN(path, string(os.PathSeparator), 4)
	fmt.Println(strs)
	path = "Group/Project/"
	path = strings.TrimRight(path, "/")
	strs = strings.SplitN(path, string(os.PathSeparator), 4)
	fmt.Println(strs)
	path = "Group/Project/Sha/Path/A/B/C/a.file"
	strs = strings.SplitN(path, string(os.PathSeparator), 4)
	fmt.Println(strs)
}
