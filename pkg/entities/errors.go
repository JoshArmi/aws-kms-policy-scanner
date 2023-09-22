package entities

type TooPermissiveError struct{}

func (e *TooPermissiveError) Error() string {
	return "action is too permissive"
}
