package xpath_test

import (
	"github.com/badzhang/due/v2/utils/xpath"
	"testing"
)

func TestSplit(t *testing.T) {
	path := "/etc/my.ini"

	t.Log(xpath.Split(path))
}
