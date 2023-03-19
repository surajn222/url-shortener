package storage

type InterfaceStorage interface {
	InsertShortenedLinks(string, string) error
	GetAllLinks() map[string]string
	UpdateDomainCount(string) error
	GetDomainCount() map[string]int
	GetLink(string) (string, error)
}
