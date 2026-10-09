package enum

type PaymentStatus string

const (
	PaymentStatusNone          PaymentStatus = "none"
	PaymentStatusPendingReview PaymentStatus = "pending_review"
	PaymentStatusPaid          PaymentStatus = "paid"
)
