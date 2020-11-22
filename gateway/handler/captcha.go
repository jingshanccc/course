package handler

import (
	"course/middleware/redis"
	"course/public"
	"github.com/gin-gonic/gin"
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
	redis.RedisClient.Set(id, value, keyExpire)
}

// customizeRdsStore implementing Get method of  Store interface
func (s *redisStore) Get(id string, clear bool) (value string) {
	val, err := redis.RedisClient.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := redis.RedisClient.Del(id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}

type Captcha struct {
	Id           string
	Type         string
	VerifyValue  string
	Base64String string
}

func init() {
	//使用自定义存储方便自定义校验逻辑
	base64Captcha.SetCustomStore(&redisStore{})
}

//GetCaptcha : 图形验证码
func GetCaptcha(ctx *gin.Context) {
	var req Captcha
	if err := ctx.Bind(&req); err == nil {
		var config interface{}
		switch req.Type {
		case "Audio":
			config = captchaAudio()
			break
		case "Digit":
			config = captchaDigit()
			break
		default:
			config = captchaCharacter()
			break
		}
		idKey, captcha := base64Captcha.GenerateCaptcha(req.Id, config)
		base64String := base64Captcha.CaptchaWriteToBase64Encoding(captcha)
		req.Id = idKey
		req.Base64String = base64String
		log.Println("id: ", idKey)
		log.Println("value: ", redis.RedisClient.Get(idKey))
		public.ResponseSuccess(ctx, req)
		//captcha.WriteTo(ctx.Writer)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//VerifyCaptcha : 校验验证码
func VerifyCaptcha(id, value string) public.BusinessException {
	// 根据验证码token去获取缓存中的验证码，和用户输入的验证码是否一致
	imageCode, _ := redis.RedisClient.Get(id).Result()
	if imageCode == "" {
		return public.NewBusinessException(public.IMAGE_CODE_EXPIRED)
	}
	if strings.ToLower(imageCode) != strings.ToLower(value) {
		return public.NewBusinessException(public.IMAGE_CODE_ERROR)
	} else {
		redis.RedisClient.Del(id)
		return public.NoException("")
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
