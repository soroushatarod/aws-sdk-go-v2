// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CertificateStatus string

// Enum values for CertificateStatus
const (
	CertificateStatusPending_validation   CertificateStatus = "PENDING_VALIDATION"
	CertificateStatusIssued               CertificateStatus = "ISSUED"
	CertificateStatusInactive             CertificateStatus = "INACTIVE"
	CertificateStatusExpired              CertificateStatus = "EXPIRED"
	CertificateStatusValidation_timed_out CertificateStatus = "VALIDATION_TIMED_OUT"
	CertificateStatusRevoked              CertificateStatus = "REVOKED"
	CertificateStatusFailed               CertificateStatus = "FAILED"
)

// Values returns all known values for CertificateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CertificateStatus) Values() []CertificateStatus {
	return []CertificateStatus{
		"PENDING_VALIDATION",
		"ISSUED",
		"INACTIVE",
		"EXPIRED",
		"VALIDATION_TIMED_OUT",
		"REVOKED",
		"FAILED",
	}
}

type CertificateTransparencyLoggingPreference string

// Enum values for CertificateTransparencyLoggingPreference
const (
	CertificateTransparencyLoggingPreferenceEnabled  CertificateTransparencyLoggingPreference = "ENABLED"
	CertificateTransparencyLoggingPreferenceDisabled CertificateTransparencyLoggingPreference = "DISABLED"
)

// Values returns all known values for CertificateTransparencyLoggingPreference.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CertificateTransparencyLoggingPreference) Values() []CertificateTransparencyLoggingPreference {
	return []CertificateTransparencyLoggingPreference{
		"ENABLED",
		"DISABLED",
	}
}

type CertificateType string

// Enum values for CertificateType
const (
	CertificateTypeImported      CertificateType = "IMPORTED"
	CertificateTypeAmazon_issued CertificateType = "AMAZON_ISSUED"
	CertificateTypePrivate       CertificateType = "PRIVATE"
)

// Values returns all known values for CertificateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CertificateType) Values() []CertificateType {
	return []CertificateType{
		"IMPORTED",
		"AMAZON_ISSUED",
		"PRIVATE",
	}
}

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusPending_validation DomainStatus = "PENDING_VALIDATION"
	DomainStatusSuccess            DomainStatus = "SUCCESS"
	DomainStatusFailed             DomainStatus = "FAILED"
)

// Values returns all known values for DomainStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"PENDING_VALIDATION",
		"SUCCESS",
		"FAILED",
	}
}

type ExtendedKeyUsageName string

// Enum values for ExtendedKeyUsageName
const (
	ExtendedKeyUsageNameTls_web_server_authentication ExtendedKeyUsageName = "TLS_WEB_SERVER_AUTHENTICATION"
	ExtendedKeyUsageNameTls_web_client_authentication ExtendedKeyUsageName = "TLS_WEB_CLIENT_AUTHENTICATION"
	ExtendedKeyUsageNameCode_signing                  ExtendedKeyUsageName = "CODE_SIGNING"
	ExtendedKeyUsageNameEmail_protection              ExtendedKeyUsageName = "EMAIL_PROTECTION"
	ExtendedKeyUsageNameTime_stamping                 ExtendedKeyUsageName = "TIME_STAMPING"
	ExtendedKeyUsageNameOcsp_signing                  ExtendedKeyUsageName = "OCSP_SIGNING"
	ExtendedKeyUsageNameIpsec_end_system              ExtendedKeyUsageName = "IPSEC_END_SYSTEM"
	ExtendedKeyUsageNameIpsec_tunnel                  ExtendedKeyUsageName = "IPSEC_TUNNEL"
	ExtendedKeyUsageNameIpsec_user                    ExtendedKeyUsageName = "IPSEC_USER"
	ExtendedKeyUsageNameAny                           ExtendedKeyUsageName = "ANY"
	ExtendedKeyUsageNameNone                          ExtendedKeyUsageName = "NONE"
	ExtendedKeyUsageNameCustom                        ExtendedKeyUsageName = "CUSTOM"
)

// Values returns all known values for ExtendedKeyUsageName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExtendedKeyUsageName) Values() []ExtendedKeyUsageName {
	return []ExtendedKeyUsageName{
		"TLS_WEB_SERVER_AUTHENTICATION",
		"TLS_WEB_CLIENT_AUTHENTICATION",
		"CODE_SIGNING",
		"EMAIL_PROTECTION",
		"TIME_STAMPING",
		"OCSP_SIGNING",
		"IPSEC_END_SYSTEM",
		"IPSEC_TUNNEL",
		"IPSEC_USER",
		"ANY",
		"NONE",
		"CUSTOM",
	}
}

type FailureReason string

// Enum values for FailureReason
const (
	FailureReasonNo_available_contacts            FailureReason = "NO_AVAILABLE_CONTACTS"
	FailureReasonAdditional_verification_required FailureReason = "ADDITIONAL_VERIFICATION_REQUIRED"
	FailureReasonDomain_not_allowed               FailureReason = "DOMAIN_NOT_ALLOWED"
	FailureReasonInvalid_public_domain            FailureReason = "INVALID_PUBLIC_DOMAIN"
	FailureReasonDomain_validation_denied         FailureReason = "DOMAIN_VALIDATION_DENIED"
	FailureReasonCaa_error                        FailureReason = "CAA_ERROR"
	FailureReasonPca_limit_exceeded               FailureReason = "PCA_LIMIT_EXCEEDED"
	FailureReasonPca_invalid_arn                  FailureReason = "PCA_INVALID_ARN"
	FailureReasonPca_invalid_state                FailureReason = "PCA_INVALID_STATE"
	FailureReasonPca_request_failed               FailureReason = "PCA_REQUEST_FAILED"
	FailureReasonPca_name_constraints_validation  FailureReason = "PCA_NAME_CONSTRAINTS_VALIDATION"
	FailureReasonPca_resource_not_found           FailureReason = "PCA_RESOURCE_NOT_FOUND"
	FailureReasonPca_invalid_args                 FailureReason = "PCA_INVALID_ARGS"
	FailureReasonPca_invalid_duration             FailureReason = "PCA_INVALID_DURATION"
	FailureReasonPca_access_denied                FailureReason = "PCA_ACCESS_DENIED"
	FailureReasonSlr_not_found                    FailureReason = "SLR_NOT_FOUND"
	FailureReasonOther                            FailureReason = "OTHER"
)

// Values returns all known values for FailureReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FailureReason) Values() []FailureReason {
	return []FailureReason{
		"NO_AVAILABLE_CONTACTS",
		"ADDITIONAL_VERIFICATION_REQUIRED",
		"DOMAIN_NOT_ALLOWED",
		"INVALID_PUBLIC_DOMAIN",
		"DOMAIN_VALIDATION_DENIED",
		"CAA_ERROR",
		"PCA_LIMIT_EXCEEDED",
		"PCA_INVALID_ARN",
		"PCA_INVALID_STATE",
		"PCA_REQUEST_FAILED",
		"PCA_NAME_CONSTRAINTS_VALIDATION",
		"PCA_RESOURCE_NOT_FOUND",
		"PCA_INVALID_ARGS",
		"PCA_INVALID_DURATION",
		"PCA_ACCESS_DENIED",
		"SLR_NOT_FOUND",
		"OTHER",
	}
}

type KeyAlgorithm string

// Enum values for KeyAlgorithm
const (
	KeyAlgorithmRsa_2048      KeyAlgorithm = "RSA_2048"
	KeyAlgorithmRsa_1024      KeyAlgorithm = "RSA_1024"
	KeyAlgorithmRsa_4096      KeyAlgorithm = "RSA_4096"
	KeyAlgorithmEc_prime256v1 KeyAlgorithm = "EC_prime256v1"
	KeyAlgorithmEc_secp384r1  KeyAlgorithm = "EC_secp384r1"
	KeyAlgorithmEc_secp521r1  KeyAlgorithm = "EC_secp521r1"
)

// Values returns all known values for KeyAlgorithm. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (KeyAlgorithm) Values() []KeyAlgorithm {
	return []KeyAlgorithm{
		"RSA_2048",
		"RSA_1024",
		"RSA_4096",
		"EC_prime256v1",
		"EC_secp384r1",
		"EC_secp521r1",
	}
}

type KeyUsageName string

// Enum values for KeyUsageName
const (
	KeyUsageNameDigital_signature   KeyUsageName = "DIGITAL_SIGNATURE"
	KeyUsageNameNon_repudation      KeyUsageName = "NON_REPUDIATION"
	KeyUsageNameKey_encipherment    KeyUsageName = "KEY_ENCIPHERMENT"
	KeyUsageNameData_encipherment   KeyUsageName = "DATA_ENCIPHERMENT"
	KeyUsageNameKey_agreement       KeyUsageName = "KEY_AGREEMENT"
	KeyUsageNameCertificate_signing KeyUsageName = "CERTIFICATE_SIGNING"
	KeyUsageNameCrl_signing         KeyUsageName = "CRL_SIGNING"
	KeyUsageNameEnchiper_only       KeyUsageName = "ENCIPHER_ONLY"
	KeyUsageNameDecipher_only       KeyUsageName = "DECIPHER_ONLY"
	KeyUsageNameAny                 KeyUsageName = "ANY"
	KeyUsageNameCustom              KeyUsageName = "CUSTOM"
)

// Values returns all known values for KeyUsageName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (KeyUsageName) Values() []KeyUsageName {
	return []KeyUsageName{
		"DIGITAL_SIGNATURE",
		"NON_REPUDIATION",
		"KEY_ENCIPHERMENT",
		"DATA_ENCIPHERMENT",
		"KEY_AGREEMENT",
		"CERTIFICATE_SIGNING",
		"CRL_SIGNING",
		"ENCIPHER_ONLY",
		"DECIPHER_ONLY",
		"ANY",
		"CUSTOM",
	}
}

type RecordType string

// Enum values for RecordType
const (
	RecordTypeCname RecordType = "CNAME"
)

// Values returns all known values for RecordType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RecordType) Values() []RecordType {
	return []RecordType{
		"CNAME",
	}
}

type RenewalEligibility string

// Enum values for RenewalEligibility
const (
	RenewalEligibilityEligible   RenewalEligibility = "ELIGIBLE"
	RenewalEligibilityIneligible RenewalEligibility = "INELIGIBLE"
)

// Values returns all known values for RenewalEligibility. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RenewalEligibility) Values() []RenewalEligibility {
	return []RenewalEligibility{
		"ELIGIBLE",
		"INELIGIBLE",
	}
}

type RenewalStatus string

// Enum values for RenewalStatus
const (
	RenewalStatusPending_auto_renewal RenewalStatus = "PENDING_AUTO_RENEWAL"
	RenewalStatusPending_validation   RenewalStatus = "PENDING_VALIDATION"
	RenewalStatusSuccess              RenewalStatus = "SUCCESS"
	RenewalStatusFailed               RenewalStatus = "FAILED"
)

// Values returns all known values for RenewalStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RenewalStatus) Values() []RenewalStatus {
	return []RenewalStatus{
		"PENDING_AUTO_RENEWAL",
		"PENDING_VALIDATION",
		"SUCCESS",
		"FAILED",
	}
}

type RevocationReason string

// Enum values for RevocationReason
const (
	RevocationReasonUnspecified            RevocationReason = "UNSPECIFIED"
	RevocationReasonKey_compromise         RevocationReason = "KEY_COMPROMISE"
	RevocationReasonCa_compromise          RevocationReason = "CA_COMPROMISE"
	RevocationReasonAffiliation_changed    RevocationReason = "AFFILIATION_CHANGED"
	RevocationReasonSuperceded             RevocationReason = "SUPERCEDED"
	RevocationReasonCessation_of_operation RevocationReason = "CESSATION_OF_OPERATION"
	RevocationReasonCertificate_hold       RevocationReason = "CERTIFICATE_HOLD"
	RevocationReasonRemove_from_crl        RevocationReason = "REMOVE_FROM_CRL"
	RevocationReasonPrivilege_withdrawn    RevocationReason = "PRIVILEGE_WITHDRAWN"
	RevocationReasonA_a_compromise         RevocationReason = "A_A_COMPROMISE"
)

// Values returns all known values for RevocationReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RevocationReason) Values() []RevocationReason {
	return []RevocationReason{
		"UNSPECIFIED",
		"KEY_COMPROMISE",
		"CA_COMPROMISE",
		"AFFILIATION_CHANGED",
		"SUPERCEDED",
		"CESSATION_OF_OPERATION",
		"CERTIFICATE_HOLD",
		"REMOVE_FROM_CRL",
		"PRIVILEGE_WITHDRAWN",
		"A_A_COMPROMISE",
	}
}

type ValidationMethod string

// Enum values for ValidationMethod
const (
	ValidationMethodEmail ValidationMethod = "EMAIL"
	ValidationMethodDns   ValidationMethod = "DNS"
)

// Values returns all known values for ValidationMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ValidationMethod) Values() []ValidationMethod {
	return []ValidationMethod{
		"EMAIL",
		"DNS",
	}
}
