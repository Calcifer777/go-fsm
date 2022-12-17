
- Steteless finite state machine framework web service
- Transactions can be request-based or event-based
- Possibility to export the workflow to image
- Load state machine configuration from file
- Validate file FSM configuration
- NO app logic implemented internally
- Open points
  - correlation-id
  - State backup in Redis?
  - how to implement sleep?


## Resources

### Golang

- https://gobyexample.com/

### Graphs & FSM

- https://pkg.go.dev/gonum.org/v1/gonum/graph?utm_source=godoc
- https://towardsdatascience.com/writing-a-finite-state-machine-in-go-e5535e89d615
- https://github.com/qmuntal/stateless

### GRPC

- https://itnext.io/build-grpc-server-with-golang-go-step-by-step-b3f5abcf9e0e
- https://grpc.io/docs/