package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)
const (
	//General
	Startup         SubCategory = ""
	ExternalService SubCategory = ""
	//Postgres
	Select   SubCategory = ""
	Rollback SubCategory = ""
	Update   SubCategory = ""
	Delete   SubCategory = ""
	Insert   SubCategory = ""
	//Internal
	Api                 SubCategory = ""
	HashPassword        SubCategory = ""
	DefaultRoleNotFound SubCategory = ""
	// validation
	MobileValidation  SubCategory = "MobileValidation"
	CustomValidation  SubCategory = "CustomValidation"
	PasswordValidator SubCategory = "PasswordValidator"
)
const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
)
