package public

import (
	"context"
	"course/middleware/redis"
	"github.com/mojocn/base64Captcha"
	"log"
	"strings"
	"time"
)

const (
	length    = 4
	keyExpire = time.Minute
)

type redisStore struct {
}

// customizeRdsStore implementing Set method of  Store interface
func (s *redisStore) Set(id string, value string) {
	redis.RedisClient.Set(context.Background(), id, value, keyExpire)
}

// customizeRdsStore implementing Get method of  Store interface
func (s *redisStore) Get(id string, clear bool) (value string) {
	val, err := redis.RedisClient.Get(context.Background(), id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := redis.RedisClient.Del(context.Background(), id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}

func init() {
	//使用自定义存储方便自定义校验逻辑
	base64Captcha.SetCustomStore(&redisStore{})
}

//GetCaptchaBase64: 获取图片验证码 以base64字符串返回
func GetCaptchaBase64(id, typ string) (string, string) {
	idKey, value, captcha := GetCaptcha(id, typ)
	base64String := base64Captcha.CaptchaWriteToBase64Encoding(captcha)
	log.Println("id: ", idKey)
	log.Println("value: ", value)
	return idKey, base64String
}

func GetCaptcha(id, typ string) (idKey string, value string, instance base64Captcha.CaptchaInterface) {
	switch typ {
	case "Audio":
		idKey, instance = base64Captcha.GenerateCaptcha(id, captchaAudio())
		value = instance.(*base64Captcha.Audio).VerifyValue
	case "Digit":
		idKey, instance = base64Captcha.GenerateCaptcha(id, captchaDigit())
		value = instance.(*base64Captcha.CaptchaImageDigit).VerifyValue
	default:
		idKey, instance = base64Captcha.GenerateCaptcha(id, captchaCharacter())
		value = instance.(*base64Captcha.CaptchaImageChar).VerifyValue
	}
	return idKey, value, instance
}

//VerifyCaptcha : 校验验证码
func VerifyCaptcha(id, value string) *BusinessException {
	// 根据验证码token去获取缓存中的验证码，和用户输入的验证码是否一致
	imageCode, _ := redis.RedisClient.Get(context.Background(), id).Result()
	if imageCode == "" {
		return NewBusinessException(VERIFY_CODE_EXPIRED)
	}
	if strings.ToLower(imageCode) != strings.ToLower(value) {
		return NewBusinessException(VERIFY_CODE_ERROR)
	} else {
		redis.RedisClient.Del(context.Background(), id)
		return nil
	}
}

// captchaCharacter : 字符公式验证码
func captchaCharacter() base64Captcha.ConfigCharacter {
	//config struct for Character
	//字符,公式,验证码配置
	return base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         length,
	}
}

//captchaAudio : 声音验证码
func captchaAudio() base64Captcha.ConfigAudio {
	//config struct for audio
	//声音验证码配置
	return base64Captcha.ConfigAudio{
		CaptchaLen: length,
		Language:   "zh",
	}
}

//captchaDigit : 数字验证码
func captchaDigit() base64Captcha.ConfigDigit {
	//config struct for digits
	//数字验证码配置
	return base64Captcha.ConfigDigit{
		Height: 60,
		Width:  240,
		//图像验证码的最大干扰系数
		MaxSkew: 0.7,
		//图像验证码干扰圆点的数量
		DotCount:   80,
		CaptchaLen: length,
	}
}
