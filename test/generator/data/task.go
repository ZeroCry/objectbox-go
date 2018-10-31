package object

//go:generate objectbox-bindings

type Task struct {
	Id       uint64 `id`
	Uid      string `index`
	Name     string // TODO value index is not supported by the native-c binding at the moment `index:"value"`
	Place    string `index:"hash"`
	Source   string `index:"hash64"`
	Text     string `nameInDb:"text"`
	Date     uint64 `date`
	tempInfo string `transient`
	//Data []byte
}