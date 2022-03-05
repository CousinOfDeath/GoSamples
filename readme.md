
# Coverage calculated and displayed in a terminal
go test -coverprofile=coverage.out

# Coverage calculated and visualized in a browser. Green indicates covered parts and red uncovered parts.
go tool cover -html=coverage.out
