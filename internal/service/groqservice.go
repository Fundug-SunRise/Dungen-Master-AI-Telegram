package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/conneroisu/groq-go"
	"github.com/joho/godotenv"
)

func test() {

	err := godotenv.Load("../../internal/config/.env")
	groqToken := os.Getenv("AITOKEN")

	client, err := groq.NewClient(groqToken)
	if err != nil {
		log.Fatal(err)
	}

	// Основной цикл
	for {
		fmt.Print("Вы: ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		// Отправка запроса
		req := groq.ChatCompletionRequest{
			Model: "llama-3.3-70b-versatile",
			Messages: []groq.ChatCompletionMessage{
				{
					Role:    groq.RoleSystem,
					Content: "Ты Dunger Master ты должен играть с игроком первоначальные статы игрока: Он эльф 2 уровня у него есть способность поле цветов",
				},
				{
					Role:    groq.RoleUser,
					Content: input,
				},
			},
		}

		response, err := client.ChatCompletion(context.Background(), req)
		if err != nil {
			log.Println("Ошибка:", err)
			continue
		}

		// Вывод ответа
		if len(response.Choices) > 0 {
			fmt.Printf("AI: %s\n", response.Choices[0].Message.Content)
		} else {
			fmt.Println("AI: Не удалось получить ответ.")
		}
	}

}

func getToken() string {
	err := godotenv.Load("../../internal/config/.env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv("AITOKEN")
}

func RequesGroq(mainCombinedPromt string, requestUser string) string {

	client, err := groq.NewClient(getToken())
	if err != nil {
		log.Fatal(err)
	}

	req := groq.ChatCompletionRequest{
		Model: "llama-3.3-70b-versatile",
		Messages: []groq.ChatCompletionMessage{
			{
				Role:    groq.RoleSystem,
				Content: mainCombinedPromt,
			},
			{
				Role:    groq.RoleUser,
				Content: requestUser,
			},
		},
	}

	response, err := client.ChatCompletion(context.Background(), req)
	if err != nil {
		log.Println("Error:", err)
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content
	} else {
		return "Не удалось получить ответ."
	}
}
