# PostBlog App

**PostBlog** is a RESTful API developed in Go that enables authenticated users to manage blog posts, including creating, reading, updating, and deleting their own posts. It uses MongoDB for data storage, Gorilla Mux for routing, and JWT for secure access control.

### Features
- **User Authentication**: Users can sign up and log in, receiving a JWT token for session-based authentication.
- **Blog Management**: Authenticated users can create posts with a title and content, read all posts or a single post by ID, update their own posts, and delete their own posts.
- **Timestamps**: Each post includes `created_at` and `updated_at` timestamps for tracking purposes.

### Setup

# Getting started
1. Install Golang on your machine.
  1.1. Make sure you have GOPATH set in your environment variables.
  1.2. Ensure it using `echo %GOPATH%`
2.**Clone the Repository**:
   ```bash
   git clone <repository-url>
3. This will take some time because it downloads this project and downloads all the imported dependencies.
4. Run go mod tidy
5. Now, run a mogodb server on your local machine, which by default runs on port :27017.
6. Run `go build` to build the go project in a executable file.
7. Make sure to create your .env file.

# Usage

>>>> Testing can be done using POSTMAN

## Endpoints
1. **Create a blog**  
   - **POST** `/api/blogs/create`
   - Enter three key values in Body (x-www-form-urlencoded): `title`, `content`, `author`

2. **Get all blogs**  
   - **GET** `/api/blogs`
   - Retrieves an array of all blogs.

3. **Get a specific blog**  
   - **GET** `/api/blog/{_id}`
   - Fetches the blog with the specified blog ID.

4. **Delete a specific blog**  
   - **DELETE** `/api/blog/{_id}/delete`
   - Deletes the blog identified by the specified ID.

5. **Update a specific blog**  
   - **PUT** `/api/blog/{_id}/update`
   - Requires three key values in Body (x-www-form-urlencoded): `title`, `content`, `author`
   - Updates the specified blog.
