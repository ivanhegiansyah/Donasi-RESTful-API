package donations

type DonationUpdate struct {
	Donation_name     string `json:"donation_name" form:"donation_name"`
	Status            string `json:"status" form:"status"`
	Short_description string `json:"short_description" form:"short_description"`
	Goal_amount       int    `json:"goal_amount" form:"goal_amount"`
	Current_amount    int    `json:"current_amount" form:"current_amount"`
	Expired_date      string `json:"expired_date" form:"expired_date"`
}
