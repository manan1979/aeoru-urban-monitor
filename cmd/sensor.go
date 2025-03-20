package main

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net"
	"sync"
	"time"
)

type SensorData struct {
	ID          int     `json:"id"`
	Temperature float64 `json:"temperature"`
	AQI         int     `json:"aqi"`
	Occupancy   int     `json:"occupancy"`
	Timestamp   string  `json:"timestamp"`
	Reliability float64 `json:"reliability"`
}

var (
	sensorDataHistory = make(map[int][]float64) // Sensor ID → Temperature readings
	sensorTimestamps  = make(map[int][]int64)   // Sensor ID → Timestamps of last readings
	mu                sync.Mutex
)

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	for {
		data := SensorData{
			Temperature: rand.Float64()*20 + 10, // Random temp between 10-30°C
			AQI:         rand.Intn(200),         // Random AQI 0-200
			Occupancy:   rand.Intn(50),          // Random occupancy 0-50
			Timestamp:   time.Now().Format(time.RFC3339),
		}

		data.Reliability = calculateReliability(data.ID, data.Temperature)

		lastInsertID, err := storeSensorData(data)
		if err != nil {
			log.Println("Database insert error:", err)
			return
		}
		data.ID = lastInsertID

		jsonData, _ := json.Marshal(data)
		_, err = conn.Write(jsonData)
		if err != nil {
			log.Println("TCP Write error:", err)
			return
		}

		broadcast <- data
		time.Sleep(3 * time.Second)
	}
}

func storeSensorData(data SensorData) (int, error) {
	res, err := db.Exec("INSERT INTO sensor_data (temperature, aqi, occupancy, timestamp, reliability) VALUES (?, ?, ?, ?, ?)",
		data.Temperature, data.AQI, data.Occupancy, data.Timestamp, data.Reliability)
	if err != nil {
		log.Println("Database insertion error:", err)
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Println("Failed to retrieve last insert ID:", err)
		return 0, err
	}

	return int(lastInsertID), nil
}

func calculateReliability(sensorID int, temperature float64) float64 {
	mu.Lock()
	defer mu.Unlock()

	// Store readings (keep only last 10)
	if len(sensorDataHistory[sensorID]) >= 10 {
		sensorDataHistory[sensorID] = sensorDataHistory[sensorID][1:]
	}
	sensorDataHistory[sensorID] = append(sensorDataHistory[sensorID], temperature)

	// Store timestamps (keep only last 10)
	currentTimestamp := time.Now().Unix()
	if len(sensorTimestamps[sensorID]) >= 10 {
		sensorTimestamps[sensorID] = sensorTimestamps[sensorID][1:]
	}
	sensorTimestamps[sensorID] = append(sensorTimestamps[sensorID], currentTimestamp)

	// Compute variance
	var variance float64
	if len(sensorDataHistory[sensorID]) > 1 {
		mean := calculateMean(sensorDataHistory[sensorID])
		var sumSquares float64
		for _, val := range sensorDataHistory[sensorID] {
			sumSquares += math.Pow(val-mean, 2)
		}
		variance = sumSquares / float64(len(sensorDataHistory[sensorID]))
	}

	// Compute update frequency
	var frequency float64
	if len(sensorTimestamps[sensorID]) > 1 {
		timeDiffs := make([]float64, len(sensorTimestamps[sensorID])-1)
		for i := 1; i < len(sensorTimestamps[sensorID]); i++ {
			timeDiffs[i-1] = float64(sensorTimestamps[sensorID][i] - sensorTimestamps[sensorID][i-1])
		}
		meanTimeDiff := calculateMean(timeDiffs)

		if meanTimeDiff > 0 {
			frequency = math.Min(1.0/meanTimeDiff, 1) // Ensure frequency is between 0-1
		}
	}

	// Compute reliability score (normalize between 0-100)
	reliability := 100 - (variance * 2) + (frequency * 50) // Reduce variance impact, increase frequency weight

	// Ensure reliability is within valid range
	if reliability > 100 {
		reliability = 100
	} else if reliability < 10 { // Set a minimum reliability value
		reliability = 10
	}

	return reliability
}

func calculateMean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}
