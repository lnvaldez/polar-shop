# Polar Shop

A full-stack e-commerce application built with Go (frontend) and Node.js (backend).

## Prerequisites

Before starting, ensure you have the following installed:

- Node.js (v14 or higher) - [Download](https://nodejs.org/)
- Go (v1.16 or higher) - [Download](https://golang.org/dl/)
- MongoDB (v4.4 or higher)
  - Option 1: [Local Installation](https://docs.mongodb.com/manual/installation/)
  - Option 2: [MongoDB Atlas](https://www.mongodb.com/cloud/atlas/register)

## Project Structure

```
polar-shop/
├── backend/         # Node.js backend
├── frontend/        # Go frontend
└── README.md
```

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/lnvaldez/polar-shop.git
cd polar-shop
```

### 2. Backend Setup

Navigate to the backend directory and install dependencies:

```bash
cd backend
npm install
```

#### Set up environment variables:

1. Locate the `.env.sample` file in the backend directory
2. Rename `.env.sample` to `.env`
3. Update the values in `.env` with your configuration:

```env
MONGO_URI=your_mongodb_connection_string
EXPRESS_PORT=5000
SESSION_SECRET=your_session_secret
JWT_SECRET=your_jwt_secret
```

> **Note**: Replace `MONGO_URI` with your MongoDB Atlas connection string if using Atlas.

### 3. Frontend Setup

Navigate to the frontend directory:

```bash
cd ../frontend
```

#### Set up environment variables:

1. Locate the `.env.sample` file in the frontend directory
2. Rename `.env.sample` to `.env`
3. Update the values in `.env` with your configuration:

```env
MONGO_URI=your_mongodb_connection_string
DB_NAME=polar_shop
PRODUCTS_COLLECTION=products
ORDERS_COLLECTION=orders
SERVER_PORT=5001
```

### 4. Start the Services

1. Start the backend server (in the `backend` directory):

```bash
npm run start
```

2. In a new terminal, start the frontend server (in the `frontend` directory):

```bash
go run main.go
```

### 5. Initial Admin Setup (Required)

1. Access the registration page:

   ```
   http://localhost:5001/register
   ```

2. Create your user account with your desired credentials

3. Set Admin Role:

   - Open MongoDB Compass or Atlas
   - Navigate to the database
   - Find the `users` collection
   - Locate your user document
   - Update the `role` field to `"admin"`

4. Login to your admin account:
   ```
   http://localhost:5001/login
   ```

## Accessing the Application

- Frontend: `http://localhost:5001/products`
- Admin Dashboard: `http://localhost:5001/admin` (requires admin login)

## Troubleshooting

### Common Issues

1. **MongoDB Connection Failed**

   - Verify your MongoDB service is running
   - Check the `MONGO_URI` in both `.env` files
   - Ensure network connectivity to MongoDB Atlas (if using cloud)

2. **Port Conflicts**

   - Ensure ports 5000 and 5001 are available
   - Check for other services using these ports
   - Modify `.env` files if you need to use different ports

3. **Admin Access Issues**
   - Confirm your user's `role` field is exactly `"admin"` in MongoDB
   - Clear browser cookies and try logging in again
   - Check backend logs for authentication errors

## Development Requirements

- Frontend:

  - Go v1.16+
  - Standard Go libraries
  - MongoDB Go driver

- Backend:
  - Node.js v14+
  - npm packages (see package.json)
  - MongoDB Node.js driver
