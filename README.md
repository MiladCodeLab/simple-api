# simple-api
## Running the Application

Ensure Go is installed (Go 1.25.1+).

### Run using Make

```sh
make run
````

The server boots on:

```
http://localhost:5001
```

---

##  Testing the API (cURL commands)

### 1. Create User

```sh
curl -X POST http://localhost:5001/api/users/ \
  -H "Content-Type: application/json" \
  -d '{"name": "Milad", "phone": "09120000000"}'
```

### 2. Get All Users

```sh
curl http://localhost:5001/api/users/
```

### 3. Get User by ID

Replace `<id>` with a real UUID returned from the create request:

```sh
curl http://localhost:5001/api/users/<id>
```

### 4. Delete User by ID

```sh
curl -X DELETE http://localhost:5001/api/users/<id>
```
