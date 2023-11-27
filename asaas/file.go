package asaas

type FileTextPlainResponse string

func (f FileTextPlainResponse) String() string {
	return string(f)
}
