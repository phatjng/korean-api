# korean-api

A 한국어 vocabulary flash card API.

🚧 This is a work in progress! 🚧

## Installation

Create a SQLite file inside the `db/` folder.

In your `.env` file, include your database URL in the `DATABASE_URL`.

`
DATABASE_URL="sqlite:db/database.sqlite3"
`

## Getting Started

| Path             | Body                        | Description                             |
|------------------|-----------------------------|-----------------------------------------|
| [POST] /api/deck | title: string               | Create a new deck                       |
| [Get] /api/deck  | id: string, title: string   | Retrieve deck based on deck id or title |
| [POST] /api/card | front: string, back: string | Create a new card                       |
| [Get] /api/card  | id: string                  | Retrieve card based on card id          |