package model

//执行数据迁移

func migration() {
	// // 	//
	// 	DB.Set("gorm:table_options", "charset=utf8mb4").
	// 		AutoMigrate(&User{}).
	// 		AutoMigrate(&Video{})
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Video{})

}
