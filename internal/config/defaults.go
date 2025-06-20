package config

func DefaultConfig() *Config {
	return &Config{
		Currency:   "GBP",
		DateFormat: "02/01/2006", // DD/MM/YYYY
		UserName:   "User",
		Categories: []string{
			"Rent",
			"Utilities",
			"Insurance",
			"Subscriptions",
			"Groceries",
			"Eating Out",
			"Coffee",
			"Public Transport",
			"Car Expenses",
			"Cycling",
			"Other Transport",
			"Work & Hobbies",
			"Activites & Socialising",
			"Drinking & Indulgences",
			"Shopping",
			"Gifts",
			"Self-Care & Health",
			"Holiday Saving",
			"Holiday Spending",
			"Other",
		},
		DataFile: "transactions.json",
	}
}
