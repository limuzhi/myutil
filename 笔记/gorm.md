如何在 gorm 中使用自定义预加载
我有这两个结构



type User struct {

gorm.Model

Name      string

CompanyID int

Company   Company

}



type Company struct {

gorm.Model

Name string

Adress string

}

我想获得用户并预加载他们的公司，但我不想得到地址字段，我尝试了自定义预加载，如波纹管，并在postman中进行了测试。查询返回了所有字段，但对于地址，我得到一个空字符串，发生这种情况的原因是，当将结果存储在用户结构中时，golang会自动初始化所有字段，并且字段地址将返回其初始值，这是一个空的刺



var user []User

 db.Table("users").Preload("Company",func(db *gorm.DB) *gorm.DB {

    return db.Select("ID" ,"Name")

}).Find(&user)


//user预加载Language,并且按语言名称Name排序
//db.Preload() 该方法请求参数可以添加func方法，在func中传入gorm.DB,
//可以在预加载过程执行其他sql操作
//类比，在gorm的其他方法中的请求参数中也可能允许传入db.*gorm参数
//注意：
//故在使用Preload时应注意顺序，如 Preload("Language")--->PreloadPreload("Language.Country")
//而不是 Preload("Language.Country")--->PreloadPreload("Language"),可能会出现未知错误，如排序失效
func QueryUser(){
    var userModel:=model.User{}
    var userData []model.User
    dao.MySQLConn.Model(process).Order("UserId").Preload("Language",func(db *gorm.DB)*gorm.DB{
        return db.Order("Name asc")
       
}).Find(&userData)
    fmt.Println(userData)
}

//拓展
//如果多级别预加载，则层级预加载
//如:user预加载Language(语言)预加载Country(国家)
//常规写法,.Preload("Language.Country"),无法实现按名称Country.Name排序
//写法如下:
 dao.MySQLConn.Model(process).Order("UserId").Preload("Language",func(db *gorm.DB)*gorm.DB{
        //再包含一层
     return db.Preload("Country",func(db *gorm.DB) *gorm.DB{
         return db.Order("Name asc")
     })
    }).Find(&userData)