package utils

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gotomsak/sconcent/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var layout = "2006-01-02 15:04:05"

func SqlConnect() (database *gorm.DB) {
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USERR")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.GetIDLog{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.AdminUser{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.AnswerResult{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.AnswerResultSection{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Questionnaire{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.OnlyQuestion{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Genre{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Season{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.GetJinsMemeTokenSave{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.SelectQuestion{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Frequency{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.ConcentrationData{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.SonConcentrationData{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

func MongoConnect() (database *mongo.Client, Context context.Context) {
	USER := os.Getenv("USERR")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOLMONGO")
	uri := "mongodb://" + USER + ":" + PASS + "@" + PROTOCOL

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	if err != nil {
		panic(err.Error())
	}
	return c, ctx
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
}

// Bytes2uint converts []byte to uint64
func Bytes2uint(bytes []byte) uint64 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint64(append(padding, bytes...))
	return i
}

// Bytes2int converts []byte to int64
func Bytes2int(bytes []byte) int64 {
	if 0x7f < bytes[0] {
		mask := uint64(1<<uint(len(bytes)*8-1) - 1)

		bytes[0] &= 0x7f
		i := Bytes2uint(bytes)
		i = (^i + 1) & mask
		return int64(-i)

	} else {
		i := Bytes2uint(bytes)
		return int64(i)
	}
}

// Uint2bytes converts uint64 to []byte
func Uint2bytes(i uint64, size int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, i)
	return bytes[8-size : 8]
}

// Int2bytes converts int to []byte
func Int2bytes(i int, size int) []byte {
	var ui uint64
	if 0 < i {
		ui = uint64(i)
	} else {
		ui = (^uint64(-i) + 1)
	}
	return Uint2bytes(ui, size)
}

// intsにsearchがあったらそれを削除してリストを返す
func Remove(ints []int, search int) []int {
	result := []int{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

// intsの中にsearchがあったらtrueを返す
func SearchIDs(ints []uint, search uint) bool {
	for _, v := range ints {
		if v == search {
			return true
		}
	}
	return false
}

// string型で受け取った数値のリストをInt型のリストにして返す
func StrToUIntList(str string) []uint {
	intList := []uint{}
	str = strings.Trim(str, "[]")
	strList := strings.Split(str, ",")
	if str != "" {
		for i := 0; i < len(strList); i++ {
			n, _ := strconv.ParseUint(strList[i], 10, 32)
			un := uint(n)
			intList = append(intList, un)
		}
	}
	return intList
}

// string型のリストをバラバラの順番にして返す
func Shuffle(a []string) {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func StringToTime(str string) time.Time {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation(layout, str, jst)
	return t
}

func StringToUint(str string) uint {
	Uint32, _ := strconv.ParseUint(str, 10, 32)
	Uint := uint(Uint32)
	return Uint
}

// io.Readerをbyteのスライスに変換
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
