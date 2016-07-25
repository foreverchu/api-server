关于API Auth的方法
====

Token vs Session:
----
1. token保存在客户端, session保存在服务端


jwt-go代码示例:
----
```
	log.SetFlags(log.Lshortfile)
    // 这是密钥
	mySigningKey := []byte("helloworld")
    // 使用的加密算法
	token := jwt.New(jwt.SigningMethodHS256)
    // claims 可以设置做生意的key-value数据
	token.Claims["foo"] = "bar"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tokenString)
    
    // 解密
	myToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		log.Println(token)
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		b := []byte(mySigningKey)
		return b, nil
	})

	log.Println(myToken)

	if err != nil {
		log.Println(err)
	}
	if myToken.Valid {
		log.Println("ok")
	} else {
		log.Println("failed")
	}
```
