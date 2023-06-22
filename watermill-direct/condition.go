package simple

type TimeRange struct {
	Start string
	End   string
}

var (
	TimeClose = []TimeRange{
		{
			Start: "22:00",
			End:   "06:00",
		},
	}

	AcceptedTopic = "my-exchange"
)

func IsTimeClose(time string) bool {
	for _, tr := range TimeClose {
		if tr.Start <= tr.End && tr.Start <= time && time <= tr.End {
			return true
		}
		if tr.Start > tr.End && (tr.Start <= time || time <= tr.End) {
			return true
		}
	}
	return false
}
