package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/smtp"
	"time"
)

const (
	host     = "smtp.yandex.ru"
	port     = "587"
	from     = "sergeevnicolas20@gmail.com"
	username = "sergeevnicolas20@gmail.com"
	password = "rqrx wvjq nanc eolr"
)

type Feedback struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func sendMail(to []string, subject, body string) error {
	start := time.Now()
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body + "\r\n")
	auth := smtp.PlainAuth("", username, password, host)

	for i := 0; i < 3; i++ {
		err := smtp.SendMail(host+":"+port, auth, from, to, msg)
		if err == nil {
			log.Printf("Письмо успешно отправлено, этап занял: %s", time.Since(start))
			return nil
		}
		log.Printf("Ошибка отправки письма, попытка %d: %v", i+1, err)
		time.Sleep(time.Second * 2)
	}

	log.Printf("Не удалось отправить письмо за 3 попытки, этап занял: %s", time.Since(start))
	return fmt.Errorf("не удалось отправить письмо")
}

func FeedbackHandlerGin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	if c.Request.Method == "OPTIONS" {
		c.Status(204)
		return
	}

	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{"error": "Method Not Allowed"})
		return
	}

	var feedback Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	htmlBody := generateEmailBody(feedback)

	to := []string{"sergeevnicolas20@gmail.com"}
	subject := "New Feedback Submission"

	if err := sendMail(to, subject, htmlBody); err != nil {
		c.JSON(500, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "Feedback sent successfully"})
}

func generateEmailBody(data Feedback) string {
	return fmt.Sprintf(
		`<!DOCTYPE html>
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; margin: 0; padding: 0;">
			<div style="background-color: #ffffff; margin: 20px auto; padding: 20px; border-radius: 8px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); max-width: 700px;">
				<div style="font-size: 24px; font-weight: bold; color: #333333; margin-bottom: 20px; text-align: center;">
					Форма обратной связи
				</div>
				<div style="font-size: 18px; color: #555555; line-height: 1.6;">
					<p>Здравствуйте!</p>
					<p>Пользователь оставил форму обратной связи. Вот его данные:</p>
					<ul style="list-style: none; padding: 0;">
						<li><strong>Имя:</strong> %s</li>
						<li><strong>Email:</strong> %s</li>
					</ul>
					
					<p style="text-align: center; font-size: 18px; color: #333; font-weight: bold; margin-bottom: 10px;">Суть обращения:</p>
					<p style="text-align: center; font-size: 18px; color: #333; margin-bottom: 10px;">%s</p>
					
					<div style="background-color: #e6f7ff; border-left: 4px solid #007acc; padding: 20px; margin-top: 30px; font-weight: bold;">
						<p>Пожалуйста, свяжитесь с ним для уточнения деталей.</p>
					</div>
				</div>
				<div style="font-size: 14px; color: #888888; text-align: center; margin-top: 20px;">
					Это письмо создано автоматически. Пожалуйста, не отвечайте на него.
				</div>
			</div>
		</body>
		</html>`, data.Name, data.Email, data.Message)
}
