# Project Overview

The task is to build an HTTP server API in Golang using the gin framework to accept and persist contact data into a relational database.
The project was designed with simplicity and modularity in mind, while also meeting the primary objectives of validating data, ensuring phone numbers are in the correct E. 164 format, and allowing data persistence.

# Code Structure

1. `main.go`
This is main function where the Gin router is initiallized, and the database connection is set up.
I created a post endpoint `/contacts` is created to handle incoming contact data.

2. `controllers/contact.go`
It contains the logic to handle HTTP request related to "Contacts".
CreateContact function handles the `POST` request to persist contact data.
I added the validataion function and the input is validated before saving any contact data.
THe includes checking for a valid email format and ensuring that phone numbers are in the #.164 format.
If any validation falis, ther user is informed with an appropriate error message but on success, the validated contact is saved to the database and the response iclude the saved data.

3. `models/contact.go`
I created the model here and it contains fields like `FullName`, `Email` and `PhoneNumber`.

4. `database/db.go`
This file is responsible for initializing and connectig to the database.
I used SQLite and the GORM's auto-migration feature is used to ensure the database schema reflects the structure of the Contact model.

# Testing
I ran serveral test using tools like curl and Postman to simulate API requests and ensure the validation logic works as expected.

# Future Enhancements
This code will be completed more because I didn't add any test like unitest.
And this is simple structure so for more complex structures could be supported by enhancing the data model and more error handling.
