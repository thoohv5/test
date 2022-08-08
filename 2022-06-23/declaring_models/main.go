package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime

	// Field1 *Field
	Field1 *Field `gorm:"column:field1;type:string;default:'';"`
	// Field2 Field
}

type Field struct {
	Name string
}

func (f *Field) Scan(src interface{}) (err error) {
	bs, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", bs)
	}

	if len(bs) == 0 {
		return nil
	}
	f.Name = string(bs)
	return
}

func (f *Field) Value() (driver.Value, error) {
	if f == nil {
		return "", nil
	}
	return f.Name, nil
}

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Migrator().DropTable(&User{})

	// 迁移 schema
	db.AutoMigrate(&User{})
	if db.Migrator().HasColumn(&User{}, "UpdatedAt2") {
		fmt.Println(db.Migrator().DropColumn(&User{}, "UpdatedAt2"))
	}

	u := &User{
		Name:         "1",
		Email:        nil,
		Age:          0,
		Birthday:     nil,
		MemberNumber: sql.NullString{},
		ActivatedAt:  sql.NullTime{},
		Field1: &Field{
			Name: "123321",
		},
		// Field2: Field{
		// 	Name: "123",
		// },
	}

	ret := db.Select("Name", "Email").Create(u)
	// ret := db.Create(u)

	fmt.Println(ret.Error, ret.RowsAffected, u)

	uu := &User{}
	db.First(uu, 1)

	// fmt.Printf("%+v, %+v\n", uu.Field1.Name, uu.Field2.Name)
	fmt.Printf("%+v\n", uu.Field1.Name)

	fmt.Println(db, err)
}
