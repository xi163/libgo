package sign

import (
	"math/rand"
	"time"

	"github.com/xi163/libgo/logs"
	"github.com/xi163/libgo/utils/codec/base64"
	"github.com/xi163/libgo/utils/codec/uri"
	"github.com/xi163/libgo/utils/crypto/aes"
	"github.com/xi163/libgo/utils/json"
	"github.com/xi163/libgo/utils/random"
)

type Sign struct {
	Rand      string `json:"rand" form:"rand"`
	Data      any    `json:"data" form:"data"`
	Timestamp int64  `json:"timestamp" form:"timestamp"`
	Expired   int64  `json:"expired" form:"expired"`
}

func Encode(data any, exp time.Time, secret []byte) string {
	token := Sign{
		Rand:      random.CharStr(rand.Int() % 10),
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
		Expired:   exp.UnixMilli(),
	}
	encrypt := aes.CBCEncryptPKCS7(json.Bytes(token), secret, secret)
	strBase64 := base64.URLEncode(encrypt)
	return uri.URLEncode(strBase64)
}

func Decode(s string, secret []byte) (v any, exp int64) {
	strBase64 := uri.URLDecode(s)
	encrypt := base64.URLDecode(strBase64)
	jsonStr := aes.CBCDecryptPKCS7(encrypt, secret, secret)
	model := Sign{}
	err := json.Parse(jsonStr, &model)
	if err != nil {
		logs.Errorf(err.Error())
		return nil, 0
	}
	return model.Data, model.Expired
}
