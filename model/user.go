package model

/**
 *NOTE当通过 struct 进行查询的时候，GORM 将会查询这些字段的非零值，
意味着你的字段包含0，''，false或者其他零值 ,将不会出现在查询语句中， 例如:
db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
//SELECT * FROM users WHERE name = "jinzhu";
你可以考虑适用指针类型或者 scanner/valuer 来避免这种情况。
// 使用指针类型
type User struct {
  gorm.Model
  Name string
  Age  *int
}

// 使用 scanner/valuer
type User struct {
  gorm.Model
  Name string
  Age  sql.NullInt64
}
 */
type User struct {
	Id string `gorm:"primaryKey" json:"id"`
	Account string `json:"account"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	IdCardNo string `json:"idCardNo"`
	EmployeeNo string `json:"employeeNo"`
	UserType string `json:"userType"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	DeptName string `json:"deptName"`
	Remark string `json:"remark"`
}

//认证参数
type AuthParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
