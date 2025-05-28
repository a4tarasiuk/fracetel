# FRaceTel - F1 2022 Add-on
### Telemetry Data Tracker, Analytics Application 

### Goal
Initial goal is creation of pet-project for practising different tech-stack (especially Go).  
From the technical perspective this is a relatively interesting task and as for pet-project 
this is the most real-world like case.   
F1 2022 Game has a functionality to output game-data to external UDP connections.  

So idea is in creating UDP server that can listen to F1 UDP packets, parse them and publish to messaging system for future processing.
Next processing is aggregating/extracting useful information for storage and future analysis.  
When there is a pretty much amount of interesting data it does not take much time to create use-cases around it.

[F1 Telemetry Spec](https://answers.ea.com/t5/F1-22/F1-22-UDP-Specification/m-p/11551319).

### Description

This project is not expected to be used by real game players.  
My real-game cases are:
* Persist session lap-times. During time-attack sessions my lap can be deleted due to crossing the track limits, and there are cases when I don't agree with the penalty and still want to fight against that lap time :)
* Track the whole history of best lap for individual tracks (within race and qualifying/time-attack sessions)
* TODO:

Features:
* Tracks user session:
  * Persists lap times
  * Persists results
* Provides an interface for monitoring telemetry data in real-time
* TODO:

## Architecture
TODO:

## Tech Stack
Used technologies:
* Golang
* NATS
* Postgres
* OpenTelemetry
* Prometheus
* Grafana
