# Project Description

People get busier and busier as they get older - it is a known fact that there is a direct correlation between age and number of responsibilities. With busy lifestyles comes absent mindedness - how are we supposed to remember to do all these things? Our app attempts to solve this issue by relieving the stress of having to remember every little task.

Our general target audience will be for people that have busy lifestyles. More specifically, millennials who need a better way to be reminded of what is due. While there are other to-do lists on the market, they usually save their most valuable features for their paying customers. Other apps are free, but they usually aren’t made specifically as todo lists. Neither of the aforementioned solutions allow the user to see their metrics. As a result, there is a real need for a simple and reliable free todo list tracker.

As developers, we want to develop an application that serves our target audience by providing a central place to keep track of their tasks, efficiency and performance. In addition, our target audience will be able to complete their tasks more efficiently by using our task completion visualization and reminder.

# Technical Description

Priority
User
Description
P0
As an email sender
I want to create a shareable link, to share it with other people.
P1
As an email recipient
I want to click on a link in the email, and it will automatically add to my reminder list.
P2
As a non-logged in user
I want to create a bunch of reminders, and log in later.
P3
A logged in user
I want to be able to view my stats (items completed in a week, items overdue, longest streaks, etc)

## Implementation Strategy
Email sender: The sender can create a shareable link to attach to their email. The shareable link will have the following resource path: /add/<random unique string>. We will store the link’s unique random string in our MySQL

Email recipient: The email recipient can click on the ‘/add/<random unique string>’ and it will add to their reminder list. If the user is logged in, then it will do it automatically.

Non-logged in user: If a user clicks on a shareable link and is not logged in, then it will ask the user for an email. Then a temporary account will be created to store the item, login credentials of this temporary account will be sent to the provided email.

Logged in user: The user can import the data from another account, including the temporary account. Import can be done by requesting ‘/import?user=id’.

## Endpoints
GET /add/<item id> - For non-logged in user
GET /import?user=id - Import non-logged in user data to a logged in user
POST /add/<item id> - For logged in user
PATCH /items?id=<item id>
DELETE /items?id=<item id>

GET /users?user=id
POST /users
POST /sessions
DELETE /sessions/mine


## Appendix

User
UserId int not null auto_increment primary key,
Username varchar(255) not null UNIQUE,
FirstName varchar(128) not null,
LastName varchar(128) not null,
PassHash varbinary(1024) not null,
Email varchar(255) not null UNIQUE,
Phone varchar(255) not null UNIQUE

Task
TaskId int not null auto_increment primary key,
UserId int not null,
TaskName varchar(255) not null UNIQUE,
TaskType varchar(255) not null,
InitTime time not null

Stats
StatsId int not null auto_increment primary key,
TaskId int not null,
CompleteTime time not null,
Completed boolean not null,
Duration time not null
