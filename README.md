# restli-go
(work in progress)

Rest.li support in Go.

This project proposes a different approach with 2 principles in mind:
 - Simplicity. For example, from the generated code, the developer should be able to send a request and get the response with 1 line of code for simple use cases
 - Small generated code: The generation should limit the number of line of codes


## Feature scope
Here is the list of feature the client will support:

 - Parallelization of asynchronous operations
 - Serialized execution for non-blocking computation
 - Query composition
 - Automatic batching
 - Retry policy
 - Backup request
 - SLA
 - Filters/middleware
 - Caching (based on Cache headers) 

 ## Run

Build and install: `go install`

Generate the files: `restli-go <output directory> <pdsc/idl directory>...`

The `output directory` is the repository where the generated files will be stored.
The `pdsc/idl directory` are the directories that contain the pdsc and idl files to scan. The code will be generated from these files and will be stored in the `output directory` path. 
If an extension file is not recognized, the file will be ignored.

example: `restli-go ./generated ./samples`

## Client

The repo comes from examples to illustrate the code generated.

To run the code generation on the examples: `restli-go ./generated ./samples`

Simple Get Request: `p, err := generated.GetProfile(context.Background(), 1)`
