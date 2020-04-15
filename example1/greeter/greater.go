package greeter

type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}

type Greeter struct {
	database DB
	lang     string
}

func NewGreeter(db DB, lang string) Greeter {
	return Greeter{db, lang}
}

func (g *Greeter) Greet() string {
	msg, _ := g.database.FetchMessage(g.lang)

	return "Message is " + msg
}

func (g *Greeter) GreetInDefaultMsg() string {
	msg, _ := g.database.FetchDefaultMessage()

	return "Message is " + msg
}
