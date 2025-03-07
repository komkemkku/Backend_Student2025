package migrations

import model "Beckend_Student2025/models"

func Models() []any {
	return []any{

		(*model.Users)(nil),
		// (*model.Checkins)(nil),
		// (*model.Staffs)(nil),
		// (*model.Events)(nil),
		// (*model.Banner)(nil),
		// (*model.Tickets)(nil),
		// (*model.Admins)(nil),
	}
}

func RawBeforeQueryMigrate() []string {
	return []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}
}

func RawAfterQueryMigrate() []string {
	return []string{}
}
