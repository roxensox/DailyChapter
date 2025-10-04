# DailyChapter - Server
## Overview
Back-end for DailyChapter app, which encourages users to read a chapter (or other specified amount) of a book each day. Reads in from a library of public domain books and feeds users a chapter of their selected book via the app, or if they don't log in to read it, via email. Will provide AI generated comprehension questions after each section is completed as a requirement to mark the chapter complete. Not for user deployment, but left public for portfolio.

## Endpoints
#### - [x] /users GET
Endpoint for registering a new user by email

```JSON format: {"email":"YOUREMAIL"}```

#### - [ ] /users GET
Endpoint to retrieve user information

#### - [ ] /users PUT
Endpoint to change user information

#### - [ ] /books POST 
Endpoint to add a new book to the library (will be exclusive to webhook with external service)

#### - [ ] /books GET
Endpoint to get book information

#### - [ ] /books PUT
Endpoint to change user progress in a book
