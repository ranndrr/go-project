# NapTime: A Simple Sleep and Health Monitoring Application

## 📋 Description
NapTime is a terminal-based backend application designed to help users record, monitor, and analyze their sleep patterns and daily health activities, supporting a healthier and more structured lifestyle. It focuses on tracking **sleep history**, **bedtime and wake-up times**, and **sleep quality**.

## ✨ Features

- ✅ Add, edit, and delete daily sleep records  
- ⏱️ Automatic sleep duration calculation  
- 📊 Weekly reports: 7-day summary & average duration  
- 🔍 Search sleep data (Sequential & Binary Search)  
- 📑 Sort sleep data by duration or date (Selection & Insertion Sort)

## 🛠️ Specifications

- Users can **add**, **edit**, and **delete** sleep records that include bedtime and wake-up time.  
- The system **automatically calculates sleep duration** and provides **healthy sleep advice**.  
- Users can **search** sleep records by date using:
  - **Sequential Search**
  - **Binary Search**
- Users can **sort** sleep records by:
  - **Duration**
  - **Date**  
  using:
  - **Selection Sort**
  - **Insertion Sort**
- Reports include:
  - 🔄 Recap of **sleep duration over the past 7 days**
  - 📈 **Average sleep duration per week**
  - ⏱️ **Longest and shortest sleep duration** within the past 7 days

## 🧾 Technical Notes

- Maximum number of stored sleep records: **100 entries** (`NMAX`)  
- Built-in validation for bedtime and wake-up time
- Only allows sleep records **from the year 2025 and onward**  
- Suggestions are shown if:
  - 🕚 Sleep starts **after 11:00 PM**
  - 😴 Total sleep is **less than 7 hours or more than 9 hours a day**

## 🚀 How to Run

Make sure you have **Go** installed on your system.  
To run the application, open your terminal and type:

```bash
