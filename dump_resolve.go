package protoss

import (
	"slices"
	"strings"
)

func resolve_package_scope(package_scope, full_name string) (scoped_name string, err error) {
	if strings.HasPrefix(full_name, ".") {
		// remove first dot
		non_prefixed_name := full_name[1:]

		// split strings
		package_scope_path := strings.Split(package_scope, ".")
		name_path := strings.Split(non_prefixed_name, ".")

		point_of_similarity := 0

		for i := 0; i < len(name_path) && i < len(package_scope_path); i++ {
			partial_scope_path := strings.Join(package_scope_path[:i+1], ".")
			partial_name_path := strings.Join(name_path[:i+1], ".")

			if partial_scope_path == partial_name_path {
				point_of_similarity = i + 1
			} else {
				break
			}
		}

		// if point_of_similarity < len(name_path) {
		// 	point_of_similarity++
		// }
		// if point_of_similarity == len(name_path) {
		// 	point_of_similarity--
		// }

		// Don't try to resolve where there may be confusion
		resolved := name_path[point_of_similarity:]
		if len(resolved) > 0 && len(package_scope_path) > 0 {
			if slices.Contains(package_scope_path, resolved[0]) {
				return full_name, nil
			}
		}

		scoped_name = strings.Join(name_path[point_of_similarity:], ".")
		return
	}

	return full_name, nil
}

// if package is main.package.id
// and full name is .main.package.other.thing
// scoped_name could successfully be other.thing
// since the main.package part is already implied
func (fd *filedumper) resolve_type_name_in_scope(full_name string) (scoped_name string, err error) {
	return resolve_package_scope(fd.package_scope, full_name)
}
