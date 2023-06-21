package simple

type TimeRange struct {
	Start string
	End   string
}

var (
	TimeClose = []TimeRange{
		{
			Start: "09:00",
			End:   "10:00",
		},
		{
			Start: "15:00",
			End:   "16:00",
		},
	}
)

func IsTimeClose(time string) bool {
	for _, tr := range TimeClose {
		if tr.Start <= time && time <= tr.End {
			return true
		}
	}
	return false
}
