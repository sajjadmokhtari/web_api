package migrations

import "GOLANG_CLEAN_WEB_API/src/data/models"

func getBodyProperties(cat int) *[]models.Property{
	var props []models.Property = []models.Property{
		{
			Name: "Height",
			CategoryId: cat,
			Description: "Height of the car",
			DataType: "int",
			Unit: "mm",

		},{

			Name: "Width",
			CategoryId: cat,
			Description: "with of the car ",
			DataType: "int",
			Unit: "mm",

		}, {

			Name: "Length",
			CategoryId: cat,
			Description: "Length of the car ",
			DataType: "int",
			Unit: "mm",
		},
	}
	return &props



}  


func getEngineProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Engine Type", CategoryId: cat, Description: "نوع پیشرانه (بنزینی، دیزلی، هیبریدی، الکتریکی)", DataType: "string", Unit: ""},
        {Name: "Displacement", CategoryId: cat, Description: "حجم موتور", DataType: "float", Unit: "cc"},
        {Name: "Cylinders", CategoryId: cat, Description: "تعداد سیلندر", DataType: "int", Unit: ""},
        {Name: "Horsepower", CategoryId: cat, Description: "قدرت موتور", DataType: "int", Unit: "hp"},
        {Name: "Torque", CategoryId: cat, Description: "گشتاور خروجی موتور", DataType: "int", Unit: "Nm"},
        {Name: "Fuel Type", CategoryId: cat, Description: "نوع سوخت مصرفی", DataType: "string", Unit: ""},
        {Name: "Compression Ratio", CategoryId: cat, Description: "نسبت تراکم موتور", DataType: "float", Unit: ""},
        {Name: "Induction Type", CategoryId: cat, Description: "نوع مکش (تنفس طبیعی، توربوشارژ، سوپرشارژ)", DataType: "string", Unit: ""},
        {Name: "Emission Standard", CategoryId: cat, Description: "استاندارد آلایندگی", DataType: "string", Unit: ""},
        {Name: "Engine Code", CategoryId: cat, Description: "کد فنی موتور", DataType: "string", Unit: ""},
    }
    return &props
}

func getDrivetrainProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Drive Type", CategoryId: cat, Description: "نوع سیستم انتقال قدرت (FWD, RWD, AWD)", DataType: "string", Unit: ""},
        {Name: "Differential Type", CategoryId: cat, Description: "نوع دیفرانسیل (باز، LSD، قفل‌شونده)", DataType: "string", Unit: ""},
        {Name: "Axle Ratio", CategoryId: cat, Description: "نسبت محور انتقال نیرو", DataType: "float", Unit: ""},
        {Name: "Transfer Case", CategoryId: cat, Description: "آیا جعبه انتقال نیرو دارد؟", DataType: "bool", Unit: ""},
        {Name: "Clutch Type", CategoryId: cat, Description: "نوع کلاچ مورد استفاده در سیستم انتقال", DataType: "string", Unit: ""},
        {Name: "Torque Split Front/Rear", CategoryId: cat, Description: "تقسیم گشتاور جلو/عقب در AWD", DataType: "string", Unit: "%"},
        {Name: "Drivetrain Control", CategoryId: cat, Description: "آیا سیستم کنترل الکترونیکی دارد؟", DataType: "bool", Unit: ""},
        {Name: "Locking Capability", CategoryId: cat, Description: "قابلیت قفل دیفرانسیل", DataType: "bool", Unit: ""},
    }
    return &props
}
func getSuspensionProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Suspension Type", CategoryId: cat, Description: "نوع سیستم تعلیق (مستقل، غیرمستقل، چند اتصالی)", DataType: "string", Unit: ""},
        {Name: "Front Suspension", CategoryId: cat, Description: "نوع سیستم تعلیق جلو", DataType: "string", Unit: ""},
        {Name: "Rear Suspension", CategoryId: cat, Description: "نوع سیستم تعلیق عقب", DataType: "string", Unit: ""},
        {Name: "Shock Absorber Type", CategoryId: cat, Description: "نوع کمک‌فنر مورد استفاده", DataType: "string", Unit: ""},
        {Name: "Spring Type", CategoryId: cat, Description: "نوع فنر (پیچشی، بادی، برگ)", DataType: "string", Unit: ""},
        {Name: "Ride Height", CategoryId: cat, Description: "ارتفاع کف خودرو از سطح زمین", DataType: "int", Unit: "mm"},
        {Name: "Suspension Adjustment", CategoryId: cat, Description: "قابلیت تنظیم ارتفاع یا سختی", DataType: "bool", Unit: ""},
        {Name: "Anti-roll Bar", CategoryId: cat, Description: "مجهز به میله ضد غلتش", DataType: "bool", Unit: ""},
        {Name: "Damping Control", CategoryId: cat, Description: "کنترل الکترونیکی میرایی", DataType: "bool", Unit: ""},
        {Name: "Suspension Mode", CategoryId: cat, Description: "حالت‌های قابل انتخاب برای تعلیق (نرمال، اسپرت، کامفورت)", DataType: "string", Unit: ""},
    }
    return &props
}


func getEquipmentProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Tool Kit", CategoryId: cat, Description: "مجموعه ابزار همراه خودرو", DataType: "bool", Unit: ""},
        {Name: "Spare Tire", CategoryId: cat, Description: "آیا خودرو لاستیک زاپاس دارد؟", DataType: "bool", Unit: ""},
        {Name: "Tow Hook", CategoryId: cat, Description: "قلاب بکسل یا اتصال یدک", DataType: "bool", Unit: ""},
        {Name: "Floor Mats", CategoryId: cat, Description: "وجود کف‌پوش‌های داخلی", DataType: "bool", Unit: ""},
        {Name: "Cargo Net", CategoryId: cat, Description: "تور و جداکننده فضای بار", DataType: "bool", Unit: ""},
        {Name: "Fire Extinguisher", CategoryId: cat, Description: "وجود کپسول آتش‌نشانی", DataType: "bool", Unit: ""},
        {Name: "First Aid Kit", CategoryId: cat, Description: "جعبه کمک‌های اولیه", DataType: "bool", Unit: ""},
        {Name: "Jack Type", CategoryId: cat, Description: "نوع جک خودرو", DataType: "string", Unit: ""},
        {Name: "Roof Rack", CategoryId: cat, Description: "باربند سقفی", DataType: "bool", Unit: ""},
        {Name: "Trunk Light", CategoryId: cat, Description: "نور درون صندوق عقب", DataType: "bool", Unit: ""},
    }
    return &props
}

func getDriverSupportProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Adaptive Cruise Control", CategoryId: cat, Description: "تنظیم سرعت براساس فاصله با خودروی جلویی", DataType: "bool", Unit: ""},
        {Name: "Lane Keep Assist", CategoryId: cat, Description: "کمک به حفظ خودرو در بین خطوط", DataType: "bool", Unit: ""},
        {Name: "Blind Spot Detection", CategoryId: cat, Description: "هشدار نقطه کور راننده", DataType: "bool", Unit: ""},
        {Name: "Automatic Emergency Braking", CategoryId: cat, Description: "ترمز اضطراری خودکار", DataType: "bool", Unit: ""},
        {Name: "Parking Assist", CategoryId: cat, Description: "کمک‌پارک اتوماتیک یا نیمه‌اتوماتیک", DataType: "bool", Unit: ""},
        {Name: "Traffic Sign Recognition", CategoryId: cat, Description: "شناسایی تابلوهای راهنمایی و رانندگی", DataType: "bool", Unit: ""},
        {Name: "Driver Drowsiness Detection", CategoryId: cat, Description: "تشخیص خواب‌آلودگی راننده", DataType: "bool", Unit: ""},
        {Name: "Surround View Camera", CategoryId: cat, Description: "دوربین 360 درجه برای دید اطراف خودرو", DataType: "bool", Unit: ""},
        {Name: "Rear Cross Traffic Alert", CategoryId: cat, Description: "هشدار برخورد از طرفین هنگام دنده عقب", DataType: "bool", Unit: ""},
        {Name: "Remote Parking", CategoryId: cat, Description: "پارک خودرو با کنترل از راه دور", DataType: "bool", Unit: ""},
    }
    return &props
}

func getLightsProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Headlight Type", CategoryId: cat, Description: "نوع چراغ جلو (هالوژن، زنون، LED)", DataType: "string", Unit: ""},
        {Name: "Automatic Headlights", CategoryId: cat, Description: "چراغ‌های جلو با فعال‌سازی خودکار", DataType: "bool", Unit: ""},
        {Name: "Daytime Running Lights", CategoryId: cat, Description: "چراغ‌های روشنایی روز", DataType: "bool", Unit: ""},
        {Name: "Fog Lights", CategoryId: cat, Description: "چراغ مه‌شکن جلو یا عقب", DataType: "bool", Unit: ""},
        {Name: "Turn Signal Type", CategoryId: cat, Description: "نوع راهنما (معمولی، دینامیک)", DataType: "string", Unit: ""},
        {Name: "Cornering Lights", CategoryId: cat, Description: "چراغ گوشه‌زن هنگام پیچ", DataType: "bool", Unit: ""},
        {Name: "High Beam Assist", CategoryId: cat, Description: "تنظیم خودکار نور بالا", DataType: "bool", Unit: ""},
        {Name: "Ambient Lighting", CategoryId: cat, Description: "نورپردازی داخلی کابین", DataType: "bool", Unit: ""},
        {Name: "Tail Light Technology", CategoryId: cat, Description: "نوع چراغ عقب", DataType: "string", Unit: ""},
    }
    return &props
}
func getMultimediaProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Screen Size", CategoryId: cat, Description: "اندازه نمایشگر مرکزی", DataType: "float", Unit: "inch"},
        {Name: "Touchscreen", CategoryId: cat, Description: "نمایشگر لمسی", DataType: "bool", Unit: ""},
        {Name: "Navigation System", CategoryId: cat, Description: "سیستم مسیریاب داخلی", DataType: "bool", Unit: ""},
        {Name: "Bluetooth Connectivity", CategoryId: cat, Description: "پشتیبانی از بلوتوث", DataType: "bool", Unit: ""},
        {Name: "USB Ports", CategoryId: cat, Description: "تعداد پورت‌های USB", DataType: "int", Unit: ""},
        {Name: "Apple CarPlay / Android Auto", CategoryId: cat, Description: "سازگاری با سیستم‌های هوشمند گوشی", DataType: "bool", Unit: ""},
    }
    return &props
}
func getSafetyProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Airbags", CategoryId: cat, Description: "تعداد کیسه هوا در خودرو", DataType: "int", Unit: ""},
        {Name: "ABS", CategoryId: cat, Description: "سیستم ترمز ضد قفل (Anti-lock Braking System)", DataType: "bool", Unit: ""},
        {Name: "ESC", CategoryId: cat, Description: "سیستم کنترل پایداری الکترونیکی", DataType: "bool", Unit: ""},
        {Name: "Hill Start Assist", CategoryId: cat, Description: "کمک به حرکت در سربالایی", DataType: "bool", Unit: ""},
        {Name: "ISOFIX", CategoryId: cat, Description: "اتصال ایمن صندلی کودک", DataType: "bool", Unit: ""},
        {Name: "Crash Sensors", CategoryId: cat, Description: "سنسورهای تشخیص تصادف", DataType: "bool", Unit: ""},
        {Name: "Rear Parking Sensors", CategoryId: cat, Description: "سنسور پارک عقب", DataType: "bool", Unit: ""},
    }
    return &props
}


func getSeatsSteeringProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Seat Material", CategoryId: cat, Description: "جنس صندلی‌ها (پارچه، چرم، ترکیبی)", DataType: "string", Unit: ""},
        {Name: "Seat Adjustment Type", CategoryId: cat, Description: "نحوه تنظیم صندلی‌ها (دستی، برقی)", DataType: "string", Unit: ""},
        {Name: "Heated Seats", CategoryId: cat, Description: "صندلی‌های دارای گرم‌کن", DataType: "bool", Unit: ""},
        {Name: "Ventilated Seats", CategoryId: cat, Description: "تهویه صندلی‌ها", DataType: "bool", Unit: ""},
        {Name: "Memory Function", CategoryId: cat, Description: "ذخیره تنظیمات صندلی راننده", DataType: "bool", Unit: ""},
        {Name: "Steering Wheel Adjustment", CategoryId: cat, Description: "تنظیم فرمان (زاویه و فاصله)", DataType: "string", Unit: ""},
    }
    return &props
}



func getWindowsMirrorsProperties(cat int) *[]models.Property {
    props := []models.Property{
        {Name: "Power Windows", CategoryId: cat, Description: "شیشه‌های برقی درها", DataType: "bool", Unit: ""},
        {Name: "Window Tint", CategoryId: cat, Description: "رنگ و میزان دودی بودن شیشه‌ها", DataType: "string", Unit: ""},
        {Name: "Side Mirror Type", CategoryId: cat, Description: "نوع آینه جانبی (معمولی، جمع‌شونده، برقی)", DataType: "string", Unit: ""},
        {Name: "Mirror Heating", CategoryId: cat, Description: "آینه‌های مجهز به گرم‌کن", DataType: "bool", Unit: ""},
        {Name: "Mirror Auto-fold", CategoryId: cat, Description: "جمع شدن خودکار آینه‌ها هنگام قفل کردن", DataType: "bool", Unit: ""},
        {Name: "Rear Window Defogger", CategoryId: cat, Description: "سیستم ضد بخار شیشه عقب", DataType: "bool", Unit: ""},
    }
    return &props
}
