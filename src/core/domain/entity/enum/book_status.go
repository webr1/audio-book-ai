package enum

type BookStatus string

const (
	BookStatusPending    BookStatus = "pending"
	BookStatusProcessing BookStatus = "processing"
	BookStatusDone       BookStatus = "done"
	BookStatusFailed     BookStatus = "failed"
)

func (s BookStatus) Display() string {
	switch s {
	case BookStatusPending:
		return "Pending"
	case BookStatusProcessing:
		return "Processing"
	case BookStatusDone:
		return "Done"
	case BookStatusFailed:
		return "Failed"
	default:
		return string(s)
	}
}
