#RestGo

this repository is using for test 

## Prerequisite


Go V1.17.xx <br>
other dependency in `go.mod` file

## How to run

* clone this repository
* move to directory
* run this command
  * `go run main.go`

## Summary
in this application, there are 4 endpoints
* `/` for health check
  * to ensure that application is running
* `/customer/login` for login
  * using basic auth username and password & return token in body response
  * will reject other authentication method
  * will check is user really exist
  * will check is password is correct
  * password hashed using bcrypt
  * for this test all password is `user` and username can be found in `mock_data/user1.json`
* `/customer/logout` for logout
  * need to bring bearer token from login process, if not will return error
  * some company didn't provide logout endpoint from backend, they usually just remove token in client side so user can't access backend data (all session thing in front end side)
  * for this case login token stored in in-memory (cache) data as session for authenticated user (must login)
  * log out process basically just remove token from in-memory (cache) data
* `/transaction/transfer` for transfer
  * transfer process with minimal required field just `destinationID` and `amount`
  * transfer can only be accessed by authenticated user (must login)
  * transfer process will write into 2 table `transaction` and `balance_movement`
  * transaction will save `source`, `destination`, `type`, `amount` etc. source data will be filled from parsed token
  * balance movement will save 2 record for each transaction, 1 for debit and 1 for credit
  * all transaction process can be seen in `trx_log/transaction.log` and `trx_log/movement.log`, this file will be reset everytime application is start/re-start

other things
* all activities are recorded to audit_log/[time].log
* dependency inversion is applied in this application with abstraction for every logical and repository layer, so if you want to make a `unit-test` or `integration-test` that will be easier to make a mock-up scenario
* if you want to test with [Postman collection](readme_assets/GoRest.postman_collection.json) is attached in `readme_assets/GoRest.postman_collection.json`
* 