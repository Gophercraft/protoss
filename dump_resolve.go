package protoss

import "strings"

func resolve_package_scope(package_scope, full_name string) (scoped_name string, err error) {
	if strings.HasPrefix(full_name, ".") {
		// remove first dot
		full_name = full_name[1:]

		// split strings
		package_scope_path := strings.Split(package_scope, ".")
		name_scope_path := strings.Split(full_name, ".")

		point_of_similarity := 0

		for i := 0; i < len(name_scope_path) && i < len(package_scope_path); i++ {
			if package_scope_path[i] == name_scope_path[i] {
				point_of_similarity = i + 1
			} else {
				break
			}
		}

		if point_of_similarity == len(name_scope_path) {
			point_of_similarity--
		}

		scoped_name = strings.Join(name_scope_path[point_of_similarity:], ".")
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
