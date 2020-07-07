# About

This project is my personal implementation of task described here <https://github.com/GoBootcamp/avatarme>. It is not perfect and alot can be done better but I think it is good enough for this task.

## How to setup

First you need to install postgresql and create there two databases: one for program itself and another for tests.

1. After you installed postgresql write  in console ```sudo -u postgres psql``` (on linux).
2. Change password of postgres user to "password" or change in code to yours. To change password of postgres user write after the command above ```ALTER USER user_name WITH PASSWORD 'new_password';```.
3. Create first db for application itself ```CREATE DATABASE indenticon```.
4. Create secod db for tests with ```CREATE DATABASE test_indenticon```.

## How to run

After setup you can run application with two commands one is for frontend and the other is for server:
1. Use ```make dev``` command in the root application folder to run server
2. Enter ```cd client``` to move to the frontend folder.
3. Run client with ```npm start``` command(for this one you will need to have npm installed on your machine)

## Testing 

To run all tests use ```make test``` command in the root folder.
