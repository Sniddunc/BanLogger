package banlogger

// Record represents all infractions a player received
type Record struct {
	PlayerID string
	Warnings []Warning
	Kicks    []Kick
	Bans     []Ban
}

// NumericRecord Represents the number of offenses a player received.NumericRecord.
// Unlike Record, this only holds counts and nothing else.
type NumericRecord struct {
	PlayerID     string
	Total        int
	WarningCount int
	KickCount    int
	BanCount     int
}

// StatService facilitates stat data exchange/manipulation between packages
type StatService interface {
	GetWarningCount(string) (int, error)
	GetKickCount(string) (int, error)
	GetBanCount(string) (int, error)
	GetTotalWarningCount() (int, error)
	GetTotalKickCount() (int, error)
	GetTotalBanCount() (int, error)
	GetRecord(string) (Record, error)
	GetTopOffender() (NumericRecord, error)
}
