# web-app
-> To run this application some libraries are required
so to get those lib just run below commands in you're CLI

#1 - go get "github.com/database/sql"
#2 - go get "github.com/gorilla/sessions"
#3 - go get "github.com/satori/go.uuid"
#4 - go get "github.com/gorilla/context"


-> PostgresSql required for DB, need to create a DB by [menu]
table name [go_users] username - aryan, password-password, if you want the username of your choice you can careate but the you have to change username, password in database.go fie.




-> it contains basic sign up and sign in part , and each user must have unique email id to sign up,
-> first create you're user id, 

Column name will be, id,email,firstname,lastname,city,state,zipcode,password all are text except id it is int

-> after starting server use 

#1 - http://localhost:8080/loginpage {to go to login page}

#2 - http://localhost:8080/signup {to go to signup page}

#3 - http://localhost:8080/logout {to log out from session}
curently we need to manually put the the address.
but after sign up it will redirect to login page so no need to open again.