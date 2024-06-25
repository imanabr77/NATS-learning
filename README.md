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
    ```

> [!NOTE]
> To use the monitoring dashboard, you must copy the Json file of the dashboard from the Grafana folder and import it into Grafana.
