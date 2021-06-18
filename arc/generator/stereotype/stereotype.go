package stereotype

type secretKey string

const (
	// e.g. by bash SET X=my_var. Value is either nil or a string.
	kEnvironmentVariable secretKey = "kEnvironmentVariable"

	// e.g. by calling ./myprogram -X=my_var. Value is either nil or a string.
	kProgramFlagVariable secretKey = "kProgramFlagVariable"

	// denotes if a struct is used for external configuration (env and program flags). Either nil, true or false.
	kConfiguration secretKey = "kConfiguration"

	// denotes if a struct is used for database configuration (env and program flags). Either nil, true or false.
	kDBConfiguration secretKey = "kDBConfiguration"

	// denotes if is mysql related.
	kMySQLRelated secretKey = "kMySQLRelated"

	// denotes a mysql table name
	kSQLTableName secretKey = "kSQLTableName"

	// denotes a mysql column name
	kSQLColumnName secretKey = "kSQLColumnName"

	// declares the default sort order
	kSQLDefaultOrder secretKey = "kSQLDefaultOrder"

	// kCMDPkg declares a package as a main package entry point. Package must be a "main" package.
	kCMDPkg secretKey = "kCMDPkg"

	// kModuleDocs declares the root of all available documentations about a module.
	kModuleDocs secretKey = "kModuleDocs"

	// kModShortName is attached to a module and contains a (hopefully global) short and
	// memorable identifier for the entire module.
	kModShortName secretKey = "kModShortName"

	// kPrjShortName is attached to a module and contains a short and
	// memorable identifier for the entire project (containing this and other modules).
	kPrjShortName secretKey = "kPrjShortName"
)
