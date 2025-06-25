package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

// UserTodo represents a single todo item
type UserTodo struct {
	Task string
	Done bool
}

// In-memory store for user todos: map[userID][]UserTodo
var userTodos = make(map[string][]UserTodo)

func addTodo(userID, task string) string {
	userTodos[userID] = append(userTodos[userID], UserTodo{Task: task, Done: false})
	return fmt.Sprintf("Added todo: %s", task)
}

func listTodos(userID string) string {
	todos := userTodos[userID]
	if len(todos) == 0 {
		return "Your todo list is empty."
	}
	result := "Your Todos:\n"
	for i, todo := range todos {
		status := "[ ]"
		if todo.Done {
			status = "[x]"
		}
		result += fmt.Sprintf("%d. %s %s\n", i+1, status, todo.Task)
	}
	return result
}

func removeTodo(userID string, index int) string {
	todos := userTodos[userID]
	if index < 0 || index >= len(todos) {
		return "Invalid todo number."
	}
	removed := todos[index].Task
	userTodos[userID] = append(todos[:index], todos[index+1:]...)
	return fmt.Sprintf("Removed todo: %s", removed)
}

func markTodoDone(userID string, index int) string {
	todos := userTodos[userID]
	if index < 0 || index >= len(todos) {
		return "Invalid todo number."
	}
	if todos[index].Done {
		return "Todo is already marked as done."
	}
	userTodos[userID][index].Done = true
	return fmt.Sprintf("Marked as done: %s", todos[index].Task)
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("--------------------------------")
	}
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")
	if botToken == "" || appToken == "" {
		log.Fatal("SLACK_BOT_TOKEN and SLACK_APP_TOKEN must be set in environment variables or .env file")
	}

	bot := slacker.NewClient(botToken, appToken)

	go printCommandEvents(bot.CommandEvents())
	bot.Command("ping", &slacker.CommandDefinition{
		Description: "Responds with pong",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong, chala ja yaha se")
		},
	})

	bot.Command("add todo <task>", &slacker.CommandDefinition{
		Description: "Add a new todo to your list",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			userID := botCtx.Event().UserID
			task := request.Param("task")
			msg := addTodo(userID, task)
			response.Reply(msg)
		},
	})

	bot.Command("list todos", &slacker.CommandDefinition{
		Description: "List your todos",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			userID := botCtx.Event().UserID
			msg := listTodos(userID)
			response.Reply(msg)
		},
	})

	bot.Command("remove todo <number>", &slacker.CommandDefinition{
		Description: "Remove a todo by its number (see list)",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			userID := botCtx.Event().UserID
			numStr := request.Param("number")
			var idx int
			_, err := fmt.Sscanf(numStr, "%d", &idx)
			if err != nil {
				response.Reply("Please provide a valid todo number.")
				return
			}
			msg := removeTodo(userID, idx-1)
			response.Reply(msg)
		},
	})

	bot.Command("done todo <number>", &slacker.CommandDefinition{
		Description: "Mark a todo as done by its number (see list)",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			userID := botCtx.Event().UserID
			numStr := request.Param("number")
			var idx int
			_, err := fmt.Sscanf(numStr, "%d", &idx)
			if err != nil {
				response.Reply("Please provide a valid todo number.")
				return
			}
			msg := markTodoDone(userID, idx-1)
			response.Reply(msg)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
