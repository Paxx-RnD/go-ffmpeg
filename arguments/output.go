package arguments

type Outputs string

func (i *Outputs) Append(path string) *Outputs {
	*i = Outputs(path)
	return i
}

func (i *Outputs) Build() []string {
	return []string{string(*i)}
}
