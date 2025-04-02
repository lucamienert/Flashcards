# Flashcards App

This is a flashcard application built with Go for the backend and React for the frontend. It provides features like user authentication, flashcard creation, and more. Users can sign up, sign in, and manage their flashcards.

## Features

- **User Authentication:**
  - Sign up with email and password.
  - Sign in with email and password.
  - Secure token-based authentication using JWT.
  - Logout functionality to invalidate tokens.

- **Flashcard Management:**
  - Create new flashcards.
  - View existing flashcards.
  - Edit or delete flashcards.

## Tech Stack

- **Frontend:**
  - React
  - Axios for API requests
  - React Router for routing

- **Backend:**
  - Go (Gin Framework)
  - GORM for database ORM
  - JWT (JSON Web Tokens) for authentication
  - SQLite database (or any other database of choice)

## Setup

### Backend (Go)

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/flashcards-app.git
    cd flashcards-app
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Configure the environment variables:
    - Create a `.env` file in the root directory.
    - Set the following variables:
      ```env
      DB_HOST=localhost
      DB_PORT=5432
      DB_USER=youruser
      DB_PASSWORD=yourpassword
      DB_NAME=flashcards
      JWT_SECRET=yourjwtsecretkey
      ACCESS_TOKEN_EXPIRY=3600 # in seconds
      REFRESH_TOKEN_EXPIRY=86400 # in seconds
      ```

4. Run database migrations (if using GORM, this can be done automatically by your application).

5. Run the Go backend server:
    ```bash
    go run main.go
    ```

   The server will run on `http://localhost:8080` by default.

### Frontend (React)

1. Navigate to the `client` directory:
    ```bash
    cd client
    ```

2. Install dependencies:
    ```bash
    npm install
    ```

3. Start the React development server:
    ```bash
    npm start
    ```

   The React app will run on `http://localhost:3000`.

### Testing

To run backend tests:

1. Install dependencies:
    ```bash
    go mod tidy
    ```

2. Run tests:
    ```bash
    go test -v
    ```

To run frontend tests:

1. Navigate to the `client` directory:
    ```bash
    cd client
    ```

2. Run tests:
    ```bash
    npm test
    ```

## API Documentation

### Auth Routes

- **POST /auth/signup**: Create a new user account.
  - Request Body:
    ```json
    {
      "name": "John Doe",
      "email": "john@example.com",
      "password": "password123",
      "passwordConfirm": "password123",
      "photo": "http://example.com/photo.jpg"
    }
    ```

- **POST /auth/signin**: Sign in with email and password.
  - Request Body:
    ```json
    {
      "email": "john@example.com",
      "password": "password123"
    }
    ```

- **POST /auth/logout**: Log out and invalidate the JWT tokens.

- **GET /auth/refresh**: Refresh the access token using the refresh token.

### Flashcard Routes (Example)

- **POST /flashcards**: Create a new flashcard.
  - Request Body:
    ```json
    {
      "front": "What is 2 + 2?",
      "back": "4"
    }
    ```

- **GET /flashcards**: Retrieve all flashcards.
- **PUT /flashcards/:id**: Update a flashcard by ID.
- **DELETE /flashcards/:id**: Delete a flashcard by ID.

## Database Schema

### Users Table

- **id**: Primary key
- **name**: User's name
- **email**: User's email (unique)
- **password**: Hashed password
- **role**: User role (default: 'user')
- **photo**: User's photo URL
- **created_at**: Timestamp of when the user was created
- **updated_at**: Timestamp of the last update

### Flashcards Table

- **id**: Primary key
- **user_id**: Foreign key to the Users table
- **front**: The front of the flashcard
- **back**: The back of the flashcard
- **created_at**: Timestamp of when the flashcard was created
- **updated_at**: Timestamp of the last update

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bugfix (`git checkout -b feature-xyz`).
3. Commit your changes (`git commit -am 'Add feature xyz'`).
4. Push to the branch (`git push origin feature-xyz`).
5. Create a new Pull Request.

---

Feel free to ask any questions or raise issues if something isn't clear! 😄
