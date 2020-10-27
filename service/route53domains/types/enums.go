// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ContactType string

// Enum values for ContactType
const (
	ContactTypePerson      ContactType = "PERSON"
	ContactTypeCompany     ContactType = "COMPANY"
	ContactTypeAssociation ContactType = "ASSOCIATION"
	ContactTypePublic_body ContactType = "PUBLIC_BODY"
	ContactTypeReseller    ContactType = "RESELLER"
)

// Values returns all known values for ContactType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ContactType) Values() []ContactType {
	return []ContactType{
		"PERSON",
		"COMPANY",
		"ASSOCIATION",
		"PUBLIC_BODY",
		"RESELLER",
	}
}

type CountryCode string

// Enum values for CountryCode
const (
	CountryCodeAd CountryCode = "AD"
	CountryCodeAe CountryCode = "AE"
	CountryCodeAf CountryCode = "AF"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAl CountryCode = "AL"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAn CountryCode = "AN"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAq CountryCode = "AQ"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAs CountryCode = "AS"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBi CountryCode = "BI"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBl CountryCode = "BL"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBr CountryCode = "BR"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeCa CountryCode = "CA"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCf CountryCode = "CF"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCh CountryCode = "CH"
	CountryCodeCi CountryCode = "CI"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCm CountryCode = "CM"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCo CountryCode = "CO"
	CountryCodeCr CountryCode = "CR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCv CountryCode = "CV"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeDe CountryCode = "DE"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEe CountryCode = "EE"
	CountryCodeEg CountryCode = "EG"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEs CountryCode = "ES"
	CountryCodeEt CountryCode = "ET"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFm CountryCode = "FM"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGb CountryCode = "GB"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGe CountryCode = "GE"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGu CountryCode = "GU"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHr CountryCode = "HR"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHu CountryCode = "HU"
	CountryCodeId CountryCode = "ID"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIn CountryCode = "IN"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJo CountryCode = "JO"
	CountryCodeJp CountryCode = "JP"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKg CountryCode = "KG"
	CountryCodeKh CountryCode = "KH"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKm CountryCode = "KM"
	CountryCodeKn CountryCode = "KN"
	CountryCodeKp CountryCode = "KP"
	CountryCodeKr CountryCode = "KR"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKy CountryCode = "KY"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLc CountryCode = "LC"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLk CountryCode = "LK"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLy CountryCode = "LY"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMf CountryCode = "MF"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMh CountryCode = "MH"
	CountryCodeMk CountryCode = "MK"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMm CountryCode = "MM"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMp CountryCode = "MP"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMu CountryCode = "MU"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMx CountryCode = "MX"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNl CountryCode = "NL"
	CountryCodeNo CountryCode = "NO"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeOm CountryCode = "OM"
	CountryCodePa CountryCode = "PA"
	CountryCodePe CountryCode = "PE"
	CountryCodePf CountryCode = "PF"
	CountryCodePg CountryCode = "PG"
	CountryCodePh CountryCode = "PH"
	CountryCodePk CountryCode = "PK"
	CountryCodePl CountryCode = "PL"
	CountryCodePm CountryCode = "PM"
	CountryCodePn CountryCode = "PN"
	CountryCodePr CountryCode = "PR"
	CountryCodePt CountryCode = "PT"
	CountryCodePw CountryCode = "PW"
	CountryCodePy CountryCode = "PY"
	CountryCodeQa CountryCode = "QA"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRs CountryCode = "RS"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSe CountryCode = "SE"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSh CountryCode = "SH"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSn CountryCode = "SN"
	CountryCodeSo CountryCode = "SO"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSv CountryCode = "SV"
	CountryCodeSy CountryCode = "SY"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTd CountryCode = "TD"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTv CountryCode = "TV"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeUa CountryCode = "UA"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUs CountryCode = "US"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVa CountryCode = "VA"
	CountryCodeVc CountryCode = "VC"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVg CountryCode = "VG"
	CountryCodeVi CountryCode = "VI"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVu CountryCode = "VU"
	CountryCodeWf CountryCode = "WF"
	CountryCodeWs CountryCode = "WS"
	CountryCodeYe CountryCode = "YE"
	CountryCodeYt CountryCode = "YT"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
)

// Values returns all known values for CountryCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CountryCode) Values() []CountryCode {
	return []CountryCode{
		"AD",
		"AE",
		"AF",
		"AG",
		"AI",
		"AL",
		"AM",
		"AN",
		"AO",
		"AQ",
		"AR",
		"AS",
		"AT",
		"AU",
		"AW",
		"AZ",
		"BA",
		"BB",
		"BD",
		"BE",
		"BF",
		"BG",
		"BH",
		"BI",
		"BJ",
		"BL",
		"BM",
		"BN",
		"BO",
		"BR",
		"BS",
		"BT",
		"BW",
		"BY",
		"BZ",
		"CA",
		"CC",
		"CD",
		"CF",
		"CG",
		"CH",
		"CI",
		"CK",
		"CL",
		"CM",
		"CN",
		"CO",
		"CR",
		"CU",
		"CV",
		"CX",
		"CY",
		"CZ",
		"DE",
		"DJ",
		"DK",
		"DM",
		"DO",
		"DZ",
		"EC",
		"EE",
		"EG",
		"ER",
		"ES",
		"ET",
		"FI",
		"FJ",
		"FK",
		"FM",
		"FO",
		"FR",
		"GA",
		"GB",
		"GD",
		"GE",
		"GH",
		"GI",
		"GL",
		"GM",
		"GN",
		"GQ",
		"GR",
		"GT",
		"GU",
		"GW",
		"GY",
		"HK",
		"HN",
		"HR",
		"HT",
		"HU",
		"ID",
		"IE",
		"IL",
		"IM",
		"IN",
		"IQ",
		"IR",
		"IS",
		"IT",
		"JM",
		"JO",
		"JP",
		"KE",
		"KG",
		"KH",
		"KI",
		"KM",
		"KN",
		"KP",
		"KR",
		"KW",
		"KY",
		"KZ",
		"LA",
		"LB",
		"LC",
		"LI",
		"LK",
		"LR",
		"LS",
		"LT",
		"LU",
		"LV",
		"LY",
		"MA",
		"MC",
		"MD",
		"ME",
		"MF",
		"MG",
		"MH",
		"MK",
		"ML",
		"MM",
		"MN",
		"MO",
		"MP",
		"MR",
		"MS",
		"MT",
		"MU",
		"MV",
		"MW",
		"MX",
		"MY",
		"MZ",
		"NA",
		"NC",
		"NE",
		"NG",
		"NI",
		"NL",
		"NO",
		"NP",
		"NR",
		"NU",
		"NZ",
		"OM",
		"PA",
		"PE",
		"PF",
		"PG",
		"PH",
		"PK",
		"PL",
		"PM",
		"PN",
		"PR",
		"PT",
		"PW",
		"PY",
		"QA",
		"RO",
		"RS",
		"RU",
		"RW",
		"SA",
		"SB",
		"SC",
		"SD",
		"SE",
		"SG",
		"SH",
		"SI",
		"SK",
		"SL",
		"SM",
		"SN",
		"SO",
		"SR",
		"ST",
		"SV",
		"SY",
		"SZ",
		"TC",
		"TD",
		"TG",
		"TH",
		"TJ",
		"TK",
		"TL",
		"TM",
		"TN",
		"TO",
		"TR",
		"TT",
		"TV",
		"TW",
		"TZ",
		"UA",
		"UG",
		"US",
		"UY",
		"UZ",
		"VA",
		"VC",
		"VE",
		"VG",
		"VI",
		"VN",
		"VU",
		"WF",
		"WS",
		"YE",
		"YT",
		"ZA",
		"ZM",
		"ZW",
	}
}

type DomainAvailability string

// Enum values for DomainAvailability
const (
	DomainAvailabilityAvailable              DomainAvailability = "AVAILABLE"
	DomainAvailabilityAvailable_reserved     DomainAvailability = "AVAILABLE_RESERVED"
	DomainAvailabilityAvailable_preorder     DomainAvailability = "AVAILABLE_PREORDER"
	DomainAvailabilityUnavailable            DomainAvailability = "UNAVAILABLE"
	DomainAvailabilityUnavailable_premium    DomainAvailability = "UNAVAILABLE_PREMIUM"
	DomainAvailabilityUnavailable_restricted DomainAvailability = "UNAVAILABLE_RESTRICTED"
	DomainAvailabilityReserved               DomainAvailability = "RESERVED"
	DomainAvailabilityDont_know              DomainAvailability = "DONT_KNOW"
)

// Values returns all known values for DomainAvailability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DomainAvailability) Values() []DomainAvailability {
	return []DomainAvailability{
		"AVAILABLE",
		"AVAILABLE_RESERVED",
		"AVAILABLE_PREORDER",
		"UNAVAILABLE",
		"UNAVAILABLE_PREMIUM",
		"UNAVAILABLE_RESTRICTED",
		"RESERVED",
		"DONT_KNOW",
	}
}

type ExtraParamName string

// Enum values for ExtraParamName
const (
	ExtraParamNameDuns_number                      ExtraParamName = "DUNS_NUMBER"
	ExtraParamNameBrand_number                     ExtraParamName = "BRAND_NUMBER"
	ExtraParamNameBirth_department                 ExtraParamName = "BIRTH_DEPARTMENT"
	ExtraParamNameBirth_date_in_yyyy_mm_dd         ExtraParamName = "BIRTH_DATE_IN_YYYY_MM_DD"
	ExtraParamNameBirth_country                    ExtraParamName = "BIRTH_COUNTRY"
	ExtraParamNameBirth_city                       ExtraParamName = "BIRTH_CITY"
	ExtraParamNameDocument_number                  ExtraParamName = "DOCUMENT_NUMBER"
	ExtraParamNameAu_id_number                     ExtraParamName = "AU_ID_NUMBER"
	ExtraParamNameAu_id_type                       ExtraParamName = "AU_ID_TYPE"
	ExtraParamNameCa_legal_type                    ExtraParamName = "CA_LEGAL_TYPE"
	ExtraParamNameCa_business_entity_type          ExtraParamName = "CA_BUSINESS_ENTITY_TYPE"
	ExtraParamNameCa_legal_representative          ExtraParamName = "CA_LEGAL_REPRESENTATIVE"
	ExtraParamNameCa_legal_representative_capacity ExtraParamName = "CA_LEGAL_REPRESENTATIVE_CAPACITY"
	ExtraParamNameEs_identification                ExtraParamName = "ES_IDENTIFICATION"
	ExtraParamNameEs_identification_type           ExtraParamName = "ES_IDENTIFICATION_TYPE"
	ExtraParamNameEs_legal_form                    ExtraParamName = "ES_LEGAL_FORM"
	ExtraParamNameFi_business_number               ExtraParamName = "FI_BUSINESS_NUMBER"
	ExtraParamNameOnwer_fi_id_number               ExtraParamName = "FI_ID_NUMBER"
	ExtraParamNameFi_nationality                   ExtraParamName = "FI_NATIONALITY"
	ExtraParamNameFi_organization_type             ExtraParamName = "FI_ORGANIZATION_TYPE"
	ExtraParamNameIt_nationality                   ExtraParamName = "IT_NATIONALITY"
	ExtraParamNameIt_pin                           ExtraParamName = "IT_PIN"
	ExtraParamNameIt_registrant_entity_type        ExtraParamName = "IT_REGISTRANT_ENTITY_TYPE"
	ExtraParamNameRu_passport_data                 ExtraParamName = "RU_PASSPORT_DATA"
	ExtraParamNameSe_id_number                     ExtraParamName = "SE_ID_NUMBER"
	ExtraParamNameSg_id_number                     ExtraParamName = "SG_ID_NUMBER"
	ExtraParamNameVat_number                       ExtraParamName = "VAT_NUMBER"
	ExtraParamNameUk_contact_type                  ExtraParamName = "UK_CONTACT_TYPE"
	ExtraParamNameUk_company_number                ExtraParamName = "UK_COMPANY_NUMBER"
)

// Values returns all known values for ExtraParamName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExtraParamName) Values() []ExtraParamName {
	return []ExtraParamName{
		"DUNS_NUMBER",
		"BRAND_NUMBER",
		"BIRTH_DEPARTMENT",
		"BIRTH_DATE_IN_YYYY_MM_DD",
		"BIRTH_COUNTRY",
		"BIRTH_CITY",
		"DOCUMENT_NUMBER",
		"AU_ID_NUMBER",
		"AU_ID_TYPE",
		"CA_LEGAL_TYPE",
		"CA_BUSINESS_ENTITY_TYPE",
		"CA_LEGAL_REPRESENTATIVE",
		"CA_LEGAL_REPRESENTATIVE_CAPACITY",
		"ES_IDENTIFICATION",
		"ES_IDENTIFICATION_TYPE",
		"ES_LEGAL_FORM",
		"FI_BUSINESS_NUMBER",
		"FI_ID_NUMBER",
		"FI_NATIONALITY",
		"FI_ORGANIZATION_TYPE",
		"IT_NATIONALITY",
		"IT_PIN",
		"IT_REGISTRANT_ENTITY_TYPE",
		"RU_PASSPORT_DATA",
		"SE_ID_NUMBER",
		"SG_ID_NUMBER",
		"VAT_NUMBER",
		"UK_CONTACT_TYPE",
		"UK_COMPANY_NUMBER",
	}
}

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusSubmitted   OperationStatus = "SUBMITTED"
	OperationStatusIn_progress OperationStatus = "IN_PROGRESS"
	OperationStatusError       OperationStatus = "ERROR"
	OperationStatusSuccessful  OperationStatus = "SUCCESSFUL"
	OperationStatusFailed      OperationStatus = "FAILED"
)

// Values returns all known values for OperationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"ERROR",
		"SUCCESSFUL",
		"FAILED",
	}
}

type OperationType string

// Enum values for OperationType
const (
	OperationTypeRegister_domain              OperationType = "REGISTER_DOMAIN"
	OperationTypeDelete_domain                OperationType = "DELETE_DOMAIN"
	OperationTypeTransfer_in_domain           OperationType = "TRANSFER_IN_DOMAIN"
	OperationTypeUpdate_domain_contact        OperationType = "UPDATE_DOMAIN_CONTACT"
	OperationTypeUpdate_nameserver            OperationType = "UPDATE_NAMESERVER"
	OperationTypeChange_privacy_protection    OperationType = "CHANGE_PRIVACY_PROTECTION"
	OperationTypeDomain_lock                  OperationType = "DOMAIN_LOCK"
	OperationTypeEnable_autorenew             OperationType = "ENABLE_AUTORENEW"
	OperationTypeDisable_autorenew            OperationType = "DISABLE_AUTORENEW"
	OperationTypeAdd_dnssec                   OperationType = "ADD_DNSSEC"
	OperationTypeRemove_dnssec                OperationType = "REMOVE_DNSSEC"
	OperationTypeExpire_domain                OperationType = "EXPIRE_DOMAIN"
	OperationTypeTransfer_out_domain          OperationType = "TRANSFER_OUT_DOMAIN"
	OperationTypeChange_domain_owner          OperationType = "CHANGE_DOMAIN_OWNER"
	OperationTypeRenew_domain                 OperationType = "RENEW_DOMAIN"
	OperationTypePush_domain                  OperationType = "PUSH_DOMAIN"
	OperationTypeInternal_transfer_out_domain OperationType = "INTERNAL_TRANSFER_OUT_DOMAIN"
	OperationTypeInternal_transfer_in_domain  OperationType = "INTERNAL_TRANSFER_IN_DOMAIN"
)

// Values returns all known values for OperationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationType) Values() []OperationType {
	return []OperationType{
		"REGISTER_DOMAIN",
		"DELETE_DOMAIN",
		"TRANSFER_IN_DOMAIN",
		"UPDATE_DOMAIN_CONTACT",
		"UPDATE_NAMESERVER",
		"CHANGE_PRIVACY_PROTECTION",
		"DOMAIN_LOCK",
		"ENABLE_AUTORENEW",
		"DISABLE_AUTORENEW",
		"ADD_DNSSEC",
		"REMOVE_DNSSEC",
		"EXPIRE_DOMAIN",
		"TRANSFER_OUT_DOMAIN",
		"CHANGE_DOMAIN_OWNER",
		"RENEW_DOMAIN",
		"PUSH_DOMAIN",
		"INTERNAL_TRANSFER_OUT_DOMAIN",
		"INTERNAL_TRANSFER_IN_DOMAIN",
	}
}

type ReachabilityStatus string

// Enum values for ReachabilityStatus
const (
	ReachabilityStatusPending ReachabilityStatus = "PENDING"
	ReachabilityStatusDone    ReachabilityStatus = "DONE"
	ReachabilityStatusExpired ReachabilityStatus = "EXPIRED"
)

// Values returns all known values for ReachabilityStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReachabilityStatus) Values() []ReachabilityStatus {
	return []ReachabilityStatus{
		"PENDING",
		"DONE",
		"EXPIRED",
	}
}

type Transferable string

// Enum values for Transferable
const (
	TransferableTransferable   Transferable = "TRANSFERABLE"
	TransferableUntransferable Transferable = "UNTRANSFERABLE"
	TransferableDont_know      Transferable = "DONT_KNOW"
)

// Values returns all known values for Transferable. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Transferable) Values() []Transferable {
	return []Transferable{
		"TRANSFERABLE",
		"UNTRANSFERABLE",
		"DONT_KNOW",
	}
}
