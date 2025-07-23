package services

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/common"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constants"
	"GOLANG_CLEAN_WEB_API/src/data/db"
	"GOLANG_CLEAN_WEB_API/src/data/models"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"GOLANG_CLEAN_WEB_API/src/pkg/service_errors"
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
	Preload []preload
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
		Preload:  []preload{{string: "Cities"}}, // این نیود تو ویدو  خودم اضافه کردم فکر کنم  الکیه
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) { //این تابع بدون اینکه مستقیم با مدل‌ها درگیر بشی، داده ورودی رو دریافت می‌کنه، تبدیل می‌کنه، داخل دیتابیس ذخیره می‌کنه، و در نهایت نسخه کامل اون رکورد رو برمی‌گردونه
	model, _ := common.TypeConverter[T](req) //رک  که یک  دی تی  او  هست وارد میشه
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.
		Create(model).
		Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	bm, _ := common.TypeConverter[models.BaseModel](model)
	return s.GetById(ctx, bm.Id)

}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, req *Tu, id int) (*Tr, error) {
	updateMap, _ := common.TypeConverter[map[string]interface{}](req)
	snakeMap := map[string]interface{}{}
	for k,v :=range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	
	(snakeMap)["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	(snakeMap)["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()

	err := tx.
		Model(model).
		Where("id = ? AND deleted_by IS NULL", id).
		Updates(snakeMap).
		Error

	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}

	tx.Commit()
	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()
	model := new(T)

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	if ctx.Value(constants.UserIdKey) == nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	if cnt := tx.
		Model(model).
		Where("id = ? AND deleted_by IS NULL", id).
		Updates(&deleteMap).
		RowsAffected; cnt == 0 {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, service_errors.PermissionDenied, nil)
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	db:= Preload(s.Database,s.Preload)
	err := db.
		Where("id = ? AND deleted_by IS NULL", id).
		First(model).
		Error
	if err != nil {
		return nil, err
	}

	return common.TypeConverter[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[Tr], error) {
	return Paginate[T, Tr](req, s.Preload, s.Database)
}

func NewPageList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PagedList[T] {
	pl := &dto.PagedList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Item:       items,
	}
	pl.TotalPage = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPage
	pl.HasPreviousPage = pl.PageNumber > 1

	return pl
}

func Paginate[T any, Tr any](pagination *dto.PaginationInputWithFilter, Preloads []preload, db *gorm.DB) (*dto.PagedList[Tr], error) {
	model := new(T)
	var items *[]T
	var rItem *[]Tr

	// 🔍 Log: Preload‌ها
	fmt.Println("📦 Preloading:", Preloads)

	// اعمال پری‌لود روی DB
	db = Preload(db, Preloads)

	// ساخت شرط‌های فیلتر و مرتب‌سازی
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)

	// 🔍 Log: WHERE و ORDER Clause‌ها
	fmt.Println("🔍 WHERE Clause:", query)
	fmt.Println("📐 ORDER Clause:", sort)

	// گرفتن تعداد کل رکوردها
	var totalRows int64 = 0
	db.Model(model).Where(query).Count(&totalRows)

	// 🔎 اجرا با Debug برای دیدن SQL نهایی
	fmt.Println("🚀 اجرای کوئری اصلی با Debug:")
	err := db.Debug().
		Where(query).
		Offset(pagination.GetOffset()).
		Limit(pagination.GetPageSize()).
		Order(sort).
		Find(&items).
		Error
	if err != nil {
		fmt.Println("❌ خطا در اجرای کوئری:", err)
		return nil, err
	}

	// 🧾 لاگ گرفتن از داده‌های اولیه
	fmt.Printf("📊 Raw Items: %+v\n", items)

	// تبدیل ساختار خروجی
	rItem, err = common.TypeConverter[[]Tr](items)
	if err != nil {
		fmt.Println("❌ خطا در تبدیل نوع:", err)
		return nil, err
	}

	// 🧾 لاگ گرفتن بعد از تبدیل
	fmt.Printf("🧾 Converted Items: %+v\n", rItem)

	// ساخت خروجی نهایی
	return NewPageList(rItem, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err
}

func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")

	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(name)
			fmt.Println("📍 تست فیلد:", name, "پیدا شد؟", ok)
			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "not contains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "endWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "equal":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEqual":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < '%s'", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > '%s'", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >=  '%s' ", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))

					} else {
						query = append(query, fmt.Sprintf("%s >=  '%s' ", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))

					}

				}
			}
		}
	}
	return strings.Join(query, " AND ")

}

func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := typeT.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {

				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}

		}
	}
	return strings.Join(sort, ", ")
}

func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)

	}
	return db
}
