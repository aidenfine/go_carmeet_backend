package meets_types

type Meet struct {
	Meet_id          string            `json:"meet_id"`
	Creator_id       string            `json:"creator_id"`
	Title            string            `json:"title"`
	Description      *string           `json:"description"`
	Location         string            `json:"location"`
	Meet_date        string            `json:"meet_date"`
	Meet_time        MeetTime          `json:"meet_time"`
	Recurring        bool              `json:"recurring"`
	RecurringOptions *RecurringOptions `json:"recurring_options"`
	Promoted         PromotedType      `json:"promoted"`
	Theme            string            `json:"theme"`
	Meet_banner      string            `json:"meet_banner"`
	Meet_thumbnail   string            `json:"meet_thumbnail"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// type Location struct {
// }

type MeetTime struct {
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
}

type RecurringOptions string

const (
	Daily       RecurringOptions = "DAILY"
	Weekly      RecurringOptions = "WEEKLY"
	Bi_weekly   RecurringOptions = "BI_WEEKLY"
	Tri_weekly  RecurringOptions = "TRI_WEEKLY"
	Monthly     RecurringOptions = "MONHTLY"
	Bi_monthly  RecurringOptions = "BI_MONHTLY"
	Tri_monthly RecurringOptions = "TRI_MONTHLY"
	Semi_year   RecurringOptions = "SEMI_YEAR"
	Yearly      RecurringOptions = "YEARLY"
)

// func (ro RecurringOption) String() string {
// 	return string(ro)
// }

type PromotedType struct {
	Is_Promoted bool    `json:"is_promoted"`
	Start_time  *string `json:"start_time"`
	End_time    *string `json:"end_time"`
}
