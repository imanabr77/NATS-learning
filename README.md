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

  Run the containers using Docker Compose:

   ```sh
    cd Infra 
    docker-compose up -d
    our 
    docker-compose up -d -f NATS-cluster.yml
   ```

  Test Connection NATS node : 

   ```sh

   telnet 127.0.0.1 4222

   ```

> [!NOTE]
> To use the monitoring dashboard, you must copy the JSON file of the dashboard from the Grafana folder and import it into Grafana.



3. **Setup and Running code**

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



**Monitoring dashboard**

![image](https://github.com/imanabr77/NATS-learning/assets/92488673/c2e6060a-e8e9-4090-bc84-e9306e724374)



ی
> [!IMPORTANT]
> Thanks to dear Parham for preparing this video and NATS training. https://github.com/1995parham

Link video : 
https://www.youtube.com/watch?v=EgiNifiCU54

