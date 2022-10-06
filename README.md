# accountService

running server with command: CompileDaemon -command="./CUAccountService

normal password before hash is either "helloWorld" or "HelloWorld"

the available api currently are

POST http://localhost:3000/user create an account

GET http://localhost:3000/users get all user

GET http://localhost:3000/user/:sid get a user by student Id

PUT http://localhost:3000/user/:sid update account info by student Id

Delete http://localhost:3000/user/:sid delete account info by student Id

POST http://localhost:3000/login login user by studentID and password(unhash)
dataRawJSON: {
"StudentId": "6230000221",
"Password": "HelloWorld"
}
