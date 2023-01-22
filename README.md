# Go Webapp Project

## Overview

This is a simple webapp written in Go which serves static HTML pages. Gorilla mux is used as the router on the backend and static HTML is served using Go's template engine.

The app provides a basic functionality to leave one-way feedback and view it. After a user registers and signs in, they can start creating feedback. Each entry is appended and the user can view all of the feedback history. An in-memory database is used to record the user information and feedback data.

## Running Locally

To run the app locally, clone the repository and start the app with the following terminal command

`$ go run .`

## Considerations

This is a very simple webapp to demonstrate functional knowledge of Go for developing web-apps and backend APIs. Here are a few thoughts on my implementation and some reasons for a few of my decisions

- Given that the webapp serves static HTML, the auth mechanism is rather basic. A production application may use JWT or other sophisticated mechanisms, however, it wasn't possible to implement it here since saving JWT headers are purely the client's responsibility
- The app uses user information stored in client session to query the database. I would have like to use URI based mechanism to serve user specific content i.e. to show user `foo`'s data, the client would send a `GET` request to e.g. `/feedback/foo` which would then fetch the appropriate resource from the server
- The app should paginate the feedback data
- There needs to be robust validation on form and input fields
- Very basic logging is implemented. I added it just for the sake of it because of the nature of the app being very simple
- Observability, metrics and tracing are a must for a production application. Lot's of options for this stack e.g. Grafana, Prometheus tend to work well in my experience
- Panic recovery middleware is missing
- I would prefer to implement logging at middleware level
- The front-end is only serviceable. I urge you to not judge me for my HTML/CSS proficiency(or for the lack of it)
- I like to structure code in accordance with [screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html)
- I didn't write any unit tests purposefully because I didn't want to spend time writing them

Finally, almost all of the above points can be attributed to the app's nature being rather simple which production applications are seldom not.
