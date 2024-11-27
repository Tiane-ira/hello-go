package code

type AppCode struct {
	Code int
	Msg  string
}

func (c AppCode) Error() string {
	return c.Msg
}
