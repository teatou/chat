package controllers

import (
	"fmt"
	"net/smtp"
	"os"
)

// func Login(c *fiber.Ctx) {
// 	var body struct {
// 		Email string
// 	}

// 	if c.Bind(&body) != nil {
// 		c.Status(fiber.StatusBadRequest)
// 		c.JSON(fiber.Map{
// 			"error": "Failed to read body",
// 		})
// 		return
// 	}
// 	// creating auth code
// 	rand.Seed(time.Now().UnixNano())
// 	code := 1000 + rand.Intn(9000)

// 	user := models.User{Email: body.Email, Nickname: "User", Code: code, Theme: "pink"}
// 	userFind := initializers.DB.First(&user, "email = ?", body.Email)

// 	// if user not found
// 	if userFind.RowsAffected == 0 {
// 		result := initializers.DB.Create(&user)
// 		if result.Error != nil {
// 			c.Status(fiber.StatusBadRequest)
// 			c.JSON(fiber.Map{
// 				"error": "Failed to read body",
// 			})
// 			return
// 		}
// 	} else {
// 		initializers.DB.Model(&user).Update("code", code)
// 	}

// 	// sending code
// 	sendEmail("Poker33 code", strconv.Itoa(code), []string{body.Email})
// }

// func VerifyCode(c *fiber.Ctx) {
// 	var body struct {
// 		Email string
// 		Code  string
// 	}

// 	if c.Bind(&body) != nil {
// 		c.Status(fiber.StatusBadRequest)
// 		c.JSON(fiber.Map{
// 			"error": "Failed to read body",
// 		})
// 		return
// 	}

// 	codeFind, _ := strconv.Atoi(body.Code)

// 	var user models.User
// 	initializers.DB.First(&user, "email = ?", body.Email)

// 	if user.Code != codeFind {
// 		c.Status(fiber.StatusBadRequest)
// 		c.JSON(fiber.Map{
// 			"error": "Failed to read body",
// 		})
// 		return
// 	}

// 	// jwt
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})

// 	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
// 	if err != nil {
// 		c.Status(fiber.StatusBadRequest)
// 		c.JSON(fiber.Map{
// 			"error": "Failed to read body",
// 		})
// 		return
// 	}

// 	// cookie
// 	c.SetSameSite(http.SameSiteLaxMode)
// 	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

// 	c.JSON(http.StatusOK, fiber.Map{
// 		"msg": "cookie set",
// 	})
// }

// func Validate(c *fiber.Ctx) {
// 	user, _ := c.Get("user")

// 	c.Status(fiber.StatusOK)
// 	c.JSON(fiber.Map{
// 		"msg": user,
// 	})
// }

// func Logout(c *fiber.Ctx) {
// 	c.SetCookie("Authorization", "", 0, "", "", false, true)
// }

func sendEmail(subj string, body string, to []string) {
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	auth := smtp.PlainAuth(
		"",
		"poker33bj@gmail.com",
		os.Getenv("GMAILPWD"),
		"smtp.gmail.com",
	)

	msg := "Subject: " + subj + "\n" + headers + "\n\n" + "<h1>" + body + "</h1>"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"poker33bj@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println("Error sending email")
	}
}
