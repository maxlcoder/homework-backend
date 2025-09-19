package repository

import "gorm.io/gorm"

func LikeScope(field string, keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" {
			return db
		}
		return db.Where(field+" LIKE ?", "%"+keyword+"%")
	}
}

func MultiLikeScope(fields []string, keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" || len(fields) == 0 {
			return db
		}

		query := ""
		args := []interface{}{}
		for i, field := range fields {
			if i > 0 {
				query += " OR "
			}
			query += field + " LIKE ?"
			args = append(args, "%"+keyword+"%")
		}
		return db.Where(query, args...)
	}

}
