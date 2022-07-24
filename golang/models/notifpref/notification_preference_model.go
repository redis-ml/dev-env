package notifpref

const (
	TableName = "NotifPref"

	ColOwner    = "Owner"
	ColPrefType = "PrefType"

	// Instrument Preferences
	ColInstrPref = "InstrPref"
	// Shared columes used in this row.
	// - ColTHld
	// - ColUpdatedAt

	// Shared columnes used by any row.
	ColThld      = "Thld"
	ColUpdatedAt = "UpdatedAt"
)

const (
	// PrefType enums.
	PrefTypePriceAbove = "PriceAbove"
)
