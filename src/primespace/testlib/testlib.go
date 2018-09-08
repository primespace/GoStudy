package testlib

var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"
	pop["John Legend"] = "All of Me"
}

func GetMusic(name string) string {
	return pop[name]
}
