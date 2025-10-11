# DailyChapter - Server
## Overview
This project is the back-end for the DailyChapter app, which encourages users to read a chapter (or other specified amount) of a book each day. Reads in from a library of public domain books and feeds users a chapter of their selected book via the app, or if they don't log in to read it, via email. Will provide AI generated comprehension questions after each section is completed as a requirement to mark the chapter complete. Not for user deployment, but left public for portfolio.

This documentation is just a roadmap for how I plan to build the project, and will be updated as portions are completed.

## Endpoints
### Status
- [x] /users POST
- [x] /login POST
- [x] /books POST
- [x] /books/{bookID}/subscribe POST
- [x] /books GET
- [ ] /users GET
- [ ] /users PUT
- [ ] /books PUT

### Endpoint Descriptions
#### /users POST

Endpoint for registering a new user by email

##### Request body format:
```json
{
    "email":"YOUREMAIL", 
    "password":"YOURPASSWORD"
}
```

#### /users GET

Endpoint to retrieve user information

#### /users PUT

Endpoint to change user information

#### /login POST

Endpoint for logging into an existing user profile; returns an access token

##### Request body format:
<u>Header must include an API Key</u>
```json
{
    "email":"YOUREMAIL",
    "password":"YOURPASSWORD"
}
```
##### Response body format:
```json
{
    "Token":"JWTSTRING"
}
```

#### /books POST 

Endpoint to add a new book to the library (will be exclusive to webhook with external service)

##### Request body format:
```json
{
    "title":"BOOK TITLE",
    "pub_date":"01/23/1999"
}
```

#### /books/{bookID}/subscribe POST 

<u>Request must include a valid access token</u> (user needs to be logged in)

##### Request body format:

No body required.

#### /books GET

Endpoint to get book information

##### Request body format:

No body required

##### Response body format:

```json
[
    {
        "title":"BOOKTITLE",
        "pub_date":"PUBLISHING DATE"
    }
]
```

#### /books PUT

Endpoint to change user progress in a book

## Database Schema

### Users table
- [x] Fully Implemented

| Field          | Data Type | Nullable |
|----------------|:----------|:--------:|
|id              |UUID       | No       |
|email           |Text       | No       |
|hashed_password |Text       | No       |
|created_at      |Timestamp  | No       |
|updated_at      |Timestamp  | No       |

### Books table
- [x] Fully Implemented

| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|chapters    |Integer    | No       |
|title       |Text       | No       |

### Authors table
- [x] Fully Implemented

| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|name        |Text       | No       |

### Refresh tokens table
- [x] Fully Implemented

| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|user_id     |UUID       | No
|token       |Text       | No       |
|valid_until |Timestamp  | No       |
|revoked_at  |Timestamp  | Yes      |

### Chapters table
- [ ] Fully Implemented

| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|book_id     |UUID       | No       |
|content     |Text       | No       |

### UserBooks table
- [ ] Fully Implemented

| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|active      |Bool       | No       |
|user_id     |UUID       | No       |
|book_id     |UUID       | No       |
|chapter_id  |UUID       | Yes      |

