=== REQUIREMENT ===
- Docker 

=== SETUP SERVER ===
1. Open terminal in the project folder 
2. Run "make test" for running unit test
3. Run "make container" for running server and database in the docker container
4. Run "make stop" for stopping and deleting server container and database container

=== API LIST ===
(GET)   http://localhost:8080/account/{id} 
(POST)  http://localhost:8080/account/{from_account_number}/transfer
{
 "to_account_number" : "555002",
 "amount" : 1000
}

=== AVAILABLE ACCOUNT LIST ===
Account_Number  | Name              | Balance
- 555001        | Bob Martin        | 10000
- 555002        | Linus Torvalds    | 15000