# Slack Bot: In-Memory Todo List

This is a simple Slack bot written in Go that allows users to manage their personal todo lists directly from Slack. Each user's todos are stored in memory and are private to them.

## Features

- Add, list, remove, and mark todo items as done
- All data is stored in memory (per bot instance)
- Simple command interface via Slack

## Setup

1. Clone this repository.
2. Create a `.env` file in this directory with the following content:

```
SLACK_BOT_TOKEN=your-bot-token-here
SLACK_APP_TOKEN=your-app-token-here
```

**Do NOT commit your `.env` file or share your tokens.**

3. Run `go run main.go` after installing dependencies.

## Usage

Tag the bot in your Slack workspace and use the following commands:

- `@your-bot ping` — Responds with "pong, chala ja yaha se"
- `@your-bot add todo <task>` — Adds a new todo item
- `@your-bot list todos` — Lists all your todos
- `@your-bot remove todo <number>` — Removes the todo at the given number (see list)
- `@your-bot done todo <number>` — Marks the todo at the given number as done

Replace `@your-bot` with your bot's actual Slack handle.

---

**Note:** Todos are stored in memory and will be lost if the bot restarts.
