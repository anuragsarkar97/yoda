package Streams

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"
	"unsafe"
	"yoda/pkg/Db"
)

type DatabaseLinks struct {
	ExpiryId		string		`bson:"expiryId"`
	Url 			string		`bson:"expiryURL"`
	Time 			time.Time	`bson:"timeStamp"`
	TimeToLive		int64		`bson:"TTL"`
	BasePath		string		`bson:"BasePathToVM"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const resourcePoolBasePath = "/Users/oyo/Downloads/yodaResources/public/"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	baseUrl = "https://localhost:5001/filehost/"
	urlSize = 30
)
const (
	mongoUrl = "localhost:27017"
	dbName = "test_db"
	userCollectionName = "data"
)
var src = rand.NewSource(time.Now().UnixNano())
var session, _ = Db.NewSession(mongoUrl)
var collection = session.GetCollection(dbName, userCollectionName)


func GenerateFileStreams(fileName string, ttl int64) string  {
	if isPresent(fileName) && isParsable() {
		filePath := getfilePath(fileName)
		expiringUrl := createExpiringUrl(urlSize)
		_ = collection.Insert(createStorableData(expiringUrl, ttl, filePath))
		return expiringUrl
	}
	return ""
}

func createStorableData(expiryUrl string, ttl int64, file string) DatabaseLinks {
	var c DatabaseLinks
	c.Url = expiryUrl
	c.ExpiryId = expiryUrl[32:]
	c.Time = time.Now()
	c.TimeToLive = ttl
	c.BasePath = file
	return c
}

func createExpiringUrl(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	dynamicUrl := *(*string)(unsafe.Pointer(&b))
	return baseUrl + dynamicUrl
}

func check(e error) bool {
	if e != nil {
		panic(e)
		return true
	}
	return false
}

func getfilePath(fileName string) string {
	files, _ := filepath.Glob(resourcePoolBasePath + "*")
	fmt.Println(files)
	for _, file := range files {
		if file == fileName {
			return fileName
		}
	}
	return ""
}

func isPresent(fileName string) bool{
	return true
}

func isParsable() bool {
	return true
}
