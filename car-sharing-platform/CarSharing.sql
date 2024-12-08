USE master;

-- Check if the database exists, and drop it if it does
IF DB_ID('CarSharingDB') IS NOT NULL
BEGIN
    ALTER DATABASE CarSharingDB SET SINGLE_USER WITH ROLLBACK IMMEDIATE;
    DROP DATABASE CarSharingDB;
END;

-- Create the database
CREATE DATABASE CarSharingDB;

-- Use the newly created database
EXEC('USE CarSharingDB;');

-- Check if the login exists, and create it only if it doesn't
IF NOT EXISTS (SELECT 1 FROM sys.server_principals WHERE name = 'CarSharingUser')
BEGIN
    CREATE LOGIN CarSharingUser WITH PASSWORD = 'CarUser@NP';
END

-- Create a database user for the login and grant permissions
EXEC('USE CarSharingDB; IF NOT EXISTS (SELECT 1 FROM sys.database_principals WHERE name = ''CarSharingUser'')
    BEGIN
        CREATE USER CarSharingUser FOR LOGIN CarSharingUser;
        ALTER ROLE db_owner ADD MEMBER CarSharingUser;
    END;');

-- Create Users table
EXEC('USE CarSharingDB;
CREATE TABLE Users (
    ID INT PRIMARY KEY IDENTITY(1,1),
    Name NVARCHAR(100) NOT NULL,
    Email NVARCHAR(255) NOT NULL UNIQUE,
    Phone NVARCHAR(15) NOT NULL UNIQUE,
    Password NVARCHAR(255) NOT NULL,
    MembershipTier NVARCHAR(50) NOT NULL DEFAULT ''Basic'',
    TotalSpending DECIMAL(10, 2) NOT NULL DEFAULT 0
);');

-- Create Vehicles table
EXEC('USE CarSharingDB;
CREATE TABLE Vehicles (
    ID INT PRIMARY KEY IDENTITY(1,1),
    LicensePlate NVARCHAR(20) NOT NULL UNIQUE,
    Model NVARCHAR(100) NOT NULL,
    ChargeLevel INT NOT NULL,
    Status NVARCHAR(50) NOT NULL DEFAULT ''Available'',
    Location NVARCHAR(255) NOT NULL
);');

-- Create Bookings table
EXEC('USE CarSharingDB;
CREATE TABLE Bookings (
    ID INT PRIMARY KEY IDENTITY(1,1),
    UserID INT NOT NULL FOREIGN KEY REFERENCES Users(ID),
    VehicleID INT NOT NULL FOREIGN KEY REFERENCES Vehicles(ID),
    StartTime DATETIME NOT NULL,
    EndTime DATETIME NOT NULL,
    TotalPrice DECIMAL(10, 2) NOT NULL
);');

-- Create RentalHistory table
EXEC('USE CarSharingDB;
CREATE TABLE RentalHistory (
    ID INT PRIMARY KEY IDENTITY(1,1),
    UserID INT NOT NULL FOREIGN KEY REFERENCES Users(ID),
    VehicleID INT NOT NULL FOREIGN KEY REFERENCES Vehicles(ID),
    RentalDate DATETIME DEFAULT GETDATE(),
    AmountSpent DECIMAL(10, 2) NOT NULL
);');

-- Create Payments table
EXEC('USE CarSharingDB;
CREATE TABLE Payments (
    ID INT PRIMARY KEY IDENTITY(1,1),
    BookingID INT NOT NULL FOREIGN KEY REFERENCES Bookings(ID),
    TransactionID NVARCHAR(100) NOT NULL,
    Amount DECIMAL(10, 2) NOT NULL,
    Status NVARCHAR(50) NOT NULL,
    CreatedAt DATETIME DEFAULT GETDATE()
);');

-- Create Invoices table
EXEC('USE CarSharingDB;
CREATE TABLE Invoices (
    ID INT PRIMARY KEY IDENTITY(1,1),
    UserID INT NOT NULL FOREIGN KEY REFERENCES Users(ID),
    VehicleID INT NOT NULL FOREIGN KEY REFERENCES Vehicles(ID),
    StartTime DATETIME NOT NULL,
    EndTime DATETIME NOT NULL,
    TotalPrice DECIMAL(10, 2) NOT NULL,
    CreatedAt DATETIME DEFAULT GETDATE()
);');

-- Insert test data into Users
EXEC('USE CarSharingDB;
INSERT INTO Users (Name, Email, Phone, Password, MembershipTier, TotalSpending)
VALUES
(''John Doe'', ''john.doe@example.com'', ''1234567890'', ''hashedpassword1'', ''Basic'', 50.00),
(''Jane Smith'', ''jane.smith@example.com'', ''9876543210'', ''hashedpassword2'', ''Premium'', 600.00),
(''Motee'', ''motee@example.com'', ''88888888'', ''hashedpassword3'', ''Basic'', 0.00);');

-- Insert test data into Vehicles
EXEC('USE CarSharingDB;
INSERT INTO Vehicles (LicensePlate, Model, ChargeLevel, Status, Location)
VALUES
(''ABC123'', ''Tesla Model 3'', 90, ''Available'', ''Downtown Station''),
(''XYZ789'', ''Nissan Leaf'', 80, ''Available'', ''Airport Parking'');');

-- Insert test data into Bookings
EXEC('USE CarSharingDB;
INSERT INTO Bookings (UserID, VehicleID, StartTime, EndTime, TotalPrice)
VALUES
(1, 1, ''2024-12-04 10:00:00'', ''2024-12-04 12:00:00'', 30.00),
(2, 2, ''2024-12-05 14:00:00'', ''2024-12-05 16:00:00'', 40.00);');

-- Insert test data into RentalHistory
EXEC('USE CarSharingDB;
INSERT INTO RentalHistory (UserID, VehicleID, AmountSpent)
VALUES
(1, 1, 30.00),
(2, 2, 40.00);');

-- Insert test data into Payments
EXEC('USE CarSharingDB;
INSERT INTO Payments (BookingID, TransactionID, Amount, Status)
VALUES
(1, ''txn_12345'', 30.00, ''Completed''),
(2, ''txn_67890'', 40.00, ''Pending'');');

-- Insert test data into Invoices
EXEC('USE CarSharingDB;
INSERT INTO Invoices (UserID, VehicleID, StartTime, EndTime, TotalPrice)
VALUES
(1, 1, ''2024-12-09 10:00:00'', ''2024-12-09 12:00:00'', 15.00),
(2, 2, ''2024-12-10 14:00:00'', ''2024-12-10 16:00:00'', 20.00);');

-- Validate Data
EXEC('USE CarSharingDB; SELECT * FROM Users;');
EXEC('USE CarSharingDB; SELECT * FROM Vehicles;');
EXEC('USE CarSharingDB; SELECT * FROM Bookings;');
EXEC('USE CarSharingDB; SELECT * FROM RentalHistory;');
EXEC('USE CarSharingDB; SELECT * FROM Payments;');
EXEC('USE CarSharingDB; SELECT * FROM Invoices;');
