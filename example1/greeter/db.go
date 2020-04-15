package greeter

type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type db struct {
}

func NewDB() DB {
	return new(db)
}

func (d *db) FetchMessage(lang string) (string, error) {
	if lang == "en" {
		return "hello", nil
	}

	if lang == "es" {
		return "holla", nil
	}

	return "selam", nil
}

func (d *db) FetchDefaultMessage() (string, error) {
	return "default message", nil
}
