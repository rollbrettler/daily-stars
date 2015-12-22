package stars

import "testing"

func TestPagesCount(t *testing.T) {

	first, _ := pagesCount(`<https://api.github.com/user/833571/starred?page=2>; rel="next", <https://api.github.com/user/833571/starred?page=14>; rel="last"`)
	if first != 14 {
		t.Error("Expected 14 Pages, got: ", first)
	}

	last, _ := pagesCount(`<https://api.github.com/user/833571/starred?page=1>; rel="first", <https://api.github.com/user/833571/starred?page=13>; rel="prev"`)
	if last != 14 {
		t.Error("Expected 14 Pages, got: ", last)
	}

	empty, _ := pagesCount("")
	if empty != 1 {
		t.Error("Expected 1 Page, got: ", empty)
	}
}
