package enumDef

type Enum struct {
	value string
}

var (
	VALUE1 = Enum{"value1"}
	VALUE2 = Enum{"value2"}
)

func (this Enum) String() string {
	return this.value
}
