package board

type typeCheckable interface {
	typeError() error
}
