module github.com/Albyne/tasksApp

go 1.17

require github.com/go-sql-driver/mysql v1.6.0 // indirect

replace github.com/Albyne/tasksApp/models => ../tasksApp/models

require github.com/Albyne/tasksApp/models v0.0.0-00010101000000-000000000000
