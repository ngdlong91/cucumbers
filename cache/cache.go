package storage

type Cache interface {
	Conn()
	Read()
	Write()
}
