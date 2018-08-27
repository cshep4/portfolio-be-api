package service

type Service interface {
	Name() string
	Path() string
}