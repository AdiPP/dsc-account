# About The Project

DSC Account Service is a service that aims to manage accounts and SSO (Single Sign-On) in the DSC ecosystem.

# Getting Started

Run docker compose.

```
docker composer up
```

Service running on port `8080`.

Run curl to ping route to check service status.

```
curl http://localhost:8080/api/ping
```

Check `resources/postman_collection.json` for more routes and example.

# Flow

### Login

![Login](./resources/flow/DSC%20Account%20Service-Login.drawio.png)

### Create User

![Create User](./resources/flow/DSC%20Account%20Service-Create%20User.drawio.png)

### Get User

![Get User](./resources/flow/DSC%20Account%20Service-Get%20User.drawio.png)

### Update User

![Get User](./resources/flow/DSC%20Account%20Service-Update%20User.drawio.png)

### Delete User

![Get User](./resources/flow/DSC%20Account%20Service-Delete%20User.drawio.png)

# Testing User Credential

Here are some default user credentials that can be used for testing purposes.

### Admin User

```
Username: admin
Password: avada_kedavra
```

### Basic User

```
Username: basic_user
Password: capacious_extremis
```

Please, check `mock/user.go` file for more information about default user.

# To do

- [ ] Enhance error handling to be more specific (validation error, service error, http error, and etc)
- [ ] Move configuration to env variable to make it more dynamic
- [ ] Implement DDD (Domain Driven Design) on project structure
- [ ] Refactor controller to Single Action Controllers to improve code readability
- [ ] Implement driver pattern for handling multiple database driver (ex: mysql, mongodb)
- [ ] Implement driver pattern for handling multiple route driver (ex: mux, gin, chi, and etc)
- [ ] ....