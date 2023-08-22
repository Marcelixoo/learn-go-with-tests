# Build An Application

The application developed along these chapters consists of a web server where users can track how many games players have won.
- `GET /players/{name}` should return a number indicating the total number of wins
- `POST /players/{name}` should record a win for that name, incrementing for every subsequent POST