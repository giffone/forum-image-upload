## forum: version image upload

forum based on: "basic forum" + "autentication" + "image upload"

In this version:

- registered users have the possibility to create a post containing an image as well as text.
- when viewing the post, users and guests should see the image associated to it.
- there are several extensions for images like: JPEG, SVG, PNG, GIF, etc. In this project we have to handle at least JPEG, PNG and GIF types.
- the max size of the images to load should be 20 mb. If there is an attempt to load an image greater than 20mb, an error message should inform the user that the image is too big.

### Objectives

This project consists in creating a web forum that allows:

- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.

### SQLite

In this project we use sqlite db.

### Run

To run project please type in command line `go run ./cmd/forumsqlite/` or `make run`

### Randomizer
Delete current db `forum/db/database-sqlite3.db`

To use random `user`, `categories` and `post` for testing, need to uncomment this lines in directory : `forum/internal/app/app.go`:

- `_, _, schema := repo.ExportSettings()`
- `repository.NewLoremIpsum().Run(db, schema)`

You can then return comment.

For example random user to login:

login: `blackbeard`

password: `12345Aa`

### Docker

Run command `make docker` in command line.
