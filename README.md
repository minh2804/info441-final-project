# Project Description

People get busier and busier as they get older - it is a known fact that there is a direct correlation between age and number of responsibilities. With busy lifestyles comes absent mindedness - how are we supposed to remember to do all these things? Our app attempts to solve this issue by relieving the stress of having to remember every little task.

Our general target audience will be for people that have busy lifestyles. More specifically, millennials who need a better way to be reminded of what is due. While there are other to-do lists on the market, they usually save their most valuable features for their paying customers. Other apps are free, but they usually aren’t made specifically as todo lists. Neither of the aforementioned solutions allow the user to see their metrics. As a result, there is a real need for a simple and reliable free todo list tracker.

As developers, we want to develop an application that serves our target audience by providing a central place to keep track of their tasks, efficiency and performance. In addition, our target audience will be able to complete their tasks more efficiently by using our task completion visualization and reminder.

# Technical Description

![technical diagram](media/diagram.png)

| Priority | User                    | Description                                                                                         |
| -------- | ----------------------- | --------------------------------------------------------------------------------------------------- |
| P0       | As an email sender      | I want to create a shareable link, to share it with other people.                                   |
| P1       | As an email recipient   | I want to click on a link in the email, and it will automatically add to my reminder list.          |
| P2       | As a non-logged in user | I want to create a bunch of reminders, and log in later.                                            |
| P3       | As a logged in user     | I want to be able to view my stats (items completed in a week, items overdue, longest streaks, etc) |

## Implementation Strategy

Email sender: The sender can create a shareable link to attach to their email. The shareable link will have the following resource path: ```/add/<item id>```. We will store the link’s unique random string in our ***MySQL***.

Email recipient: The email recipient can click on the ```/add/<item id>``` and it will add to their reminder list. If the user is logged in, then it will do it automatically.

P2: If a user clicks on a shareable link and is not logged in, then it will ask the user for an email through an ***HTML*** form. Then a temporary account will be created in our ***MySQL database*** to store the item, login credentials of this temporary account will be sent to the provided email.

P3: The user can import the data from another account, including the temporary account, through our ***SQL database***. Reminders will also be held in our ***SQL database*** and then populated on the front end using ***React***. Import can be done by requesting ‘/import?user=id’.

Stats: The stats will be pulled from the ***SQL database*** and then transferred onto the front end using ***React***

## Endpoints

* ```/tasks```: refers all tasks
  * ```GET```: Reponse with a JSON array of the current user's todo list and status code ```200```. If the user is logged in, then the todo list is retrieved from the database. If the user is not logged in, then it will retrieve the todo list from current session. If the user is not found, then it will response with status code ```404```.
  * ```POST```: Create a new task, then response with a JSON object of the newly-created task and status code ```201```. If the request's content type is not ```application/json```, then it will response with status code ```415```. If the user is logged in, then the task is added to the database. If the user is not logged in, then it will add to the current session. If the task already existed or an invalid post is provided, then it will repsonse with status code ```400```.
* ```/tasks/{taskID}```: refers to a specific task
  * ```GET```: Reponse with a JSON object of the requested task and status code ```200```. If the task is not found, then it will response with status code ```404```. If the task is hidden and the currently logged user is not the owner, then it will response with status code ```401```.
  * ```PATCH```: Update an existing task, then response with a JSON object of the newly-updated task and status code ```200```. If the request's content type is not ```application/json```, then it will response with status code ```415```. If the user is logged in, then the task is updated to the database. If the user is not logged in, then it will update to the current session. If the task is not found, then it will response with status code ```404```. If the update object is invalid, then it will response with status code ```400```.
  * ```DELETE```: Delete an existing task, then response with status code ```200```. If the task is not found, then it will response with status code ```404```.
* ```/tasks/import/{userID}```: refers to importing another user's todo list
  * ```GET```: Reponse with a JSON array of the requested user's todo list and status code ```200```. If the user is not found, then it will response with status code ```404```. If the user is logged in, then the requested user's todo list will be added to the current user's todo list on the database. If the user is not logged in, then the requested user's todo list will be added to the current user's todo list on the current session. Any tasks that is marked as hidden from the requested user's todo list will not be added to the current user's todo list.
* ```/stats```: refers to all stats of the current user
  * ```GET```: Response with a JSON object of the current user's stats. If the user not logged in, then it will response with status code ```401```.
* ```/stats/{property}?start={startDate}&end={endDate}```: refers to specific properties of a stat of the current user
  * ```GET```: Reponse with a JSON object of the requested stats property. ```property``` can be ```year```, ```month```, ```week```, and ```custom```. Only ```custom``` property accepts ```start``` and ```end``` query arguments. If the user it not logged in, then it will response with status code ```401```.
* ```/users```: refers to all users
  * ```POST```: Create a new user, then response with a JSON object of the newly-created user and status code ```201```. If the user already existed or an invalid user is provided, then it will response with status code ```400```.
* ```/users/{userID}```: refers to a specific user
  * ```GET```: Response with a JSON object of the requested user and status code ```200```. If the user is not logged in, then it will response with status code ```401```. If the user is not found, then it will response with status code ```404```. The ```userID``` can be ```me``` which refers to the current user.
  * ```PATCH```: Update an existing user, then response with a JSON object of the newly-updated user and status code ```200```. If the request's content type is not ```application/json```, then it will response with status code ```415```. If the requested user is not the current user or ```userID``` is not equal to ```me```, then it will response with status code ```403```. If the update object is invalid, then it will response with status code ```400```.
* ```/sessions```: refers to all sessions
  * ```POST```: Create a new session, then response with a JSON object of the logged-in user and status code ```201```. If the request's content type is not ```application/json```, then it will response with status code ```415```. If the credentials is invalid, then it will response with status code ```401```.
* ```/sessions/{sessionID}```: refers to a specific session
  * ```DELETE```: Delete the requested session, then response with status code ```200```. If the ```sessionID``` is not equal to ```mine```, then it will return with status code ```403```.

## Appendix

```
CREATE TABLE IF NOT EXISTS User (
	ID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	Username  VARCHAR(255) NOT NULL UNIQUE,
	PassHash  CHAR(72) NOT NULL,
	FirstName VARCHAR(255),
	LastName  VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS TodoList (
	ID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	UserID BIGINT NOT NULL,
	Name VARCHAR(255) NOT NULL,
	Description VARCHAR(1000),
	IsComplete BOOL NOT NULL,
	IsHidden BOOL NOT NULL,
	CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	EditedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
	FOREIGN KEY (UserID) REFERENCES User (ID)
);

DELIMITER //
CREATE PROCEDURE DeleteUser (IN p_ID BIGINT)
BEGIN
	IF NOT EXISTS (SELECT * FROM User WHERE ID = p_ID) THEN
		SIGNAL SQLSTATE '45000'
			SET MESSAGE_TEXT = 'user not found';
	END IF;
	START TRANSACTION;
	DELETE FROM TodoList WHERE UserID = p_ID;
	DELETE FROM User WHERE ID = p_ID;
	COMMIT;
END;//
DELIMITER ;
```
