# Project Journey And Thought Process during my Journey 


# Challange faced  and solution during development journey -

## How the Reliability Score is Calculated(challange) with solution -:
The `calculateReliability` function is designed to measure the reliability of a sensor based on two key factors:
1. **Variance in Temperature Readings**: How much the temperature fluctuates over time.
2. **Frequency of Readings**: How consistently new temperature data is received.

#### 1. **Storing the Latest 10 Temperature Readings**
- The function keeps track of the last 10 temperature values for each sensor in `sensorDataHistory`.
- If more than 10 values exist, the oldest value is removed to make space for the new one.

#### 2. **Tracking the Timestamps of Readings**
- A separate list, `sensorTimestamps`, maintains the timestamps of the last 10 readings.
- Like the temperature values, it ensures that only the latest 10 timestamps are stored.

#### 3. **Calculating Variance (Data Stability Check)**
- Variance helps in determining how much the temperature values fluctuate.
- If the variance is high, it means the sensor readings are inconsistent.
- That's why we are generating temprature with less fluctuations between 10 to 30 degree.

#### 4. **Calculating Frequency (Consistency of Data Arrival)**
- The function calculates the average time difference between consecutive readings.
- If readings are coming at regular intervals, the frequency is high (closer to 1).
- If readings are irregular or missing, the frequency decreases.

#### 5. **Combining Variance and Frequency to Compute Reliability Score**
- The reliability formula is:
  
  **reliability = 100 - (variance * 2) + (frequency * 50)**
  
  - **Variance negatively impacts reliability** (higher variance reduces the score).
  - **Higher frequency increases reliability**, ensuring regular readings contribute to a better score.
  - The score is capped at **100** (max) and **10** (min) to prevent extreme values.


## What will i improve with more time - 

### 1. **Better Logging for Debugging and Analysis**
-Should  add structured logging to track when a sensor sends data and any anomalies in readings.
- Log when a sensor stops sending data, indicating possible failures.


### 2. **Optimize Data Storage and Access**
- Instead of using in-memory maps, consider a **time-series database** (like InfluxDB or Prometheus) for handling large-scale sensor data.
- This allows efficient querying and long-term analytics with large-scale.


### 3. **User-Friendly Dashboard for Visualization**
- Display sensor reliability scores in a UI with **color-coded indicators** (green = reliable, red = inconsistent).
- Show **historical trends** to analyze sensor performance over time.


# One key technical decision that i have made  - 


Key Technical Decision: Using Configuration and Makefile for Better Readability and Clean Code Practices

To keep the project well-structured and maintainable, I chose to use configuration files and a Makefile.

Why?

Separation of Concerns: Instead of hardcoding database credentials or server details, they are defined in a YAML configuration file, making the system more flexible and environment-agnostic.

Improved Readability: A Makefile helps standardize build, run, and deployment commands, reducing manual errors and ensuring a consistent development workflow.


## Conclusion
The current reliability calculation effectively captures **data consistency and frequency**, making it a good indicator of sensor performance.