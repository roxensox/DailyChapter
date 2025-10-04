# DailyChapter - Server
## Overview
This project is the back-end for the DailyChapter app, which encourages users to read a chapter (or other specified amount) of a book each day. Reads in from a library of public domain books and feeds users a chapter of their selected book via the app, or if they don't log in to read it, via email. Will provide AI generated comprehension questions after each section is completed as a requirement to mark the chapter complete. Not for user deployment, but left public for portfolio.

This documentation is just a roadmap for how I plan to build the project, and will be updated as portions are completed.

## Endpoints
### Status
- [x] /users POST
- [ ] /users GET
- [ ] /users PUT
- [ ] /books POST
- [ ] /books GET
- [ ] /books PUT

### Endpoint Descriptions
#### /users POST

Endpoint for registering a new user by email

```JSON format: {"email":"YOUREMAIL"}```

#### /users GET

Endpoint to retrieve user information

#### /users PUT

Endpoint to change user information

#### /books POST 

Endpoint to add a new book to the library (will be exclusive to webhook with external service)

#### /books GET

Endpoint to get book information

#### /books PUT

Endpoint to change user progress in a book

## Database Schema
### Users table
| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|email       |Text       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |

### Books table
| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|book_id     |UUID       | No       |
|chapters    |Integer    | No       |
|title       |Text       | No       |

### Chapters table
| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|book_id     |UUID       | No       |
|content     |Text       | No       |

### UserBooks table
| Field      | Data Type | Nullable |
|------------|:----------|:--------:|
|id          |UUID       | No       |
|created_at  |Timestamp  | No       |
|updated_at  |Timestamp  | No       |
|active      |Bool       | No       |
|user_id     |UUID       | No       |
|book_id     |UUID       | No       |
|chapter_id  |UUID       | Yes      |

