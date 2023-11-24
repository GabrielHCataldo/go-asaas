package asaas

type Env int

const (
	SANDBOX    Env = iota
	PRODUCTION Env = iota
)

func (a Env) IsEnumValid() bool {
	switch a {
	case SANDBOX, PRODUCTION:
		return true
	}
	return false
}

func (a Env) BaseURL() string {
	return []string{"https://sandbox.asaas.com/api/", "https://api.asaas.com/"}[a]
}
