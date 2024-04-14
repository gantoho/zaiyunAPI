package service

import (
	"github.com/google/uuid"
	"time"
	"zaiyun.app/app/models"

	"github.com/dgrijalva/jwt-go"
)

// IsTokenExpired 检查给定的JWT令牌是否已过期
func IsTokenExpired(tokenString string) bool {
	// 解析JWT令牌，仅验证过期时间（exp）而不验证签名
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		// 如果解析出错，认为令牌已过期（或格式不正确）
		return true
	}

	// 获取Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// 如果Claims类型不匹配，认为令牌已过期
		return true
	}

	// 获取过期时间（exp）
	expirationTime, ok := claims["exp"].(float64)
	if !ok {
		// 如果exp键不存在或类型不匹配，认为令牌已过期
		return true
	}

	// 将过期时间转换为时间戳
	expirationTimestamp := time.Unix(int64(expirationTime), 0)

	// 检查当前时间是否晚于过期时间
	return time.Now().After(expirationTimestamp)
}

// FindUserByUsername 根据用户名查找用户
func FindUserByUsername(username string) (error, *models.User) {
	// 在这里实现查询数据库或其他数据源的逻辑
	// 假设db是您的数据库连接或ORM实例
	user := &models.User{}
	err := models.Conn.Where("username = ?", username).First(user).Error
	if err != nil {
		return err, nil
	}
	return nil, user
}

// JsonInBlacklist 将JWT添加到黑名单，并设置过期时间
func JsonInBlacklist(blacklist models.JwtBlacklist) error {
	blacklist.ExpireAt = time.Now().Add(time.Hour * 24) // 示例：设置黑名单有效期为24小时
	blacklist.CreatedTime = time.Now()
	blacklist.UpdatedTime = time.Now()
	// 在这里实现将blacklist插入数据库或其他数据源的逻辑
	// 假设db是您的数据库连接或ORM实例
	err := models.Conn.Create(&blacklist).Error
	if err != nil {
		return err
	}
	return nil
}

// FindUserByUuid 根据UUID查找用户
func FindUserByUuid(uuidStr string) (error, *models.User) {
	// 解析UUID字符串
	uuidObj, err := uuid.Parse(uuidStr)
	if err != nil {
		return err, nil
	}

	// 在这里实现查询数据库或其他数据源的逻辑
	// 假设db是您的数据库连接或ORM实例
	user := &models.User{}
	err = models.Conn.Where("uuid = ?", uuidObj).First(user).Error
	if err != nil {
		return err, nil
	}
	return nil, user
}
