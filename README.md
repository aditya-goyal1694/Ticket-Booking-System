# Ticket Booking System

A high-performance, concurrent ticket booking system built in Go. This system ensures efficient ticket allocation using Goroutines for asynchronous ticket confirmation while maintaining data integrity.

## Features
- **Fast & Concurrent**: Uses Goroutines for sending tickets asynchronously.
- **User-Friendly Input Handling**: Ensures valid names, emails, and ticket requests.
- **Dynamic Seat Management**: Tracks remaining tickets in real-time.
- **Thread-Safe Operations**: Prevents race conditions while updating bookings.
- **Graceful Exit**: Closes when all tickets are sold out.

## Installation & Setup
### Prerequisites
- Install [Go](https://golang.org/dl/) (version 1.21 or later recommended)

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/ticket-booking-system.git
   cd ticket-booking-system
   ```
2. Initialize a Go module (if not already initialized):
   ```sh
   go mod init ticket-booking-system
   ```
3. Build and run the application:
   ```sh
   go run main.go
   ```

## Usage
1. Enter your details when prompted.
2. Book the required number of tickets.
3. Receive a confirmation email (simulated).

### Example Booking Flow:
```sh
Welcome to Go Conference booking app!
Get your tickets here. Only 50 tickets are available.
Enter your first name: John
Enter your last name: Doe
Enter your email: john.doe@example.com
Enter the no. of seats you want to book: 2
Thank you John Doe for booking 2 tickets. You will receive a confirmation email shortly.
```

## Concurrency & Performance
- **Goroutines**: Handles ticket confirmation asynchronously.
- **WaitGroup**: Ensures all operations complete before program exit.

## Why Choose This Project?
- Optimized for speed and efficiency  
- Thread-safe operations ensure data consistency  
- Built using best practices in Go

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
