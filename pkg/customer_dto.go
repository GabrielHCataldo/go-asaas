package asaas

type CreateCustomerRequest struct {
	Name                 string `json:"name,omitempty" validate:"required,full_name"`
	CpfCnpj              string `json:"cpfCnpj,omitempty" validate:"required,document"`
	Email                string `json:"email,omitempty" validate:"omitempty,email"`
	Phone                string `json:"phone,omitempty" validate:"omitempty,phone"`
	MobilePhone          string `json:"mobilePhone,omitempty" validate:"omitempty,phone"`
	Address              string `json:"address,omitempty"`
	AddressNumber        string `json:"addressNumber,omitempty"`
	Complement           string `json:"complement,omitempty"`
	Province             string `json:"province,omitempty"`
	PostalCode           string `json:"postalCode,omitempty" validate:"omitempty,postal_code"`
	ExternalReference    string `json:"externalReference,omitempty"`
	NotificationDisabled bool   `json:"notificationDisabled,omitempty"`
	AdditionalEmails     string `json:"additionalEmails,omitempty"`
	MunicipalInscription string `json:"municipalInscription,omitempty"`
}

type CreateCustomerResponse struct {
	ID                   string          `json:"id,omitempty"`
	Name                 string          `json:"name,omitempty"`
	Email                string          `json:"email,omitempty"`
	Phone                string          `json:"phone,omitempty"`
	MobilePhone          string          `json:"mobilePhone,omitempty"`
	Address              string          `json:"address,omitempty"`
	AddressNumber        string          `json:"addressNumber,omitempty"`
	Complement           string          `json:"complement,omitempty"`
	Province             string          `json:"province,omitempty"`
	PostalCode           string          `json:"postalCode,omitempty"`
	CpfCnpj              string          `json:"cpfCnpj,omitempty"`
	PersonType           string          `json:"personType,omitempty"`
	Deleted              bool            `json:"deleted,omitempty"`
	AdditionalEmails     string          `json:"additionalEmails,omitempty"`
	ExternalReference    string          `json:"externalReference,omitempty"`
	NotificationDisabled bool            `json:"notificationDisabled,omitempty"`
	City                 int             `json:"city,omitempty"`
	State                string          `json:"state,omitempty"`
	Country              string          `json:"country,omitempty"`
	Observations         string          `json:"observations,omitempty"`
	Errors               []ErrorResponse `json:"errors,omitempty"`
	DateCreated          Date            `json:"dateCreated,omitempty"`
}
