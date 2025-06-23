# Testing Vuetify

## Cara run

1. Copy file .env.example ke file .env di directory yang sama

2. Run docker

   ```sh
   docker compose down
   docker compose up --build
   ```

3. Contoh Frontend

   ```sh
   http://localhost:3000
   ```

4. Contoh Backend

   ```sh
   http://localhost:8080/api/users
   ```

   ```sh
   http://localhost:8080/api/users/1
   ```

## Cek DB PostgreSQL

```sh
docker compose exec -it db psql -U myuser -d mydb
```

## Endpoints

### Request Formats

1. POST /api/users (CreateUser)
   - Accepts JSON body that maps to User struct
   - Example: {"username": "john_doe", "full_name": "John Doe", "password": "secret123"}
2. GET /api/users/:id (GetUser)
   - Requires ID parameter in URL path
   - Example: /api/users/1
3. GET /api/users (GetAllUsers)
   - No request body needed
4. PUT /api/users/:id (UpdateUser)
   - Requires ID parameter in URL path
   - Accepts JSON body with fields to update
   - Example: {"full_name": "John Smith"}
5. DELETE /api/users/:id (DeleteUser)
   - Requires ID parameter in URL path

### Response Formats

1. Success Responses
   - JSON objects for user data
   - Password field is excluded (json:"-" tag)
   - Example:

   ```json
   {
     "id": 1,
     "username": "john_doe",
     "full_name": "John Doe"
   }
   ```

2. Error Responses
   - JSON with error message
   - Example:

     ```json
     {
       "error": "User not found"
     }
     ```

### Status Codes

1. CreateUser
   - 201 Created (success)
   - 400 Bad Request (invalid input)
   - 500 Internal Server Error (database error)
2. GetUser
   - 200 OK (success)
   - 400 Bad Request (invalid ID)
   - 404 Not Found (user not found)
3. GetAllUsers
   - 200 OK (success)
   - 500 Internal Server Error (database error)
4. UpdateUser
   - 200 OK (success)
   - 400 Bad Request (invalid ID/data)
   - 404 Not Found (user not found)
   - 500 Internal Server Error (update failed)
5. DeleteUser
   - 204 No Content (success)
   - 400 Bad Request (invalid ID)
   - 404 Not Found (user not found)
