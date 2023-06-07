weavestore is an in-memory key-value store.

### API:

#### Insert (/insert)
``` To insert a single object into the store```
* /read -> Retrieve a value by it's key
* /delete -> Delete a key
* /update -> Update a value for a key
* /updateBulk -> Update a bulk of keys

### Implementation Specifics:
* Update operation is implemented as an Upsert. Update existing values and inserting new ones.
* I have kept the eviction policy very simple for the assignment. It's a basic linked-list based lru where we evict the objects upon hitting the cache capacity threshold.

### How to run the project:
* Clone the project on your local machine. Go into the root of the project and run `make setup_and_run` to build the docker image and start the server.
* Additionally, you can run `make docs` to build and start a documentation server which will give you a UI link to refer the documentation in a web format. (NOTE: this automation is prone to not work on macOS, since i dont have a mac i am unable to debug the issue.)
* Other automations that are useful -
    - `make test` to run all the tests
    - `make check` to run linters on the entire project
