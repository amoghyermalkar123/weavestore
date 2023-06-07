weavestore is an in-memory key-value store. It has an HTTP API that can be used programatically to perform CRUD operations on the store. Refer to the API below. You can find the documentation reference at the bottom of the readme as well.

### API:

#### Insert (http://localhost:8000/insert/)
``` To insert a single object into the store```
#### Read (http://localhost:8000/read/)
``` Retrieve a value by it's key ```
#### Delete (http://localhost:8000/delete/)
``` Delete a key ```
#### Update (http://localhost:8000/update/)
``` Update a value for a key ```
#### BulkUpdate (http://localhost:8000/updateBulk/)
``` Update a bulk of keys ```

### Implementation Specifics:
* Update operation is implemented as an Upsert. Update existing values and inserting new ones.
* I have kept the eviction policy very simple for the assignment. It's a basic linked-list based lru where we evict the objects upon hitting the cache capacity threshold.
* For the tests, I have used the ginkgo framework, the thought here is to mimick behavior based testing. I have kept the tests straightforward and have covered the core set of features since it would take alot more time to come up with edge cases and write tests for the same.
* I have also provided a docker image to demonstrate a complete version of the assignment.
* On the response codes, I have taken a very basic inspiration from redis's resp protocol.
* for the eviction policy, there might be some edge cases or bugs which I might not have covered since those would take some amount of brainstorming. Hence have used go's List structure to implement a naive LRU based eviction policy.

### How to run the project:
* Clone the project on your local machine. Go into the root of the project and run `make setup_and_run` to build the docker image and start the server.
* Additionally, you can run `make docs` to build and start a documentation server which will give you a UI link to refer the documentation in a web format. (NOTE: this automation is prone to not work on macOS, since i dont have a mac i am unable to debug the issue.)
* Other automations that are useful -
    - `make test` to run all the tests
    - `make check` to run linters on the entire project
