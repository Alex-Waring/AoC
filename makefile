default:
	go run main.go i $(YEAR) $(DAY)
	go run main.go f $(YEAR) $(DAY)

rust:
	go run main.go i $(YEAR) $(DAY)
	go run main.go rf $(YEAR) $(DAY)
