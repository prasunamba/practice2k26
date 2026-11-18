ref link :https://gqlgen.com/getting-started/?utm_source=chatgpt.com

graphQL implementation

Start the GraphQL server
Run:
go mod tidy
go run server.go


Your server will start on:
http://localhost:8080/


Test 1: Create a user (Mutation)
mutation {
  createUser(input: {
    name: "xxxxxxx"
    email: "xxxx@gmail.com"
  }) {
    id
    name
    email
  }
}

Test 2: Get all users (Query)
query {
  users {
    id
    name
    email
  }
}

Test 3: Get user by ID
query {
  user(id: "1") {
    id
    name
    email
  }
}

