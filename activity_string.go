// generated by stringer -type=activity; DO NOT EDIT

package gopher

import "fmt"

const _activity_name = "WaitingWonderingBoringLoving"

var _activity_index = [...]uint8{0, 7, 16, 22, 28}

func (i activity) String() string {
	if i+1 >= activity(len(_activity_index)) {
		return fmt.Sprintf("activity(%d)", i)
	}
	return _activity_name[_activity_index[i]:_activity_index[i+1]]
}
