# Sensor Dashboard

## Overview
The **Sensor Dashboard** is a full-stack application that monitors sensor data, such as temperature, AQI (Air Quality Index), and occupancy. The backend is built with **Golang** and utilizes **MySQL** as the database, while the frontend is developed using **React**.

---

## Project Structure
```
📦 sensor-dashboard
├── 📂 sensor-dashboard-backend  # Golang Backend
├── 📂 sensor-dashboard-frontend # React Frontend
├── 📄 README.md                 # Project documentation
```

---

## Backend - Golang
### Prerequisites
Ensure you have the following installed:
- **Go** (latest version recommended)
- **MySQL** (ensure the database is running)
- **Git** (for cloning the repository)

### Configuration
The backend uses a `config.yaml` file to store database credentials and server configurations. Ensure your `config.yaml` is correctly set up before running the backend.

Example `config.yaml`:
```yaml
db:
  host: "localhost"
  port: 3306
  user: "root"
  password: "yourpassword"
  database: "sensor_data"

bind:
  http: ":9000"
```

### Installation & Running the Backend
```sh
# Clone the repository
git clone https://github.com/yourusername/sensor-dashboard.git
cd sensor-dashboard/sensor-dashboard-backend
go mod tidy 


# Run the backend service
go run ./cmd/ cfg=./config.yaml
```

---

## Frontend - React
### Prerequisites
Ensure you have the following installed:
- **Node.js** (latest LTS version recommended)
- **npm** or **yarn**

### Installation & Running the Frontend
```sh
# Navigate to the frontend directory
cd sensor-dashboard-frontend

# Install dependencies
npm install  # or yarn install

# Start the frontend application
npm start  # or yarn start
```

The React application will be available at `http://localhost:3000/` by default.



---

## API Endpoints
### WebSocket Connection
- **URL:** `ws://localhost:8080/ws`
- **Description:** Real-time WebSocket connection for receiving sensor data updates.

### Database Schema (MySQL)
Ensure the following table is created before running the backend:
```sql
CREATE TABLE sensor_data (
    id INT AUTO_INCREMENT PRIMARY KEY,
    temperature FLOAT NOT NULL,
    aqi INT NOT NULL,
    occupancy INT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    reliability FLOAT NOT NULL
);
```

---

## Contribution
1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m "Add new feature"`).
4. Push the branch (`git push origin feature-name`).
5. Open a Pull Request.

---


---

## Contact
For any inquiries or support, please contact [your-email@example.com](manansaini.aza.0999.gmail.com).

