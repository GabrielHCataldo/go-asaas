package asaas

type CreateInvoiceSettingRequest struct {
	MunicipalServiceId   string                   `json:"municipalServiceId,omitempty"`
	MunicipalServiceCode string                   `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName string                   `json:"municipalServiceName,omitempty"`
	UpdatePayment        bool                     `json:"updatePayment,omitempty"`
	Deductions           float64                  `json:"deductions,omitempty"`
	EffectiveDatePeriod  InvoiceDatePeriod        `json:"effectiveDatePeriod,omitempty" validate:"omitempty,enum"`
	ReceivedOnly         bool                     `json:"receivedOnly,omitempty"`
	DaysBeforeDueDate    InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty" validate:"omitempty,enum"`
	Observations         string                   `json:"observations,omitempty"`
	Taxes                *InvoiceTaxesRequest     `json:"taxes,omitempty"`
}

type UpdateInvoiceSettingRequest struct {
	Deductions          float64                  `json:"deductions,omitempty"`
	EffectiveDatePeriod InvoiceDatePeriod        `json:"effectiveDatePeriod,omitempty" validate:"omitempty,enum"`
	ReceivedOnly        bool                     `json:"receivedOnly,omitempty"`
	DaysBeforeDueDate   InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty" validate:"omitempty,enum"`
	Observations        string                   `json:"observations,omitempty"`
	Taxes               *InvoiceTaxesRequest     `json:"taxes,omitempty"`
}

type InvoiceTaxesRequest struct {
	RetainIss bool    `json:"retainIss,omitempty"`
	Iss       float64 `json:"iss,omitempty" validate:"required,gt=0"`
	Confins   float64 `json:"cofins,omitempty" validate:"required,gt=0"`
	Csll      float64 `json:"csll,omitempty" validate:"required,gt=0"`
	Inss      float64 `json:"inss,omitempty" validate:"required,gt=0"`
	Ir        float64 `json:"ir,omitempty" validate:"required,gt=0"`
	Pis       float64 `json:"pis,omitempty" validate:"required,gt=0"`
}

type GetAllInvoicesRequest struct {
	EffectiveDateGE   *Date         `json:"effectiveDate[ge],omitempty"`
	EffectiveDateLE   *Date         `json:"effectiveDate[le],omitempty"`
	ExternalReference string        `json:"externalReference,omitempty"`
	Status            InvoiceStatus `json:"status,omitempty"`
	Customer          string        `json:"customer,omitempty"`
	Offset            int           `json:"offset,omitempty"`
	Limit             int           `json:"limit,omitempty"`
}

type InvoiceSettingResponse struct {
	MunicipalServiceId    string                   `json:"municipalServiceId,omitempty"`
	MunicipalServiceCode  string                   `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName  string                   `json:"municipalServiceName,omitempty"`
	Deductions            float64                  `json:"deductions,omitempty"`
	InvoiceCreationPeriod string                   `json:"invoiceCreationPeriod,omitempty"`
	DaysBeforeDueDate     InvoiceDaysBeforeDueDate `json:"daysBeforeDueDate,omitempty"`
	ReceivedOnly          bool                     `json:"receivedOnly,omitempty"`
	Observations          string                   `json:"observations,omitempty"`
	Taxes                 *InvoiceTaxesResponse    `json:"taxes,omitempty"`
	Errors                []ErrorResponse          `json:"errors,omitempty"`
}

type InvoiceResponse struct {
	ID                        string                `json:"id,omitempty"`
	Status                    InvoiceStatus         `json:"status,omitempty"`
	Customer                  string                `json:"customer,omitempty"`
	Type                      string                `json:"type,omitempty"`
	StatusDescription         string                `json:"statusDescription,omitempty"`
	ServiceDescription        string                `json:"serviceDescription,omitempty"`
	PdfUrl                    string                `json:"pdfUrl,omitempty"`
	XmlUrl                    string                `json:"xmlUrl,omitempty"`
	RpsSerie                  string                `json:"rpsSerie,omitempty"`
	RpsNumber                 string                `json:"rpsNumber,omitempty"`
	Number                    string                `json:"number,omitempty"`
	ValidationCode            string                `json:"validationCode,omitempty"`
	Value                     float64               `json:"value,omitempty"`
	Deductions                float64               `json:"deductions,omitempty"`
	EffectiveDate             *Date                 `json:"effectiveDate,omitempty"`
	Observations              string                `json:"observations,omitempty"`
	EstimatedTaxesDescription string                `json:"estimatedTaxesDescription,omitempty"`
	Payment                   string                `json:"payment,omitempty"`
	Installment               string                `json:"installment,omitempty"`
	ExternalReference         string                `json:"externalReference,omitempty"`
	Taxes                     *InvoiceTaxesResponse `json:"taxes,omitempty"`
	MunicipalServiceCode      string                `json:"municipalServiceCode,omitempty"`
	MunicipalServiceName      string                `json:"municipalServiceName,omitempty"`
}

type InvoiceTaxesResponse struct {
	RetainIss bool    `json:"retainIss,omitempty"`
	Iss       float64 `json:"iss,omitempty"`
	Confins   float64 `json:"cofins,omitempty"`
	Csll      float64 `json:"csll,omitempty"`
	Inss      float64 `json:"inss,omitempty"`
	Ir        float64 `json:"ir,omitempty"`
	Pis       float64 `json:"pis,omitempty"`
}
