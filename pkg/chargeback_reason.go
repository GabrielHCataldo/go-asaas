package asaas

type ChargebackReason string

const (
	ABSENCE_OF_PRINT                          ChargebackReason = "ABSENCE_OF_PRINT"
	ABSENT_CARD_FRAUD                         ChargebackReason = "ABSENT_CARD_FRAUD"
	CARD_ACTIVATED_PHONE_TRANSACTION          ChargebackReason = "CARD_ACTIVATED_PHONE_TRANSACTION"
	CARD_FRAUD                                ChargebackReason = "CARD_FRAUD"
	CARD_RECOVERY_BULLETIN                    ChargebackReason = "CARD_RECOVERY_BULLETIN"
	COMMERCIAL_DISAGREEMENT                   ChargebackReason = "COMMERCIAL_DISAGREEMENT"
	COPY_NOT_RECEIVED                         ChargebackReason = "COPY_NOT_RECEIVED"
	CREDIT_OR_DEBIT_PRESENTATION_ERROR        ChargebackReason = "CREDIT_OR_DEBIT_PRESENTATION_ERROR"
	DIFFERENT_PAY_METHOD                      ChargebackReason = "DIFFERENT_PAY_METHOD"
	FRAUD                                     ChargebackReason = "FRAUD"
	INCORRECT_TRANSACTION_VALUE               ChargebackReason = "INCORRECT_TRANSACTION_VALUE"
	INVALID_CURRENCY                          ChargebackReason = "INVALID_CURRENCY"
	INVALID_DATA                              ChargebackReason = "INVALID_DATA"
	LATE_PRESENTATION                         ChargebackReason = "LATE_PRESENTATION"
	LOCAL_REGULATORY_OR_LEGAL_DISPUTE         ChargebackReason = "LOCAL_REGULATORY_OR_LEGAL_DISPUTE"
	MULTIPLE_ROCS                             ChargebackReason = "MULTIPLE_ROCS"
	ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED  ChargebackReason = "ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED"
	OTHER_ABSENT_CARD_FRAUD                   ChargebackReason = "OTHER_ABSENT_CARD_FRAUD"
	PROCESS_ERROR                             ChargebackReason = "PROCESS_ERROR"
	RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE     ChargebackReason = "RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE"
	RECURRENCE_CANCELED                       ChargebackReason = "RECURRENCE_CANCELED"
	REQUIRED_AUTHORIZATION_NOT_GRANTED        ChargebackReason = "REQUIRED_AUTHORIZATION_NOT_GRANTED"
	RIGHT_OF_FULL_RECOURSE_FOR_FRAUD          ChargebackReason = "RIGHT_OF_FULL_RECOURSE_FOR_FRAUD"
	SALE_CANCELED                             ChargebackReason = "SALE_CANCELED"
	SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT ChargebackReason = "SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT"
	SERVICE_NOT_RECEIVED                      ChargebackReason = "SERVICE_NOT_RECEIVED"
	SPLIT_SALE                                ChargebackReason = "SPLIT_SALE"
	TRANSFERS_OF_DIVERSE_RESPONSIBILITIES     ChargebackReason = "TRANSFERS_OF_DIVERSE_RESPONSIBILITIES"
	UNQUALIFIED_CAR_RENTAL_DEBIT              ChargebackReason = "UNQUALIFIED_CAR_RENTAL_DEBIT"
	USA_CARDHOLDER_DISPUTE                    ChargebackReason = "USA_CARDHOLDER_DISPUTE"
	VISA_FRAUD_MONITORING_PROGRAM             ChargebackReason = "VISA_FRAUD_MONITORING_PROGRAM"
	WARNING_BULLETIN_FILE                     ChargebackReason = "WARNING_BULLETIN_FILE"
)

func (c ChargebackReason) IsEnumValid() bool {
	switch c {
	case ABSENCE_OF_PRINT, ABSENT_CARD_FRAUD, CARD_ACTIVATED_PHONE_TRANSACTION, CARD_FRAUD, CARD_RECOVERY_BULLETIN,
		COMMERCIAL_DISAGREEMENT, COPY_NOT_RECEIVED, CREDIT_OR_DEBIT_PRESENTATION_ERROR, DIFFERENT_PAY_METHOD, FRAUD,
		INCORRECT_TRANSACTION_VALUE, INVALID_CURRENCY, INVALID_DATA, LATE_PRESENTATION, LOCAL_REGULATORY_OR_LEGAL_DISPUTE,
		MULTIPLE_ROCS, ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED, OTHER_ABSENT_CARD_FRAUD, PROCESS_ERROR,
		RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE, RECURRENCE_CANCELED, REQUIRED_AUTHORIZATION_NOT_GRANTED,
		RIGHT_OF_FULL_RECOURSE_FOR_FRAUD, SALE_CANCELED, SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT, SERVICE_NOT_RECEIVED,
		SPLIT_SALE, TRANSFERS_OF_DIVERSE_RESPONSIBILITIES, UNQUALIFIED_CAR_RENTAL_DEBIT, USA_CARDHOLDER_DISPUTE,
		VISA_FRAUD_MONITORING_PROGRAM, WARNING_BULLETIN_FILE:
		return true
	}
	return false
}

func (c ChargebackReason) String() string {
	switch c {
	case ABSENCE_OF_PRINT:
		return "Ausência de impressão"
	case ABSENT_CARD_FRAUD:
		return "Fraude em ambiente de cartão não presente"
	case CARD_ACTIVATED_PHONE_TRANSACTION:
		return "Transação telefônica ativada por cartão"
	case CARD_FRAUD:
		return "Fraude em ambiente de cartão presente"
	case CARD_RECOVERY_BULLETIN:
		return "Boletim de negativação de cartões"
	case COMMERCIAL_DISAGREEMENT:
		return "Desacordo comercial"
	case COPY_NOT_RECEIVED:
		return "Cópia não atendida"
	case CREDIT_OR_DEBIT_PRESENTATION_ERROR:
		return "Erro de apresentação de crédito / débito"
	case DIFFERENT_PAY_METHOD:
		return "Pagamento por outros meios"
	case FRAUD:
		return "Sem autorização do portador do cartão"
	case INCORRECT_TRANSACTION_VALUE:
		return "Valor da transação é diferente"
	case INVALID_CURRENCY:
		return "Moeda inválida"
	case INVALID_DATA:
		return "Dados inválidos"
	case LATE_PRESENTATION:
		return "Apresentação tardia"
	case LOCAL_REGULATORY_OR_LEGAL_DISPUTE:
		return "Contestação regulatória / legal local"
	case MULTIPLE_ROCS:
		return "ROCs múltiplos"
	case ORIGINAL_CREDIT_TRANSACTION_NOT_ACCEPTED:
		return "Transação de crédito original não aceita"
	case OTHER_ABSENT_CARD_FRAUD:
		return "Outras fraudes - Cartão ausente"
	case PROCESS_ERROR:
		return "Erro de processamento"
	case RECEIVED_COPY_ILLEGIBLE_OR_INCOMPLETE:
		return "Cópia atendida ilegível / incompleta"
	case RECURRENCE_CANCELED:
		return "Recorrência cancelada"
	case REQUIRED_AUTHORIZATION_NOT_GRANTED:
		return "Autorização requerida não obtida"
	case RIGHT_OF_FULL_RECOURSE_FOR_FRAUD:
		return "Direito de regresso integral por fraude"
	case SALE_CANCELED:
		return "Mercadoria / serviços cancelado"
	case SERVICE_DISAGREEMENT_OR_DEFECTIVE_PRODUCT:
		return "Mercadoria / serviço com defeito ou em desacordo"
	case SERVICE_NOT_RECEIVED:
		return "Mercadoria / serviços não recebidos"
	case SPLIT_SALE:
		return "Desmembramento de venda"
	case TRANSFERS_OF_DIVERSE_RESPONSIBILITIES:
		return "Transf. de responsabilidades diversas"
	case UNQUALIFIED_CAR_RENTAL_DEBIT:
		return "Débito de aluguel de carro não qualificado"
	case USA_CARDHOLDER_DISPUTE:
		return "Contestação do portador de cartão (EUA)"
	case VISA_FRAUD_MONITORING_PROGRAM:
		return "Programa Visa de monitoramento de fraude"
	case WARNING_BULLETIN_FILE:
		return "Arquivo boletim de advertência"
	}
	return ""
}
