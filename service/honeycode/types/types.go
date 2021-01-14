// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// An object that represents a single cell in a table.
type Cell struct {

	// The format of the cell. If this field is empty, then the format is either not
	// specified in the workbook or the format is set to AUTO.
	Format Format

	// The formatted value of the cell. This is the value that you see displayed in the
	// cell in the UI. Note that the formatted value of a cell is always represented as
	// a string irrespective of the data that is stored in the cell. For example, if a
	// cell contains a date, the formatted value of the cell is the string
	// representation of the formatted date being shown in the cell in the UI. See
	// details in the rawValue field below for how cells of different formats will have
	// different raw and formatted values.
	FormattedValue *string

	// The formula contained in the cell. This field is empty if a cell does not have a
	// formula.
	Formula *string

	// The raw value of the data contained in the cell. The raw value depends on the
	// format of the data in the cell. However the attribute in the API return value is
	// always a string containing the raw value. Cells with format DATE, DATE_TIME or
	// TIME have the raw value as a floating point number where the whole number
	// represents the number of days since 1/1/1900 and the fractional part represents
	// the fraction of the day since midnight. For example, a cell with date 11/3/2020
	// has the raw value "44138". A cell with the time 9:00 AM has the raw value
	// "0.375" and a cell with date/time value of 11/3/2020 9:00 AM has the raw value
	// "44138.375". Notice that even though the raw value is a number in all three
	// cases, it is still represented as a string. Cells with format NUMBER, CURRENCY,
	// PERCENTAGE and ACCOUNTING have the raw value of the data as the number
	// representing the data being displayed. For example, the number 1.325 with two
	// decimal places in the format will have it's raw value as "1.325" and formatted
	// value as "1.33". A currency value for $10 will have the raw value as "10" and
	// formatted value as "$10.00". A value representing 20% with two decimal places in
	// the format will have its raw value as "0.2" and the formatted value as "20.00%".
	// An accounting value of -$25 will have "-25" as the raw value and "$ (25.00)" as
	// the formatted value. Cells with format TEXT will have the raw text as the raw
	// value. For example, a cell with text "John Smith" will have "John Smith" as both
	// the raw value and the formatted value. Cells with format CONTACT will have the
	// name of the contact as a formatted value and the email address of the contact as
	// the raw value. For example, a contact for John Smith will have "John Smith" as
	// the formatted value and "john.smith@example.com" as the raw value. Cells with
	// format ROWLINK (aka picklist) will have the first column of the linked row as
	// the formatted value and the row id of the linked row as the raw value. For
	// example, a cell containing a picklist to a table that displays task status might
	// have "Completed" as the formatted value and
	// "row:dfcefaee-5b37-4355-8f28-40c3e4ff5dd4/ca432b2f-b8eb-431d-9fb5-cbe0342f9f03"
	// as the raw value. Cells with format AUTO or cells without any format that are
	// auto-detected as one of the formats above will contain the raw and formatted
	// values as mentioned above, based on the auto-detected formats. If there is no
	// auto-detected format, the raw and formatted values will be the same as the data
	// in the cell.
	RawValue *string
}

// CellInput object contains the data needed to create or update cells in a table.
type CellInput struct {

	// Fact represents the data that is entered into a cell. This data can be free text
	// or a formula. Formulas need to start with the equals (=) sign.
	Fact *string
}

// Metadata for column in the table.
type ColumnMetadata struct {

	// The format of the column.
	//
	// This member is required.
	Format Format

	// The name of the column.
	//
	// This member is required.
	Name *string
}

// Data needed to create a single row in a table as part of the
// BatchCreateTableRows request.
type CreateRowData struct {

	// An external identifier that represents the single row that is being created as
	// part of the BatchCreateTableRows request. This can be any string that you can
	// use to identify the row in the request. The BatchCreateTableRows API puts the
	// batch item id in the results to allow you to link data in the request to data in
	// the results.
	//
	// This member is required.
	BatchItemId *string

	// A map representing the cells to create in the new row. The key is the column id
	// of the cell and the value is the CellInput object that represents the data to
	// set in that cell.
	//
	// This member is required.
	CellsToCreate map[string]CellInput
}

// The data in a particular data cell defined on the screen.
type DataItem struct {

	// The formatted value of the data. e.g. John Smith.
	FormattedValue *string

	// The overrideFormat is optional and is specified only if a particular row of data
	// has a different format for the data than the default format defined on the
	// screen or the table.
	OverrideFormat Format

	// The raw value of the data. e.g. jsmith@example.com
	RawValue *string
}

// An object that contains the options relating to parsing delimited text as part
// of an import request.
type DelimitedTextImportOptions struct {

	// The delimiter to use for separating columns in a single row of the input.
	//
	// This member is required.
	Delimiter *string

	// The encoding of the data in the input file.
	DataCharacterEncoding ImportDataCharacterEncoding

	// Indicates whether the input file has a header row at the top containing the
	// column names.
	HasHeaderRow bool

	// A parameter to indicate whether empty rows should be ignored or be included in
	// the import.
	IgnoreEmptyRows bool
}

// An object that contains the options relating to the destination of the import
// request.
type DestinationOptions struct {

	// A map of the column id to the import properties for each column.
	ColumnMap map[string]SourceDataColumnProperties
}

// A single item in a batch that failed to perform the intended action because of
// an error preventing it from succeeding.
type FailedBatchItem struct {

	// The error message that indicates why the batch item failed.
	//
	// This member is required.
	ErrorMessage *string

	// The id of the batch item that failed. This is the batch item id for the
	// BatchCreateTableRows and BatchUpsertTableRows operations and the row id for the
	// BatchUpdateTableRows and BatchDeleteTableRows operations.
	//
	// This member is required.
	Id *string
}

// An object that represents a filter formula along with the id of the context row
// under which the filter function needs to evaluate.
type Filter struct {

	// A formula representing a filter function that returns zero or more matching rows
	// from a table. Valid formulas in this field return a list of rows from a table.
	// The most common ways of writing a formula to return a list of rows are to use
	// the FindRow() or Filter() functions. Any other formula that returns zero or more
	// rows is also acceptable. For example, you can use a formula that points to a
	// cell that contains a filter function.
	//
	// This member is required.
	Formula *string

	// The optional contextRowId attribute can be used to specify the row id of the
	// context row if the filter formula contains unqualified references to table
	// columns and needs a context row to evaluate them successfully.
	ContextRowId *string
}

// An object that has details about the source of the data that was submitted for
// import.
type ImportDataSource struct {

	// The configuration parameters for the data source of the import
	//
	// This member is required.
	DataSourceConfig *ImportDataSourceConfig
}

// An object that contains the configuration parameters for the data source of an
// import request.
type ImportDataSourceConfig struct {

	// The URL from which source data will be downloaded for the import request.
	DataSourceUrl *string
}

// An object that contains the attributes of the submitter of the import job.
type ImportJobSubmitter struct {

	// The email id of the submitter of the import job, if available.
	Email *string

	// The AWS user ARN of the submitter of the import job, if available.
	UserArn *string
}

// An object that contains the options specified by the sumitter of the import
// request.
type ImportOptions struct {

	// Options relating to parsing delimited text. Required if dataFormat is
	// DELIMITED_TEXT.
	DelimitedTextOptions *DelimitedTextImportOptions

	// Options relating to the destination of the import request.
	DestinationOptions *DestinationOptions
}

// A single row in the ResultSet.
type ResultRow struct {

	// List of all the data cells in a row.
	//
	// This member is required.
	DataItems []DataItem

	// The ID for a particular row.
	RowId *string
}

// ResultSet contains the results of the request for a single block or list defined
// on the screen.
type ResultSet struct {

	// List of headers for all the data cells in the block. The header identifies the
	// name and default format of the data cell. Data cells appear in the same order in
	// all rows as defined in the header. The names and formats are not repeated in the
	// rows. If a particular row does not have a value for a data cell, a blank value
	// is used. For example, a task list that displays the task name, due date and
	// assigned person might have headers [ { "name": "Task Name"}, {"name": "Due
	// Date", "format": "DATE"}, {"name": "Assigned", "format": "CONTACT"} ]. Every row
	// in the result will have the task name as the first item, due date as the second
	// item and assigned person as the third item. If a particular task does not have a
	// due date, that row will still have a blank value in the second element and the
	// assigned person will still be in the third element.
	//
	// This member is required.
	Headers []ColumnMetadata

	// List of rows returned by the request. Each row has a row Id and a list of data
	// cells in that row. The data cells will be present in the same order as they are
	// defined in the header.
	//
	// This member is required.
	Rows []ResultRow
}

// An object that contains the properties for importing data to a specific column
// in a table.
type SourceDataColumnProperties struct {

	// The index of the column in the input file.
	ColumnIndex int32
}

// An object representing the properties of a table in a workbook.
type Table struct {

	// The id of the table.
	TableId *string

	// The name of the table.
	TableName *string
}

// An object that contains attributes about a single column in a table
type TableColumn struct {

	// The column level format that is applied in the table. An empty value in this
	// field means that the column format is the default value 'AUTO'.
	Format Format

	// The id of the column in the table.
	TableColumnId *string

	// The name of the column in the table.
	TableColumnName *string
}

// The metadata associated with the table data import job that was submitted.
type TableDataImportJobMetadata struct {

	// The source of the data that was submitted for import.
	//
	// This member is required.
	DataSource *ImportDataSource

	// The options that was specified at the time of submitting the import request.
	//
	// This member is required.
	ImportOptions *ImportOptions

	// The timestamp when the job was submitted for import.
	//
	// This member is required.
	SubmitTime *time.Time

	// Details about the submitter of the import request.
	//
	// This member is required.
	Submitter *ImportJobSubmitter
}

// An object that contains attributes about a single row in a table
type TableRow struct {

	// A list of cells in the table row. The cells appear in the same order as the
	// columns of the table.
	//
	// This member is required.
	Cells []Cell

	// The id of the row in the table.
	//
	// This member is required.
	RowId *string
}

// Data needed to create a single row in a table as part of the
// BatchCreateTableRows request.
type UpdateRowData struct {

	// A map representing the cells to update in the given row. The key is the column
	// id of the cell and the value is the CellInput object that represents the data to
	// set in that cell.
	//
	// This member is required.
	CellsToUpdate map[string]CellInput

	// The id of the row that needs to be updated.
	//
	// This member is required.
	RowId *string
}

// Data needed to upsert rows in a table as part of a single item in the
// BatchUpsertTableRows request.
type UpsertRowData struct {

	// An external identifier that represents a single item in the request that is
	// being upserted as part of the BatchUpsertTableRows request. This can be any
	// string that you can use to identify the item in the request. The
	// BatchUpsertTableRows API puts the batch item id in the results to allow you to
	// link data in the request to data in the results.
	//
	// This member is required.
	BatchItemId *string

	// A map representing the cells to update for the matching rows or an appended row.
	// The key is the column id of the cell and the value is the CellInput object that
	// represents the data to set in that cell.
	//
	// This member is required.
	CellsToUpdate map[string]CellInput

	// The filter formula to use to find existing matching rows to update. The formula
	// needs to return zero or more rows. If the formula returns 0 rows, then a new row
	// will be appended in the target table. If the formula returns one or more rows,
	// then the returned rows will be updated. Note that the filter formula needs to
	// return rows from the target table for the upsert operation to succeed. If the
	// filter formula has a syntax error or it doesn't evaluate to zero or more rows in
	// the target table for any one item in the input list, then the entire
	// BatchUpsertTableRows request fails and no updates are made to the table.
	//
	// This member is required.
	Filter *Filter
}

// An object that represents the result of a single upsert row request.
type UpsertRowsResult struct {

	// The list of row ids that were changed as part of an upsert row operation. If the
	// upsert resulted in an update, this list could potentially contain multiple rows
	// that matched the filter and hence got updated. If the upsert resulted in an
	// append, this list would only have the single row that was appended.
	//
	// This member is required.
	RowIds []string

	// The result of the upsert action.
	//
	// This member is required.
	UpsertAction UpsertAction
}

// The input variables to the app to be used by the InvokeScreenAutomation action
// request.
type VariableValue struct {

	// Raw value of the variable.
	//
	// This member is required.
	RawValue *string
}