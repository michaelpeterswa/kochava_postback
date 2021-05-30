# kochava_postback

Simulation of a method Kochava uses to distribute data to third parties in real time.

## Planning

The first step was to set up my development environment. Based upon the criteria, I initially created one monorepo with a folder for the delivery and ingest agents. Normally, I would separate these into two different repositories (like how I did with the GU AR Walking Tour Project). The monorepo solution ended up working well for this project. Inside the ingestion and delivery folders are Dockerfile's specific to that agent.

Central to the entire project is the Redis queue. This is managed using Redis Lists. The operations LPush and BRPop create a FIFO queue going from left to right. BRPop ensures that access is blocked to the list until the operation is complete. This helps prevent any access timing issues for the specific key being popped from. The ingest program (PHP) LPush'es the postback object when it is recieved from a POST request. As fast as possible, the delivery agent (Go) pulls objects from out of the queue and quickly processes them, finally sending a request to the formatted response URL.

## Problems Encountered

The first major issue that required me to make a design choice was the Unmatched Key specification (in Extra Credit). In a language like Python, JSON objects can be natively converted to Dictionaries, whereas Go requires a stricter Type-safe implementation using structs. The only other way in Go is to map the values, which increases complexity and removes the type-safety of the application. Given that, and further information from https://bencane.com/2020/12/08/maps-vs-structs-for-json/, I chose to implement a struct that represents the ingested JSON. If more fields are needed in the future (beyond "location" and "masccot") then they should be added to the struct. The value for unmatched keys are now set to the default as specified in the constructURL() utility. This ensures that other values can be safely passed through the ingest and out through delivery while maintaining safe and compliant URL's.

Another problem was encountered in setting up PHP's Docker container. I have a very small amount of familiarity with PHP and the installation on my computer took a fair amount of effort and troubleshooting. I wanted to streamline the install in Docker though, and eventually settled on a PHP image that includes the Apache webserver. This made the install much smoother, and only required installing the Redis extension on top of the base image. I have found that this works well, especially now that I am utilizing a multi-container architecture (See Docker section below).

## Docker

At the beginning I was slightly confused how I would be able to structure and distribute this project from a single Docker container, like Ubuntu:14.04. The Docker documentation says, "It is generally recommended that you separate areas of concern by using one service per container". This is the method that I prefer, as it allows for better horizontal scaling and much safer failures (that don't take the entire application down). In the end, I settled for using three separate containers, one for delivery, one for ingestion, and the other for Redis. These three containers can be easily managed by the scripts in the "scripts" folder as well as through the Makefile.
