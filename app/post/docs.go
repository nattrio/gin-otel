package post

// swagger:route POST /posts post createPost
// Create a new post.
// responses:
//   200: createPostResponse

// swagger:parameters createPost
type createPostRequest struct {
	// Request body for creating a new post.
	// in: body
	Body CreatePost
}

// Create a new post successfully.
// swagger:response createPostResponse
type createPostResponse struct {
	// in: body
	Body struct {
		// Example: post created
		Message string `json:"message"`
	}
}

// swagger:route GET /posts post getPosts
// Get all posts.
// responses:
//   200: getPostsResponse

// Get all posts successfully.
// swagger:response getPostsResponse
type getPostsResponse struct {
	// in: body
	Body []Post
}

// swagger:route GET /posts/{id} post getPost
// Get a post by ID.
// responses:
//   200: getPostResponse

// swagger:parameters getPost
type getPostRequest struct {
	// Post ID.
	// in: path
	// required: true
	ID string `json:"id"`
}

// Get a post by ID successfully.
// swagger:response getPostResponse
type getPostResponse struct {
	// in: body
	Body Post
}

// swagger:route PATCH /posts/{id} post updatePost
// Update a post by ID.
// responses:
//   200: updatePostResponse

// swagger:parameters updatePost
type updatePostRequest struct {
	// in: path
	ID string `json:"id"`

	// in: body
	Body CreatePost
}

// Update a post by ID successfully.
// swagger:response updatePostResponse
type updatePostResponse struct {
	// in: body
	Body struct {
		// Example: post updated
		Message string `json:"message"`
	}
}

// swagger:route DELETE /posts/{id} post deletePost
// Delete a post by ID.
// responses:
//   200: deletePostResponse

// swagger:parameters deletePost
type deletePostRequest struct {
	// in: path
	ID string `json:"id"`
}

// Delete a post by ID successfully.
// swagger:response deletePostResponse
type deletePostResponse struct {
	// in: body
	Body struct {
		// Example: post deleted
		Message string `json:"message"`
	}
}
