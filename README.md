├───cmd: Entrypoint of the application
├───external: undecided
├───internal
│   ├───adapters
│   │   ├───app: undecided
│   │   │   ├───cli
│   │   │   └───integrations
│   │   ├───core: undecided
│   │   │   ├───domain
│   │   │   ├───entities
│   │   │   ├───repositories
│   │   │   ├───services
│   │   │   └───usecases
│   │   ├───frameworks
│   │   │   ├───left
│   │   │   │   ├───fasthttp: server endpoints and handlers
│   │   │   │   └───grpc: undecided
│   │   │   │       ├───pb
│   │   │   │       └───proto
│   │   │   └───right
│   │   │       └───databases: implementation of database operations interface for each option
│   │   │           ├───mongodb
│   │   │           ├───mysql
│   │   │           └───redis
│   │   └───plugins: for each plugin, implement core logic using the internal\adapters\app and internal\adapters\core structure
│   │       ├───chesscom
│   │       └───github
│   └───ports
│       ├───app: undecided
│       ├───core: undecided
│       ├───frameworks
│       │   ├───left
│       │   │   ├───fasthttp: handler interfaces
│       │   │   └───grpc: undecided
│       │   └───right
│       │       └───database: database interface
│       └───plugins: common plugin interface
└───tests
|-- requirements.txt
|-- # Dependencies and libraries required for the application
|-- README.md
|-- # Documentation and instructions
