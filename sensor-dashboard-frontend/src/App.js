import React, { useState, useEffect } from "react";
import { Line } from "react-chartjs-2";
import { Chart as ChartJS, LineElement, PointElement, CategoryScale, LinearScale } from "chart.js";

ChartJS.register(LineElement, PointElement, CategoryScale, LinearScale);

const WebSocketURL = "ws://localhost:8080/ws";

const Dashboard = () => {
  const [sensorData, setSensorData] = useState([]);

  useEffect(() => {
    const socket = new WebSocket(WebSocketURL);
    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setSensorData((prev) => [...prev.slice(-9), data]); // Keep last 10 records
    };
    return () => socket.close();
  }, []);

  const chartData = {
    labels: sensorData.map((data) => new Date(data.timestamp).toLocaleTimeString()),
    datasets: [
      {
        label: "Temperature (°C)",
        data: sensorData.map((data) => data.temperature),
        borderColor: "red",
        fill: false,
      },
      {
        label: "AQI",
        data: sensorData.map((data) => data.aqi),
        borderColor: "blue",
        fill: false,
      },
      {
        label: "Occupancy",
        data: sensorData.map((data) => data.occupancy),
        borderColor: "green",
        fill: false,
      },
      {
        label: "Reliability Score",
        data: sensorData.map((data) => data.reliability),
        borderColor: "purple",
        fill: false,
      },
    ],
  };

  return (
  <div style={{ width: "90%", maxWidth: "1600px", margin: "auto", padding: "20px", textAlign: "center" }}>
  <h2 style={{ marginBottom: "20px" }}>Real-Time Sensor Data</h2>

  <hr style={{ border: "1px solid #ccc", marginBottom: "20px" }} />

  <div style={{ display: "flex", flexDirection: "row", gap: "20px", alignItems: "center", justifyContent: "center", flexWrap: "wrap" }}>
    <div style={{ flex: "2", minWidth: "600px" }}>
      <Line data={chartData} />
      {/* Legend for colors */}
      <div style={{ display: "flex", justifyContent: "center", gap: "20px", marginTop: "10px" }}>
        <div style={{ display: "flex", alignItems: "center", gap: "5px" }}>
          <span style={{ width: "12px", height: "12px", backgroundColor: "red", display: "inline-block", borderRadius: "50%" }}></span>
          <span>Temperature</span>
        </div>
        <div style={{ display: "flex", alignItems: "center", gap: "5px" }}>
          <span style={{ width: "12px", height: "12px", backgroundColor: "green", display: "inline-block", borderRadius: "50%" }}></span>
          <span>Occupancy</span>
        </div>
        <div style={{ display: "flex", alignItems: "center", gap: "5px" }}>
          <span style={{ width: "12px", height: "12px", backgroundColor: "blue", display: "inline-block", borderRadius: "50%" }}></span>
          <span>AQI</span>
        </div>
        <div style={{ display: "flex", alignItems: "center", gap: "5px" }}>
          <span style={{ width: "12px", height: "12px", backgroundColor: "purple", display: "inline-block", borderRadius: "50%" }}></span>
          <span>Reliability</span>
        </div>
      </div>
    </div>

    <hr style={{ border: "1px solid #ccc", height: "100%", margin: "0 20px" }} />

    <div style={{ flex: "1", minWidth: "300px", textAlign: "center", padding: "20px", borderRadius: "10px", background: "#f8f8f8", boxShadow: "0 4px 8px rgba(0,0,0,0.1)" }}>
      <h3 style={{ color: "#333", marginBottom: "15px" }}>Reliability Score</h3>
      {sensorData.length > 0 ? (
        <p style={{ fontSize: "24px", fontWeight: "bold", color: "purple", margin: "0" }}>
          {sensorData[sensorData.length - 1].reliability.toFixed(2)} / 100
        </p>
      ) : (
        <p>Waiting for data...</p>
      )}
    </div>
  </div>

  <hr style={{ border: "1px solid #ccc", marginTop: "20px" }} />
</div>

  );
};

export default Dashboard;
