Backend for [my portfolio website](ssnk.in)
<br>
This backend has several packages:

- server: primary server accepting external requests
  <br>
  <i>Endpoints:

1. biodata: GET
2. githubdata: GET
3. login: POST
4. todos: POST
5. logout: POST <b>planned</b>
6. graphql: POST
7. trigger: POST
8. schedule: POST
9. register: POST
10. temptoken: POST <b>planned</b>
11. resetpassword: POST <b>planned</b>
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

1. trigger: POST
   </i><br>

- auth: handles authorization
  <br>
  <i>Endpoints:

1. signup: POST
2. login: POST <b>planned</b>
3. logout: POST <b>planned</b>
4. temptoken: POST <b>planned</b>
5. resetpassword: POST <b>planned</b>
   </i><br>
   <br>
   <i>Packages:
   1. session: <b>planned</b>
      - limit sessions for user,
      - fetch list of open sesssions,
      - allow session termination before login,
      - invalidate session in a week
        <br>
        <i>Endpoints:
      1. </i><br>

</i><br>

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
