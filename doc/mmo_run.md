# MMO Run

Approach:
* [X] Build image
* [ ] Export image name to values.yml
* [ ] Deploy base chart that contains other services as dependencies
    * If base chart is already deployed then just make upgrade with new values.yml
    * If not then install base chart
    
## Base chart
Consists of:
* values.yml - contains values of the all services
* requirements.yml - contains application's services as dependencies

## Problems
* One database for all services in dev env vs one database for each of the services in prod env