package protoss

import "testing"

func TestResolvePath(t *testing.T) {
	var tcases = [][3]string{
		{"package.scope.part", ".package.scope.part", "part"},
		{"package.scope.v1", ".package.scope.v3", "v3"},
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
