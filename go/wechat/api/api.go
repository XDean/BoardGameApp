package api

type (
	UserState interface {
		Handle(string) string
	}
)
