# Healthcare Patient Management API

[![Go](https://img.shields.io/badge/Go-1.x-blue?logo=go)](https://golang.org)  
[![MongoDB Atlas](https://img.shields.io/badge/MongoDB-Atlas-Cloud-green?logo=mongodb)](https://www.mongodb.com/cloud/atlas)  


## Table of Contents
- [Overview](#overview)  
- [Features](#features)  
- [Technology Stack](#technology-stack)  
- [Getting Started](#getting-started)  
  - [Prerequisites](#prerequisites)  
  - [Setup Instructions](#setup-instructions)  
- [API Endpoints](#api-endpoints)  
- [Request & Response Examples](#request--response-examples)  
- [Project Structure](#project-structure)   

---

## Overview  
This project is a RESTful API built in Go (Golang) for managing basic healthcare patient records. It supports Create, Read, Update, and Delete (CRUD) operations for patients and uses MongoDB Atlas for storage. The API is designed to be clean, maintainable, and ready for extension (e.g., adding authentication, search, pagination).

---

## Features  
- Create new patient records  
- Retrieve all patient records  
- Retrieve a specific patient by ID  
- Update an existing patient record  
- Delete a patient record   
- Docker-free, lightweight backend using Gin + MongoDB Atlas  

---

## Technology Stack  
- **Backend:** Go (Golang)  
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)  
- **Database:** [MongoDB Atlas (M0 Free Tier)](https://www.mongodb.com/cloud/atlas)  

---

## Getting Started  

### Prerequisites  
- Go installed (1.18+ recommended)  
- A MongoDB Atlas free cluster and connection URI  
- Git installed  

### Setup Instructions  
1. Clone this repository:  
   ```bash
   git clone https://github.com/Mizbain-Fathima/Healthcare-patient-management.git
   cd Healthcare-patient-management
```

2. Create a `.env` file in the project root and add your MongoDB URI:

   ```env
   MONGO_URI=mongodb+srv://<username>:<password>@<cluster>.mongodb.net/?retryWrites=true&w=majority
   ```

3. Ignore the `.env` file by adding it to `.gitignore` so credentials don’t get committed.

4. Install dependencies:

   ```bash
   go mod tidy
   ```

5. Generate Swagger documentation (if not generated yet):

   ```bash
   swag init
   ```

6. Run the API server:

   ```bash
   go run main.go
   ```

7. Visit the API:
`
API Root: `http://localhost:8080`

---

## API Endpoints

| Method | Endpoint         | Description           |
| ------ | ---------------- | --------------------- |
| POST   | `/patients`      | Create a new patient  |
| GET    | `/patients`      | Get all patients      |
| GET    | `/patients/{id}` | Get one patient by ID |
| PUT    | `/patients/{id}` | Update patient by ID  |
| DELETE | `/patients/{id}` | Delete patient by ID  |

---

## Request & Response Examples

### Create Patient

**Request**

```json
{
  "name": "Jane Doe",
  "age": 30,
  "diagnosis": "Flu",
  "phone_number": "9876543210"
}
```

**Response (201 Created)**

```json
{
  "id": "653a2f3a4b9c45d1234abcd5",
  "name": "Jane Doe",
  "age": 30,
  "diagnosis": "Flu",
  "phone_number": "9876543210"
}
```

### Get All Patients

**Response (200 OK)**

```json
[
  {
    "id": "653a2f3a4b9c45d1234abcd5",
    "name": "Jane Doe",
    "age": 30,
    "diagnosis": "Flu",
    "phone_number": "9876543210"
  }
]
```

---

## Project Structure

```
.
├── controllers/
│   └── patient_controller.go
├── database/
│   └── connection.go
├── docs/
│   └── (generated Swagger files)
├── models/
│   └── patient.go
├── routes/
│   └── patient_routes.go
├── .env
├── .gitignore
├── go.mod
└── main.go
```

---