# Concurrency and Reservation Strategy

## Handling Reservations
[cite_start]When `available_copies` reaches 0, the system prevents further checkouts and informs the user that they have been added to a reservation queue.

## Preventing Race Conditions
[cite_start]To prevent "over-booking" (two people taking the last book at once), we use a logic check on the `AvailableCopies` count before saving the checkout record to the database[cite: 30, 35].