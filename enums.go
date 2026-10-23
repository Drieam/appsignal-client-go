package appsignal

// LogSourceFormat is the log line format of a log source. It maps onto the
// SourceFormatEnum GraphQL enum.
type LogSourceFormat string

const (
	LogSourceFormatPlaintext  LogSourceFormat = "PLAINTEXT"
	LogSourceFormatLogfmt     LogSourceFormat = "LOGFMT"
	LogSourceFormatJSON       LogSourceFormat = "JSON"
	LogSourceFormatAutodetect LogSourceFormat = "AUTODETECT"
)

// LogTriggerActionType is the action a log trigger performs when it matches.
type LogTriggerActionType string

const (
	ActionTypeTrigger LogTriggerActionType = "TRIGGER"
	ActionTypeFilter  LogTriggerActionType = "FILTER"
	ActionTypeMetrics LogTriggerActionType = "METRICS"
)

// LogTriggerNotificationOption controls how often a log trigger notifies.
type LogTriggerNotificationOption string

const (
	NotificationOptionAlways          LogTriggerNotificationOption = "ALWAYS"
	NotificationOptionNever           LogTriggerNotificationOption = "NEVER"
	NotificationOptionFirstInDeploy   LogTriggerNotificationOption = "FIRST_IN_DEPLOY"
	NotificationOptionFirstAfterClose LogTriggerNotificationOption = "FIRST_AFTER_CLOSE"
	NotificationOptionNthInHour       LogTriggerNotificationOption = "NTH_IN_HOUR"
	NotificationOptionNthInDay        LogTriggerNotificationOption = "NTH_IN_DAY"
)
