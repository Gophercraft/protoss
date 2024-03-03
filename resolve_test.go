package protoss

import "testing"

func TestResolvePath(t *testing.T) {
	var tcases = [][3]string{
		{"package.scope.part", ".package.scope.part.Something", "Something"},
		{"package.scope.part", ".package.scope.part", ""},
		{"package.scope.v1", ".package.scope.v3.Item", "scope.v3.Item"},
		{"package.what.v2.thing.scope", ".package.other.v2.thing_alt.alternate", "package.other.v2.thing_alt.alternate"},
	}

	for _, tcase := range tcases {
		package_scope := tcase[0]
		name_scope := tcase[1]
		correct_result := tcase[2]

		result, err := resolve_package_scope(package_scope, name_scope)
		if err != nil {
			t.Fatal(err)
		}

		if correct_result != result {
			t.Fatalf("result for resolve_package_scope(%s, %s) should have been %s, was %s", package_scope, name_scope, correct_result, result)
		}
	}
}
