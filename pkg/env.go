package asaas

type AsaasEnv int

const (
	SANDBOX    AsaasEnv = iota
	PRODUCTION AsaasEnv = iota
)

func (a AsaasEnv) IsEnumValid() bool {
	switch a {
	case SANDBOX, PRODUCTION:
		return true
	}
	return false
}

func (a AsaasEnv) BaseURL() string {
	return []string{"https://sandbox.asaas.com/api/", "https://api.asaas.com/"}[a]
}
