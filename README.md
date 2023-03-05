Backend for [my portfolio website](ssnk.in)
<br>
This backend has several packages:

- server: primary server accepting external requests
  <br>
  <i>Endpoints:

1. biodata: GET
2. githubdata: GET
3. graphql: POST
4. trigger: POST
5. todos: POST
6. schedule: POST
   </i>
   <br>

- todos: returns open todos list from MongoDB
  <br>
  <i>Endpoints:

1. list: POST
2. new: POST
3. done: POST
   </i><br>

- ghintegration: use GitHub REST API and aggregate data and push to MongoDB
  <br>
  <i>Endpoints:

1. trigger
   </i><br>

- auth: handles authorization

- db: handles connection to sql & nosql db
  <br>
  <i>Packages:

1. mongoconnection: MongoDB connection
2. sqlconnection: MySQL connection
   </i><br>

- config: reads and stores all environment variables
- middleware
- common
  <br>

  <b>Note:

  - Only the primary server accepts external requests, todos and ghintegration micro-services accept requests from the server only
  - All REST endpoints will eventually disallow external connections, only /graphql will be available
  - Internal communication between micro services will migrate from REST to gRPC
  - middleware package might be removed, dependent on requirements after moving to gin
  - [gin-contrib/logger](https://github.com/gin-contrib/logger) for path logging, [gin-contrib/cors](https://github.com/gin-contrib/cors) for CORS
    </b>
