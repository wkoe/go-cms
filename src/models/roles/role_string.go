// Code generated by "stringer -type=Role src/models/roles/role.go"; DO NOT EDIT

package roles

import "fmt"

const _Role_name = "AdministratorAuthorEditorSubscriber"

var _Role_index = [...]uint8{0, 13, 19, 25, 35}

func (i Role) String() string {
	i -= 1
	if i >= Role(len(_Role_index)-1) {
		return fmt.Sprintf("Role(%d)", i+1)
	}
	return _Role_name[_Role_index[i]:_Role_index[i+1]]
}
