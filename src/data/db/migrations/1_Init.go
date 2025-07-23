package migrations

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constants"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()
	CreateTables(database)
	CreateDefaultUserInformation(database)
	CreateCountry(database)
	CreatePropertyCategory(database)
}

func CreateTables(database *gorm.DB) {
	tables := []interface{}{}

	// SYSTEM
	tables = addNewTable(database, models.Country{}, tables)
	tables = addNewTable(database, models.City{}, tables)
	tables = addNewTable(database, models.PersianYear{}, tables)
	tables = addNewTable(database, models.Color{}, tables)
	tables = addNewTable(database, models.File{}, tables)

	// USER
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	// CAR
	tables = addNewTable(database, models.Gearbox{}, tables)
	tables = addNewTable(database, models.CarType{}, tables)
	tables = addNewTable(database, models.Company{}, tables)
	tables = addNewTable(database, models.CarModel{}, tables)
	tables = addNewTable(database, models.CarModelColor{}, tables)
	tables = addNewTable(database, models.CarModelYear{}, tables)
	tables = addNewTable(database, models.CarModelImage{}, tables)
	tables = addNewTable(database, models.CarModelPriceHistory{}, tables)
	tables = addNewTable(database, models.CarModelComment{}, tables)

	// PROPERTY (اصلاح ترتیب)
	tables = addNewTable(database, models.PropertyCategory{}, tables)
	tables = addNewTable(database, models.Property{}, tables)
	tables = addNewTable(database, models.CarModelProperty{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, "Migration failed", nil)
	}

	logger.Info(logging.Postgres, logging.Migration, "Table create seccesfuly", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables

}

func CreateDefaultUserInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName} //نقش ادمین
	CreateRoleIfNotExists(database, &adminRole)             //برای جلوگیری از ساخته شدن دوباره

	defaultRole := models.Role{Name: constants.DefaultRoleName} //ساخت نقش دیفالت
	CreateRoleIfNotExists(database, &defaultRole)               //جلوگیزی ازتکرار

	u := models.User{UserName: constants.DefaultUserName, FirstName: "Test", LastName: "Test", MobileNumber: "09911732328", Email: "admin@admin.com"} // ساخت ادمین پیش فرض
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	CreateAdminUserIfNotExists(database, &u, adminRole.Id) //در صورت وجود نداشتن میسازه
}

func CreateRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)

	if exists == 0 {
		database.Create(r)
	}
}

func CreateAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("user_name = ?", u.UserName).
		First(&exists)

	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func CreateCountry(database *gorm.DB) {
	var count int
	database.Model(&models.Country{}).Select("count(*)").Find(&count)

	if count == 0 {

		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahvaz"},
		}})

		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}})

		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}})

		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}})

		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}})

		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Rome"},
			{Name: "Milan"},
		}})

		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Busan"},
		}})

	} else {
		fmt.Println("ℹ️ [Seed] قبلاً کشورهایی ثبت شده‌اند؛ نیاز به درج مجدد نیست")
	}
}

func CreatePropertyCategory(database *gorm.DB) {
	count := 0
	database.
		Model(&models.PropertyCategory{}).
		Select("count(*)").
		Find(&count)

	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})
		database.Create(&models.PropertyCategory{Name: "Engin"})
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})
		database.Create(&models.PropertyCategory{Name: "Suspension"})
		database.Create(&models.PropertyCategory{Name: "Equipment"})
		database.Create(&models.PropertyCategory{Name: "Driver Support system"})
		database.Create(&models.PropertyCategory{Name: "Lights"})
		database.Create(&models.PropertyCategory{Name: "Multimedia"})
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})
		database.Create(&models.PropertyCategory{Name: "seats and steering wheel"})
		database.Create(&models.PropertyCategory{Name: "windows and mirrors"})
	}

	CreateProperty(database, "Body")
	CreateProperty(database, "Engin")
	CreateProperty(database, "Drivetrain")
	CreateProperty(database, "Suspension")
	CreateProperty(database, "Equipment")
	CreateProperty(database, "Driver Support system")
	CreateProperty(database, "Lights")
	CreateProperty(database, "Multimedia")
	CreateProperty(database, "Safety equipment")
	CreateProperty(database, "seats and steering wheel")
	CreateProperty(database, "windows and mirrors")
}

func CreateProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}
	database.
		Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(&catModel)

	database.
		Model(&models.Property{}).
		Select("count(*)").
		Where("category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}

	var props *[]models.Property

	switch cat {
	case "Body":
		props = getBodyProperties(catModel.Id)
	case "Engin":
		props = getEngineProperties(catModel.Id)
	case "Drivetrain":
		props = getDrivetrainProperties(catModel.Id)
	case "Suspension":
		props = getSuspensionProperties(catModel.Id)
	case "Equipment":
		props = getEquipmentProperties(catModel.Id)
	case "Driver Support system":
		props = getDriverSupportProperties(catModel.Id)
	case "Lights":
		props = getLightsProperties(catModel.Id)
	case "Multimedia":
		props = getMultimediaProperties(catModel.Id)
	case "Safety equipment":
		props = getSafetyProperties(catModel.Id)
	case "seats and steering wheel":
		props = getSeatsSteeringProperties(catModel.Id)
	case "windows and mirrors":
		props = getWindowsMirrorsProperties(catModel.Id)
	default:
		fmt.Printf("⚠️ دسته \"%s\" شناخته‌شده نیست یا تابع ویژگی‌هاش تعریف نشده\n", cat)
		return
	}

	database.Create(&props)
}

func Down_1() {

}
