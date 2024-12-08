# Car Sharing Platform

This is a Car Sharing Platform that allows users to rent vehicles, view available cars, manage bookings, and make payments. The system is designed with a microservices architecture to handle various functionalities like user management, vehicle management, booking management, and payment processing.

---

## **Design Considerations**

### **Microservices Architecture**

- **User Service**: Manages user registration, login, authentication, and user profile management. It interacts with the `Users` table in the database and ensures that users' details are kept secure.
- **Vehicle Service**: Handles the vehicle inventory, manages car availability, vehicle booking, and related actions. It communicates with the `Vehicles` and `Bookings` tables to manage vehicle status and reservations.
- **Billing Service**: Manages the pricing model, calculates billing based on membership tiers, generates invoices, and processes payments. This service is integrated into the `Users`, `Bookings`, and `Payments` tables.
- **Payment Service**: Processes payments and manages payment methods, though for this project, the system currently only generates invoices without handling actual payment transactions.

### **Database Structure**

- The system uses a centralized database (`CarSharingDB`) which includes tables for `Users`, `Vehicles`, `Bookings`, `RentalHistory`, and `Invoices`.
- The relationships between these tables are handled with foreign keys to ensure data consistency.
  
---

## **Architecture Diagram**


- **User Service**: Handles registration, login, and account management.
- **Vehicle Service**: Manages vehicle availability and booking.

---

## **Instructions for Setting Up and Running the Microservices**

### **Prerequisites**

1. **Go**: Make sure Go is installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2. **MySQL** or **MSSQL**: Make sure you have a working instance of MySQL or MSSQL for your database.
3. **Postman**: For testing the API endpoints.
3. **Curl**: For testing the API endpoints.


### **Steps to Set Up**

1. **Clone the repository**:
   ```bash
   git clone 
   cd car-sharing-platform
2. **Run the User Service**:
   Navigate to the user-service directory.
   Run the service: go run main.go
   This service will run on http://localhost:8081.
3. **Run the Vehicle Service**:
   Navigate to the vehicle-service directory.
   Run the service: go run main.go
   This service will run on http://localhost:8082.
4. **Testing the API Endpoints**:
   Use Postman or curl commands to test the following:
   Register a new user.
   Login and get the session cookie.
   Create and manage bookings for vehicles.
   Generate invoices for completed bookings.
5. **Shutdown**:
   To stop the services, simply press Ctrl + C in the terminal running each service.

## **Running the Services Locally**
**The user-service and vehicle-service run on different ports**:
   User Service: http://localhost:8081
   Vehicle Service: http://localhost:8082
   Make sure to run both services simultaneously for the system to function properly. Use Postman or curl for interaction with the services.


## **Test Data**

1. **Register a new user**
**This command registers a new user with their name, email, phone, and password.**
curl -X POST http://localhost:8081/auth/register -H "Content-Type: application/json" -d '{
  "name": "Kai",
  "email": "ka@example.com",
  "phone": "88888888",
  "password": "kj"
}'

2. **Login**
**This command logs the user in by providing email and password. The session is saved in a cookies file.**
curl -X POST http://localhost:8081/auth/login -H "Content-Type: application/json" -d '{
  "email": "ka@example.com",
  "password": "kj"
}' -c ../cookies.txt

3. **Logout**
**This command logs the user out and clears the session.**
curl -X GET http://localhost:8081/auth/logout -b ../cookies.txt -c ../cookies.txt

4. **Access Account**
**This command accesses the account information of the logged-in user.**
curl -X GET http://localhost:8081/account -b ../cookies.txt

5. **Access Rental History**
**This command fetches the rental history of the logged-in user.**
curl -X GET http://localhost:8081/rental-history -b ../cookies.txt

6. **Get Vehicles**
**This command fetches a list of available vehicles from the vehicle service.**
curl -X GET http://localhost:8082/vehicles

7. **Create Booking**
**This command creates a booking for a vehicle by providing user ID, vehicle ID, start time, end time, and total price.**
curl -X POST http://localhost:8082/bookings -H "Content-Type: application/json" -d '{
  "user_id": 3,
  "vehicle_id": 1,
  "start_time": "2024-12-09T10:00:00Z",
  "end_time": "2024-12-09T12:00:00Z",
  "total_price": 30.00
}'

8. **Edit Booking**
**This command modifies an existing booking by changing the vehicle or the start time.**
curl -X PUT http://localhost:8082/bookings/1 -H "Content-Type: application/json" -d '{
  "vehicle_id": 2,
  "start_time": "2024-12-10T12:00:00Z"
}'

9. **Cancel Booking**
**This command cancels an existing booking based on the booking ID.**
curl -X DELETE http://localhost:8082/bookings/2

11. Generate Invoice
**This command generates an invoice after a successful booking, which includes details like user ID, vehicle ID, and total price.**
curl -X POST http://localhost:8082/billing/invoice -H "Content-Type: application/json" -d '{
  "user_id": 3,
  "vehicle_id": 1,
  "start_time": "2024-12-09T10:00:00Z",
  "end_time": "2024-12-09T12:00:00Z",
  "total_price": 30.00
}'

13. **View All Bookings for a User**
**This command retrieves all bookings made by the logged-in user.**
curl -X GET http://localhost:8082/bookings -b ../cookies.txt


### **Conclusion**
This Car Sharing Platform is built with a microservices architecture to ensure scalability and modularity. Each service is designed to handle specific functionality, ensuring maintainability and extensibility. The platform is ready for further enhancements, including payment processing, real-time billing updates, and user-specific vehicle recommendations.

