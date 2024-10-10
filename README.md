# go-web-app-template
Web application project template for new go projects

## Structure
This project structure aligns well with Go best practices and with the principles of hexagonal architecture (or the ports and adapters pattern), emphasizing the separation of core business logic from external dependencies.

The file tree below highlignts the structure used in this template:
```
/my-project
│
├── /cmd                    # Main entry points (applications)
│   └── /main                # Contains the application entry point (e.g., main.go)
│
├── /internal               # Internal logic that is private to your project
│   ├── /core               # Core domain logic and services
│   │   └── /example        # Example domain logic (can be user, order, etc.)
│   │
│   ├── /ports              # Interfaces defining contracts for external components
│   │   └── /example_repo    # Example repository interface (can be used to persist data)
│   │   
│   └── /adapters           # Implementations of ports for external systems
│       └── /example_repo   # Implementation of the example repository (e.g., in-memory, database, etc.)
│
├── /web                    # Web application components
│   ├── /views              # HTML templates or view files
│   ├── /assets             # Static files (CSS, JS, images)
│   ├── /router             # Routing configuration
│   └── /handlers           # HTTP handler functions
│
├── /pkg                 # Shared utilities or libraries
│   └── /example_util       # Example utility functions
│
├── /config                 # Configuration files (YAML, JSON, etc.)
├── /scripts                # Deployment, build, or helper scripts
├── /api                    # API or protocol definitions (if applicable)
├── go.mod
└── go.sum
```
