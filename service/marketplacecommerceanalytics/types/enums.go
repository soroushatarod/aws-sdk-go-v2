// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DataSetType string

// Enum values for DataSetType
const (
	DataSetTypeCustomer_subscriber_hourly_monthly_subscriptions   DataSetType = "customer_subscriber_hourly_monthly_subscriptions"
	DataSetTypeCustomer_subscriber_annual_subscriptions           DataSetType = "customer_subscriber_annual_subscriptions"
	DataSetTypeDaily_business_usage_by_instance_type              DataSetType = "daily_business_usage_by_instance_type"
	DataSetTypeDaily_business_fees                                DataSetType = "daily_business_fees"
	DataSetTypeDaily_business_free_trial_conversions              DataSetType = "daily_business_free_trial_conversions"
	DataSetTypeDaily_business_new_instances                       DataSetType = "daily_business_new_instances"
	DataSetTypeDaily_business_new_product_subscribers             DataSetType = "daily_business_new_product_subscribers"
	DataSetTypeDaily_business_canceled_product_subscribers        DataSetType = "daily_business_canceled_product_subscribers"
	DataSetTypeMonthly_revenue_billing_and_revenue_data           DataSetType = "monthly_revenue_billing_and_revenue_data"
	DataSetTypeMonthly_revenue_annual_subscriptions               DataSetType = "monthly_revenue_annual_subscriptions"
	DataSetTypeMonthly_revenue_field_demonstration_usage          DataSetType = "monthly_revenue_field_demonstration_usage"
	DataSetTypeMonthly_revenue_flexible_payment_schedule          DataSetType = "monthly_revenue_flexible_payment_schedule"
	DataSetTypeDisbursed_amount_by_product                        DataSetType = "disbursed_amount_by_product"
	DataSetTypeDisbursed_amount_by_product_with_uncollected_funds DataSetType = "disbursed_amount_by_product_with_uncollected_funds"
	DataSetTypeDisbursed_amount_by_instance_hours                 DataSetType = "disbursed_amount_by_instance_hours"
	DataSetTypeDisbursed_amount_by_customer_geo                   DataSetType = "disbursed_amount_by_customer_geo"
	DataSetTypeDisbursed_amount_by_age_of_uncollected_funds       DataSetType = "disbursed_amount_by_age_of_uncollected_funds"
	DataSetTypeDisbursed_amount_by_age_of_disbursed_funds         DataSetType = "disbursed_amount_by_age_of_disbursed_funds"
	DataSetTypeDisbursed_amount_by_age_of_past_due_funds          DataSetType = "disbursed_amount_by_age_of_past_due_funds"
	DataSetTypeDisbursed_amount_by_uncollected_funds_breakdown    DataSetType = "disbursed_amount_by_uncollected_funds_breakdown"
	DataSetTypeCustomer_profile_by_industry                       DataSetType = "customer_profile_by_industry"
	DataSetTypeCustomer_profile_by_revenue                        DataSetType = "customer_profile_by_revenue"
	DataSetTypeCustomer_profile_by_geography                      DataSetType = "customer_profile_by_geography"
	DataSetTypeSales_compensation_billed_revenue                  DataSetType = "sales_compensation_billed_revenue"
	DataSetTypeUs_sales_and_use_tax_records                       DataSetType = "us_sales_and_use_tax_records"
)

// Values returns all known values for DataSetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DataSetType) Values() []DataSetType {
	return []DataSetType{
		"customer_subscriber_hourly_monthly_subscriptions",
		"customer_subscriber_annual_subscriptions",
		"daily_business_usage_by_instance_type",
		"daily_business_fees",
		"daily_business_free_trial_conversions",
		"daily_business_new_instances",
		"daily_business_new_product_subscribers",
		"daily_business_canceled_product_subscribers",
		"monthly_revenue_billing_and_revenue_data",
		"monthly_revenue_annual_subscriptions",
		"monthly_revenue_field_demonstration_usage",
		"monthly_revenue_flexible_payment_schedule",
		"disbursed_amount_by_product",
		"disbursed_amount_by_product_with_uncollected_funds",
		"disbursed_amount_by_instance_hours",
		"disbursed_amount_by_customer_geo",
		"disbursed_amount_by_age_of_uncollected_funds",
		"disbursed_amount_by_age_of_disbursed_funds",
		"disbursed_amount_by_age_of_past_due_funds",
		"disbursed_amount_by_uncollected_funds_breakdown",
		"customer_profile_by_industry",
		"customer_profile_by_revenue",
		"customer_profile_by_geography",
		"sales_compensation_billed_revenue",
		"us_sales_and_use_tax_records",
	}
}

type SupportDataSetType string

// Enum values for SupportDataSetType
const (
	SupportDataSetTypeCustomer_support_contacts_data      SupportDataSetType = "customer_support_contacts_data"
	SupportDataSetTypeTest_customer_support_contacts_data SupportDataSetType = "test_customer_support_contacts_data"
)

// Values returns all known values for SupportDataSetType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SupportDataSetType) Values() []SupportDataSetType {
	return []SupportDataSetType{
		"customer_support_contacts_data",
		"test_customer_support_contacts_data",
	}
}
