# NATS-learning
This repo is for learning the basic concepts of NATS-Server.

> [!NOTE]
> In this repo, three strategies, request-response, queue-subscribe, and subscription, have been implemented to learn NATS concepts.



## Setup Infra

1. **Clone the Repository**

    ```sh
    git clone git@github.com:imanabr77/NATS-learning.git
    cd NATS-learning
    ```
2. **Build and Run the Containers**

    Build the Docker images and run the containers using Docker Compose:

    ```sh
    cd Infra 
    docker-compose up -d
    our 
    docker-compose up -d -f NATS-cluster.yml
    ```


> [!NOTE]
> To use the monitoring dashboard, you must copy the JSON file of the dashboard from the Grafana folder and import it into Grafana.



2. **Setup and Running code**
   ```sh
   go get 'github.com/nats-io/nats.go'

   # Ensure the go.mod file is up-to-date
   go mod tidy

   # List all the dependencies
   go list -m all

   # Create a vendor directory with all the dependencies
   go mod vendor

   go run queue-subscribe.go
   go run request-response.go
   go run subscription.go

    ```

