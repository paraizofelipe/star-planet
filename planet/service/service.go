package service

import "github.com/paraizofelipe/star-planet/planet/repository"

type Service interface {
	repository.Reader
	repository.Writer
}
