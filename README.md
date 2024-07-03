# Simple Shopping List using GoLang

This repository contains a GoLang project developed during my 4th semester as part of an extra course on Go. The project demonstrates various aspects of Go programming, including its concurrency model, standard libraries, and best practices for structuring Go applications.

## Features

- **Concurrency:** Implementation of Go routines and channels to handle multiple tasks simultaneously.
- **Standard Libraries:** Utilization of Go's rich standard libraries for tasks such as file handling, HTTP requests, and data manipulation.
- **Project Structure:** Organized code following Go's idiomatic project structure for easy maintenance and scalability.
- **Error Handling:** Robust error handling to ensure smooth execution and debugging.

## Code Overview

The core functionality of this project is a simple task management system implemented using the [Gin](https://github.com/gin-gonic/gin) web framework. The project includes the following key features:

1. **Task Struct**: Defines a `Task` struct with fields for ID, name, and completion status.
2. **In-Memory Storage**: Stores tasks in an in-memory slice, simulating a simple database.
3. **HTTP Endpoints**:
   - `GET /tasks`: Retrieves the list of all tasks.
   - `POST /tasks`: Adds a new task to the list.
   - `PUT /tasks/:id`: Updates the completion status of a specific task based on its ID.
4. **Static File Serving**: Serves a static HTML file and related assets for the front-end interface.

## Frontend Overview

The frontend is a simple shopping list application implemented with HTML, CSS, and JavaScript. Key features include:

1. **Task List Container**: Displays the list of tasks dynamically.
2. **Task Input Field**: Allows users to input new tasks.
3. **Add Item Button**: Adds new tasks to the list.
4. **Task Item**: Displays individual tasks with checkboxes for marking completion and buttons for deleting tasks.

### HTML Structure

The HTML file includes:
- A container for the task list.
- An input field for new tasks.
- A button to add new tasks.

### CSS Styling

The CSS provides:
- General styling for the body and headings.
- Styling for the task list container, input field, buttons, and task items.

### JavaScript Functionality

The JavaScript handles:
- Adding new tasks to the list.
- Deleting tasks from the list.
- Updating the task list display.

